package timewheel

import (
	"context"
	"time"
)

// Job 是时间轮任务到期后的执行逻辑。
type Job func(context.Context, *Task) error

// Task 表示时间轮上的一个待触发任务。
type Task struct {
	ID      string
	DueAt   time.Time
	Payload any
	Job     Job

	slot   int64
	rounds int64
}
