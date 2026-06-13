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

func (u *TotpUsecase) BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error) {
	current, err := u.totpClient.GetCurrentTotp(ctx, &bbsuserv1.GetCurrentTotp_Request{})
	if err != nil {
		return nil, err
	}
	if current.GetTotp().GetEnable() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED)
	}
	return u.totpClient.BeginEnableTotp(ctx, req)
}

func (u *TotpUsecase) ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error) {
	current, err := u.totpClient.GetCurrentTotp(ctx, &bbsuserv1.GetCurrentTotp_Request{})
	if err != nil {
		return nil, err
	}
	if current.GetTotp().GetEnable() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_ENABLED)
	}
	verified, err := u.totpClient.CheckEnableCodeTotp(ctx, req)
	if err != nil {
		return nil, err
	}
	if !verified {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}
	return u.totpClient.ConfirmEnableTotp(ctx, req)
}

func (u *TotpUsecase) DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error) {
	current, err := u.totpClient.GetCurrentTotp(ctx, &bbsuserv1.GetCurrentTotp_Request{})
	if err != nil {
		return nil, err
	}
	if !current.GetTotp().GetEnable() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_ALREADY_DISABLED)
	}
	verified, err := u.totpClient.ValidateTotp(ctx, req)
	if err != nil {
		return nil, err
	}
	if !verified {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOTP_CODE_INVALID)
	}
	return u.totpClient.DisableTotp(ctx, req)
}

func (u *TotpUsecase) GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error) {
	return u.totpClient.GetCurrentTotp(ctx, req)
}
