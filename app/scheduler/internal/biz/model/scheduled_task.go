package model

import (
	schedulerenum "scheduler/internal/enum"
	"time"
)

type ScheduledTask struct {
	ID            int64
	TaskKey       string
	HandlerName   schedulerenum.TaskHandlerName
	Title         string
	Description   string
	Enabled       bool
	CronSpec      string
	Payload       string
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
	AllowOverlap  bool
	Version       int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
