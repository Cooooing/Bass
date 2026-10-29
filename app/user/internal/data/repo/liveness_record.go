package repo

import (
	utilent "common/pkg/util/ent"
	"context"
	"time"
	"user/internal/biz/model"
	bizrepo "user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/livenessrecord"
	"user/internal/enum"
)

var _ bizrepo.LivenessRecordRepo = (*LivenessRecordRepo)(nil)

type LivenessRecordRepo struct{ db *gen.Client }

func NewLivenessRecordRepo(db *gen.Client) bizrepo.LivenessRecordRepo {
	return &LivenessRecordRepo{db: db}
}
func (r *LivenessRecordRepo) client(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}
func (r *LivenessRecordRepo) Get(ctx context.Context, userID int64, date time.Time) (*model.LivenessRecord, error) {
	row, err := r.client(ctx).LivenessRecord.Query().Where(livenessrecord.UserIDEQ(userID), livenessrecord.DateEQ(date)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}
func (r *LivenessRecordRepo) Save(ctx context.Context, item *model.LivenessRecord) (*model.LivenessRecord, error) {
	client := r.client(ctx)
	old, err := r.Get(ctx, item.UserID, *item.Date)
	if err != nil {
		return nil, err
	}
	if old != nil {
		return old, nil
	}
	row, err := client.LivenessRecord.Create().SetUserID(item.UserID).SetDate(*item.Date).SetScore(item.Score).SetRewardStatus(livenessrecord.RewardStatus(item.RewardStatus)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}
func (r *LivenessRecordRepo) UpdateRewardStatus(ctx context.Context, id int64, status enum.LivenessRewardStatus, claimedAt *time.Time) (*model.LivenessRecord, error) {
	update := r.client(ctx).LivenessRecord.UpdateOneID(id).SetRewardStatus(livenessrecord.RewardStatus(status))
	if claimedAt != nil {
		update.SetClaimedAt(*claimedAt)
	}
	row, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}
func (r *LivenessRecordRepo) model(row *gen.LivenessRecord) *model.LivenessRecord {
	return &model.LivenessRecord{ID: row.ID, UserID: row.UserID, Date: new(row.Date), Score: row.Score, RewardStatus: enum.LivenessRewardStatus(row.RewardStatus), ClaimedAt: row.ClaimedAt}
}
