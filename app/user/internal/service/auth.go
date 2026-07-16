package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/config"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	conf        *config.Bootstrap
	authUsecase *usecase.AuthUsecase
}

func NewAuthService(conf *config.Bootstrap, authUsecase *usecase.AuthUsecase) *AuthService {
	return &AuthService{
		conf:        conf,
		authUsecase: authUsecase,
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *v1.StartEmailRegistration_Request) (*v1.StartEmailRegistration_Response, error) {
	res, err := s.authUsecase.StartEmailRegistration(ctx, &usecase.StartEmailRegistrationReq{Account: &model.Account{
		Email:    &req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	}})
	reply := &v1.StartEmailRegistration_Response{}
	if res != nil {
		reply.CodeToken = res.Token
	}
	if err == nil && res != nil && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = res.Code
	}
	return reply, err
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *v1.VerifyEmailRegistration_Request) (*v1.VerifyEmailRegistration_Response, error) {
	err := s.authUsecase.VerifyEmailRegistration(ctx, &usecase.VerifyEmailRegistrationReq{CodeToken: req.CodeToken, Code: req.Code})
	return &v1.VerifyEmailRegistration_Response{}, err
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *v1.StartPhoneRegistration_Request) (*v1.StartPhoneRegistration_Response, error) {
	res, err := s.authUsecase.StartPhoneRegistration(ctx, &usecase.StartPhoneRegistrationReq{Account: &model.Account{
		Phone:    &req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	}})
	reply := &v1.StartPhoneRegistration_Response{}
	if res != nil {
		reply.CodeToken = res.Token
	}
	if err == nil && res != nil && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = res.Code
	}
	return reply, err
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *v1.VerifyPhoneRegistration_Request) (*v1.VerifyPhoneRegistration_Response, error) {
	err := s.authUsecase.VerifyPhoneRegistration(ctx, &usecase.VerifyPhoneRegistrationReq{CodeToken: req.CodeToken, Code: req.Code})
	return &v1.VerifyPhoneRegistration_Response{}, err
}

func (s *AuthService) LoginByPassword(ctx context.Context, req *v1.LoginByPassword_Request) (*v1.LoginByPassword_Response, error) {
	res, err := s.authUsecase.LoginByPassword(ctx, &usecase.LoginByPasswordReq{
		Account:  req.Account,
		Password: req.Password,
		LoginContext: &model.LoginContext{
			IP:          req.GetIp(),
			Country:     req.GetCountry(),
			CountryCode: req.GetCountryCode(),
			Province:    req.GetProvince(),
			City:        req.GetCity(),
			ISP:         req.GetIsp(),
			UserAgent:   req.GetUserAgent(),
			DeviceID:    req.GetDeviceId(),
			Platform:    req.GetPlatform(),
			RequestID:   req.GetRequestId(),
		},
	})
	if err != nil {
		if cerrors.IsInternalServerError(err) {
			return nil, err
		}
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS).WithCause(err)
	}
	account := res.Account
	basic := &v1.LoginByPassword_Response_AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarUrl:     account.AvatarURL,
		Introduction:  account.Introduction,
		FollowCount:   account.FollowCount,
		FollowerCount: account.FollowerCount,
	}
	if account.Mbti != nil {
		basic.Mbti = enum.MBTIMap.MustToProto(*account.Mbti)
	}
	if account.Status != nil {
		basic.Status = enum.AccountStatusMap.MustToProto(*account.Status)
	}
	if account.CreatedAt != nil {
		basic.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		basic.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return &v1.LoginByPassword_Response{
		Token: res.Token,
		Account: &v1.LoginByPassword_Response_Account{
			Basic: basic,
			Contact: &v1.LoginByPassword_Response_AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		},
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *v1.Logout_Request) (*v1.Logout_Response, error) {
	err := s.authUsecase.Logout(ctx, &usecase.LogoutReq{Token: req.GetToken()})
	return &v1.Logout_Response{}, err
}

func (s *AuthService) ParseToken(ctx context.Context, req *v1.ParseToken_Request) (*v1.ParseToken_Response, error) {
	res, err := s.authUsecase.ParseToken(ctx, &usecase.ParseTokenReq{Token: req.GetToken()})
	if err != nil {
		return nil, err
	}
	user := res.User
	return &v1.ParseToken_Response{
		User: &v1.ParseToken_Response_TokenUser{
			Id:       user.ID,
			Name:     user.Name,
			Nickname: user.Nickname,
			Language: user.Language,
			Timezone: user.Timezone,
		},
	}, nil
}
