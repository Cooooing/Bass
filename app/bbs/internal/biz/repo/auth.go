package repo

import (
	"bbs/internal/enum"
	"context"
	"time"
)

type AuthRepo interface {
	Register(ctx context.Context, req *RegisterReq) error
	Login(ctx context.Context, req *LoginReq) (*LoginResp, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResp, error)
	Logout(ctx context.Context, accessToken string) error
	CancelAccount(ctx context.Context, req *CancelAccountReq) error
}

type RegisterReq struct {
	Type     enum.RegisterType
	Name     string
	Password string
	Nickname *string
	Email    string
	Phone    string
	Code     string
}
type LoginReq struct {
	Type     enum.LoginType
	Account  string
	Password string
	Email    string
	Phone    string
	Code     string
}

type TokenResp struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
	SessionExpiresAt      *time.Time
}

type LoginResp struct {
	Token   TokenResp
	Account *Account
}
type CancelAccountReq struct {
	UserID   int64
	Password string
	Code     string
}
