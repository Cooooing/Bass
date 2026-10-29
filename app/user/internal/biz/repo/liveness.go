package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/enum"
)

// LivenessRecordRepo 保存每日活跃度结算和领取状态。
type LivenessRecordRepo interface {
	Get(context.Context, int64, time.Time) (*model.LivenessRecord, error)
	Save(context.Context, *model.LivenessRecord) (*model.LivenessRecord, error)
	UpdateRewardStatus(context.Context, int64, enum.LivenessRewardStatus, *time.Time) (*model.LivenessRecord, error)
}
