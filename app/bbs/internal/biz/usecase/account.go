package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"common/api/gen/common"
	"context"
)

type AccountUsecase struct {
	accountRepo repo.AccountRepo
}

func NewAccountUsecase(accountRepo repo.AccountRepo) *AccountUsecase {
	return &AccountUsecase{accountRepo: accountRepo}
}

func (u *AccountUsecase) GetCurrentAccount(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	return u.accountRepo.GetCurrentAccount(ctx, req)
}

func (u *AccountUsecase) GetProfileAccount(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	return u.accountRepo.GetProfileAccount(ctx, req)
}

func (u *AccountUsecase) UpdateProfileAccount(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	return u.accountRepo.UpdateProfileAccount(ctx, req)
}

func (u *AccountUsecase) AvatarAccount(ctx context.Context, req *bbsuserv1.AvatarAccount_Request) (*common.ImageReply, error) {
	return u.accountRepo.AvatarAccount(ctx, req)
}
