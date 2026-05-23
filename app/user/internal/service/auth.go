package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/conf"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	conf        *conf.Bootstrap
	log         *log.Helper
	verify      *VerifyService
	authUsecase *usecase.AuthUsecase
}

func NewAuthService(conf *conf.Bootstrap, logger log.Logger, verifyService *VerifyService, authUsecase *usecase.AuthUsecase) *AuthService {
	return &AuthService{
		conf:        conf,
		log:         log.NewHelper(logger),
		verify:      verifyService,
		authUsecase: authUsecase,
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {}

func (s *AuthService) RegisterEmail(ctx context.Context, req *v1.RegisterEmail_Request) (*v1.RegisterEmail_Reply, error) {
	if !s.verify.VerifyName(req.Name) {
		return nil, cerrors.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.verify.VerifyNickname(*req.Nickname) {
		return nil, cerrors.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters")
	}
	if !s.verify.VerifyPassword(req.Password) {
		return nil, cerrors.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number")
	}
	_, token, err := s.authUsecase.RegisterEmail(ctx, &model.Account{
		Email:    &req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	return &v1.RegisterEmail_Reply{CodeToken: token}, err
}

func (s *AuthService) VerifyEmailRegister(ctx context.Context, req *v1.VerifyEmailRegister_Request) (*v1.VerifyEmailRegister_Reply, error) {
	err := s.authUsecase.RegisterEmailVerify(ctx, req.CodeToken, req.Code)
	return &v1.VerifyEmailRegister_Reply{}, err
}

func (s *AuthService) RegisterPhone(ctx context.Context, req *v1.RegisterPhone_Request) (*v1.RegisterPhone_Reply, error) {
	if !s.verify.VerifyName(req.Name) {
		return nil, cerrors.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.verify.VerifyNickname(*req.Nickname) {
		return nil, cerrors.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters")
	}
	if !s.verify.VerifyPassword(req.Password) {
		return nil, cerrors.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number")
	}
	_, token, err := s.authUsecase.RegisterPhone(ctx, &model.Account{
		Phone:    &req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	return &v1.RegisterPhone_Reply{CodeToken: token}, err
}

func (s *AuthService) VerifyPhoneRegister(ctx context.Context, req *v1.VerifyPhoneRegister_Request) (*v1.VerifyPhoneRegister_Reply, error) {
	err := s.authUsecase.RegisterPhoneVerify(ctx, req.CodeToken, req.Code)
	return &v1.VerifyPhoneRegister_Reply{}, err
}

func (s *AuthService) LoginPassword(ctx context.Context, req *v1.LoginPassword_Request) (*v1.LoginPassword_Reply, error) {
	token, account, err := s.authUsecase.LoginAccount(ctx, req.Account, req.Password)
	if err != nil {
		return nil, cerrors.ErrorBadRequest("account not exist or password is incorrect").WithCause(err)
	}
	basic := &v1.AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarUrl:     account.AvatarURL,
		Introduction:  account.Introduction,
		Mbti:          account.Mbti,
		GroupName:     account.GroupName,
		FollowCount:   account.FollowCount,
		FollowerCount: account.FollowerCount,
		BlockCount:    account.BlockCount,
		BlockedCount:  account.BlockedCount,
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
	return &v1.LoginPassword_Reply{
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
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err := s.authUsecase.Logout(ctx, token)
	return &v1.Logout_Reply{}, err
}
