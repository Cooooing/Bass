package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
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

func (u *CharacterAbilityUsecase) Map(
	ctx context.Context,
	characterID int64,
) (map[enum.Ability]*model.CharacterAbility, error) {
	rows, err := u.characterAbilityRepo.Map(ctx, &repo.CharacterAbilityMapReq{
		CharacterID: characterID,
	})
	if err != nil {
		return nil, err
	}
	for _, abilityID := range enum.AbilityValues() {
		ability := enum.Ability(abilityID)
		if rows[ability] == nil {
			rows[ability] = &model.CharacterAbility{
				CharacterID:  characterID,
				AbilityID:    ability,
				Level:        1,
				Exp:          0,
				NextLevelExp: 100,
			}
		}
	}
	return rows, nil
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ABILITY_LEVEL_INSUFFICIENT)
	}
	if ability.Level < requiredLevel {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ABILITY_LEVEL_INSUFFICIENT)
	}
	return nil
}
