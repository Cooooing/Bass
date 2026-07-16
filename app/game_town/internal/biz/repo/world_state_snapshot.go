package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type WorldStateSnapshotRepo interface {
	GetLatestWorldState(ctx context.Context, req *GetLatestWorldStateReq) (*GetLatestWorldStateResponse, error)
	CreateState(ctx context.Context, req *CreateStateReq) (*CreateStateResponse, error)
}

type GetLatestWorldStateReq struct {
	WorldID int64
}

type GetLatestWorldStateResponse struct {
	Row *model.WorldStateSnapshot
}

type CreateStateReq struct {
	Row *model.WorldStateSnapshot
}

type CreateStateResponse struct {
	Row *model.WorldStateSnapshot
}
