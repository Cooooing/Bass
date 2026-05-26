package repo

import (
	"context"

	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	bizrepo "content/internal/biz/repo"
)

var _ bizrepo.UserClient = (*UserClient)(nil)

type UserClient struct {
	userClient *rpc.UserClient
}

func NewUserClient(userClient *rpc.UserClient) bizrepo.UserClient {
	return &UserClient{userClient: userClient}
}

func (c *UserClient) BatchGetBasicAccounts(ctx context.Context, userIDs []int64) (map[int64]*userv1.AccountBasic, error) {
	if len(userIDs) == 0 {
		return map[int64]*userv1.AccountBasic{}, nil
	}
	reply, err := c.userClient.Account.BatchGetBasic(ctx, &userv1.BatchGetBasicAccount_Request{UserIds: userIDs})
	if err != nil {
		return nil, err
	}
	if reply == nil || reply.Accounts == nil {
		return map[int64]*userv1.AccountBasic{}, nil
	}
	return reply.Accounts, nil
}
