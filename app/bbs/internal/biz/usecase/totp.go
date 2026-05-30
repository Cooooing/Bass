package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type TotpUsecase struct {
	totpRepo repo.TotpRepo
}

func NewTotpUsecase(totpRepo repo.TotpRepo) *TotpUsecase {
	return &TotpUsecase{totpRepo: totpRepo}
}

func (u *TotpUsecase) BeginEnableTotp(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error) {
	return u.totpRepo.BeginEnableTotp(ctx, req)
}

func (u *TotpUsecase) ConfirmEnableTotp(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error) {
	return u.totpRepo.ConfirmEnableTotp(ctx, req)
}

func (u *TotpUsecase) DisableTotp(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error) {
	return u.totpRepo.DisableTotp(ctx, req)
}

func (u *TotpUsecase) GetCurrentTotp(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error) {
	return u.totpRepo.GetCurrentTotp(ctx, req)
}
