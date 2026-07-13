package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1"
)

// TaskExecutionStatus 定义任务执行状态。
type TaskExecutionStatus string

const (
	TaskExecutionStatusRunning        TaskExecutionStatus = "running"
	TaskExecutionStatusSuccess        TaskExecutionStatus = "success"
	TaskExecutionStatusFailed         TaskExecutionStatus = "failed"
	TaskExecutionStatusTimeout        TaskExecutionStatus = "timeout"
	TaskExecutionStatusCanceled       TaskExecutionStatus = "canceled"
	TaskExecutionStatusOverlapSkipped TaskExecutionStatus = "overlap_skipped"
	TaskExecutionStatusUnknown        TaskExecutionStatus = "unknown"
)

// TaskExecutionStatusMap 将内部任务执行状态映射到 proto 枚举。
var TaskExecutionStatusMap = commonenum.NewMapping[TaskExecutionStatus, schedulerv1.SchedulerTaskExecutionStatus](map[TaskExecutionStatus]commonenum.Entry[TaskExecutionStatus, schedulerv1.SchedulerTaskExecutionStatus]{
	TaskExecutionStatusRunning:        {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_RUNNING},
	TaskExecutionStatusSuccess:        {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_SUCCESS},
	TaskExecutionStatusFailed:         {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_FAILED},
	TaskExecutionStatusTimeout:        {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_TIMEOUT},
	TaskExecutionStatusCanceled:       {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_CANCELED},
	TaskExecutionStatusOverlapSkipped: {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_OVERLAP_SKIPPED},
	TaskExecutionStatusUnknown:        {Proto: schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNKNOWN},
})

func (TaskExecutionStatus) Values() []string {
	return TaskExecutionStatusMap.EnumValues()
}

func (e TaskExecutionStatus) String() string {
	return string(e)
}
