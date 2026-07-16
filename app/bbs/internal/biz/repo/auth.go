package repo

import "context"

type AuthRepo interface {
	StartEmailRegistration(ctx context.Context, req *StartEmailRegistrationReq) (*StartEmailRegistrationResponse, error)
	VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) (*VerifyEmailRegistrationResponse, error)
	StartPhoneRegistration(ctx context.Context, req *StartPhoneRegistrationReq) (*StartPhoneRegistrationResponse, error)
	VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) (*VerifyPhoneRegistrationResponse, error)
	LoginByPassword(ctx context.Context, req *LoginByPasswordReq) (*LoginByPasswordResponse, error)
	Logout(ctx context.Context, req *LogoutReq) (*LogoutResponse, error)
}

type StartEmailRegistrationReq struct {
	Email    string
	Password string
	Name     string
	Nickname *string
}

type StartEmailRegistrationResponse struct {
	CodeToken string
	Code      string
}

type VerifyEmailRegistrationReq struct {
	Code      string
	CodeToken string
}

type VerifyEmailRegistrationResponse struct{}

type StartPhoneRegistrationReq struct {
	Phone    string
	Password string
	Name     string
	Nickname *string
}

type StartPhoneRegistrationResponse struct {
	CodeToken string
	Code      string
}

type VerifyPhoneRegistrationReq struct {
	Code      string
	CodeToken string
}

type VerifyPhoneRegistrationResponse struct{}

type LoginByPasswordReq struct {
	Account  string
	Password string
}

type LoginByPasswordResponse struct {
	Token   string
	Account *Account
}

type LogoutReq struct {
	Token string
}

type LogoutResponse struct{}
