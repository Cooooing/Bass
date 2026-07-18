package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type MemoryRepo interface {
	ListMemories(ctx context.Context, query *MemoryQuery) ([]*model.Memory, error)
	CreateMemory(ctx context.Context, row *model.Memory) (*model.Memory, error)
}

type MemoryQuery struct {
	WorldID  int64
	PlayerID int64
	NpcID    *int64
	Type     *string
}
