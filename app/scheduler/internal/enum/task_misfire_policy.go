package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/enum/v1"
)

// TaskMisfirePolicy 定义错过计划时间后的处理策略。
type TaskMisfirePolicy string

const (
	TaskMisfirePolicySkip          TaskMisfirePolicy = "skip"           // 跳过错过的执行
	TaskMisfirePolicyExecuteLatest TaskMisfirePolicy = "execute_latest" // 仅补执行最近一次
	TaskMisfirePolicyExecuteAll    TaskMisfirePolicy = "execute_all"    // 补执行所有错过任务
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
