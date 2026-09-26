package repo

import (
	"context"
	"time"
)

type OtpClient interface {
	BeginEnableTotp(ctx context.Context, req *BeginEnableTotpReq) (*BeginEnableTotpResp, error)
	ConfirmEnableTotp(ctx context.Context, req *ConfirmEnableTotpReq) error
	DisableTotp(ctx context.Context, req *DisableTotpReq) error
	GetCurrentTotp(ctx context.Context, userID int64) (*Totp, error)
	SendEmailOtp(ctx context.Context, req *SendEmailOtpReq) (*OtpCodeResp, error)
	SendPhoneOtp(ctx context.Context, req *SendPhoneOtpReq) (*OtpCodeResp, error)
}

type Totp struct {
	UserID     int64
	Enable     bool
	EnableTime *time.Time
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

type SendEmailOtpReq struct {
	UserID *int64
	Email  string
}

type SendPhoneOtpReq struct {
	UserID *int64
	Phone  string
}

type OtpCodeResp struct {
	Code string
}
