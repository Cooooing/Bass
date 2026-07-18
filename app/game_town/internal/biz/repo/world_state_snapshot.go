package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type WorldStateSnapshotRepo interface {
	GetLatestWorldState(ctx context.Context, worldID int64) (*model.WorldStateSnapshot, error)
	CreateState(ctx context.Context, row *model.WorldStateSnapshot) (*model.WorldStateSnapshot, error)
}
