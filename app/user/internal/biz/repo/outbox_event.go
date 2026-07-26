package repo

import (
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"context"
	"time"
	"user/internal/biz/model"
)

// OutboxEventSave defines an outbox event to persist.
type OutboxEventSave struct {
	Event *commonenums.Event
}

type OutboxEventRepo interface {
	Save(ctx context.Context, req *OutboxEventSave) error
	Get(ctx context.Context, req *OutboxEventGetReq) (*model.OutboxEvent, error)
	List(ctx context.Context, req *OutboxEventGetReq) ([]*model.OutboxEvent, error)
	Map(ctx context.Context, req *OutboxEventGetReq) (map[int64]*model.OutboxEvent, error)
	Count(ctx context.Context, req *OutboxEventGetReq) (int, error)
	Page(ctx context.Context, req *OutboxEventPageReq) (*OutboxEventPageResp, error)
	ClaimForPublish(ctx context.Context, req *OutboxEventClaimForPublishReq) ([]*model.OutboxEvent, error)
	MarkPublished(ctx context.Context, req *OutboxEventMarkPublishedReq) error
	MarkFailed(ctx context.Context, req *OutboxEventMarkFailedReq) error
}

type OutboxEventGetReq struct {
	ID       *int64
	IDs      []int64
	EventID  *string
	EventIDs []string
	Subject  *commonenums.EventSubject
	Status   *commonenum.OutboxEventStatus
}

type OutboxEventPageReq struct {
	Page  PageReq
	Query OutboxEventGetReq
}

type OutboxEventPageResp struct {
	Rows []*model.OutboxEvent
	Page PageResp
}

type OutboxEventClaimForPublishReq struct {
	Limit       int
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
	Subject string
	Payload string
	Headers map[string]string
}

type EventClient interface {
	Publish(ctx context.Context, message *EventClientMessage) error
}
