package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

type CharacterRepo interface {
	Create(ctx context.Context, req *CreateCharacterReq) (*model.Character, error)
	List(ctx context.Context, req *ListCharacterReq) ([]*model.Character, error)
	Online(ctx context.Context, req *OnlineCharacterReq) (*model.CharacterOnlineSession, error)
	Ping(ctx context.Context, req *PingCharacterReq) (*model.CharacterOnlineSession, error)
	Offline(ctx context.Context, req *OfflineCharacterReq) error
}

type CreateCharacterReq struct {
	UserID int64
	Name   string
}

type ListCharacterReq struct {
	UserID      int64
	CharacterID int64
}

type OnlineCharacterReq struct {
	CharacterID int64
}

type PingCharacterReq struct {
	CharacterID int64
	SessionID   string
}

type OfflineCharacterReq struct {
	CharacterID int64
	SessionID   string
	Timeout     bool
}
