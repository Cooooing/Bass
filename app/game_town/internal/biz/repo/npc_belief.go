package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
)

type NpcBeliefRepo interface {
	Save(context.Context, *model.NpcBelief) (*model.NpcBelief, error)
	Get(context.Context, *NpcBeliefQuery) (*model.NpcBelief, error)
	List(context.Context, *NpcBeliefQuery) ([]*model.NpcBelief, error)
	Map(context.Context, *NpcBeliefQuery) (map[int64]*model.NpcBelief, error)
	Count(context.Context, *NpcBeliefQuery) (int, error)
	Page(context.Context, *NpcBeliefPageReq) (*NpcBeliefPageResp, error)
}

type NpcBeliefQuery struct {
	ID, WorldID, NpcID, ClaimID *int64
	MinConfidence               *float64
}
type NpcBeliefPageReq struct {
	Page  base.PageRequest
	Query NpcBeliefQuery
}
type NpcBeliefPageResp struct {
	Rows []*model.NpcBelief
	Page base.PageResp
}
