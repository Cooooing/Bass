package repo

import "context"

type TotpClient interface {
	CheckEnableCodeTotp(ctx context.Context, req *CheckEnableCodeTotpReq) (bool, error)
	ValidateTotp(ctx context.Context, req *ValidateTotpReq) (bool, error)
	BeginEnableTotp(ctx context.Context, req *BeginEnableTotpReq) (*BeginEnableTotpResp, error)
	ConfirmEnableTotp(ctx context.Context, req *ConfirmEnableTotpReq) error
	DisableTotp(ctx context.Context, req *DisableTotpReq) error
	GetCurrentTotp(ctx context.Context, userID int64) (*Totp, error)
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

type ValidateTotpReq struct {
	UserID int64
	Code   string
}

type BeginEnableTotpReq struct {
	UserID      int64
	AccountName string
}

type BeginEnableTotpResp struct {
	URL    string
	QRCode []byte
}

type ConfirmEnableTotpReq struct {
	UserID int64
	Code   string
}

type DisableTotpReq struct {
	UserID int64
	Code   string
}
