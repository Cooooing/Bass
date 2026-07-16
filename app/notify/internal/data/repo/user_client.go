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

func NewUserClient(userClient *rpc.UserClient) bizrepo.UserClient {
	return &UserClient{
		userClient: userClient,
	}
}

func (c *UserClient) MapAccounts(ctx context.Context, req *bizrepo.UserMapAccountsReq) (*bizrepo.UserMapAccountsResponse, error) {
	if req == nil || len(req.UserIDs) == 0 {
		return &bizrepo.UserMapAccountsResponse{Rows: map[int64]*model.UserAccount{}}, nil
	}
	reply, err := c.userClient.Account.Map(ctx, &userv1.MapAccounts_Request{
		Query: &userv1.MapAccounts_Request_AccountQuery{UserIds: req.UserIDs},
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
	return &bizrepo.UserMapAccountsResponse{Rows: result}, nil
}

func (c *UserClient) ListFollowerIDs(ctx context.Context, req *bizrepo.UserListFollowerIDsReq) (*bizrepo.UserListFollowerIDsResponse, error) {
	if req == nil || req.UserID == 0 {
		return &bizrepo.UserListFollowerIDsResponse{}, nil
	}
	page := uint32(1)
	size := uint32(500)
	userIDs := make([]int64, 0)
	for {
		reply, err := c.userClient.Relation.ListFollowers(ctx, &userv1.ListFollowersRelations_Request{
			UserId: req.UserID,
			Page:   &common.PageRequest{Page: page, Size: size},
		})
		if err != nil {
			return nil, err
		}
		rows := reply.GetRows()
		for _, row := range rows {
			userIDs = append(userIDs, row.GetActorId())
		}
		pageResponse := reply.GetPage()
		if pageResponse == nil || len(rows) == 0 || pageResponse.GetPage()*pageResponse.GetSize() >= pageResponse.GetTotal() {
			break
		}
		page++
	}
	return &bizrepo.UserListFollowerIDsResponse{UserIDs: userIDs}, nil
}
