package enum

import event "common/proto/gen/common/enums/event"

// OutboxEventStatus 定义生产者侧本地 outbox 记录的投递状态。
type OutboxEventStatus string

const (
	// OutboxEventStatusPending 表示事件已持久化，等待投递。
	OutboxEventStatusPending OutboxEventStatus = "pending"
	// OutboxEventStatusPublishing 表示分发器已锁定事件并正在投递。
	OutboxEventStatusPublishing OutboxEventStatus = "publishing"
	// OutboxEventStatusPublished 表示事件已被消息系统确认。
	OutboxEventStatusPublished OutboxEventStatus = "published"
	// OutboxEventStatusFailed 表示最近一次投递失败，可以重试。
	OutboxEventStatusFailed OutboxEventStatus = "failed"
	// OutboxEventStatusDead 表示重试已耗尽，需要人工处理。
	OutboxEventStatusDead OutboxEventStatus = "dead"
)

// OutboxEventStatusMap 将持久化的 outbox 状态映射到 proto 值。
var OutboxEventStatusMap = NewMapping[OutboxEventStatus, event.OutboxEventStatus](map[OutboxEventStatus]Entry[OutboxEventStatus, event.OutboxEventStatus]{
	OutboxEventStatusPending:    {Proto: event.OutboxEventStatus_OUTBOX_EVENT_STATUS_PENDING},
	OutboxEventStatusPublishing: {Proto: event.OutboxEventStatus_OUTBOX_EVENT_STATUS_PUBLISHING},
	OutboxEventStatusPublished:  {Proto: event.OutboxEventStatus_OUTBOX_EVENT_STATUS_PUBLISHED},
	OutboxEventStatusFailed:     {Proto: event.OutboxEventStatus_OUTBOX_EVENT_STATUS_FAILED},
	OutboxEventStatusDead:       {Proto: event.OutboxEventStatus_OUTBOX_EVENT_STATUS_DEAD},
})

func (e OutboxEventStatus) String() string {
	return string(e)
}
