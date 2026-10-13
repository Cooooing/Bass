package model

import (
	"math"
	"time"

	schedulerenum "scheduler/internal/enum"
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

func (r *ScheduledTaskExecutionRecord) TimeoutSeconds() int32 {
	seconds := r.Timeout / time.Second
	if seconds > time.Duration(math.MaxInt32) {
		return math.MaxInt32
	}
	return int32(seconds) // #nosec G115 -- bounded by MaxInt32 above.
}
func (r *ScheduledTaskExecutionRecord) StaleAfterSeconds() *int32 {
	if r.StaleAfter == nil {
		return nil
	}
	seconds := *r.StaleAfter / time.Second
	if seconds > time.Duration(math.MaxInt32) {
		value := int32(math.MaxInt32)
		return &value
	}
	value := int32(seconds) // #nosec G115 -- bounded by MaxInt32 above.
	return &value
}
