package usecase

import (
	"context"
	"log/slog"
	"maps"
	"time"

	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/enum"
)

// HolidayUsecase 负责读取持久化的节假日状态，并按范围刷新外部数据。
type HolidayUsecase struct {
	log         *slog.Logger
	client      repo.HolidayClient
	holidayRepo repo.HolidayRepo
	cacheRepo   repo.HolidayConfigCacheRepo
	location    *time.Location
	configTTL   time.Duration
}

func NewHolidayUsecase(
	logger *slog.Logger,
	client repo.HolidayClient,
	holidayRepo repo.HolidayRepo,
	cacheRepo repo.HolidayConfigCacheRepo,
) *HolidayUsecase {
	return &HolidayUsecase{
		log:         logger,
		client:      client,
		holidayRepo: holidayRepo,
		cacheRepo:   cacheRepo,
		location:    time.FixedZone("Asia/Shanghai", 8*60*60),
		configTTL:   12 * time.Hour,
	}
}

func (u *HolidayUsecase) Get(ctx context.Context) (*model.HolidayInfo, error) {
	date := u.dayStart(time.Now())
	info, err := u.holidayRepo.GetByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return info, nil
	}
	return &model.HolidayInfo{Date: date}, nil
}

// Refresh 写入 [start, end] 范围内的每一天，范围由调用方给出。
func (u *HolidayUsecase) Refresh(ctx context.Context, start time.Time, end time.Time) error {
	begin := u.dayStart(start)
	finish := u.dayStart(end)
	if finish.Before(begin) {
		finish = begin
	}
	// 多加载一年，nextHoliday 向前扫描时才不会跨年查空。
	entries, err := u.loadEntries(ctx, begin, finish.AddDate(0, 0, 1))
	if err != nil {
		if u.log != nil {
			u.log.ErrorContext(ctx, "fetch holiday info failed", "error", err)
		}
		return err
	}
	for day := begin; !day.After(finish); day = day.AddDate(0, 0, 1) {
		info := u.buildDay(day, entries)
		if _, err = u.holidayRepo.Upsert(ctx, info); err != nil {
			return err
		}
	}
	return nil
}

func (u *HolidayUsecase) loadEntries(ctx context.Context, begin time.Time, end time.Time) (map[string]*model.HolidayEntry, error) {
	entries := make(map[string]*model.HolidayEntry)
	for year := begin.Year(); year <= end.Year(); year++ {
		// 跨年时下一年配置可能尚未发布，缺失时先跳过，等发布后再刷新。
		yearEntries, err := u.yearEntries(ctx, year)
		if err != nil {
			continue
		}
		maps.Copy(entries, yearEntries)
	}
	return entries, nil
}

// yearEntries 从 Redis 读取整年配置，缺失或过期时回源并写回缓存。
func (u *HolidayUsecase) yearEntries(ctx context.Context, year int) (map[string]*model.HolidayEntry, error) {
	if u.cacheRepo != nil {
		cached, err := u.cacheRepo.Get(ctx, year)
		if err != nil {
			return nil, err
		}
		if cached != nil && time.Since(cached.FetchedAt) < u.configTTL {
			return cached.Entries, nil
		}
	}
	entries, err := u.client.FetchYear(ctx, year)
	if err != nil {
		return nil, err
	}
	u.writeCache(ctx, year, entries)
	return entries, nil
}

func (u *HolidayUsecase) writeCache(ctx context.Context, year int, entries map[string]*model.HolidayEntry) {
	if u.cacheRepo == nil {
		return
	}
	if err := u.cacheRepo.SetIfFresher(ctx, year, &model.HolidayConfigCache{
		Entries:   entries,
		FetchedAt: time.Now(),
	}); err != nil && u.log != nil {
		u.log.WarnContext(ctx, "write holiday config cache failed", "error", err)
	}
}

func (u *HolidayUsecase) buildDay(day time.Time, entries map[string]*model.HolidayEntry) *model.HolidayInfo {
	entry := entries[day.Format("2006-01-02")]
	info := &model.HolidayInfo{Date: day, Type: enum.HolidayTypeWorkday, DayName: u.weekdayName(day)}
	switch {
	case entry != nil && entry.Holiday:
		info.Type = enum.HolidayTypeFestival
		info.DayName = entry.Name
	case entry != nil:
		info.Type = enum.HolidayTypeMakeupWorkday
		info.DayName = entry.Name
	case u.isWeekend(day):
		info.Type = enum.HolidayTypeWeekend
	}
	switch info.Type {
	case enum.HolidayTypeWorkday, enum.HolidayTypeMakeupWorkday:
		info.HolidayName, info.DaysToHoliday = u.nextRestDay(day, entries)
	case enum.HolidayTypeWeekend, enum.HolidayTypeFestival:
		info.DaysRemaining = u.daysUntilWorkday(day, entries)
	}
	return info
}

// nextRestDay 返回下一个休息日（含周末和节日）的名称与相距天数，调休工作日不算休息日。
func (u *HolidayUsecase) nextRestDay(day time.Time, entries map[string]*model.HolidayEntry) (string, int32) {
	for offset := 1; offset <= 60; offset++ {
		next := day.AddDate(0, 0, offset)
		if !u.isRestDay(next, entries) {
			continue
		}
		if entry := entries[next.Format("2006-01-02")]; entry != nil {
			return entry.Name, int32(offset)
		}
		return u.weekdayName(next), int32(offset)
	}
	return "", 0
}

// daysUntilWorkday 返回距离下一个工作日的天数，用于假期余额展示。
func (u *HolidayUsecase) daysUntilWorkday(day time.Time, entries map[string]*model.HolidayEntry) int32 {
	for offset := 1; offset <= 60; offset++ {
		next := day.AddDate(0, 0, offset)
		if u.isRestDay(next, entries) {
			continue
		}
		return int32(offset)
	}
	return 0
}

func (u *HolidayUsecase) isRestDay(day time.Time, entries map[string]*model.HolidayEntry) bool {
	entry := entries[day.Format("2006-01-02")]
	if entry != nil {
		return entry.Holiday
	}
	return u.isWeekend(day)
}

func (*HolidayUsecase) isWeekend(day time.Time) bool {
	return day.Weekday() == time.Saturday || day.Weekday() == time.Sunday
}

func (u *HolidayUsecase) weekdayName(day time.Time) string {
	names := [...]string{"周日", "周一", "周二", "周三", "周四", "周五", "周六"}
	return names[int(day.In(u.location).Weekday())]
}

func (u *HolidayUsecase) dayStart(date time.Time) time.Time {
	if date.IsZero() {
		date = time.Now()
	}
	local := date.In(u.location)
	return time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, u.location)
}
