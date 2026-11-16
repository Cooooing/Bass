package repo

import (
	"context"
	"time"

	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
)

type ScheduledTaskScheduleRepo interface {
	Ensure(ctx context.Context) error
	Schedule(ctx context.Context, req *ScheduledTaskScheduleReq) error
	Cancel(ctx context.Context, subject string) error
	Consume(ctx context.Context, handler func(context.Context, *ScheduledTaskScheduleMessage) (*MessageHandleResult, error)) error
	Stop(ctx context.Context) error
}

type ScheduledTaskScheduleReq struct {
	ScheduledTask *model.ScheduledTask
	Subject       string
	Target        string
}
type ScheduledTaskScheduleMessage struct {
	ScheduledTaskID      int64                         `json:"scheduled_task_id"`
	ScheduledTaskVersion int64                         `json:"scheduled_task_version"`
	TriggerType          schedulerenum.TaskTriggerType `json:"trigger_type"`
	ScheduleKey          string                        `json:"-"`
	ScheduledAt          time.Time                     `json:"-"`
	NumDelivered         uint64                        `json:"-"`
	StreamSequence       uint64                        `json:"-"`
	LatestForSubject     bool                          `json:"-"`
}
