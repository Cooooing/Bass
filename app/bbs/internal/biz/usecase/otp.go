package usecase

import (
	"bbs/internal/biz/model"
	"bbs/internal/biz/repo"
	"context"
)

type OtpUsecase struct {
	otpClient repo.OtpClient
}

func NewOtpUsecase(
	otpClient repo.OtpClient,
) *OtpUsecase {
	return &OtpUsecase{
		otpClient: otpClient,
	}
}

type BeginEnableTotpReq struct {
	UserID      int64
	AccountName string
}

type BeginEnableTotpResp struct {
	URL    string
	QRCode []byte
}

func (u *OtpUsecase) BeginEnableTotp(ctx context.Context, req *BeginEnableTotpReq) (*BeginEnableTotpResp, error) {
	reply, err := u.otpClient.BeginEnableTotp(ctx, &repo.BeginEnableTotpReq{
		UserID:      req.UserID,
		AccountName: req.AccountName,
	})
	if err != nil {
		return nil, err
	}
	return &BeginEnableTotpResp{
		URL:    reply.URL,
		QRCode: reply.QRCode,
	}, nil
}

type ConfirmEnableTotpReq struct {
	UserID int64
	Code   string
}

func (u *OtpUsecase) ConfirmEnableTotp(ctx context.Context, req *ConfirmEnableTotpReq) error {
	return u.otpClient.ConfirmEnableTotp(ctx, &repo.ConfirmEnableTotpReq{
		UserID: req.UserID,
		Code:   req.Code,
	})
}

type DisableTotpReq struct {
	UserID int64
	Code   string
}

func (u *OtpUsecase) DisableTotp(ctx context.Context, req *DisableTotpReq) error {
	return u.otpClient.DisableTotp(ctx, &repo.DisableTotpReq{
		UserID: req.UserID,
		Code:   req.Code,
	})
}

func (u *OtpUsecase) GetCurrentTotp(ctx context.Context, userID int64) (*model.Totp, error) {
	reply, err := u.otpClient.GetCurrentTotp(ctx, userID)
	if err != nil {
		return nil, err
	}
	if reply == nil {
		return nil, nil
	}
	return &model.Totp{
		UserID:     reply.UserID,
		Enable:     reply.Enable,
		EnableTime: reply.EnableTime,
	}, nil
}

type SendOtpResp struct {
	Code string
}

type SendEmailOtpReq struct {
	UserID *int64
	Email  string
}

func (u *OtpUsecase) SendEmailOtp(ctx context.Context, req *SendEmailOtpReq) (*SendOtpResp, error) {
	reply, err := u.otpClient.SendEmailOtp(ctx, &repo.SendEmailOtpReq{
		UserID: req.UserID,
		Email:  req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &SendOtpResp{
		Code: reply.Code,
	}, nil
}

type SendPhoneOtpReq struct {
	UserID *int64
	Phone  string
}

func (u *OtpUsecase) SendPhoneOtp(ctx context.Context, req *SendPhoneOtpReq) (*SendOtpResp, error) {
	reply, err := u.otpClient.SendPhoneOtp(ctx, &repo.SendPhoneOtpReq{
		UserID: req.UserID,
		Phone:  req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &SendOtpResp{
		Code: reply.Code,
	}, nil
}
