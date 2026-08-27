package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

type CharacterAbilityRepo interface {
	Map(ctx context.Context, characterID int64) (map[string]*model.CharacterAbility, error)
}
