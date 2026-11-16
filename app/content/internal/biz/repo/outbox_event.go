package repo

import (
	commonenum "common/pkg/enum"
	commonmodel "common/pkg/model"
	commonevent "common/proto/gen/common/enums/event"
	"context"
	"time"
)

type OutboxEvent struct {
	ID         int64
	EventID    string
	Subject    commonenum.EventSubject
	Payload    string
	Headers    map[string]string
	Status     commonenum.OutboxEventStatus
	RetryCount int32
	LastError  *string
	UpdatedAt  *time.Time
}

type OutboxEventRepo interface {
	Save(ctx context.Context, event *commonevent.Event) (*OutboxEvent, error)
	Get(ctx context.Context, req *OutboxEventGetReq) (*OutboxEvent, error)
	List(ctx context.Context, req *OutboxEventGetReq) ([]*OutboxEvent, error)
	Map(ctx context.Context, req *OutboxEventGetReq) (map[int64]*OutboxEvent, error)
	Count(ctx context.Context, req *OutboxEventGetReq) (int, error)
	Page(ctx context.Context, req *OutboxEventGetReq) (*OutboxEventPageResp, error)
	ClaimOneForPublish(ctx context.Context, req *OutboxEventClaimOneForPublishReq) (*OutboxEvent, error)
	ClaimForPublish(ctx context.Context, req *OutboxEventClaimForPublishReq) ([]*OutboxEvent, error)
	MarkPublished(ctx context.Context, req *OutboxEventMarkPublishedReq) error
	MarkFailed(ctx context.Context, req *OutboxEventMarkFailedReq) error
}

type OutboxEventPageResp struct {
	Rows []*OutboxEvent
	Page *commonmodel.PageResp
}

type OutboxEventGetReq struct {
	Page     *commonmodel.PageReq
	ID       *int64
	IDs      []int64
	EventID  *string
	EventIDs []string
	Subject  *commonevent.EventSubject
	Status   *commonenum.OutboxEventStatus
}

type OutboxEventClaimForPublishReq struct {
	Limit       int
	StaleBefore *time.Time
}

type OutboxEventClaimOneForPublishReq struct {
	ID          int64
	StaleBefore *time.Time
}

type OutboxEventMarkPublishedReq struct {
	ID          int64
	PublishedAt *time.Time
}

type OutboxEventMarkFailedReq struct {
	ID        int64
	LastError string
	MaxRetry  int32
}

type EventClientMessage struct {
	EventID string
	Subject string
	Payload string
	Headers map[string]string
}

type EventClient interface {
	Ensure(context.Context) error
	Publish(ctx context.Context, message *EventClientMessage) error
}
