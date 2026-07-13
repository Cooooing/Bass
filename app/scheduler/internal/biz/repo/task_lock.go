package repo

import (
	"context"
	schedulerenum "scheduler/internal/enum"
	"time"
)

// TaskLockRepo 封装 scheduler 的 Redis 协调能力；数据库执行记录仍是唯一事实源。
type TaskLockRepo interface {
	TryAcquireSchedule(ctx context.Context, req *TaskScheduleAcquireReq) (*TaskScheduleAcquireResult, error)
	RegisterRunning(ctx context.Context, taskID int64, executionRecordID int64, runningToken string, exclusive bool, ttl time.Duration) (bool, error)
	RefreshRunning(ctx context.Context, taskID int64, executionRecordID int64, runningToken string, exclusive bool, ttl time.Duration) (bool, error)
	ReleaseRunning(ctx context.Context, taskID int64, executionRecordID int64, runningToken string, exclusive bool) error
	MapRunning(ctx context.Context, taskID int64, executionRecordIDs []int64) (map[int64]bool, error)
}

// TaskScheduleAcquireReq 描述一次本地 cron 或手动触发进入 Redis 协调层所需的参数。
type TaskScheduleAcquireReq struct {
	TaskID            int64
	ScheduledAt       time.Time
	AllowOverlap      bool
	SchedulePeriodTTL time.Duration
	RunningLockTTL    time.Duration
}

// TaskScheduleAcquireResult 返回调度准入结果；RunningToken 只在 run 且需要互斥运行时有值。
type TaskScheduleAcquireResult struct {
	Decision     schedulerenum.TaskScheduleDecision
	RunningToken string
}
