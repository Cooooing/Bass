package enum

import event "common/proto/gen/common/enums/event/v1"

// EventConsumer identifies a JetStream durable consumer.
type EventConsumer string

const (
	EventConsumerNotify                 EventConsumer = "notify"
	EventConsumerSchedulerScheduledTask EventConsumer = "scheduler-scheduled-task-executor"
	EventConsumerSchedulerDelayedTask   EventConsumer = "scheduler-delayed-task-executor"
	EventConsumerPlatformAssetUpload    EventConsumer = "platform-asset-upload-recorder"
)

// EventConsumerMap 将内部消费组映射到 proto 枚举。
var EventConsumerMap = NewMapping[EventConsumer, event.EventConsumer](map[EventConsumer]Entry[EventConsumer, event.EventConsumer]{
	EventConsumerNotify:                 {Proto: event.EventConsumer_EVENT_CONSUMER_NOTIFY},
	EventConsumerSchedulerScheduledTask: {Proto: event.EventConsumer_EVENT_CONSUMER_SCHEDULER_SCHEDULED_TASK},
	EventConsumerSchedulerDelayedTask:   {Proto: event.EventConsumer_EVENT_CONSUMER_SCHEDULER_DELAYED_TASK},
	EventConsumerPlatformAssetUpload:    {Proto: event.EventConsumer_EVENT_CONSUMER_PLATFORM_ASSET_UPLOAD},
})

func (e EventConsumer) String() string {
	return string(e)
}
