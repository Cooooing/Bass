package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

type DelayedTaskStatus string

const (
	DelayedTaskStatusPending   DelayedTaskStatus = "pending"
	DelayedTaskStatusRunning   DelayedTaskStatus = "running"
	DelayedTaskStatusSuccess   DelayedTaskStatus = "success"
	DelayedTaskStatusFailed    DelayedTaskStatus = "failed"
	DelayedTaskStatusCancelled DelayedTaskStatus = "cancelled"
)

var DelayedTaskStatusMap = commonenum.NewMapping[DelayedTaskStatus, schedulerv1.SchedulerDelayedTaskStatus](map[DelayedTaskStatus]commonenum.Entry[DelayedTaskStatus, schedulerv1.SchedulerDelayedTaskStatus]{
	DelayedTaskStatusPending:   {Proto: schedulerv1.SchedulerDelayedTaskStatus_SCHEDULER_DELAYED_TASK_STATUS_PENDING},
	DelayedTaskStatusRunning:   {Proto: schedulerv1.SchedulerDelayedTaskStatus_SCHEDULER_DELAYED_TASK_STATUS_RUNNING},
	DelayedTaskStatusSuccess:   {Proto: schedulerv1.SchedulerDelayedTaskStatus_SCHEDULER_DELAYED_TASK_STATUS_SUCCESS},
	DelayedTaskStatusFailed:    {Proto: schedulerv1.SchedulerDelayedTaskStatus_SCHEDULER_DELAYED_TASK_STATUS_FAILED},
	DelayedTaskStatusCancelled: {Proto: schedulerv1.SchedulerDelayedTaskStatus_SCHEDULER_DELAYED_TASK_STATUS_CANCELLED},
})

func (DelayedTaskStatus) Values() []string {
	return DelayedTaskStatusMap.EnumValues()
}

func (s DelayedTaskStatus) String() string {
	return string(s)
}
