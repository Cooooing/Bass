package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type SessionRepo interface {
	StartSession(ctx context.Context, req *StartSessionReq) (*StartSessionResponse, error)
	EndSession(ctx context.Context, req *EndSessionReq) (*EndSessionResponse, error)
	GetSession(ctx context.Context, req *GetSessionReq) (*GetSessionResponse, error)
	UpdateSessionWorld(ctx context.Context, req *UpdateSessionWorldReq) (*UpdateSessionWorldResponse, error)
}

type StartSessionReq struct {
	PlayerID   *int64
	ClientType string
}

type StartSessionResponse struct {
	Row *model.Session
}

type EndSessionReq struct {
	ID int64
}

type EndSessionResponse struct {
	Row *model.Session
}

type GetSessionReq struct {
	ID int64
}

type GetSessionResponse struct {
	Row *model.Session
}

type UpdateSessionWorldReq struct {
	ID       int64
	PlayerID int64
	WorldID  int64
}

type UpdateSessionWorldResponse struct {
	Row *model.Session
}
