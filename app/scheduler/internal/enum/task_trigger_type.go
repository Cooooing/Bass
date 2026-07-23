package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

// TaskTriggerType 定义任务触发类型。
type TaskTriggerType string

const (
	TaskTriggerTypeSchedule TaskTriggerType = "schedule"
	TaskTriggerTypeManual   TaskTriggerType = "manual"
)

// TaskTriggerTypeMap 将内部任务触发类型映射到 proto 枚举。
var TaskTriggerTypeMap = commonenum.NewMapping[TaskTriggerType, schedulerv1.SchedulerTaskTriggerType](map[TaskTriggerType]commonenum.Entry[TaskTriggerType, schedulerv1.SchedulerTaskTriggerType]{
	TaskTriggerTypeSchedule: {Proto: schedulerv1.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_SCHEDULE},
	TaskTriggerTypeManual:   {Proto: schedulerv1.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_MANUAL},
})

func (TaskTriggerType) Values() []string {
	return TaskTriggerTypeMap.EnumValues()
}
