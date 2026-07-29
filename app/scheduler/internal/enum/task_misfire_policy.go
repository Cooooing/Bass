package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

// TaskMisfirePolicy 定义错过计划时间后的处理策略。
type TaskMisfirePolicy string

const (
	TaskMisfirePolicySkip          TaskMisfirePolicy = "skip"
	TaskMisfirePolicyExecuteLatest TaskMisfirePolicy = "execute_latest"
	TaskMisfirePolicyExecuteAll    TaskMisfirePolicy = "execute_all"
)

// TaskMisfirePolicyMap 将内部错过调度策略映射到 proto 枚举。
var TaskMisfirePolicyMap = commonenum.NewMapping[TaskMisfirePolicy, schedulerv1.SchedulerTaskMisfirePolicy](
	map[TaskMisfirePolicy]commonenum.Entry[TaskMisfirePolicy, schedulerv1.SchedulerTaskMisfirePolicy]{
		TaskMisfirePolicySkip:          {Proto: schedulerv1.SchedulerTaskMisfirePolicy_SCHEDULER_TASK_MISFIRE_POLICY_SKIP},
		TaskMisfirePolicyExecuteLatest: {Proto: schedulerv1.SchedulerTaskMisfirePolicy_SCHEDULER_TASK_MISFIRE_POLICY_EXECUTE_LATEST},
		TaskMisfirePolicyExecuteAll:    {Proto: schedulerv1.SchedulerTaskMisfirePolicy_SCHEDULER_TASK_MISFIRE_POLICY_EXECUTE_ALL},
	},
)

func (TaskMisfirePolicy) Values() []string {
	return TaskMisfirePolicyMap.EnumValues()
}

func (e TaskMisfirePolicy) String() string {
	return string(e)
}
