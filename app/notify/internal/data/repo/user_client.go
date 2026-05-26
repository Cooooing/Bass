package repo

import (
	"common/api/gen/common"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"
	"notify/internal/biz/usecase"
)

var _ usecase.UserClient = (*UserClient)(nil)

type UserClient struct {
	userClient *rpc.UserClient
}

func NewUserClient(userClient *rpc.UserClient) usecase.UserClient {
	return &UserClient{
		userClient: userClient,
	}
}

func (c *UserClient) GetContacts(ctx context.Context, userIDs []int64) (map[int64]*usecase.UserContact, error) {
	reply, err := c.userClient.Account.BatchGetContact(ctx, &userv1.BatchGetContactAccount_Request{UserIds: userIDs})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*usecase.UserContact, len(reply.GetContacts()))
	for userID, contact := range reply.GetContacts() {
		if contact == nil {
			continue
		}
		result[userID] = &usecase.UserContact{
			Email: contact.GetEmail(),
			Phone: contact.GetPhone(),
		}
	}
	return result, nil
}

func (c *UserClient) ListFollowerIDs(ctx context.Context, userID int64) ([]int64, error) {
	page := uint32(1)
	size := uint32(500)
	userIDs := make([]int64, 0)
	for {
		reply, err := c.userClient.Relation.ListFollowerIds(ctx, &userv1.ListFollowerIdsRelation_Request{
			UserId: userID,
			Page:   &common.PageRequest{Page: page, Size: size},
		})
		if err != nil {
			return nil, err
		}
		userIDs = append(userIDs, reply.GetUserIds()...)
		pageReply := reply.GetPage()
		if pageReply == nil || len(reply.GetUserIds()) == 0 || pageReply.GetPage()*pageReply.GetSize() >= pageReply.GetTotal() {
			break
		}
		page++
	}
	return userIDs, nil
}
