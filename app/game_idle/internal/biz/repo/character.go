package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// CharacterRepo 管理玩家角色的持久化。
type CharacterRepo interface {
	Save(ctx context.Context, character *model.Character) (*model.Character, error)
	Get(ctx context.Context, characterID int64) (*model.Character, error)
	ListByUserID(ctx context.Context, userID int64) ([]*model.Character, error)
}
