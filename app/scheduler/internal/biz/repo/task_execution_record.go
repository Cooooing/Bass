package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type TaskExecutionRecordRepo interface {
	Get(ctx context.Context, req *TaskExecutionRecordGetReq) (*TaskExecutionRecordGetResponse, error)
	List(ctx context.Context, req *TaskExecutionRecordGetReq) (*TaskExecutionRecordListResponse, error)
	Map(ctx context.Context, req *TaskExecutionRecordGetReq) (*TaskExecutionRecordMapResponse, error)
	Count(ctx context.Context, req *TaskExecutionRecordGetReq) (*TaskExecutionRecordCountResponse, error)
	Page(ctx context.Context, req *TaskExecutionRecordPageReq) (*TaskExecutionRecordPageResponse, error)

	ExistsPeriod(ctx context.Context, req *TaskExecutionRecordExistsPeriodReq) (*TaskExecutionRecordExistsPeriodResponse, error)
	Create(ctx context.Context, req *TaskExecutionRecordCreateReq) (*TaskExecutionRecordCreateResponse, error)
	HasUnexpiredRunning(ctx context.Context, req *TaskExecutionRecordHasUnexpiredRunningReq) (*TaskExecutionRecordHasUnexpiredRunningResponse, error)
	MarkUnknown(ctx context.Context, req *TaskExecutionRecordMarkUnknownReq) (*TaskExecutionRecordMarkUnknownResponse, error)
	MarkFinished(ctx context.Context, req *TaskExecutionRecordMarkFinishedReq) (*TaskExecutionRecordMarkFinishedResponse, error)
}

type TaskExecutionRecordGetReq struct {
	ID          *int64
	IDs         []int64
	TaskID      *int64
	ScheduledAt *time.Time
	Status      *schedulerenum.TaskExecutionStatus
	TriggerType *schedulerenum.TaskTriggerType
}

type TaskExecutionRecordGetResponse struct {
	Row *model.TaskExecutionRecord
}

type TaskExecutionRecordListResponse struct {
	Rows []*model.TaskExecutionRecord
}

type TaskExecutionRecordMapResponse struct {
	Rows map[int64]*model.TaskExecutionRecord
}

type TaskExecutionRecordCountResponse struct {
	Count int
}

type TaskExecutionRecordPageReq struct {
	Page *common.PageRequest
	TaskExecutionRecordGetReq
}

type TaskExecutionRecordPageResponse struct {
	Rows []*model.TaskExecutionRecord
	Page *common.PageResponse
}

type TaskExecutionRecordExistsPeriodReq struct {
	TaskID      int64
	ScheduledAt time.Time
}

type TaskExecutionRecordExistsPeriodResponse struct {
	Exists bool
}

type TaskExecutionRecordCreateReq struct {
	Record *model.TaskExecutionRecord
	Status schedulerenum.TaskExecutionStatus
}

type TaskExecutionRecordCreateResponse struct {
	Row      *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

type TaskExecutionRecordHasUnexpiredRunningReq struct {
	TaskID       int64
	StartedAfter time.Time
}

type TaskExecutionRecordHasUnexpiredRunningResponse struct {
	Exists bool
}

type TaskExecutionRecordMarkUnknownReq struct {
	IDs        []int64
	FinishedAt time.Time
	LastError  string
}

type TaskExecutionRecordMarkUnknownResponse struct {
	Rows []*model.TaskExecutionRecord
}

type TaskExecutionRecordMarkFinishedReq struct {
	ID         int64
	Status     schedulerenum.TaskExecutionStatus
	FinishedAt time.Time
	DurationMS int64
	LastError  string
}

type TaskExecutionRecordMarkFinishedResponse struct {
	Row     *model.TaskExecutionRecord
	Updated bool
}
