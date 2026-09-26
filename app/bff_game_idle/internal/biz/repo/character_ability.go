package repo

import (
	"bff_game_idle/internal/biz/model"
	"context"
)

type CharacterAbilityRepo interface {
	Map(ctx context.Context, characterID int64) (map[string]*model.CharacterAbility, error)
}
