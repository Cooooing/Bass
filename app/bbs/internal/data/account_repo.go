package data

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"common/api/gen/common"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.AccountRepo = (*AccountRepo)(nil)

type AccountRepo struct {
	userClient *rpc.UserClient
}

func NewAccountRepo(userClient *rpc.UserClient) repo.AccountRepo {
	return &AccountRepo{userClient: userClient}
}

func (r *AccountRepo) GetCurrentAccount(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Request{UserId: userID})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var out *bbsuserv1.Account
	if account != nil {
		out = &bbsuserv1.Account{}
		if basic := account.GetBasic(); basic != nil {
			out.Profile = &bbsuserv1.AccountProfile{
				Id:            basic.GetId(),
				Name:          basic.GetName(),
				Nickname:      basic.Nickname,
				Url:           basic.Url,
				AvatarUrl:     basic.AvatarUrl,
				Introduction:  basic.Introduction,
				Status:        bbsuserv1.AccountStatus(basic.GetStatus()),
				Mbti:          bbsuserv1.MBTI(basic.GetMbti()),
				FollowCount:   basic.FollowCount,
				FollowerCount: basic.FollowerCount,
				CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
				UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
			}
		}
		if contact := account.GetContact(); contact != nil {
			out.Contact = &bbsuserv1.AccountContact{
				UserId: contact.GetUserId(),
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &bbsuserv1.GetCurrentAccount_Reply{Account: out}, nil
}

func (r *AccountRepo) GetProfileAccount(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Request{UserId: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount().GetBasic()
	var profile *bbsuserv1.AccountProfile
	if account != nil {
		profile = &bbsuserv1.AccountProfile{
			Id:            account.GetId(),
			Name:          account.GetName(),
			Nickname:      account.Nickname,
			Url:           account.Url,
			AvatarUrl:     account.AvatarUrl,
			Introduction:  account.Introduction,
			Status:        bbsuserv1.AccountStatus(account.GetStatus()),
			Mbti:          bbsuserv1.MBTI(account.GetMbti()),
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
			CreatedAt:     formatProtoTime(account.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(account.GetUpdatedAt()),
		}
	}
	return &bbsuserv1.GetProfileAccount_Reply{Profile: profile}, nil
}

func (r *AccountRepo) UpdateProfileAccount(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	updateReq := &userv1.UpdateProfileAccount_Request{
		UserId:       userID,
		AvatarUrl:    req.AvatarUrl,
		Nickname:     req.Nickname,
		Url:          req.Url,
		Introduction: req.Introduction,
	}
	if req.Mbti != nil {
		updateReq.Mbti = new(userv1.MBTI(*req.Mbti))
	}
	reply, err := r.userClient.Account.UpdateProfile(ctx, updateReq)
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var profile *bbsuserv1.AccountProfile
	if account != nil {
		profile = &bbsuserv1.AccountProfile{
			Id:            account.GetId(),
			Name:          account.GetName(),
			Nickname:      account.Nickname,
			Url:           account.Url,
			AvatarUrl:     account.AvatarUrl,
			Introduction:  account.Introduction,
			Status:        bbsuserv1.AccountStatus(account.GetStatus()),
			Mbti:          bbsuserv1.MBTI(account.GetMbti()),
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
			CreatedAt:     formatProtoTime(account.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(account.GetUpdatedAt()),
		}
	}
	return &bbsuserv1.UpdateProfileAccount_Reply{Profile: profile}, nil
}

func (r *AccountRepo) AvatarAccount(ctx context.Context, req *bbsuserv1.AvatarAccount_Request) (*common.ImageReply, error) {
	return r.userClient.Account.Avatar(ctx, &userv1.AvatarAccount_Request{Name: req.GetName()})
}
