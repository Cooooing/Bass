package repo

import "context"

type TotpClient interface {
	CheckEnableCodeTotp(ctx context.Context, req *CheckEnableCodeTotpReq) (*CheckEnableCodeTotpResponse, error)
	ValidateTotp(ctx context.Context, req *ValidateTotpReq) (*ValidateTotpResponse, error)
	BeginEnableTotp(ctx context.Context, req *BeginEnableTotpReq) (*BeginEnableTotpResponse, error)
	ConfirmEnableTotp(ctx context.Context, req *ConfirmEnableTotpReq) (*ConfirmEnableTotpResponse, error)
	DisableTotp(ctx context.Context, req *DisableTotpReq) (*DisableTotpResponse, error)
	GetCurrentTotp(ctx context.Context, req *GetCurrentTotpReq) (*GetCurrentTotpResponse, error)
}

type Totp struct {
	UserID     int64
	Enable     bool
	EnableTime string
}

type CheckEnableCodeTotpReq struct {
	UserID int64
	Code   string
}

type CheckEnableCodeTotpResponse struct {
	Verified bool
}

type ValidateTotpReq struct {
	UserID int64
	Code   string
}

type ValidateTotpResponse struct {
	Verified bool
}

type BeginEnableTotpReq struct {
	UserID      int64
	AccountName string
}

type BeginEnableTotpResponse struct {
	URL    string
	QRCode []byte
}

type ConfirmEnableTotpReq struct {
	UserID int64
	Code   string
}

type ConfirmEnableTotpResponse struct{}

type DisableTotpReq struct {
	UserID int64
	Code   string
}

type DisableTotpResponse struct{}

type GetCurrentTotpReq struct {
	UserID int64
}

type GetCurrentTotpResponse struct {
	Totp *Totp
}
