package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type NpcRepo interface {
	ListNpcs(ctx context.Context, req *ListNpcsReq) ([]*model.Npc, error)
	GetNpc(ctx context.Context, id int64) (*model.Npc, error)
	GetNpcByCode(ctx context.Context, req *GetNpcByCodeReq) (*model.Npc, error)
}

type ListNpcsReq struct {
	WorldID    int64
	LocationID *int64
}

type GetNpcByCodeReq struct {
	WorldID int64
	Code    string
}
