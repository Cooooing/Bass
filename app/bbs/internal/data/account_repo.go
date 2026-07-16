package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.AccountClient = (*AccountClient)(nil)

type AccountClient struct {
	userClient *rpc.UserClient
}

func NewAccountClient(userClient *rpc.UserClient) repo.AccountClient {
	return &AccountClient{userClient: userClient}
}

func (r *AccountClient) GetCurrentAccount(ctx context.Context, req *repo.GetCurrentAccountReq) (*repo.GetCurrentAccountResponse, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Request{UserId: req.UserID})
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
				CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
				UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
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
	return &repo.GetCurrentAccountResponse{Account: out}, nil
}

func (r *AccountClient) GetProfileAccount(ctx context.Context, req *repo.GetProfileAccountReq) (*repo.GetProfileAccountResponse, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Request{UserId: req.UserID})
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
			CreatedAt:     formatProtoTime(account.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(account.GetUpdatedAt()),
		}
	}
	return &repo.GetProfileAccountResponse{Profile: profile}, nil
}

func (r *AccountClient) UpdateProfileAccount(ctx context.Context, req *repo.UpdateProfileAccountReq) (*repo.UpdateProfileAccountResponse, error) {
	updateReq := &userv1.UpdateProfileAccount_Request{
		UserId:       req.UserID,
		AvatarUrl:    req.AvatarURL,
		Nickname:     req.Nickname,
		Url:          req.URL,
		Introduction: req.Introduction,
	}
	if req.MBTI != nil {
		mbti := userv1.MBTI(*req.MBTI)
		updateReq.Mbti = &mbti
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
			CreatedAt:     formatProtoTime(account.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(account.GetUpdatedAt()),
		}
	}
	return &repo.UpdateProfileAccountResponse{Profile: profile}, nil
}

func (r *AccountClient) AvatarAccount(ctx context.Context, req *repo.AvatarAccountReq) (*repo.AvatarAccountResponse, error) {
	reply, err := r.userClient.Account.Avatar(ctx, &userv1.AvatarAccount_Request{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &repo.AvatarAccountResponse{Data: reply.GetData(), ContentType: reply.GetContentType()}, nil
}
