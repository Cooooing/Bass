package data

import (
	"context"

	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"

	"bff_bbs/internal/biz/model"

	"github.com/samber/lo"
)

type accountProfileLoader struct {
	userClient *rpc.UserClient
}

func (r *accountProfileLoader) loadAccountProfiles(ctx context.Context, userIDs ...int64) (map[int64]*model.AccountProfile, error) {
	if r.userClient == nil {
		return map[int64]*model.AccountProfile{}, nil
	}
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool {
		return userID != 0
	})
	if len(ids) == 0 {
		return map[int64]*model.AccountProfile{}, nil
	}
	reply, err := r.userClient.Account.Map(ctx, &userv1.MapAccounts_Req{
		Query: &userv1.MapAccounts_Req_AccountQuery{UserIds: ids},
	})
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.AccountProfile, len(reply.GetAccounts()))
	for userID, account := range reply.GetAccounts() {
		basic := account.GetBasic()
		if basic == nil {
			continue
		}
		out[userID] = &model.AccountProfile{
			ID:            basic.GetId(),
			Name:          basic.GetName(),
			Nickname:      basic.Nickname,
			URL:           basic.Url,
			AvatarAssetID: basic.AvatarAssetId,
			Introduction:  basic.Introduction,
			MBTI:          basic.Mbti,
			Status:        int32(basic.GetStatus()),
			FollowCount:   basic.FollowCount,
			FollowerCount: basic.FollowerCount,
			CreatedAt:     new(basic.GetCreatedAt().AsTime()),
			UpdatedAt:     new(basic.GetUpdatedAt().AsTime()),
		}
	}
	return out, nil
}
