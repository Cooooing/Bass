package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/server"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.AuthRepo = (*AuthRepo)(nil)

type AuthRepo struct {
	userClient *rpc.UserClient
}

func NewAuthRepo(userClient *rpc.UserClient) repo.AuthRepo {
	return &AuthRepo{userClient: userClient}
}

func NewAuthClient(userClient *rpc.UserClient) repo.AuthRepo {
	return NewAuthRepo(userClient)
}

func (r *AuthRepo) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Reply, error) {
	reply, err := r.userClient.Auth.StartEmailRegistration(ctx, &userv1.StartEmailRegistration_Request{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.StartEmailRegistration_Reply{CodeToken: reply.GetCodeToken(), Code: reply.GetCode()}, nil
}

func (r *AuthRepo) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Reply, error) {
	_, err := r.userClient.Auth.VerifyEmailRegistration(ctx, &userv1.VerifyEmailRegistration_Request{
		Code:      req.GetCode(),
		CodeToken: req.GetCodeToken(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyEmailRegistration_Reply{}, nil
}

func (r *AuthRepo) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Reply, error) {
	reply, err := r.userClient.Auth.StartPhoneRegistration(ctx, &userv1.StartPhoneRegistration_Request{
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.StartPhoneRegistration_Reply{CodeToken: reply.GetCodeToken(), Code: reply.GetCode()}, nil
}

func (r *AuthRepo) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Reply, error) {
	_, err := r.userClient.Auth.VerifyPhoneRegistration(ctx, &userv1.VerifyPhoneRegistration_Request{
		Code:      req.GetCode(),
		CodeToken: req.GetCodeToken(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyPhoneRegistration_Reply{}, nil
}

func (r *AuthRepo) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Reply, error) {
	loginReq := &userv1.LoginByPassword_Request{
		Account:   req.GetAccount(),
		Password:  req.GetPassword(),
		UserAgent: server.GetHeader(ctx, constant.HeaderUserAgent),
		DeviceId:  server.GetHeader(ctx, constant.HeaderDeviceID),
		Platform:  server.GetHeader(ctx, constant.HeaderPlatform),
		RequestId: server.GetHeader(ctx, constant.HeaderRequestID),
	}
	if loginReq.RequestId == "" {
		loginReq.RequestId = server.GetHeader(ctx, constant.HeaderTraceID)
	}
	if ipInfo, ok := util.GetContextValue[*commonmodel.IpInfo](ctx, constant.CtxIpInfo); ok && ipInfo != nil {
		loginReq.Ip = ipInfo.Ip
		loginReq.Country = ipInfo.Country
		loginReq.CountryCode = ipInfo.CountryCode
		loginReq.Province = ipInfo.Province
		loginReq.City = ipInfo.City
		loginReq.Isp = ipInfo.ISP
	}
	if loginReq.Ip == "" {
		loginReq.Ip = server.ClientIP(ctx)
	}
	reply, err := r.userClient.Auth.LoginByPassword(ctx, loginReq)
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
	return &bbsuserv1.LoginByPassword_Reply{Token: reply.GetToken(), Account: out}, nil
}

func (r *AuthRepo) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	token, err := currentToken(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Auth.Logout(ctx, &userv1.Logout_Request{Token: token})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.Logout_Reply{}, nil
}
