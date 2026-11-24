package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/enum/v1"
)

// TaskExecutionStatus 定义任务执行状态。
type TaskExecutionStatus string

const (
	TaskExecutionStatusPending      TaskExecutionStatus = "pending"       // 等待开始执行
	TaskExecutionStatusRunning      TaskExecutionStatus = "running"       // 正在执行任务
	TaskExecutionStatusSuccess      TaskExecutionStatus = "success"       // 任务执行成功
	TaskExecutionStatusFailed       TaskExecutionStatus = "failed"        // 任务执行失败
	TaskExecutionStatusTimeout      TaskExecutionStatus = "timeout"       // 任务执行超时
	TaskExecutionStatusCanceled     TaskExecutionStatus = "canceled"      // 任务被取消
	TaskExecutionStatusSkipped      TaskExecutionStatus = "skipped"       // 任务被跳过
	TaskExecutionStatusRetryPending TaskExecutionStatus = "retry_pending" // 等待下一次重试
)

// TaskExecutionStatusMap 将内部任务执行状态映射到 proto 枚举。
var TaskExecutionStatusMap = commonenum.NewMapping[TaskExecutionStatus, schedulerv1.SchedulerTaskExecutionStatus](
	map[TaskExecutionStatus]commonenum.Entry[TaskExecutionStatus, schedulerv1.SchedulerTaskExecutionStatus]{
		TaskExecutionStatusPending:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_PENDING},
		TaskExecutionStatusRunning:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_RUNNING},
		TaskExecutionStatusSuccess:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_SUCCESS},
		TaskExecutionStatusFailed:       {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_FAILED},
		TaskExecutionStatusTimeout:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_TIMEOUT},
		TaskExecutionStatusCanceled:     {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_CANCELED},
		TaskExecutionStatusSkipped:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_SKIPPED},
		TaskExecutionStatusRetryPending: {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_RETRY_PENDING},
	},
)

func (TaskExecutionStatus) Values() []string {
	return TaskExecutionStatusMap.EnumValues()
}

func (e TaskExecutionStatus) String() string {
	return string(e)
}
