package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1"
)

type TaskExecutionRuntimeState string

const (
	TaskExecutionRuntimeStateActive   TaskExecutionRuntimeState = "active"
	TaskExecutionRuntimeStateStale    TaskExecutionRuntimeState = "stale"
	TaskExecutionRuntimeStateUnknown  TaskExecutionRuntimeState = "unknown"
	TaskExecutionRuntimeStateTerminal TaskExecutionRuntimeState = "terminal"
)

var TaskExecutionRuntimeStateMap = commonenum.NewMapping[TaskExecutionRuntimeState, schedulerv1.SchedulerTaskExecutionRuntimeState](map[TaskExecutionRuntimeState]commonenum.Entry[TaskExecutionRuntimeState, schedulerv1.SchedulerTaskExecutionRuntimeState]{
	TaskExecutionRuntimeStateActive:   {Proto: schedulerv1.SchedulerTaskExecutionRuntimeState_SCHEDULER_TASK_EXECUTION_RUNTIME_STATE_ACTIVE},
	TaskExecutionRuntimeStateStale:    {Proto: schedulerv1.SchedulerTaskExecutionRuntimeState_SCHEDULER_TASK_EXECUTION_RUNTIME_STATE_STALE},
	TaskExecutionRuntimeStateUnknown:  {Proto: schedulerv1.SchedulerTaskExecutionRuntimeState_SCHEDULER_TASK_EXECUTION_RUNTIME_STATE_UNKNOWN},
	TaskExecutionRuntimeStateTerminal: {Proto: schedulerv1.SchedulerTaskExecutionRuntimeState_SCHEDULER_TASK_EXECUTION_RUNTIME_STATE_TERMINAL},
})

func (TaskExecutionRuntimeState) Values() []string {
	return TaskExecutionRuntimeStateMap.EnumValues()
}

func (e TaskExecutionRuntimeState) String() string {
	return string(e)
}
