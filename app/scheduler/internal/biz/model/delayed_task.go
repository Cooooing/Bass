package model

import (
	schedulerenum "scheduler/internal/enum"
	"time"
)

type DelayedTask struct {
	ID             int64
	IdempotencyKey string
	TaskName       string
	Payload        string
	ExecuteAt      time.Time
	NextRunAt      time.Time
	Status         schedulerenum.DelayedTaskStatus
	Attempt        int32
	MaxAttempts    int32
	TimeoutSeconds int32
	LockedBy       string
	LockExpiresAt  *time.Time
	StartedAt      *time.Time
	FinishedAt     *time.Time
	LastError      string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
