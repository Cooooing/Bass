package usecase

import (
	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
	"context"
)

type BackpackUsecase struct {
	backpackRepo repo.BackpackRepo
}

func NewBackpackUsecase(
	backpackRepo repo.BackpackRepo,
) *BackpackUsecase {
	return &BackpackUsecase{
		backpackRepo: backpackRepo,
	}
}

type BackpackMapReq struct {
	CharacterID int64
	ItemIDs     []string
}

func (u *BackpackUsecase) Map(ctx context.Context, req *BackpackMapReq) (map[string]*model.CharacterItem, error) {
	return u.backpackRepo.Map(ctx, &repo.BackpackMapReq{
		CharacterID: req.CharacterID,
		ItemIDs:     req.ItemIDs,
	})
}
