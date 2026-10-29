package task

import (
	"context"
	"time"

	schedulerenum "scheduler/internal/enum"
)

// Execution 是 Scheduler 传给任务处理器的一次显式执行元数据。
type Execution struct {
	Payload     string
	ScheduledAt time.Time
}

type Task interface {
	HandlerName() schedulerenum.TaskHandlerName
	Title() string
	Description() string
	DefaultScheduledTasks() []*DefaultScheduledTask
	DefaultDelayedTasks() []*DefaultDelayedTask
	Execute(ctx context.Context, execution Execution) error
}

type DefaultScheduledTask struct {
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
	Title         string
	Description   string
	Enabled       bool
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
}
