package timewheel

import (
	"time"
)

// Task 表示时间轮上的一个待触发任务。
type Task struct {
	ID      string
	DueAt   time.Time
	Payload any

	slot   int64
	rounds int64
}
