package usecase

import (
	"context"
	"log/slog"
	"time"

	"platform/internal/biz/model"
	"platform/internal/biz/repo"
)

// HolidayUsecase 负责读取持久化的节假日状态，并执行外部数据刷新。
type HolidayUsecase struct {
	log         *slog.Logger
	client      repo.HolidayClient
	holidayRepo repo.HolidayRepo
}

func NewHolidayUsecase(
	logger *slog.Logger,
	client repo.HolidayClient,
	holidayRepo repo.HolidayRepo,
) *HolidayUsecase {
	return &HolidayUsecase{
		log:         logger,
		client:      client,
		holidayRepo: holidayRepo,
	}
}

func (u *HolidayUsecase) Get(ctx context.Context) (*model.HolidayInfo, error) {
	date := u.normalizeDate(time.Now())
	info, err := u.holidayRepo.GetByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return info, nil
	}
	return &model.HolidayInfo{Date: date}, nil
}

func (u *HolidayUsecase) Refresh(ctx context.Context, date time.Time) error {
	target := u.normalizeDate(date)
	info, err := u.client.Fetch(ctx, target)
	if err != nil {
		if u.log != nil {
			u.log.ErrorContext(ctx, "fetch holiday info failed", "error", err)
		}
		return err
	}
	if _, err = u.holidayRepo.Upsert(ctx, info); err != nil {
		return err
	}
	return nil
}

func (u *HolidayUsecase) normalizeDate(date time.Time) time.Time {
	if date.IsZero() {
		date = time.Now()
	}
	location := time.FixedZone("Asia/Shanghai", 8*60*60)
	local := date.In(location)
	return time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, location)
}
