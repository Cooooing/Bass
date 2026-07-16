package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"time"
)

type InboxEventRepo interface {
	SaveProcessing(ctx context.Context, req *InboxEventSaveProcessingReq) (*InboxEventSaveProcessingResponse, error)
	Get(ctx context.Context, req *InboxEventGetReq) (*InboxEventGetResponse, error)
	List(ctx context.Context, req *InboxEventListReq) (*InboxEventListResponse, error)
	Map(ctx context.Context, req *InboxEventMapReq) (*InboxEventMapResponse, error)
	Count(ctx context.Context, req *InboxEventCountReq) (*InboxEventCountResponse, error)
	Page(ctx context.Context, req *InboxEventPageReq) (*InboxEventPageResponse, error)
	ClaimRetry(ctx context.Context, req *InboxEventClaimRetryReq) (*InboxEventClaimRetryResponse, error)
	MarkProcessed(ctx context.Context, req *InboxEventMarkProcessedReq) (*InboxEventMarkProcessedResponse, error)
	MarkFailed(ctx context.Context, req *InboxEventMarkFailedReq) (*InboxEventMarkFailedResponse, error)
}

type InboxEventSaveProcessingReq struct {
	EventID   string
	EventType enums.EventType
	Subject   commonenum.EventSubject
	Payload   string
	Now       time.Time
}

type InboxEventSaveProcessingResponse struct {
	Event   *model.InboxEvent
	Claimed bool
}

type InboxEventGetReq struct {
	Query *InboxEventQuery
}

type InboxEventGetResponse struct {
	Item *model.InboxEvent
}

type InboxEventListReq struct {
	Query *InboxEventQuery
}

type InboxEventListResponse struct {
	Rows []*model.InboxEvent
}

type InboxEventMapReq struct {
	Query *InboxEventQuery
}

type InboxEventMapResponse struct {
	Rows map[int64]*model.InboxEvent
}

type InboxEventCountReq struct {
	Query *InboxEventQuery
}

type InboxEventCountResponse struct {
	Count int
}

type InboxEventPageReq struct {
	Query *InboxEventQuery
}

type InboxEventPageResponse struct {
	Rows []*model.InboxEvent
	Page *base.PageResponse
}

type InboxEventClaimRetryReq struct {
	EventID           string
	Now               time.Time
	ProcessingTimeout time.Duration
	MaxRetry          int32
}

type InboxEventClaimRetryResponse struct {
	Claimed bool
}

type InboxEventMarkProcessedReq struct {
	EventID string
	Now     time.Time
}

type InboxEventMarkProcessedResponse struct{}

type InboxEventMarkFailedReq struct {
	EventID   string
	LastError string
	MaxRetry  int32
}

type InboxEventMarkFailedResponse struct{}

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
