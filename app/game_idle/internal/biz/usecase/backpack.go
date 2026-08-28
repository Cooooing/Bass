package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
)

type BackpackUsecase struct {
	characterRepo repo.CharacterRepo
	backpackRepo  repo.BackpackRepo
}

func NewBackpackUsecase(
	characterRepo repo.CharacterRepo,
	backpackRepo repo.BackpackRepo,
) *BackpackUsecase {
	return &BackpackUsecase{
		characterRepo: characterRepo,
		backpackRepo:  backpackRepo,
	}
}

type BackpackMapReq struct {
	CharacterID int64
	// ItemIDs 为空时返回全部库存，非空时只返回指定物品。
	ItemIDs []string
}

func (u *BackpackUsecase) Map(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error) {
	if req.CharacterID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	character, err := u.characterRepo.Get(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	if character.Status != enum.CharacterStatusActive {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	return u.backpackRepo.MapItems(ctx, &repo.BackpackMapReq{
		CharacterID: req.CharacterID,
		ItemIDs:     req.ItemIDs,
	})
}
