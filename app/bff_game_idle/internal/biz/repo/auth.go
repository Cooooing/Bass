package repo

import (
	"context"

	"bff_game_idle/internal/biz/model"
)

type AuthRepo interface {
	Register(ctx context.Context, req *RegisterReq) error
	Login(ctx context.Context, req *LoginReq) (*model.LoginToken, error)
}

type RegisterReq struct {
	Password string
	Email    string
}

type LoginReq struct {
	Email    string
	Password string
}
