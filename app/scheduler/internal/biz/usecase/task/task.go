package task

import (
	"context"
	schedulerenum "scheduler/internal/enum"
	"time"
)

type Task interface {
	HandlerName() schedulerenum.TaskHandlerName
	Title() string
	Description() string
	DefaultScheduledTasks() []*DefaultScheduledTask
	DefaultDelayedTasks() []*DefaultDelayedTask
	Execute(ctx context.Context, payload string) error
}

type DefaultScheduledTask struct {
	TaskKey       schedulerenum.TaskKey
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
}

type DefaultDelayedTask struct {
	TaskKey       schedulerenum.TaskKey
	Title         string
	Description   string
	Enabled       bool
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
}
