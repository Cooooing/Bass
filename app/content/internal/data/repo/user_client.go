package repo

import (
	"context"

	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"content/internal/biz/model"
	bizrepo "content/internal/biz/repo"
)

var _ bizrepo.UserClient = (*UserClient)(nil)

type UserClient struct {
	userClient *rpc.UserClient
}

func NewUserClient(
	userClient *rpc.UserClient,
) bizrepo.UserClient {
	return &UserClient{
		userClient: userClient,
	}
}

func (c *UserClient) MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.UserAccountBasic, error) {
	if len(userIDs) == 0 {
		return map[int64]*model.UserAccountBasic{}, nil
	}
	reply, err := c.userClient.Account.Map(ctx, &userv1.MapAccounts_Req{
		Query: &userv1.MapAccounts_Req_AccountQuery{
			UserIds: userIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	accounts := reply.GetAccounts()
	if len(accounts) == 0 {
		return map[int64]*model.UserAccountBasic{}, nil
	}
	result := make(map[int64]*model.UserAccountBasic, len(accounts))
	for userID, account := range accounts {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		result[userID] = &model.UserAccountBasic{
			UserID:        userID,
			Username:      basic.GetName(),
			Nickname:      basic.GetNickname(),
			AvatarAssetID: basic.AvatarAssetId,
		}
	}
	return result, nil
}
