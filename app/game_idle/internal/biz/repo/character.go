package repo

import (
	"context"
	"game_idle/internal/biz/model"
	"time"
)

// CharacterRepo 管理玩家角色的持久化。
type CharacterRepo interface {
	Save(ctx context.Context, character *model.Character) (*model.Character, error)
	Get(ctx context.Context, characterID int64) (*model.Character, error)
	GetName(ctx context.Context, characterID int64) (string, error)
	List(ctx context.Context, req *ListCharacterReq) ([]*model.Character, error)
	UpdateLastOfflineAt(ctx context.Context, characterID int64, at time.Time) error
}

type ListCharacterReq struct {
	UserID      *int64
	CharacterID *int64
}
