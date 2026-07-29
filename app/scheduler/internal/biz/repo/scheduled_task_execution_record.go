package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type ScheduledTaskExecutionRecordRepo interface {
	Get(ctx context.Context, req *ScheduledTaskExecutionRecordGetReq) (*model.ScheduledTaskExecutionRecord, error)
	List(ctx context.Context, req *ScheduledTaskExecutionRecordGetReq) ([]*model.ScheduledTaskExecutionRecord, error)
	Page(ctx context.Context, req *ScheduledTaskExecutionRecordPageReq) (*ScheduledTaskExecutionRecordPageResp, error)
	HasRunning(ctx context.Context, req *ScheduledTaskExecutionRecordHasRunningReq) (bool, error)
	Create(ctx context.Context, req *ScheduledTaskExecutionRecordCreateReq) (*ScheduledTaskExecutionRecordCreateResp, error)
	Claim(ctx context.Context, req *ScheduledTaskExecutionRecordClaimReq) (*ScheduledTaskExecutionRecordClaimResp, error)
	MarkFinished(ctx context.Context, req *ScheduledTaskExecutionRecordMarkFinishedReq) (*ScheduledTaskExecutionRecordMarkFinishedResp, error)
	MarkCanceled(ctx context.Context, id int64, finishedAt time.Time) (*model.ScheduledTaskExecutionRecord, error)
}

type ScheduledTaskExecutionRecordGetReq struct {
	ID              *int64
	IDs             []int64
	ScheduledTaskID *int64
	ScheduledAt     *time.Time
	ScheduleKey     *string
	Status          *schedulerenum.TaskExecutionStatus
	TriggerType     *schedulerenum.TaskTriggerType
}

type ScheduledTaskExecutionRecordPageReq struct {
	Page *common.PageReq
	ScheduledTaskExecutionRecordGetReq
}

type ScheduledTaskExecutionRecordPageResp struct {
	Rows []*model.ScheduledTaskExecutionRecord
	Page *common.PageResp
}

type ScheduledTaskExecutionRecordHasRunningReq struct {
	ScheduledTaskID int64
}

type ScheduledTaskExecutionRecordCreateReq struct {
	Record *model.ScheduledTaskExecutionRecord
	Status schedulerenum.TaskExecutionStatus
}

type ScheduledTaskExecutionRecordCreateResp struct {
	Row      *model.ScheduledTaskExecutionRecord
	Created  bool
	Conflict bool
}

type ScheduledTaskExecutionRecordClaimReq struct {
	ID        int64
	WorkerID  string
	StartedAt time.Time
}

type ScheduledTaskExecutionRecordClaimResp struct {
	Row     *model.ScheduledTaskExecutionRecord
	Claimed bool
}

type ScheduledTaskExecutionRecordMarkFinishedReq struct {
	ID         int64
	WorkerID   string
	Attempt    int32
	Status     schedulerenum.TaskExecutionStatus
	FinishedAt time.Time
	Duration   time.Duration
	LastError  string
}

type ScheduledTaskExecutionRecordMarkFinishedResp struct {
	Row     *model.ScheduledTaskExecutionRecord
	Updated bool
}
