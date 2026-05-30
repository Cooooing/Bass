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
		TotalOnlineMinutes: new(s.TotalOnlineMinutes),
		CurrentStreak:      new(s.CurrentStreak),
		LongestStreak:      new(s.LongestStreak),
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
		update := tx.CheckinRecord.UpdateOneID(existing.ID).
			SetChecked(record.Checked)
		if record.OnlineMinutes != nil {
			update.SetOnlineMinutes(*record.OnlineMinutes)
		}
		if record.Activity != nil {
			update.SetActivity(*record.Activity)
		}
		saved, err := update.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.CheckinRecord{
			ID:            saved.ID,
			UserID:        saved.UserID,
			Date:          new(saved.Date),
			OnlineMinutes: new(saved.OnlineMinutes),
			Activity:      new(saved.Activity),
			Checked:       saved.Checked,
		}, nil
	}
	create := tx.CheckinRecord.Create().
		SetUserID(record.UserID).
		SetDate(*record.Date).
		SetChecked(record.Checked)
	if record.OnlineMinutes != nil {
		create.SetOnlineMinutes(*record.OnlineMinutes)
	}
	if record.Activity != nil {
		create.SetActivity(*record.Activity)
	}
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CheckinRecord{
		ID:            saved.ID,
		UserID:        saved.UserID,
		Date:          new(saved.Date),
		OnlineMinutes: new(saved.OnlineMinutes),
		Activity:      new(saved.Activity),
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
		create := tx.CheckinStat.Create().
			SetUserID(stat.UserID)
		if stat.TotalOnlineMinutes != nil {
			create.SetTotalOnlineMinutes(*stat.TotalOnlineMinutes)
		}
		if stat.CurrentStreak != nil {
			create.SetCurrentStreak(*stat.CurrentStreak)
		}
		if stat.LongestStreak != nil {
			create.SetLongestStreak(*stat.LongestStreak)
		}
		saved, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.CheckinStat{
			ID:                 saved.ID,
			UserID:             saved.UserID,
			TotalOnlineMinutes: new(saved.TotalOnlineMinutes),
			CurrentStreak:      new(saved.CurrentStreak),
			LongestStreak:      new(saved.LongestStreak),
		}, nil
	}
	if stat.TotalOnlineMinutes == nil && stat.CurrentStreak == nil && stat.LongestStreak == nil {
		return existing, nil
	}
	stat.ID = existing.ID
	tx := r.getClient(ctx)
	update := tx.CheckinStat.UpdateOneID(stat.ID)
	if stat.TotalOnlineMinutes != nil {
		update.SetTotalOnlineMinutes(*stat.TotalOnlineMinutes)
	}
	if stat.CurrentStreak != nil {
		update.SetCurrentStreak(*stat.CurrentStreak)
	}
	if stat.LongestStreak != nil {
		update.SetLongestStreak(*stat.LongestStreak)
	}
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.CheckinStat{
		ID:                 saved.ID,
		UserID:             saved.UserID,
		TotalOnlineMinutes: new(saved.TotalOnlineMinutes),
		CurrentStreak:      new(saved.CurrentStreak),
		LongestStreak:      new(saved.LongestStreak),
	}, nil
}
