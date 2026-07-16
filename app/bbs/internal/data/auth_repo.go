package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/server"
	"common/pkg/util"
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

func (r *AuthRepo) StartEmailRegistration(ctx context.Context, req *repo.StartEmailRegistrationReq) (*repo.StartEmailRegistrationResponse, error) {
	reply, err := r.userClient.Auth.StartEmailRegistration(ctx, &userv1.StartEmailRegistration_Request{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &repo.StartEmailRegistrationResponse{CodeToken: reply.GetCodeToken(), Code: reply.GetCode()}, nil
}

func (r *AuthRepo) VerifyEmailRegistration(ctx context.Context, req *repo.VerifyEmailRegistrationReq) (*repo.VerifyEmailRegistrationResponse, error) {
	_, err := r.userClient.Auth.VerifyEmailRegistration(ctx, &userv1.VerifyEmailRegistration_Request{
		Code:      req.Code,
		CodeToken: req.CodeToken,
	})
	if err != nil {
		return nil, err
	}
	return &repo.VerifyEmailRegistrationResponse{}, nil
}

func (r *AuthRepo) StartPhoneRegistration(ctx context.Context, req *repo.StartPhoneRegistrationReq) (*repo.StartPhoneRegistrationResponse, error) {
	reply, err := r.userClient.Auth.StartPhoneRegistration(ctx, &userv1.StartPhoneRegistration_Request{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &repo.StartPhoneRegistrationResponse{CodeToken: reply.GetCodeToken(), Code: reply.GetCode()}, nil
}

func (r *AuthRepo) VerifyPhoneRegistration(ctx context.Context, req *repo.VerifyPhoneRegistrationReq) (*repo.VerifyPhoneRegistrationResponse, error) {
	_, err := r.userClient.Auth.VerifyPhoneRegistration(ctx, &userv1.VerifyPhoneRegistration_Request{
		Code:      req.Code,
		CodeToken: req.CodeToken,
	})
	if err != nil {
		return nil, err
	}
	return &repo.VerifyPhoneRegistrationResponse{}, nil
}

func (r *AuthRepo) LoginByPassword(ctx context.Context, req *repo.LoginByPasswordReq) (*repo.LoginByPasswordResponse, error) {
	loginReq := &userv1.LoginByPassword_Request{
		Account:   req.Account,
		Password:  req.Password,
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
	return &repo.LoginByPasswordResponse{Token: reply.GetToken(), Account: out}, nil
}

func (r *AuthRepo) Logout(ctx context.Context, req *repo.LogoutReq) (*repo.LogoutResponse, error) {
	_, err := r.userClient.Auth.Logout(ctx, &userv1.Logout_Request{Token: req.Token})
	if err != nil {
		return nil, err
	}
	return &repo.LogoutResponse{}, nil
}
