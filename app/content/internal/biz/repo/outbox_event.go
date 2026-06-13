package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	commonenums "common/proto/gen/common/enums"
	"context"
	"time"
)

// OutboxEventSave 定义一条待持久化的 outbox 事件。
type OutboxEventSave struct {
	Event *commonenums.Event
}

type OutboxEvent struct {
	ID         int64
	EventID    string
	EventType  commonenum.EventType
	Subject    commonenum.EventSubject
	Payload    string
	Headers    map[string]string
	Status     commonenum.OutboxEventStatus
	RetryCount int32
	LastError  *string
	UpdatedAt  *time.Time
}

type OutboxEventRepo interface {
	Save(ctx context.Context, req *OutboxEventSave) error
	Get(ctx context.Context, req *OutboxEventGetReq) (*OutboxEvent, error)
	List(ctx context.Context, req *OutboxEventGetReq) ([]*OutboxEvent, error)
	Map(ctx context.Context, req *OutboxEventGetReq) (map[int64]*OutboxEvent, error)
	Count(ctx context.Context, req *OutboxEventGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *OutboxEventGetReq) ([]*OutboxEvent, *common.PageReply, error)
	ClaimForPublish(ctx context.Context, limit int, staleBefore time.Time) ([]*OutboxEvent, error)
	MarkPublished(ctx context.Context, id int64, publishedAt time.Time) error
	MarkFailed(ctx context.Context, id int64, lastError string, maxRetry int32) error
}

type OutboxEventGetReq struct {
	ID       *int64
	IDs      []int64
	EventID  *string
	EventIDs []string
	Subject  *commonenums.EventSubject
	Status   *commonenum.OutboxEventStatus
}

type EventClientMessage struct {
	Subject string
	Payload string
	Headers map[string]string
}

type EventClient interface {
	Publish(ctx context.Context, msg *EventClientMessage) error
}
