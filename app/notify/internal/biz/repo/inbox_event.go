package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"time"
)

type InboxEventRepo interface {
	SaveProcessing(ctx context.Context, req *InboxEventSaveProcessingReq) (*InboxEventSaveProcessingResp, error)
	Get(ctx context.Context, query *InboxEventQuery) (*model.InboxEvent, error)
	List(ctx context.Context, query *InboxEventQuery) ([]*model.InboxEvent, error)
	Map(ctx context.Context, query *InboxEventQuery) (map[int64]*model.InboxEvent, error)
	Count(ctx context.Context, query *InboxEventQuery) (int, error)
	Page(ctx context.Context, query *InboxEventQuery) (*InboxEventPageResp, error)
	ClaimRetry(ctx context.Context, req *InboxEventClaimRetryReq) (bool, error)
	MarkProcessed(ctx context.Context, req *InboxEventMarkProcessedReq) error
	MarkFailed(ctx context.Context, req *InboxEventMarkFailedReq) error
}

type InboxEventSaveProcessingReq struct {
	EventID   string
	EventType commonenum.EventType
	Subject   commonenum.EventSubject
	Payload   string
	Now       time.Time
}

type InboxEventSaveProcessingResp struct {
	Event   *model.InboxEvent
	Claimed bool
}

type InboxEventPageResp struct {
	Rows []*model.InboxEvent
	Page *base.PageResp
}

type InboxEventClaimRetryReq struct {
	EventID           string
	Now               time.Time
	ProcessingTimeout time.Duration
	MaxRetry          int32
}

type InboxEventMarkProcessedReq struct {
	EventID string
	Now     time.Time
}

type InboxEventMarkFailedReq struct {
	EventID   string
	LastError string
	MaxRetry  int32
}

type InboxEventQuery struct {
	Page      *base.PageRequest
	ID        *int64
	IDs       []int64
	EventID   *string
	EventIDs  []string
	EventType *commonenum.EventType
	Subject   *commonenum.EventSubject
	Status    *commonenum.InboxEventStatus
}
