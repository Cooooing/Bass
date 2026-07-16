package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type LocationRepo interface {
	GetLocationByCode(ctx context.Context, req *GetLocationByCodeReq) (*GetLocationByCodeResponse, error)
	GetLocation(ctx context.Context, req *GetLocationReq) (*GetLocationResponse, error)
}

type GetLocationByCodeReq struct {
	WorldID int64
	Code    string
}

type GetLocationByCodeResponse struct {
	Row *model.Location
}

type GetLocationReq struct {
	ID int64
}

type GetLocationResponse struct {
	Row *model.Location
}
