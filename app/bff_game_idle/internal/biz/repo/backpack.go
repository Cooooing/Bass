package repo

import (
	"bff_game_idle/internal/biz/model"
	"context"
)

type BackpackRepo interface {
	Map(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error)
}

type BackpackMapReq struct {
	CharacterID int64
	ItemIDs     []string
}
