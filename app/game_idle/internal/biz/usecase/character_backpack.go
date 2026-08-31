package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
)

type CharacterBackpackUsecase struct {
	characterRepo repo.CharacterRepo
	backpackRepo  repo.CharacterBackpackRepo
	stateUsecase  *CharacterStateUsecase
}

func NewCharacterBackpackUsecase(
	characterRepo repo.CharacterRepo,
	backpackRepo repo.CharacterBackpackRepo,
	stateUsecase *CharacterStateUsecase,
) *CharacterBackpackUsecase {
	return &CharacterBackpackUsecase{
		characterRepo: characterRepo,
		backpackRepo:  backpackRepo,
		stateUsecase:  stateUsecase,
	}
}

type BackpackMapReq struct {
	CharacterID int64
	// ItemIDs 为空时返回全部库存，非空时只返回指定物品。
	ItemIDs []string
}

func (u *CharacterBackpackUsecase) Map(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error) {
	if req.CharacterID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	character, err := u.characterRepo.Get(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	if character == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
	}
	if character.Status != enum.CharacterStatusActive {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	var items map[string]*model.CharacterItem
	err = u.stateUsecase.Execute(ctx, req.CharacterID, func(runCtx context.Context) error {
		var err error
		items, err = u.backpackRepo.MapItems(runCtx, &repo.BackpackMapReq{
			CharacterID: req.CharacterID,
			ItemIDs:     req.ItemIDs,
		})
		return err
	})
	return items, err
}

func (u *CharacterBackpackUsecase) Persist(ctx context.Context, characterID int64) error {
	return u.stateUsecase.Persist(ctx, characterID)
}
