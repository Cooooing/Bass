package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type EventRepo interface {
	Save(ctx context.Context, event *model.Event) (*model.Event, error)
	Get(ctx context.Context, req *EventQuery) (*model.Event, error)
	List(ctx context.Context, req *EventQuery) ([]*model.Event, error)
	Map(ctx context.Context, req *EventQuery) (map[int64]*model.Event, error)
	Count(ctx context.Context, req *EventQuery) (int, error)
	Page(ctx context.Context, req *EventPageReq) (*EventPageResp, error)
}

type EventQuery struct {
	ID               *int64
	IDs              []int64
	WorldID          *int64
	AfterSequence    *uint64
	Type             *enum.EventType
	ActorPlayerID    *int64
	NpcID            *int64
	RecentLimit      int
	Limit            int
	CausationEventID *int64
}

type EventPageReq struct {
	Page  base.PageRequest
	Query EventQuery
}

type EventPageResp struct {
	Rows []*model.Event
	Page base.PageResp
}
