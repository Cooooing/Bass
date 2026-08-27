package usecase

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

type CharacterAbilityUsecase struct {
	characterAbilityRepo repo.CharacterAbilityRepo
}

func NewCharacterAbilityUsecase(
	characterAbilityRepo repo.CharacterAbilityRepo,
) *CharacterAbilityUsecase {
	return &CharacterAbilityUsecase{
		characterAbilityRepo: characterAbilityRepo,
	}
}

func (u *CharacterAbilityUsecase) Map(
	ctx context.Context,
	characterID int64,
) (map[string]*model.CharacterAbility, error) {
	return u.characterAbilityRepo.Map(ctx, characterID)
}
