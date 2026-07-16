package repo

import (
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/model"
)

type EventRepo interface {
	CreateEvent(ctx context.Context, req *CreateEventReq) (*CreateEventResponse, error)
	Page(ctx context.Context, req *EventPageReq) (*EventPageResponse, error)
	ListRecentEvents(ctx context.Context, req *ListRecentEventsReq) (*ListRecentEventsResponse, error)
}

type CreateEventReq struct {
	Row *model.Event
}

type CreateEventResponse struct {
	Row *model.Event
}

type EventPageReq struct {
	Page  *common.PageRequest
	Query EventQuery
}

type EventPageResponse struct {
	Rows []*model.Event
	Page *common.PageResponse
}

type ListRecentEventsReq struct {
	WorldID int64
	Limit   int
}

type ListRecentEventsResponse struct {
	Rows []*model.Event
}

type EventQuery struct {
	WorldID       int64
	ActorPlayerID *int64
	TargetNpcID   *int64
	Type          *string
}
