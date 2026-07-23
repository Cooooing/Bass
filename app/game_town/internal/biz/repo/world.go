package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type WorldRepo interface {
	Save(ctx context.Context, world *model.World) (*model.World, error)
	Update(ctx context.Context, world *model.World) (*model.World, error)
	Get(ctx context.Context, req *WorldQuery) (*model.World, error)
	List(ctx context.Context, req *WorldQuery) ([]*model.World, error)
	Map(ctx context.Context, req *WorldQuery) (map[int64]*model.World, error)
	Count(ctx context.Context, req *WorldQuery) (int, error)
	Page(ctx context.Context, req *WorldPageReq) (*WorldPageResp, error)
}

type WorldQuery struct {
	ID              *int64
	IDs             []int64
	Code            *string
	CreatorPlayerID *int64
	Status          *enum.WorldStatus
}

type WorldPageReq struct {
	Page  base.PageRequest
	Query WorldQuery
}

type WorldPageResp struct {
	Rows []*model.World
	Page base.PageResp
}
