package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/checkinrecord"
	"user/internal/data/gen/checkinstat"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.CheckinRepo = (*CheckinRepo)(nil)

type CheckinRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewCheckinRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.CheckinRepo {
	return &CheckinRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *CheckinRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func checkinStatToDomain(s *gen.CheckinStat) *model.CheckinStat {
	return &model.CheckinStat{
		ID:                 s.ID,
		UserID:             s.UserID,
		TotalOnlineMinutes: s.TotalOnlineMinutes,
		CurrentStreak:      s.CurrentStreak,
		LongestStreak:      s.LongestStreak,
	}
}

func checkinRecordToDomain(r *gen.CheckinRecord) *model.CheckinRecord {
	return &model.CheckinRecord{
		ID:            r.ID,
		UserID:        r.UserID,
		Date:          r.Date,
		OnlineMinutes: r.OnlineMinutes,
		Activity:      r.Activity,
		Checked:       r.Checked,
	}
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
	return checkinStatToDomain(s), nil
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
		return checkinRecordToDomain(saved), nil
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
	return checkinRecordToDomain(saved), nil
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
		return checkinStatToDomain(saved), nil
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
	return checkinStatToDomain(saved), nil
}
