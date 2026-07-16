package repo

import (
	"context"
	schedulerenum "scheduler/internal/enum"
	"time"
)

// TaskLockRepo 封装 scheduler 的 Redis 协调能力；数据库执行记录仍是唯一事实源。
type TaskLockRepo interface {
	TryAcquireSchedule(ctx context.Context, req *TaskScheduleAcquireReq) (*TaskScheduleAcquireResponse, error)
	RegisterRunning(ctx context.Context, req *TaskRunningLockReq) (*TaskRunningLockResponse, error)
	RefreshRunning(ctx context.Context, req *TaskRunningLockReq) (*TaskRunningLockResponse, error)
	ReleaseRunning(ctx context.Context, req *TaskRunningLockReq) (*TaskRunningReleaseResponse, error)
	MapRunning(ctx context.Context, req *TaskRunningMapReq) (*TaskRunningMapResponse, error)
}

// TaskScheduleAcquireReq 描述一次本地 cron 或手动触发进入 Redis 协调层所需的参数。
type TaskScheduleAcquireReq struct {
	TaskID            int64
	ScheduledAt       time.Time
	AllowOverlap      bool
	SchedulePeriodTTL time.Duration
	RunningLockTTL    time.Duration
}

// TaskScheduleAcquireResponse 返回调度准入结果；RunningToken 只在 run 且需要互斥运行时有值。
type TaskScheduleAcquireResponse struct {
	Decision     schedulerenum.TaskScheduleDecision
	RunningToken string
}

type TaskRunningLockReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
	TTL               time.Duration
}

type TaskRunningLockResponse struct {
	OK bool
}

type TaskRunningReleaseResponse struct{}

type TaskRunningMapReq struct {
	TaskID             int64
	ExecutionRecordIDs []int64
}

type TaskRunningMapResponse struct {
	Rows map[int64]bool
}
