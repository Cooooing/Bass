package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

// TaskExecutionStatus 定义任务执行状态。
type TaskExecutionStatus string

const (
	TaskExecutionStatusPending      TaskExecutionStatus = "pending"
	TaskExecutionStatusRunning      TaskExecutionStatus = "running"
	TaskExecutionStatusRetryPending TaskExecutionStatus = "retry_pending"
	TaskExecutionStatusSuccess      TaskExecutionStatus = "success"
	TaskExecutionStatusFailed       TaskExecutionStatus = "failed"
	TaskExecutionStatusTimeout      TaskExecutionStatus = "timeout"
	TaskExecutionStatusCanceled     TaskExecutionStatus = "canceled"
	TaskExecutionStatusSkipped      TaskExecutionStatus = "skipped"
)

// TaskExecutionStatusMap 将内部任务执行状态映射到 proto 枚举。
var TaskExecutionStatusMap = commonenum.NewMapping[TaskExecutionStatus, schedulerv1.SchedulerTaskExecutionStatus](
	map[TaskExecutionStatus]commonenum.Entry[TaskExecutionStatus, schedulerv1.SchedulerTaskExecutionStatus]{
		TaskExecutionStatusPending:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_PENDING},
		TaskExecutionStatusRunning:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_RUNNING},
		TaskExecutionStatusRetryPending: {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_RETRY_PENDING},
		TaskExecutionStatusSuccess:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_SUCCESS},
		TaskExecutionStatusFailed:       {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_FAILED},
		TaskExecutionStatusTimeout:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_TIMEOUT},
		TaskExecutionStatusCanceled:     {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_CANCELED},
		TaskExecutionStatusSkipped:      {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_SKIPPED},
	},
)

func (TaskExecutionStatus) Values() []string {
	return TaskExecutionStatusMap.EnumValues()
}

func (e TaskExecutionStatus) String() string {
	return string(e)
}
