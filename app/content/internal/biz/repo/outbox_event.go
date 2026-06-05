package repo

import (
	commonenums "common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"time"
)

// OutboxEventSave 定义一条待持久化的 outbox 事件。
type OutboxEventSave struct {
	Event *commonenums.Event
}

type OutboxEvent struct {
	ID      int64
	EventID string
	Subject commonenum.EventSubject
	Payload []byte
	Headers map[string]string
}

type OutboxEventRepo interface {
	Save(ctx context.Context, req *OutboxEventSave) error
	ClaimForPublish(ctx context.Context, limit int, staleBefore time.Time) ([]*OutboxEvent, error)
	MarkPublished(ctx context.Context, id int64, publishedAt time.Time) error
	MarkFailed(ctx context.Context, id int64) error
}

type EventClientMessage struct {
	Subject string
	Payload []byte
	Headers map[string]string
}

type EventClient interface {
	Publish(ctx context.Context, msg *EventClientMessage) error
}
