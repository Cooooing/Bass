package repo

import (
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"content/internal/biz/base"
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
	Save(ctx context.Context, event *commonenums.Event) error
	Get(ctx context.Context, req *OutboxEventGetReq) (*OutboxEvent, error)
	List(ctx context.Context, req *OutboxEventGetReq) ([]*OutboxEvent, error)
	Map(ctx context.Context, req *OutboxEventGetReq) (map[int64]*OutboxEvent, error)
	Count(ctx context.Context, req *OutboxEventGetReq) (int, error)
	Page(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventPageResp, error)
	ClaimForPublish(ctx context.Context, req *OutboxEventClaimForPublishReq) ([]*OutboxEvent, error)
	MarkPublished(ctx context.Context, req *OutboxEventMarkPublishedReq) error
	MarkFailed(ctx context.Context, req *OutboxEventMarkFailedReq) error
}

type OutboxEventPageResp struct {
	Rows []*OutboxEvent
	Page *base.PageResp
}

type OutboxEventGetReq struct {
	Page     *base.PageRequest
	ID       *int64
	IDs      []int64
	EventID  *string
	EventIDs []string
	Subject  *commonenums.EventSubject
	Status   *commonenum.OutboxEventStatus
}

type OutboxEventClaimForPublishReq struct {
	Limit       int
	StaleBefore time.Time
}

type OutboxEventMarkPublishedReq struct {
	ID          int64
	PublishedAt time.Time
}

type OutboxEventMarkFailedReq struct {
	ID        int64
	LastError string
	MaxRetry  int32
}

type EventClientMessage struct {
	Subject string
	Payload string
	Headers map[string]string
}

type EventClient interface {
	Publish(ctx context.Context, message *EventClientMessage) error
}
