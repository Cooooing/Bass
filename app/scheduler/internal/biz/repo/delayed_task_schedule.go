package repo

import (
	"context"
	"time"

	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
)

type DelayedTaskScheduleRepo interface {
	Ensure(ctx context.Context) error
	Schedule(ctx context.Context, req *DelayedTaskScheduleReq) error
	Publish(ctx context.Context, req *DelayedTaskScheduleReq) error
	Cancel(ctx context.Context, subject string) error
	Consume(ctx context.Context, handler func(context.Context, *DelayedTaskScheduleMessage) (*MessageHandleResult, error)) error
	Stop(ctx context.Context) error
}

type DelayedTaskScheduleReq struct {
	Record  *model.DelayedTaskExecutionRecord
	Subject string
	Target  string
}

type DelayedTaskScheduleMessage struct {
	ExecutionRecordID int64                         `json:"execution_record_id"`
	TriggerType       schedulerenum.TaskTriggerType `json:"trigger_type"`
	ScheduleKey       string                        `json:"-"`
	ScheduledAt       time.Time                     `json:"-"`
	NumDelivered      uint64                        `json:"-"`
}
