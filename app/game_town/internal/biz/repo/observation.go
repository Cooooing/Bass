package repo

import (
	"context"
	"time"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type ObservationRepo interface {
	Save(context.Context, *model.Observation) (*model.Observation, error)
	Get(context.Context, *ObservationQuery) (*model.Observation, error)
	List(context.Context, *ObservationQuery) ([]*model.Observation, error)
	Map(context.Context, *ObservationQuery) (map[int64]*model.Observation, error)
	Count(context.Context, *ObservationQuery) (int, error)
	Page(context.Context, *ObservationPageReq) (*ObservationPageResp, error)
}

type ObservationQuery struct {
	ID, WorldID, EventID, NpcID, PlayerID *int64
	EventIDs                              []int64
	AfterEventID                          *int64
	AfterEventSequence                    *uint64
	AfterWorldTime                        *time.Time
	EventType                             *enum.EventType
}
type ObservationPageReq struct {
	Page      base.PageRequest
	Query     ObservationQuery
	SkipTotal bool
}
type ObservationPageResp struct {
	Rows []*model.Observation
	Page base.PageResp
}
