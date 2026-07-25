package repo

import (
	"bbs/internal/enum"
	"context"
	"time"
)

type AuthRepo interface {
	StartEmailRegistration(ctx context.Context, req *StartEmailRegistrationReq) (*StartEmailRegistrationResp, error)
	VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) error
	StartPhoneRegistration(ctx context.Context, req *StartPhoneRegistrationReq) (*StartPhoneRegistrationResp, error)
	VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) error
	StartEmailLogin(ctx context.Context, email string) (*StartEmailLoginResp, error)
	StartPhoneLogin(ctx context.Context, phone string) (*StartPhoneLoginResp, error)
	Login(ctx context.Context, req *LoginReq) (*LoginResp, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResp, error)
	Logout(ctx context.Context, accessToken string) error
	CancelAccount(ctx context.Context, req *CancelAccountReq) error
}

type StartEmailRegistrationReq struct {
	Email    string
	Password string
	Name     string
	Nickname *string
}
type StartEmailRegistrationResp struct{ Code string }
type VerifyEmailRegistrationReq struct {
	Email string
	Code  string
}
type StartPhoneRegistrationReq struct {
	Phone    string
	Password string
	Name     string
	Nickname *string
}
type StartPhoneRegistrationResp struct{ Code string }
type VerifyPhoneRegistrationReq struct {
	Phone string
	Code  string
}
type StartEmailLoginResp struct{ Code string }
type StartPhoneLoginResp struct{ Code string }

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
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
	SessionExpiresAt      time.Time
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
