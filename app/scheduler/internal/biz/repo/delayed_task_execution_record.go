package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type DelayedTaskExecutionRecordRepo interface {
	Get(ctx context.Context, req *DelayedTaskExecutionRecordGetReq) (*model.DelayedTaskExecutionRecord, error)
	List(ctx context.Context, req *DelayedTaskExecutionRecordGetReq) ([]*model.DelayedTaskExecutionRecord, error)
	Page(ctx context.Context, req *DelayedTaskExecutionRecordPageReq) (*DelayedTaskExecutionRecordPageResp, error)
	CreatePending(ctx context.Context, record *model.DelayedTaskExecutionRecord) (*DelayedTaskExecutionRecordCreateResp, error)
	Claim(ctx context.Context, req *DelayedTaskExecutionRecordClaimReq) (*DelayedTaskExecutionRecordClaimResp, error)
	MarkFinished(ctx context.Context, req *DelayedTaskExecutionRecordMarkFinishedReq) (*model.DelayedTaskExecutionRecord, error)
	MarkCanceled(ctx context.Context, req *DelayedTaskExecutionRecordGetReq, finishedAt time.Time) (*model.DelayedTaskExecutionRecord, error)
}

type DelayedTaskExecutionRecordGetReq struct {
	ID             *int64
	IDs            []int64
	DelayedTaskID  *int64
	IdempotencyKey *string
	Status         *schedulerenum.TaskExecutionStatus
	TriggerType    *schedulerenum.TaskTriggerType
}

type DelayedTaskExecutionRecordPageReq struct {
	Page *common.PageReq
	DelayedTaskExecutionRecordGetReq
}

type DelayedTaskExecutionRecordPageResp struct {
	Rows []*model.DelayedTaskExecutionRecord
	Page *common.PageResp
}

type DelayedTaskExecutionRecordCreateResp struct {
	Row      *model.DelayedTaskExecutionRecord
	Created  bool
	Conflict bool
}

type DelayedTaskExecutionRecordClaimReq struct {
	ID        int64
	WorkerID  string
	StartedAt time.Time
}

type DelayedTaskExecutionRecordClaimResp struct {
	Row     *model.DelayedTaskExecutionRecord
	Claimed bool
}

type DelayedTaskExecutionRecordMarkFinishedReq struct {
	ID         int64
	WorkerID   string
	Attempt    int32
	Status     schedulerenum.TaskExecutionStatus
	FinishedAt time.Time
	Duration   time.Duration
	LastError  string
}
