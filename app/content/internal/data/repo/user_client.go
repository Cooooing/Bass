package repo

import (
	"context"

	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	bizrepo "content/internal/biz/repo"
)

var _ bizrepo.UserClient = (*UserClient)(nil)

type UserClient struct {
	userClient *rpc.UserClient
}

func NewUserClient(userClient *rpc.UserClient) bizrepo.UserClient {
	return &UserClient{userClient: userClient}
}

func (c *UserClient) MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*userv1.AccountBasic, error) {
	if len(userIDs) == 0 {
		return map[int64]*userv1.AccountBasic{}, nil
	}
	reply, err := c.userClient.Account.Map(ctx, &userv1.MapAccounts_Request{
		Query: &userv1.AccountQuery{UserIds: userIDs},
	})
	if err != nil {
		return nil, err
	}
	accounts := reply.GetAccounts()
	if len(accounts) == 0 {
		return map[int64]*userv1.AccountBasic{}, nil
	}
	result := make(map[int64]*userv1.AccountBasic, len(accounts))
	for userID, account := range accounts {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		result[userID] = basic
	}
	return result, nil
}
