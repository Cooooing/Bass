package usecase

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
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
