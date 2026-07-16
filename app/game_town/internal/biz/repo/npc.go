package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type NpcRepo interface {
	ListNpcs(ctx context.Context, req *ListNpcsReq) (*ListNpcsResponse, error)
	GetNpc(ctx context.Context, req *GetNpcReq) (*GetNpcResponse, error)
	GetNpcByCode(ctx context.Context, req *GetNpcByCodeReq) (*GetNpcByCodeResponse, error)
}

type ListNpcsReq struct {
	WorldID    int64
	LocationID *int64
}

type ListNpcsResponse struct {
	Rows []*model.Npc
}

type GetNpcReq struct {
	ID int64
}

type GetNpcResponse struct {
	Row *model.Npc
}

type GetNpcByCodeReq struct {
	WorldID int64
	Code    string
}

type GetNpcByCodeResponse struct {
	Row *model.Npc
}
