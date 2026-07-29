package model

import (
	schedulerenum "scheduler/internal/enum"
	"time"
)

type ScheduledTaskExecutionRecord struct {
	ID                   int64
	ScheduledTaskID      int64
	ScheduledTaskVersion int64
	TriggerType          schedulerenum.TaskTriggerType
	ScheduleKey          string
	ScheduledAt          time.Time
	StartedAt            *time.Time
	FinishedAt           *time.Time
	Duration             *time.Duration
	Status               schedulerenum.TaskExecutionStatus
	Attempt              int32
	MaxAttempts          int32
	Timeout              time.Duration
	StaleAfter           *time.Duration
	MisfirePolicy        schedulerenum.TaskMisfirePolicy
	WorkerID             string
	Payload              string
	LastError            string
	TraceID              string
	CreatedAt            *time.Time
	UpdatedAt            *time.Time
}
