package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
)

type CharacterAbilityUsecase struct {
	characterAbilityRepo repo.CharacterAbilityRepo
}

func NewCharacterAbilityUsecase(characterAbilityRepo repo.CharacterAbilityRepo) *CharacterAbilityUsecase {
	return &CharacterAbilityUsecase{
		characterAbilityRepo: characterAbilityRepo,
	}
}

func (u *CharacterAbilityUsecase) CheckLevel(
	ctx context.Context,
	characterID int64,
	abilityID enum.Ability,
	requiredLevel int32,
) error {
	if abilityID == "" || requiredLevel <= 0 {
		return nil
	}
	abilities, err := u.characterAbilityRepo.Map(ctx, &repo.CharacterAbilityMapReq{
		CharacterID: characterID,
		AbilityIDs:  []enum.Ability{abilityID},
	})
	if err != nil {
		return err
	}
	ability := abilities[abilityID]
	if ability == nil {
		if requiredLevel <= 1 {
			return nil
		}
		return model.ErrAbilityLevelInsufficient
	}
	if ability.Level < requiredLevel {
		return model.ErrAbilityLevelInsufficient
	}
	return nil
}
