package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/v1/enum"
)

// TaskKey 定义代码内默认任务配置键。
type TaskKey string

const (
	TaskKeyNoopScheduledDefault                    TaskKey = "noop.scheduled.default"
	TaskKeyNoopDelayedDefault                      TaskKey = "noop.delayed.default"
	TaskKeyUserOutboxPublishBatchDefault           TaskKey = "user.outbox_publish_batch.default"
	TaskKeyUserOutboxPublishBatchDelayedDefault    TaskKey = "user.outbox_publish_batch.delayed.default"
	TaskKeyUserUnbanAccountsDefault                TaskKey = "user.unban_accounts.default"
	TaskKeyContentOutboxPublishBatchDefault        TaskKey = "content.outbox_publish_batch.default"
	TaskKeyContentOutboxPublishBatchDelayedDefault TaskKey = "content.outbox_publish_batch.delayed.default"
)

// TaskKeyMap 将内部默认任务配置键映射到 proto 枚举。
var TaskKeyMap = commonenum.NewMapping[TaskKey, schedulerv1.SchedulerTaskKey](
	map[TaskKey]commonenum.Entry[TaskKey, schedulerv1.SchedulerTaskKey]{
		TaskKeyNoopScheduledDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_NOOP_SCHEDULED_DEFAULT,
		},
		TaskKeyNoopDelayedDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_NOOP_DELAYED_DEFAULT,
		},
		TaskKeyUserOutboxPublishBatchDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_USER_OUTBOX_PUBLISH_BATCH_DEFAULT,
		},
		TaskKeyUserOutboxPublishBatchDelayedDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_USER_OUTBOX_PUBLISH_BATCH_DELAYED_DEFAULT,
		},
		TaskKeyUserUnbanAccountsDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_USER_UNBAN_ACCOUNTS_DEFAULT,
		},
		TaskKeyContentOutboxPublishBatchDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_CONTENT_OUTBOX_PUBLISH_BATCH_DEFAULT,
		},
		TaskKeyContentOutboxPublishBatchDelayedDefault: {
			Proto: schedulerv1.SchedulerTaskKey_SCHEDULER_TASK_KEY_CONTENT_OUTBOX_PUBLISH_BATCH_DELAYED_DEFAULT,
		},
	},
)

func (TaskKey) Values() []string {
	return TaskKeyMap.EnumValues()
}

func (e TaskKey) String() string {
	return string(e)
}
