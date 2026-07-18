package repo

import "context"

type AuthRepo interface {
	StartEmailRegistration(ctx context.Context, req *StartEmailRegistrationReq) (*StartEmailRegistrationResp, error)
	VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) error
	StartPhoneRegistration(ctx context.Context, req *StartPhoneRegistrationReq) (*StartPhoneRegistrationResp, error)
	VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) error
	LoginByPassword(ctx context.Context, req *LoginByPasswordReq) (*LoginByPasswordResp, error)
	Logout(ctx context.Context, token string) error
}

type StartEmailRegistrationReq struct {
	Email    string
	Password string
	Name     string
	Nickname *string
}

type StartEmailRegistrationResp struct {
	CodeToken string
	Code      string
}

type VerifyEmailRegistrationReq struct {
	Code      string
	CodeToken string
}

type StartPhoneRegistrationReq struct {
	Phone    string
	Password string
	Name     string
	Nickname *string
}

type StartPhoneRegistrationResp struct {
	CodeToken string
	Code      string
}

type VerifyPhoneRegistrationReq struct {
	Code      string
	CodeToken string
}

type LoginByPasswordReq struct {
	Account  string
	Password string
}

type LoginByPasswordResp struct {
	Token   string
	Account *Account
}
