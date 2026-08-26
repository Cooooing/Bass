package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

type BackpackRepo interface {
	Map(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error)
}

type BackpackMapReq struct {
	CharacterID int64
	ItemIDs     []string
}
