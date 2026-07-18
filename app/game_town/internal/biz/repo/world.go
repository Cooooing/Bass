package repo

import (
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/agent"
	"game_town/internal/biz/model"
)

type WorldRepo interface {
	CreateWorld(ctx context.Context, req *CreateWorldReq) (*CreateWorldResp, error)
	Get(ctx context.Context, id int64) (*model.World, error)
	Page(ctx context.Context, req *WorldPageReq) (*WorldPageResp, error)
}

type CreateWorldReq struct {
	CreatorPlayerID int64
	Description     string
	NpcCount        uint32
	LocationCount   uint32
	Scale           string
	Seed            int64
	StyleTags       []string
	AgentConfigID   *int64
	Generated       *agent.GenerateWorldOutput
}

type CreateWorldResp struct {
	World           *model.World
	DefaultLocation *model.Location
	Npcs            []*model.Npc
	State           *model.WorldStateSnapshot
	Events          []*model.Event
}

type WorldPageReq struct {
	Page  *common.PageReq
	Query WorldQuery
}

type WorldPageResp struct {
	Rows []*model.World
	Page *common.PageResp
}

type WorldQuery struct {
	CreatorPlayerID *int64
	Status          *string
}
