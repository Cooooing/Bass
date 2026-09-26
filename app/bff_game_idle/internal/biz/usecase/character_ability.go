package usecase

import (
	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
	"context"
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
