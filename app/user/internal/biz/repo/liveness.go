package repo

import (
	"context"
	"time"

	"user/internal/biz/model"
	"user/internal/enum"
)

// LivenessRecordRepo 保存每日活跃度结算和领取状态。
type LivenessRecordRepo interface {
	Get(ctx context.Context, userID int64, date time.Time) (*model.LivenessRecord, error)
	Save(ctx context.Context, item *model.LivenessRecord) (*model.LivenessRecord, error)
	UpdateRewardStatus(ctx context.Context, id int64, status enum.LivenessRewardStatus, claimedAt *time.Time) (*model.LivenessRecord, error)
}
