package repo

import (
	"context"
	"fmt"
	"time"

	utilent "common/pkg/util/ent"

	"platform/internal/biz/model"
	bizrepo "platform/internal/biz/repo"
	"platform/internal/data/gen"
	"platform/internal/data/gen/holiday"
	"platform/internal/enum"

	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.HolidayRepo = (*HolidayRepo)(nil)

type HolidayRepo struct {
	db *gen.Client
}

func NewHolidayRepo(db *gen.Client) bizrepo.HolidayRepo {
	return &HolidayRepo{
		db: db,
	}
}

func (r *HolidayRepo) client(ctx context.Context) *gen.Client {
	if client, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return client
	}
	return r.db
}

func (r *HolidayRepo) GetByDate(ctx context.Context, date time.Time) (*model.HolidayInfo, error) {
	entity, err := r.client(ctx).Holiday.Query().Where(holiday.DateEQ(r.normalizeDate(date))).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.holidayModel(entity), nil
}

func (r *HolidayRepo) Upsert(ctx context.Context, info *model.HolidayInfo) (*model.HolidayInfo, error) {
	if info == nil {
		return nil, fmt.Errorf("holiday info is required")
	}
	var holidayName *string
	if info.HolidayName != "" {
		holidayName = &info.HolidayName
	}
	err := r.client(ctx).Holiday.Create().
		SetDate(r.normalizeDate(info.Date)).
		SetType(holiday.Type(info.Type)).
		SetDayName(info.DayName).
		SetNillableHolidayName(holidayName).
		SetDaysToHoliday(info.DaysToHoliday).
		SetDaysRemaining(info.DaysRemaining).
		OnConflict(entsql.ConflictColumns(holiday.FieldDate)).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.GetByDate(ctx, info.Date)
}

func (r *HolidayRepo) holidayModel(entity *gen.Holiday) *model.HolidayInfo {
	if entity == nil {
		return nil
	}
	holidayName := ""
	if entity.HolidayName != nil {
		holidayName = *entity.HolidayName
	}
	return &model.HolidayInfo{
		Date:          r.normalizeDate(entity.Date),
		Type:          enum.HolidayType(entity.Type),
		DayName:       entity.DayName,
		HolidayName:   holidayName,
		DaysToHoliday: entity.DaysToHoliday,
		DaysRemaining: entity.DaysRemaining,
	}
}

func (r *HolidayRepo) normalizeDate(date time.Time) time.Time {
	location := time.FixedZone("Asia/Shanghai", 8*60*60)
	local := date.In(location)
	return time.Date(local.Year(), local.Month(), local.Day(), 0, 0, 0, 0, location)
}
