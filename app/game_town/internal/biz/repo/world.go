package repo

import (
	"common/proto/gen/common"
	"context"
	"game_town/internal/biz/agent"
	"game_town/internal/biz/model"
)

type WorldRepo interface {
	CreateWorld(ctx context.Context, req *CreateWorldReq) (*CreateWorldResponse, error)
	Get(ctx context.Context, req *WorldGetReq) (*WorldGetResponse, error)
	Page(ctx context.Context, req *WorldPageReq) (*WorldPageResponse, error)
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

type CreateWorldResponse struct {
	World           *model.World
	DefaultLocation *model.Location
	Npcs            []*model.Npc
	State           *model.WorldStateSnapshot
	Events          []*model.Event
}

type WorldGetReq struct {
	ID int64
}

type WorldGetResponse struct {
	Row *model.World
}

type WorldPageReq struct {
	Page  *common.PageRequest
	Query WorldQuery
}

type WorldPageResponse struct {
	Rows []*model.World
	Page *common.PageResponse
}

type WorldQuery struct {
	CreatorPlayerID *int64
	Status          *string
}
