package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/conf"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	conf        *conf.Bootstrap
	authUsecase *usecase.AuthUsecase
}

func NewAuthService(conf *conf.Bootstrap, authUsecase *usecase.AuthUsecase) *AuthService {
	return &AuthService{
		conf:        conf,
		authUsecase: authUsecase,
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *v1.StartEmailRegistration_Request) (*v1.StartEmailRegistration_Reply, error) {
	code, token, err := s.authUsecase.StartEmailRegistration(ctx, &model.Account{
		Email:    &req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	reply := &v1.StartEmailRegistration_Reply{CodeToken: token}
	if err == nil && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = code
	}
	return reply, err
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *v1.VerifyEmailRegistration_Request) (*v1.VerifyEmailRegistration_Reply, error) {
	err := s.authUsecase.VerifyEmailRegistration(ctx, req.CodeToken, req.Code)
	return &v1.VerifyEmailRegistration_Reply{}, err
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *v1.StartPhoneRegistration_Request) (*v1.StartPhoneRegistration_Reply, error) {
	code, token, err := s.authUsecase.StartPhoneRegistration(ctx, &model.Account{
		Phone:    &req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	reply := &v1.StartPhoneRegistration_Reply{CodeToken: token}
	if err == nil && s.conf.GetServer().GetMode() != constant.Prod {
		reply.Code = code
	}
	return reply, err
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *v1.VerifyPhoneRegistration_Request) (*v1.VerifyPhoneRegistration_Reply, error) {
	err := s.authUsecase.VerifyPhoneRegistration(ctx, req.CodeToken, req.Code)
	return &v1.VerifyPhoneRegistration_Reply{}, err
}

func (s *AuthService) LoginByPassword(ctx context.Context, req *v1.LoginByPassword_Request) (*v1.LoginByPassword_Reply, error) {
	token, account, err := s.authUsecase.LoginByPassword(ctx, req.Account, req.Password, &model.LoginContext{
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
	})
	if err != nil {
		if cerrors.IsInternalServerError(err) {
			return nil, err
		}
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS).WithCause(err)
	}
	basic := &v1.AccountBasic{
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
	return &v1.LoginByPassword_Reply{
		Token: token,
		Account: &v1.Account{
			Basic: basic,
			Contact: &v1.AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		},
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *v1.Logout_Request) (*v1.Logout_Reply, error) {
	err := s.authUsecase.Logout(ctx, req.GetToken())
	return &v1.Logout_Reply{}, err
}

func (s *AuthService) ParseToken(ctx context.Context, req *v1.ParseToken_Request) (*v1.ParseToken_Reply, error) {
	user, err := s.authUsecase.ParseToken(ctx, req.GetToken())
	if err != nil {
		return nil, err
	}
	return &v1.ParseToken_Reply{
		User: &v1.TokenUser{
			Id:       user.ID,
			Name:     user.Name,
			Nickname: user.Nickname,
			Language: user.Language,
			Timezone: user.Timezone,
		},
	}, nil
}
