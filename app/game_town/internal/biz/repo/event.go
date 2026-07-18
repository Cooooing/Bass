package repo

import (
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/model"
)

type EventRepo interface {
	CreateEvent(ctx context.Context, row *model.Event) (*model.Event, error)
	Page(ctx context.Context, req *EventPageReq) (*EventPageResp, error)
	ListRecentEvents(ctx context.Context, req *ListRecentEventsReq) ([]*model.Event, error)
}

type EventPageReq struct {
	Page  *common.PageReq
	Query EventQuery
}

type EventPageResp struct {
	Rows []*model.Event
	Page *common.PageResp
}

type ListRecentEventsReq struct {
	WorldID int64
	Limit   int
}

type EventQuery struct {
	WorldID       int64
	ActorPlayerID *int64
	TargetNpcID   *int64
	Type          *string
}
