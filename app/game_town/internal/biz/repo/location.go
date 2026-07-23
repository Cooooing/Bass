package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type LocationRepo interface {
	Save(context.Context, *model.Location) (*model.Location, error)
	UpdateState(context.Context, *LocationStateUpdateReq) (*model.Location, error)
	Get(context.Context, *LocationQuery) (*model.Location, error)
	List(context.Context, *LocationQuery) ([]*model.Location, error)
	Map(context.Context, *LocationQuery) (map[int64]*model.Location, error)
	Count(context.Context, *LocationQuery) (int, error)
	Page(context.Context, *LocationPageReq) (*LocationPageResp, error)
}

type LocationStateUpdateReq struct {
	LocationID           int64
	Version              int64
	Status               *enum.LocationStatus
	Description          string
	Accessible           *bool
	ControllingFactionID *int64
	EnvironmentTags      []string
	Attributes           map[string]any
}

type LocationQuery struct {
	ID      *int64
	IDs     []int64
	WorldID *int64
	Code    *string
}

type LocationPageReq struct {
	Page  base.PageRequest
	Query LocationQuery
}

type LocationPageResp struct {
	Rows []*model.Location
	Page base.PageResp
}
