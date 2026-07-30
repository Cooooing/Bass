package repo

import (
	"context"
	"scheduler/internal/biz/model"
	schedulerenum "scheduler/internal/enum"
	"time"
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
	ScheduledTaskID      int64
	ScheduledTaskKey     string
	ScheduledTaskVersion int64
	ScheduleKey          string
	ScheduledAt          time.Time
	TriggerType          schedulerenum.TaskTriggerType
	NumDelivered         uint64
	StreamSequence       uint64
	LatestForSubject     bool
}
