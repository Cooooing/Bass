package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	userv1enum "common/proto/gen/user/v1/enum"
	"context"
)

var _ repo.AccountClient = (*AccountClient)(nil)

type AccountClient struct {
	userClient *rpc.UserClient
}

func NewAccountClient(
	userClient *rpc.UserClient,
) repo.AccountClient {
	return &AccountClient{
		userClient: userClient,
	}
}

func (r *AccountClient) GetCurrentAccount(ctx context.Context, userID int64) (*repo.Account, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var out *repo.Account
	if account != nil {
		out = &repo.Account{}
		if basic := account.GetBasic(); basic != nil {
			out.Profile = &repo.AccountProfile{
				ID:            basic.GetId(),
				Name:          basic.GetName(),
				Nickname:      basic.Nickname,
				URL:           basic.Url,
				AvatarURL:     basic.AvatarUrl,
				Introduction:  basic.Introduction,
				Status:        int32(basic.GetStatus()),
				MBTI:          int32(basic.GetMbti()),
				FollowCount:   basic.FollowCount,
				FollowerCount: basic.FollowerCount,
				CreatedAt:     new(basic.GetCreatedAt().AsTime()),
				UpdatedAt:     new(basic.GetUpdatedAt().AsTime()),
			}
		}
		if contact := account.GetContact(); contact != nil {
			out.Contact = &repo.AccountContact{
				UserID: contact.GetUserId(),
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return out, nil
}

func (r *AccountClient) GetProfileAccount(ctx context.Context, userID int64) (*repo.AccountProfile, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount().GetBasic()
	var profile *repo.AccountProfile
	if account != nil {
		profile = &repo.AccountProfile{
			ID:            account.GetId(),
			Name:          account.GetName(),
			Nickname:      account.Nickname,
			URL:           account.Url,
			AvatarURL:     account.AvatarUrl,
			Introduction:  account.Introduction,
			Status:        int32(account.GetStatus()),
			MBTI:          int32(account.GetMbti()),
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
			CreatedAt:     new(account.GetCreatedAt().AsTime()),
			UpdatedAt:     new(account.GetUpdatedAt().AsTime()),
		}
	}
	return profile, nil
}

func (r *AccountClient) UpdateProfileAccount(ctx context.Context, req *repo.UpdateProfileAccountReq) (*repo.AccountProfile, error) {
	updateReq := &userv1.UpdateProfileAccount_Req{
		UserId:       req.UserID,
		AvatarUrl:    req.AvatarURL,
		Nickname:     req.Nickname,
		Url:          req.URL,
		Introduction: req.Introduction,
	}
	if req.MBTI != nil {
		updateReq.Mbti = new(userv1enum.MBTI(*req.MBTI))
	}
	reply, err := r.userClient.Account.UpdateProfile(ctx, updateReq)
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var profile *repo.AccountProfile
	if account != nil {
		profile = &repo.AccountProfile{
			ID:            account.GetId(),
			Name:          account.GetName(),
			Nickname:      account.Nickname,
			URL:           account.Url,
			AvatarURL:     account.AvatarUrl,
			Introduction:  account.Introduction,
			Status:        int32(account.GetStatus()),
			MBTI:          int32(account.GetMbti()),
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
			CreatedAt:     new(account.GetCreatedAt().AsTime()),
			UpdatedAt:     new(account.GetUpdatedAt().AsTime()),
		}
	}
	return profile, nil
}

func (r *AccountClient) AvatarAccount(ctx context.Context, name string) (*repo.AvatarAccountResp, error) {
	reply, err := r.userClient.Account.Avatar(ctx, &userv1.AvatarAccount_Req{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return &repo.AvatarAccountResp{
		Data:        reply.GetData(),
		ContentType: reply.GetContentType(),
	}, nil
}
