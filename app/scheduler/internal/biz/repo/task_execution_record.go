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
	Page(ctx context.Context, page *common.PageRequest, req *TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, *common.PageReply, error)

	ExistsPeriod(ctx context.Context, taskID int64, scheduledAt time.Time) (bool, error)
	Create(ctx context.Context, record *model.TaskExecutionRecord, status schedulerenum.TaskExecutionStatus) (*model.TaskExecutionRecord, bool, bool, error)
	HasUnexpiredRunning(ctx context.Context, taskID int64, startedAfter time.Time) (bool, error)
	MarkUnknown(ctx context.Context, ids []int64, finishedAt time.Time, lastError string) ([]*model.TaskExecutionRecord, error)
	MarkFinished(ctx context.Context, id int64, status schedulerenum.TaskExecutionStatus, finishedAt time.Time, durationMS int64, lastError string) (*model.TaskExecutionRecord, bool, error)
}

type TaskExecutionRecordGetReq struct {
	ID          *int64
	IDs         []int64
	TaskID      *int64
	ScheduledAt *time.Time
	Status      *schedulerenum.TaskExecutionStatus
	TriggerType *schedulerenum.TaskTriggerType
}
