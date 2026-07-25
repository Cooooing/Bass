package repo

import (
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
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

func (c *UserClient) MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.
	UserAccount, error) {
	if len(userIDs) == 0 {
		return map[int64]*model.UserAccount{}, nil
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
	result := make(map[int64]*model.UserAccount, len(accounts))
	for userID, account := range accounts {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		item := &model.UserAccount{
			ID:       userID,
			Name:     basic.GetName(),
			Nickname: basic.GetNickname(),
		}
		if contact := account.GetContact(); contact != nil {
			item.Email = contact.GetEmail()
			item.Phone = contact.GetPhone()
		}
		result[item.ID] = item
	}
	return result, nil
}

func (c *UserClient) ListFollowerIDs(ctx context.Context, userID int64) ([]int64, error) {
	if userID == 0 {
		return []int64{}, nil
	}
	page := uint32(1)
	size := uint32(500)
	userIDs := make([]int64, 0)
	for {
		reply, err := c.userClient.Relation.ListFollowers(ctx, &userv1.ListFollowersRelations_Req{
			UserId: userID,
			Page: &common.PageReq{
				Page: page,
				Size: size,
			},
		})
		if err != nil {
			return nil, err
		}
		rows := reply.GetRows()
		for _, row := range rows {
			userIDs = append(userIDs, row.GetActorId())
		}
		pageResp := reply.GetPage()
		if pageResp == nil || len(rows) == 0 || pageResp.GetPage()*pageResp.GetSize() >= pageResp.GetTotal() {
			break
		}
		page++
	}
	return userIDs, nil
}
