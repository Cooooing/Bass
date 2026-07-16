package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type TotpUsecase struct {
	totpClient repo.TotpClient
}

func NewTotpUsecase(totpClient repo.TotpClient) *TotpUsecase {
	return &TotpUsecase{totpClient: totpClient}
}

type BeginEnableTotpReq struct {
	UserID      int64
	AccountName string
}

type BeginEnableTotpResponse struct {
	URL    string
	QRCode []byte
}

func (u *TotpUsecase) BeginEnableTotp(ctx context.Context, req *BeginEnableTotpReq) (*BeginEnableTotpResponse, error) {
	current, err := u.totpClient.GetCurrentTotp(ctx, &repo.GetCurrentTotpReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	if current.Totp != nil && current.Totp.Enable {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED)
	}
	reply, err := u.totpClient.BeginEnableTotp(ctx, &repo.BeginEnableTotpReq{UserID: req.UserID, AccountName: req.AccountName})
	if err != nil {
		return nil, err
	}
	return &BeginEnableTotpResponse{URL: reply.URL, QRCode: reply.QRCode}, nil
}

type ConfirmEnableTotpReq struct {
	UserID int64
	Code   string
}

func (u *TotpUsecase) ConfirmEnableTotp(ctx context.Context, req *ConfirmEnableTotpReq) error {
	current, err := u.totpClient.GetCurrentTotp(ctx, &repo.GetCurrentTotpReq{UserID: req.UserID})
	if err != nil {
		return err
	}
	if current.Totp != nil && current.Totp.Enable {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED)
	}
	checkReq := &repo.CheckEnableCodeTotpReq{UserID: req.UserID, Code: req.Code}
	verified, err := u.totpClient.CheckEnableCodeTotp(ctx, checkReq)
	if err != nil {
		return err
	}
	if !verified.Verified {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}
	_, err = u.totpClient.ConfirmEnableTotp(ctx, &repo.ConfirmEnableTotpReq{UserID: req.UserID, Code: req.Code})
	return err
}

type DisableTotpReq struct {
	UserID int64
	Code   string
}

func (u *TotpUsecase) DisableTotp(ctx context.Context, req *DisableTotpReq) error {
	current, err := u.totpClient.GetCurrentTotp(ctx, &repo.GetCurrentTotpReq{UserID: req.UserID})
	if err != nil {
		return err
	}
	if current.Totp == nil || !current.Totp.Enable {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_DISABLED)
	}
	disableReq := &repo.ValidateTotpReq{UserID: req.UserID, Code: req.Code}
	verified, err := u.totpClient.ValidateTotp(ctx, disableReq)
	if err != nil {
		return err
	}
	if !verified.Verified {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}
	_, err = u.totpClient.DisableTotp(ctx, &repo.DisableTotpReq{UserID: req.UserID, Code: req.Code})
	return err
}

type GetCurrentTotpReq struct {
	UserID int64
}

type GetCurrentTotpResponse struct {
	Totp *bbsuserv1.GetCurrentTotp_Response_Totp
}

func (u *TotpUsecase) GetCurrentTotp(ctx context.Context, req *GetCurrentTotpReq) (*GetCurrentTotpResponse, error) {
	reply, err := u.totpClient.GetCurrentTotp(ctx, &repo.GetCurrentTotpReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var totp *bbsuserv1.GetCurrentTotp_Response_Totp
	if row := reply.Totp; row != nil {
		totp = &bbsuserv1.GetCurrentTotp_Response_Totp{UserId: row.UserID, Enable: row.Enable, EnableTime: row.EnableTime}
	}
	return &GetCurrentTotpResponse{Totp: totp}, nil
}
