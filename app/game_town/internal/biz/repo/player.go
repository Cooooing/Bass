package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type PlayerRepo interface {
	Save(ctx context.Context, player *model.Player) (*model.Player, error)
	Get(ctx context.Context, req *PlayerQuery) (*model.Player, error)
	List(ctx context.Context, req *PlayerQuery) ([]*model.Player, error)
	Map(ctx context.Context, req *PlayerQuery) (map[int64]*model.Player, error)
	Count(ctx context.Context, req *PlayerQuery) (int, error)
	Page(ctx context.Context, req *PlayerPageReq) (*PlayerPageResp, error)
}

type PlayerQuery struct {
	ID     *int64
	IDs    []int64
	Name   *string
	Status *enum.PlayerStatus
}

type PlayerPageReq struct {
	Page  base.PageRequest
	Query PlayerQuery
}

type PlayerPageResp struct {
	Rows []*model.Player
	Page base.PageResp
}
