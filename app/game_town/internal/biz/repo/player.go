package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type PlayerRepo interface {
	GetPlayer(ctx context.Context, id int64) (*model.Player, error)
	GetPlayerByName(ctx context.Context, name string) (*model.Player, error)
	CreatePlayer(ctx context.Context, req *CreatePlayerReq) (*model.Player, error)
}

type CreatePlayerReq struct {
	Name        string
	DisplayName string
}
