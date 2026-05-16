package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/usercheckinrecord"
	"user/internal/data/ent/gen/usercheckinstat"

	utilent "common/pkg/util/ent"
)

type UserCheckinRepo struct {
	*base.BaseData
}

func NewUserCheckinRepo(repo *base.BaseData) repo.UserCheckinRepo {
	return &UserCheckinRepo{
		BaseData: repo,
	}
}

func (r *UserCheckinRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.Db
}

func checkinStatToDomain(s *gen.UserCheckinStat) *model.UserCheckinStat {
	return &model.UserCheckinStat{
		ID:                 s.ID,
		UserID:             s.UserID,
		TotalOnlineMinutes: s.TotalOnlineMinutes,
		CurrentStreak:      s.CurrentStreak,
		LongestStreak:      s.LongestStreak,
	}
}

func checkinRecordToDomain(r *gen.UserCheckinRecord) *model.UserCheckinRecord {
	return &model.UserCheckinRecord{
		ID:            r.ID,
		UserID:        r.UserID,
		Date:          r.Date,
		OnlineMinutes: r.OnlineMinutes,
		Activity:      r.Activity,
		Checked:       r.Checked,
	}
}

func (r *UserCheckinRepo) GetStatByUserID(ctx context.Context, userID int64) (*model.UserCheckinStat, error) {
	tx := r.getClient(ctx)
	s, err := tx.UserCheckinStat.Query().Where(usercheckinstat.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return checkinStatToDomain(s), nil
}

func (r *UserCheckinRepo) UpsertRecord(ctx context.Context, record *model.UserCheckinRecord) (*model.UserCheckinRecord, error) {
	tx := r.getClient(ctx)
	existing, err := tx.UserCheckinRecord.Query().
		Where(usercheckinrecord.UserID(record.UserID)).
		Where(usercheckinrecord.DateEQ(*record.Date)).
		Only(ctx)
	if err != nil && !gen.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		saved, err := tx.UserCheckinRecord.UpdateOneID(existing.ID).
			SetNillableOnlineMinutes(record.OnlineMinutes).
			SetNillableActivity(record.Activity).
			SetChecked(record.Checked).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return checkinRecordToDomain(saved), nil
	}
	saved, err := tx.UserCheckinRecord.Create().
		SetUserID(record.UserID).
		SetDate(*record.Date).
		SetNillableOnlineMinutes(record.OnlineMinutes).
		SetNillableActivity(record.Activity).
		SetChecked(record.Checked).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return checkinRecordToDomain(saved), nil
}

func (r *UserCheckinRepo) UpdateStat(ctx context.Context, stat *model.UserCheckinStat) (*model.UserCheckinStat, error) {
	tx := r.getClient(ctx)
	saved, err := tx.UserCheckinStat.UpdateOneID(stat.ID).
		SetNillableTotalOnlineMinutes(stat.TotalOnlineMinutes).
		SetNillableCurrentStreak(stat.CurrentStreak).
		SetNillableLongestStreak(stat.LongestStreak).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return checkinStatToDomain(saved), nil
}
