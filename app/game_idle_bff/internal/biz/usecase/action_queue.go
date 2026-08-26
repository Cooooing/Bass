package usecase

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

type ActionQueueUsecase struct {
	actionQueueRepo repo.ActionQueueRepo
}

func NewActionQueueUsecase(
	actionQueueRepo repo.ActionQueueRepo,
) *ActionQueueUsecase {
	return &ActionQueueUsecase{
		actionQueueRepo: actionQueueRepo,
	}
}

func (u *ActionQueueUsecase) List(ctx context.Context, characterID int64) (*model.ActionQueue, error) {
	return u.actionQueueRepo.List(ctx, characterID)
}

type AddActionReq struct {
	CharacterID int64
	ActionID    string
	Times       int64
	Position    *int32
}

func (u *ActionQueueUsecase) Add(ctx context.Context, req *AddActionReq) error {
	return u.actionQueueRepo.Add(ctx, &repo.AddActionReq{
		CharacterID: req.CharacterID,
		ActionID:    req.ActionID,
		Times:       req.Times,
		Position:    req.Position,
	})
}

type MoveActionReq struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (u *ActionQueueUsecase) Move(ctx context.Context, req *MoveActionReq) error {
	return u.actionQueueRepo.Move(ctx, &repo.MoveActionReq{
		CharacterID:     req.CharacterID,
		CurrentPosition: req.CurrentPosition,
		TargetPosition:  req.TargetPosition,
	})
}

type RemoveActionReq struct {
	CharacterID int64
	Position    int32
}

func (u *ActionQueueUsecase) Remove(ctx context.Context, req *RemoveActionReq) error {
	return u.actionQueueRepo.Remove(ctx, &repo.RemoveActionReq{
		CharacterID: req.CharacterID,
		Position:    req.Position,
	})
}

func (u *ActionQueueUsecase) Clear(ctx context.Context, characterID int64) error {
	return u.actionQueueRepo.Clear(ctx, characterID)
}
