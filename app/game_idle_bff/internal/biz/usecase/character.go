package usecase

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

type CharacterUsecase struct {
	characterRepo repo.CharacterRepo
}

func NewCharacterUsecase(
	characterRepo repo.CharacterRepo,
) *CharacterUsecase {
	return &CharacterUsecase{
		characterRepo: characterRepo,
	}
}

func (u *CharacterUsecase) List(ctx context.Context, userID int64) ([]*model.Character, error) {
	return u.characterRepo.List(ctx, userID)
}

type CreateCharacterReq struct {
	UserID int64
	Name   string
}

func (u *CharacterUsecase) Create(ctx context.Context, req *CreateCharacterReq) (*model.Character, error) {
	return u.characterRepo.Create(ctx, &repo.CreateCharacterReq{
		UserID: req.UserID,
		Name:   req.Name,
	})
}
