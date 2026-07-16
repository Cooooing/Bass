package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type PlayerRepo interface {
	GetPlayer(ctx context.Context, req *GetPlayerReq) (*GetPlayerResponse, error)
	GetPlayerByName(ctx context.Context, req *GetPlayerByNameReq) (*GetPlayerByNameResponse, error)
	CreatePlayer(ctx context.Context, req *CreatePlayerReq) (*CreatePlayerResponse, error)
}

type GetPlayerReq struct {
	ID int64
}

type GetPlayerResponse struct {
	Row *model.Player
}

type GetPlayerByNameReq struct {
	Name string
}

type GetPlayerByNameResponse struct {
	Row *model.Player
}

type CreatePlayerReq struct {
	Name        string
	DisplayName string
}

type CreatePlayerResponse struct {
	Row *model.Player
}
