package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type DelayedTaskRepo interface {
	Register(ctx context.Context, row *model.DelayedTask) (*model.DelayedTask, error)
	Get(ctx context.Context, req *DelayedTaskGetReq) (*model.DelayedTask, error)
	Page(ctx context.Context, req *DelayedTaskPageReq) (*DelayedTaskPageResp, error)
	Cancel(ctx context.Context, req *DelayedTaskGetReq) (bool, error)
	ListDue(ctx context.Context, now time.Time, limit int) ([]*model.DelayedTask, error)
	MarkRunning(ctx context.Context, id int64, workerID string, lockExpiresAt time.Time) (bool, *model.DelayedTask, error)
	MarkSuccess(ctx context.Context, id int64, finishedAt time.Time) (*model.DelayedTask, error)
	MarkFailed(ctx context.Context, id int64, attempt int32, final bool, nextRunAt time.Time, lastError string) (*model.DelayedTask, error)
}

type DelayedTaskGetReq struct {
	ID             *int64
	IdempotencyKey *string
	TaskName       *string
	Status         *schedulerenum.DelayedTaskStatus
}

type DelayedTaskPageReq struct {
	Page *common.PageReq
	DelayedTaskGetReq
}

type DelayedTaskPageResp struct {
	Rows []*model.DelayedTask
	Page *common.PageResp
}
