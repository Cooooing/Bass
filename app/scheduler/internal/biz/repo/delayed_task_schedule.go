package repo

import (
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type DelayedTaskScheduleRepo interface {
	Ensure(ctx context.Context) error
	Schedule(ctx context.Context, req *DelayedTaskScheduleReq) error
	Cancel(ctx context.Context, subject string) error
	Consume(ctx context.Context, handler func(context.Context, *DelayedTaskScheduleMessage) (*MessageHandleResult, error)) error
	Stop(ctx context.Context) error
}

type DelayedTaskScheduleReq struct {
	DelayedTask *model.DelayedTask
	Record      *model.DelayedTaskExecutionRecord
	Subject     string
	Target      string
}

type DelayedTaskScheduleMessage struct {
	ExecutionRecordID int64
	DelayedTaskKey    string
	ScheduleKey       string
	ScheduledAt       time.Time
	TriggerType       schedulerenum.TaskTriggerType
	NumDelivered      uint64
}
