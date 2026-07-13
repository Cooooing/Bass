package model

import (
	schedulerenum "scheduler/internal/enum"
	"time"
)

type TaskExecutionRecord struct {
	ID          int64
	TaskID      int64
	ScheduledAt time.Time
	StartedAt   *time.Time
	FinishedAt  *time.Time
	DurationMS  *int64
	Status      schedulerenum.TaskExecutionStatus
	TriggerType schedulerenum.TaskTriggerType
	TaskVersion int64
	WorkerID    string
	Payload     string
	LastError   string
	TraceID     string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
