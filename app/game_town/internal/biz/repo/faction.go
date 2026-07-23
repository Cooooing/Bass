package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type FactionRepo interface {
	Save(context.Context, *model.Faction) (*model.Faction, error)
	UpdateState(context.Context, *FactionStateUpdateReq) (*model.Faction, error)
	Get(context.Context, *FactionQuery) (*model.Faction, error)
	List(context.Context, *FactionQuery) ([]*model.Faction, error)
	Map(context.Context, *FactionQuery) (map[int64]*model.Faction, error)
	Count(context.Context, *FactionQuery) (int, error)
	Page(context.Context, *FactionPageReq) (*FactionPageResp, error)
}

type FactionStateUpdateReq struct {
	FactionID   int64
	Version     int64
	Status      *enum.FactionStatus
	Description string
	PublicGoal  string
	Attributes  map[string]any
}

type FactionQuery struct {
	ID      *int64
	WorldID *int64
	IDs     []int64
	Code    *string
}

type FactionPageReq struct {
	Page  base.PageRequest
	Query FactionQuery
}

type FactionPageResp struct {
	Rows []*model.Faction
	Page base.PageResp
}
