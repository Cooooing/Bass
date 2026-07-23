package repo

import (
	"context"
	"time"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type NpcRepo interface {
	Save(context.Context, *model.Npc) (*model.Npc, error)
	UpdateState(context.Context, *NpcStateUpdateReq) (*model.Npc, error)
	UpdateContext(context.Context, int64, int64, string) (*model.Npc, error)
	UpdateAutonomy(context.Context, *NpcAutonomyUpdateReq) (*model.Npc, error)
	Get(context.Context, *NpcQuery) (*model.Npc, error)
	List(context.Context, *NpcQuery) ([]*model.Npc, error)
	Map(context.Context, *NpcQuery) (map[int64]*model.Npc, error)
	Count(context.Context, *NpcQuery) (int, error)
	Page(context.Context, *NpcPageReq) (*NpcPageResp, error)
}

type NpcAutonomyUpdateReq struct {
	NpcID          int64
	Version        int64
	Goal           string
	ContextSummary string
	NextDecisionAt *time.Time
	LastPlannedAt  *time.Time
}

type NpcStateUpdateReq struct {
	NpcID             int64
	Version           int64
	CurrentLocationID *int64
	LifeStatus        *enum.NpcLifeStatus
	StateTags         []string
	Attributes        map[string]any
	DeathWorldTime    *time.Time
}

type NpcQuery struct {
	ID                 *int64
	WorldID            *int64
	LocationID         *int64
	IDs                []int64
	Code               *string
	NextDecisionBefore *time.Time
}

type NpcPageReq struct {
	Page  base.PageRequest
	Query NpcQuery
}

type NpcPageResp struct {
	Rows []*model.Npc
	Page base.PageResp
}
