package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"common/proto/gen/common"
	"context"
)

type AccountUsecase struct {
	accountClient repo.AccountClient
}

func NewAccountUsecase(accountClient repo.AccountClient) *AccountUsecase {
	return &AccountUsecase{accountClient: accountClient}
}

func (u *AccountUsecase) GetCurrentAccount(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	return u.accountClient.GetCurrentAccount(ctx, req)
}

func (u *AccountUsecase) GetProfileAccount(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	return u.accountClient.GetProfileAccount(ctx, req)
}

func (u *AccountUsecase) UpdateProfileAccount(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	return u.accountClient.UpdateProfileAccount(ctx, req)
}

func (u *AccountUsecase) AvatarAccount(ctx context.Context, req *bbsuserv1.AvatarAccount_Request) (*common.ImageReply, error) {
	return u.accountClient.AvatarAccount(ctx, req)
}
