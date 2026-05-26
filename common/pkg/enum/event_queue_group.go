package enum

import "common/api/gen/common/enums"

// EventQueueGroup 定义跨服务事件消费的 NATS queue group。
type EventQueueGroup string

const (
	EventQueueGroupNotify EventQueueGroup = "notify"
)

// EventQueueGroupMap 将内部消费队列组映射到 proto 枚举。
var EventQueueGroupMap = NewMapping[EventQueueGroup, enums.EventQueueGroup](map[EventQueueGroup]Entry[EventQueueGroup, enums.EventQueueGroup]{
	EventQueueGroupNotify: {Proto: enums.EventQueueGroup_EVENT_QUEUE_GROUP_NOTIFY},
})
