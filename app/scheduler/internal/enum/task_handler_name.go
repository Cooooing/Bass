package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

// TaskHandlerName 定义代码内任务处理器名称。
type TaskHandlerName string

const (
	TaskHandlerNameNoop                      TaskHandlerName = "noop"
	TaskHandlerNameUserOutboxPublishBatch    TaskHandlerName = "user.outbox_publish_batch"
	TaskHandlerNameUserUnbanAccounts         TaskHandlerName = "user.unban_accounts"
	TaskHandlerNameContentOutboxPublishBatch TaskHandlerName = "content.outbox_publish_batch"
)

// TaskHandlerNameMap 将内部任务处理器名称映射到 proto 枚举。
var TaskHandlerNameMap = commonenum.NewMapping[TaskHandlerName, schedulerv1.SchedulerTaskHandlerName](
	map[TaskHandlerName]commonenum.Entry[TaskHandlerName, schedulerv1.SchedulerTaskHandlerName]{
		TaskHandlerNameNoop: {
			Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_NOOP,
		},
		TaskHandlerNameUserOutboxPublishBatch: {
			Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_USER_OUTBOX_PUBLISH_BATCH,
		},
		TaskHandlerNameUserUnbanAccounts: {
			Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_USER_UNBAN_ACCOUNTS,
		},
		TaskHandlerNameContentOutboxPublishBatch: {
			Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_OUTBOX_PUBLISH_BATCH,
		},
	},
)

func (TaskHandlerName) Values() []string {
	return TaskHandlerNameMap.EnumValues()
}

func (e TaskHandlerName) String() string {
	return string(e)
}
