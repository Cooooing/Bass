package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type TaskExecutionRecordRepo interface {
	Get(ctx context.Context, req *TaskExecutionRecordGetReq) (*model.TaskExecutionRecord, error)
	List(ctx context.Context, req *TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, error)
	Map(ctx context.Context, req *TaskExecutionRecordGetReq) (map[int64]*model.TaskExecutionRecord, error)
	Count(ctx context.Context, req *TaskExecutionRecordGetReq) (int, error)
	Page(ctx context.Context, req *TaskExecutionRecordPageReq) (*TaskExecutionRecordPageResp, error)

	ExistsPeriod(ctx context.Context, req *TaskExecutionRecordExistsPeriodReq) (bool, error)
	Create(ctx context.Context, req *TaskExecutionRecordCreateReq) (*TaskExecutionRecordCreateResp, error)
	HasUnexpiredRunning(ctx context.Context, req *TaskExecutionRecordHasUnexpiredRunningReq) (bool, error)
	MarkUnknown(ctx context.Context, req *TaskExecutionRecordMarkUnknownReq) ([]*model.TaskExecutionRecord, error)
	MarkFinished(ctx context.Context, req *TaskExecutionRecordMarkFinishedReq) (*TaskExecutionRecordMarkFinishedResp, error)
}

type TaskExecutionRecordGetReq struct {
	ID          *int64
	IDs         []int64
	TaskID      *int64
	ScheduledAt *time.Time
	Status      *schedulerenum.TaskExecutionStatus
	TriggerType *schedulerenum.TaskTriggerType
}

type TaskExecutionRecordPageReq struct {
	Page *common.PageReq
	TaskExecutionRecordGetReq
}

type TaskExecutionRecordPageResp struct {
	Rows []*model.TaskExecutionRecord
	Page *common.PageResp
}

type TaskExecutionRecordExistsPeriodReq struct {
	TaskID      int64
	ScheduledAt time.Time
}

type TaskExecutionRecordCreateReq struct {
	Record *model.TaskExecutionRecord
	Status schedulerenum.TaskExecutionStatus
}

type TaskExecutionRecordCreateResp struct {
	Row      *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

type TaskExecutionRecordHasUnexpiredRunningReq struct {
	TaskID       int64
	StartedAfter time.Time
}

type TaskExecutionRecordMarkUnknownReq struct {
	IDs        []int64
	FinishedAt time.Time
	LastError  string
}

type TaskExecutionRecordMarkFinishedReq struct {
	ID         int64
	Status     schedulerenum.TaskExecutionStatus
	FinishedAt time.Time
	DurationMS int64
	LastError  string
}

type TaskExecutionRecordMarkFinishedResp struct {
	Row     *model.TaskExecutionRecord
	Updated bool
}
