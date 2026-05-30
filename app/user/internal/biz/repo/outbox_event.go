package repo

import (
	commonenums "common/api/gen/common/enums"
	"context"
	"time"
	"user/internal/biz/model"
)

// OutboxEventSave 定义一条待持久化的 outbox 事件。
type OutboxEventSave struct {
	Event *commonenums.Event
}

type OutboxEventRepo interface {
	Save(ctx context.Context, req *OutboxEventSave) error
	ClaimForPublish(ctx context.Context, limit int, staleBefore time.Time) ([]*model.OutboxEvent, error)
	MarkPublished(ctx context.Context, id int64, publishedAt time.Time) error
	MarkFailed(ctx context.Context, id int64) error
}

type EventClientMessage struct {
	Subject string
	Payload string
	Headers map[string]string
}

type EventClient interface {
	Publish(ctx context.Context, msg *EventClientMessage) error
}
