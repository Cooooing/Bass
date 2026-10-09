package model

import (
	"time"

	schedulerenum "scheduler/internal/enum"
)

type DelayedTask struct {
	ID            int64
	TaskKey       string
	HandlerName   schedulerenum.TaskHandlerName
	Title         string
	Description   string
	Enabled       bool
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
	Version       int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
