package enum

import "common/proto/gen/common/enums"

// OutboxEventStatus 定义生产者侧本地 outbox 记录的投递状态。
// 分发器后续会按这些状态推进记录。
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
var OutboxEventStatusMap = NewMapping[OutboxEventStatus, enums.OutboxEventStatus](map[OutboxEventStatus]Entry[OutboxEventStatus, enums.OutboxEventStatus]{
	OutboxEventStatusPending:    {Proto: enums.OutboxEventStatus_OUTBOX_EVENT_STATUS_PENDING},
	OutboxEventStatusPublishing: {Proto: enums.OutboxEventStatus_OUTBOX_EVENT_STATUS_PUBLISHING},
	OutboxEventStatusPublished:  {Proto: enums.OutboxEventStatus_OUTBOX_EVENT_STATUS_PUBLISHED},
	OutboxEventStatusFailed:     {Proto: enums.OutboxEventStatus_OUTBOX_EVENT_STATUS_FAILED},
	OutboxEventStatusDead:       {Proto: enums.OutboxEventStatus_OUTBOX_EVENT_STATUS_DEAD},
})

// InboxEventStatus 定义消费者侧本地 inbox 记录的处理状态。
type InboxEventStatus string

const (
	// InboxEventStatusReceived 表示消息已记录但尚未处理。
	InboxEventStatusReceived InboxEventStatus = "received"
	// InboxEventStatusProcessing 表示消费者已锁定消息并正在处理。
	InboxEventStatusProcessing InboxEventStatus = "processing"
	// InboxEventStatusProcessed 表示消息已处理成功。
	InboxEventStatusProcessed InboxEventStatus = "processed"
	// InboxEventStatusFailed 表示最近一次处理失败，可以重试。
	InboxEventStatusFailed InboxEventStatus = "failed"
	// InboxEventStatusDead 表示重试已耗尽，需要人工处理。
	InboxEventStatusDead InboxEventStatus = "dead"
)

// InboxEventStatusMap 将持久化的 inbox 状态映射到 proto 值。
var InboxEventStatusMap = NewMapping[InboxEventStatus, enums.InboxEventStatus](map[InboxEventStatus]Entry[InboxEventStatus, enums.InboxEventStatus]{
	InboxEventStatusReceived:   {Proto: enums.InboxEventStatus_INBOX_EVENT_STATUS_RECEIVED},
	InboxEventStatusProcessing: {Proto: enums.InboxEventStatus_INBOX_EVENT_STATUS_PROCESSING},
	InboxEventStatusProcessed:  {Proto: enums.InboxEventStatus_INBOX_EVENT_STATUS_PROCESSED},
	InboxEventStatusFailed:     {Proto: enums.InboxEventStatus_INBOX_EVENT_STATUS_FAILED},
	InboxEventStatusDead:       {Proto: enums.InboxEventStatus_INBOX_EVENT_STATUS_DEAD},
})
