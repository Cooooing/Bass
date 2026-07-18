package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type SessionRepo interface {
	StartSession(ctx context.Context, req *StartSessionReq) (*model.Session, error)
	EndSession(ctx context.Context, id int64) (*model.Session, error)
	GetSession(ctx context.Context, id int64) (*model.Session, error)
	UpdateSessionWorld(ctx context.Context, req *UpdateSessionWorldReq) (*model.Session, error)
}

type StartSessionReq struct {
	PlayerID   *int64
	ClientType string
}

type UpdateSessionWorldReq struct {
	ID       int64
	PlayerID int64
	WorldID  int64
}
