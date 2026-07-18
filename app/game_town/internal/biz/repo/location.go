package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type LocationRepo interface {
	GetLocationByCode(ctx context.Context, req *GetLocationByCodeReq) (*model.Location, error)
	GetLocation(ctx context.Context, id int64) (*model.Location, error)
}

type GetLocationByCodeReq struct {
	WorldID int64
	Code    string
}
