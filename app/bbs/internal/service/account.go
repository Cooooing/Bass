package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AccountService struct {
	bbsuserv1.UnimplementedAccountServiceServer
	userClient *rpc.UserClient
}

func NewAccountService(userClient *rpc.UserClient) *AccountService {
	return &AccountService{userClient: userClient}
}

func (s *AccountService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterAccountServiceServer(gs, s)
}

func (s *AccountService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAccountServiceHTTPServer(hs, s)
}

func (s *AccountService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	reply, err := s.userClient.Account.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentAccount_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentAccount_Reply{Account: toBFFAccount(reply.GetAccount())}, nil
}

func (s *AccountService) GetProfile(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	reply, err := s.userClient.Account.GetBasic(ctx, &userv1.GetBasicAccount_Request{UserId: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetProfileAccount_Reply{Profile: toBFFAccountProfile(reply.GetAccount())}, nil
}

func (s *AccountService) BatchGetProfile(ctx context.Context, req *bbsuserv1.BatchGetProfileAccount_Request) (*bbsuserv1.BatchGetProfileAccount_Reply, error) {
	reply, err := s.userClient.Account.BatchGetBasic(ctx, &userv1.BatchGetBasicAccount_Request{UserIds: req.GetUserIds()})
	if err != nil {
		return nil, err
	}
	profiles := make(map[int64]*bbsuserv1.AccountProfile, len(reply.GetAccounts()))
	for id, profile := range reply.GetAccounts() {
		profiles[id] = toBFFAccountProfile(profile)
	}
	return &bbsuserv1.BatchGetProfileAccount_Reply{Profiles: profiles}, nil
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	reply, err := s.userClient.Account.UpdateProfile(forwardAuth(ctx), &userv1.UpdateProfileAccount_Request{
		AvatarUrl:    req.AvatarUrl,
		Nickname:     req.Nickname,
		Url:          req.Url,
		Introduction: req.Introduction,
		Mbti:         req.Mbti,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateProfileAccount_Reply{Profile: toBFFAccountProfile(reply.GetAccount())}, nil
}

func toBFFAccount(in *userv1.Account) *bbsuserv1.Account {
	if in == nil {
		return nil
	}
	return &bbsuserv1.Account{
		Profile: toBFFAccountProfile(in.GetBasic()),
		Contact: toBFFAccountContact(in.GetContact()),
	}
}

func toBFFAccountProfile(in *userv1.AccountBasic) *bbsuserv1.AccountProfile {
	if in == nil {
		return nil
	}
	return &bbsuserv1.AccountProfile{
		Id:            in.GetId(),
		Name:          in.GetName(),
		Nickname:      in.Nickname,
		Url:           in.Url,
		AvatarUrl:     in.AvatarUrl,
		Introduction:  in.Introduction,
		Mbti:          in.Mbti,
		Status:        bbsuserv1.AccountStatus(in.GetStatus()),
		GroupName:     in.GetGroupName(),
		FollowCount:   in.FollowCount,
		FollowerCount: in.FollowerCount,
		BlockCount:    in.BlockCount,
		BlockedCount:  in.BlockedCount,
		CreatedAt:     formatProtoTime(in.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(in.GetUpdatedAt()),
	}
}

func toBFFAccountContact(in *userv1.AccountContact) *bbsuserv1.AccountContact {
	if in == nil {
		return nil
	}
	return &bbsuserv1.AccountContact{
		UserId: in.GetUserId(),
		Email:  in.Email,
		Phone:  in.Phone,
	}
}
