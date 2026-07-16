package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type MemoryRepo interface {
	ListMemories(ctx context.Context, req *ListMemoriesReq) (*ListMemoriesResponse, error)
	CreateMemory(ctx context.Context, req *CreateMemoryReq) (*CreateMemoryResponse, error)
}

type ListMemoriesReq struct {
	Query MemoryQuery
}

type ListMemoriesResponse struct {
	Rows []*model.Memory
}

type CreateMemoryReq struct {
	Row *model.Memory
}

type CreateMemoryResponse struct {
	Row *model.Memory
}

type MemoryQuery struct {
	WorldID  int64
	PlayerID int64
	NpcID    *int64
	Type     *string
}
