package enum

import event "common/proto/gen/common/enums/event"

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
var InboxEventStatusMap = NewMapping[InboxEventStatus, event.InboxEventStatus](map[InboxEventStatus]Entry[InboxEventStatus, event.InboxEventStatus]{
	InboxEventStatusReceived:   {Proto: event.InboxEventStatus_INBOX_EVENT_STATUS_RECEIVED},
	InboxEventStatusProcessing: {Proto: event.InboxEventStatus_INBOX_EVENT_STATUS_PROCESSING},
	InboxEventStatusProcessed:  {Proto: event.InboxEventStatus_INBOX_EVENT_STATUS_PROCESSED},
	InboxEventStatusFailed:     {Proto: event.InboxEventStatus_INBOX_EVENT_STATUS_FAILED},
	InboxEventStatusDead:       {Proto: event.InboxEventStatus_INBOX_EVENT_STATUS_DEAD},
})

func (e InboxEventStatus) String() string {
	return string(e)
}
