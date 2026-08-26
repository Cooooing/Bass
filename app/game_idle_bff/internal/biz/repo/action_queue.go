package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

type ActionQueueRepo interface {
	List(ctx context.Context, characterID int64) (*model.ActionQueue, error)
	Add(ctx context.Context, req *AddActionReq) error
	Move(ctx context.Context, req *MoveActionReq) error
	Remove(ctx context.Context, req *RemoveActionReq) error
	Clear(ctx context.Context, characterID int64) error
}

type AddActionReq struct {
	CharacterID int64
	ActionID    string
	Times       int64
	Position    *int32
}

type MoveActionReq struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

type RemoveActionReq struct {
	CharacterID int64
	Position    int32
}
