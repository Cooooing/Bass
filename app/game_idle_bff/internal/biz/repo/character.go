package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

type CharacterRepo interface {
	Create(ctx context.Context, req *CreateCharacterReq) (*model.Character, error)
	List(ctx context.Context, userID int64) ([]*model.Character, error)
}

type CreateCharacterReq struct {
	UserID int64
	Name   string
}
