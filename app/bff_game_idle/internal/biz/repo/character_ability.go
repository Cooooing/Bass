package repo

import (
	"context"

	"bff_game_idle/internal/biz/model"
)

type CharacterAbilityRepo interface {
	Map(ctx context.Context, characterID int64) (map[string]*model.CharacterAbility, error)
}
