package repo

import (
	"context"

	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"

	"im/internal/biz/repo"
)

type RelationClient struct{ user *rpc.UserClient }

func NewRelationClient(user *rpc.UserClient) repo.RelationClient { return &RelationClient{user: user} }

func (r *RelationClient) BlockedCounterparts(ctx context.Context, userID int64) ([]int64, error) {
	response, err := r.user.Relation.ListBlockedCounterparts(ctx, &userv1.ListBlockedCounterparts_Req{UserId: userID})
	if err != nil {
		return nil, err
	}
	return response.GetUserIds(), nil
}
