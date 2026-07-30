package model

import (
	schedulerenum "scheduler/internal/enum"
	"time"
)

type DelayedTaskVersion struct {
	ID            int64
	DelayedTaskID int64
	Version       int64
	TaskKey       string
	HandlerName   schedulerenum.TaskHandlerName
	Title         string
	Description   string
	Enabled       bool
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
