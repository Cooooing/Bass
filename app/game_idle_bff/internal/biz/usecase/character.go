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
	UserID      int64
	CharacterID int64
}

func (u *CharacterUsecase) Online(ctx context.Context, req *OnlineCharacterReq) (*model.WebSocketSession, error) {
	return u.characterRepo.Online(ctx, &repo.OnlineCharacterReq{
		UserID:      req.UserID,
		CharacterID: req.CharacterID,
	})
}

type PingCharacterReq struct {
	CharacterID int64
	SessionID   string
}

func (u *CharacterUsecase) Ping(ctx context.Context, req *PingCharacterReq) (*model.WebSocketSession, error) {
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
