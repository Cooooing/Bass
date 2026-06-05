package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/checkinrecord"
	"user/internal/data/gen/checkinstat"

	utilent "common/pkg/util/ent"
)

var _ repo.CheckinRepo = (*CheckinRepo)(nil)

type CheckinRepo struct {
	db *gen.Client
}

func NewCheckinRepo(db *gen.Client) repo.CheckinRepo {
	return &CheckinRepo{
		db: db,
	}
}

func (r *CheckinRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *CheckinRepo) FindStatByUserID(ctx context.Context, userID int64) (*model.CheckinStat, error) {
	tx := r.getClient(ctx)
	s, err := tx.CheckinStat.Query().Where(checkinstat.UserID(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.CheckinStat{
		ID:                 s.ID,
		UserID:             s.UserID,
		TotalOnlineMinutes: s.TotalOnlineMinutes,
		CurrentStreak:      s.CurrentStreak,
		LongestStreak:      s.LongestStreak,
	}, nil
}

func (r *CheckinRepo) UpsertRecord(ctx context.Context, record *model.CheckinRecord) (*model.CheckinRecord, error) {
	tx := r.getClient(ctx)
	existing, err := tx.CheckinRecord.Query().
		Where(checkinrecord.UserID(record.UserID)).
		Where(checkinrecord.DateEQ(*record.Date)).
		Only(ctx)
	if err != nil && !gen.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		saved, err := tx.CheckinRecord.UpdateOneID(existing.ID).
			SetNillableOnlineMinutes(record.OnlineMinutes).
			SetNillableActivity(record.Activity).
			SetChecked(record.Checked).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.CheckinRecord{
			ID:            saved.ID,
			UserID:        saved.UserID,
			Date:          new(saved.Date),
			OnlineMinutes: saved.OnlineMinutes,
			Activity:      saved.Activity,
			Checked:       saved.Checked,
		}, nil
	}
	saved, err := tx.CheckinRecord.Create().
		SetUserID(record.UserID).
		SetDate(*record.Date).
		SetNillableOnlineMinutes(record.OnlineMinutes).
		SetNillableActivity(record.Activity).
		SetChecked(record.Checked).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CheckinRecord{
		ID:            saved.ID,
		UserID:        saved.UserID,
		Date:          new(saved.Date),
		OnlineMinutes: saved.OnlineMinutes,
		Activity:      saved.Activity,
		Checked:       saved.Checked,
	}, nil
}

func (r *CheckinRepo) UpsertStat(ctx context.Context, stat *model.CheckinStat) (*model.CheckinStat, error) {
	existing, err := r.FindStatByUserID(ctx, stat.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		saved, err := tx.CheckinStat.Create().
			SetUserID(stat.UserID).
			SetNillableTotalOnlineMinutes(stat.TotalOnlineMinutes).
			SetNillableCurrentStreak(stat.CurrentStreak).
			SetNillableLongestStreak(stat.LongestStreak).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.CheckinStat{
			ID:                 saved.ID,
			UserID:             saved.UserID,
			TotalOnlineMinutes: saved.TotalOnlineMinutes,
			CurrentStreak:      saved.CurrentStreak,
			LongestStreak:      saved.LongestStreak,
		}, nil
	}
	stat.ID = existing.ID
	tx := r.getClient(ctx)
	saved, err := tx.CheckinStat.UpdateOneID(stat.ID).
		SetNillableTotalOnlineMinutes(stat.TotalOnlineMinutes).
		SetNillableCurrentStreak(stat.CurrentStreak).
		SetNillableLongestStreak(stat.LongestStreak).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CheckinStat{
		ID:                 saved.ID,
		UserID:             saved.UserID,
		TotalOnlineMinutes: saved.TotalOnlineMinutes,
		CurrentStreak:      saved.CurrentStreak,
		LongestStreak:      saved.LongestStreak,
	}, nil
}
