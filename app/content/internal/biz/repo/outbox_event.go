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
	Save(ctx context.Context, req *OutboxEventSave) (*OutboxEventSaveResponse, error)
	Get(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventGetResponse, error)
	List(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventListResponse, error)
	Map(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventMapResponse, error)
	Count(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventCountResponse, error)
	Page(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventPageResponse, error)
	ClaimForPublish(ctx context.Context, req *OutboxEventClaimForPublishReq) (*OutboxEventClaimForPublishResponse, error)
	MarkPublished(ctx context.Context, req *OutboxEventMarkPublishedReq) (*OutboxEventMarkPublishedResponse, error)
	MarkFailed(ctx context.Context, req *OutboxEventMarkFailedReq) (*OutboxEventMarkFailedResponse, error)
}

type OutboxEventSaveResponse struct{}

type OutboxEventGetResponse struct {
	Event *OutboxEvent
}

type OutboxEventListResponse struct {
	Rows []*OutboxEvent
}

type OutboxEventMapResponse struct {
	Rows map[int64]*OutboxEvent
}

type OutboxEventCountResponse struct {
	Count int
}

type OutboxEventPageResponse struct {
	Rows []*OutboxEvent
	Page *base.PageResponse
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

type OutboxEventClaimForPublishResponse struct {
	Rows []*OutboxEvent
}

type OutboxEventMarkPublishedReq struct {
	ID          int64
	PublishedAt time.Time
}

type OutboxEventMarkPublishedResponse struct{}

type OutboxEventMarkFailedReq struct {
	ID        int64
	LastError string
	MaxRetry  int32
}

type OutboxEventMarkFailedResponse struct{}

type EventClientMessage struct {
	Subject string
	Payload string
	Headers map[string]string
}

type EventClient interface {
	Publish(ctx context.Context, req *EventClientPublishReq) (*EventClientPublishResponse, error)
}

type EventClientPublishReq struct {
	Message *EventClientMessage
}

type EventClientPublishResponse struct{}
