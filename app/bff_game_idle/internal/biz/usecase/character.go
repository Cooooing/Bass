package usecase

import (
	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
	"context"
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

type ListCharacterReq struct {
	UserID      int64
	CharacterID int64
}

func (u *CharacterUsecase) List(ctx context.Context, req *ListCharacterReq) ([]*model.Character, error) {
	return u.characterRepo.List(ctx, &repo.ListCharacterReq{
		UserID:      req.UserID,
		CharacterID: req.CharacterID,
	})
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

type OnlineCharacterReq struct {
	CharacterID int64
}

func (u *CharacterUsecase) Online(ctx context.Context, req *OnlineCharacterReq) (*model.CharacterOnlineSession, error) {
	return u.characterRepo.Online(ctx, &repo.OnlineCharacterReq{
		CharacterID: req.CharacterID,
	})
}

type PingCharacterReq struct {
	CharacterID int64
	SessionID   string
}

func (u *CharacterUsecase) Ping(ctx context.Context, req *PingCharacterReq) (*model.CharacterOnlineSession, error) {
	return u.characterRepo.Ping(ctx, &repo.PingCharacterReq{
		CharacterID: req.CharacterID,
		SessionID:   req.SessionID,
	})
}

type OfflineCharacterReq struct {
	CharacterID int64
	SessionID   string
	Timeout     bool
}

func (u *CharacterUsecase) Offline(ctx context.Context, req *OfflineCharacterReq) error {
	return u.characterRepo.Offline(ctx, &repo.OfflineCharacterReq{
		CharacterID: req.CharacterID,
		SessionID:   req.SessionID,
		Timeout:     req.Timeout,
	})
}
