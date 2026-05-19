package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	"common/pkg/util"

	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/biz/usecase"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthService struct {
	v1.UnimplementedUserAuthServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	*VerifyService
	authUsecase         *usecase.AuthUsecase
	userRepo            repo.UserRepo
	userPreferencesRepo repo.UserPreferencesRepo
	userPrivacyRepo     repo.UserPrivacyRepo
	userLocationRepo    repo.UserLocationRepo
	userTfaRepo         repo.UserTfaRepo
	userCheckinRepo     repo.UserCheckinRepo
}

func NewAuthService(conf *conf.Bootstrap, logger log.Logger, verifyService *VerifyService, authUsecase *usecase.AuthUsecase,
	userRepo repo.UserRepo,
	userPreferencesRepo repo.UserPreferencesRepo,
	userPrivacyRepo repo.UserPrivacyRepo,
	userLocationRepo repo.UserLocationRepo,
	userTfaRepo repo.UserTfaRepo,
	userCheckinRepo repo.UserCheckinRepo) *AuthService {
	return &AuthService{
		conf:                conf,
		log:                 log.NewHelper(logger),
		VerifyService:       verifyService,
		authUsecase:         authUsecase,
		userRepo:            userRepo,
		userPreferencesRepo: userPreferencesRepo,
		userPrivacyRepo:     userPrivacyRepo,
		userLocationRepo:    userLocationRepo,
		userTfaRepo:         userTfaRepo,
		userCheckinRepo:     userCheckinRepo,
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {}

func (s *AuthService) RegisterEmail(ctx context.Context, req *v1.RegisterEmailAuth_Request) (rsp *v1.RegisterEmailAuth_Reply, err error) {
	if !s.VerifyName(req.Name) {
		return nil, cerrors.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.VerifyNickname(*req.Nickname) {
		return nil, cerrors.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters (emoji, Chinese, etc.)")
	}
	if !s.VerifyPassword(req.Password) {
		return nil, cerrors.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number, and may include letters, numbers, and special symbols @#$%^&*!()_+-=[]{};:'\",.<>/?`~|\\")
	}
	code, token, err := s.authUsecase.RegisterEmail(ctx, &model.User{
		Email:    &req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	return &v1.RegisterEmailAuth_Reply{Code: code, CodeToken: token}, err
}

func (s *AuthService) RegisterEmailVerify(ctx context.Context, req *v1.RegisterEmailVerifyAuth_Request) (rsp *v1.RegisterEmailVerifyAuth_Reply, err error) {
	err = s.authUsecase.RegisterEmailVerify(ctx, req.CodeToken, req.Code)
	return &v1.RegisterEmailVerifyAuth_Reply{}, err
}

func (s *AuthService) RegisterPhone(ctx context.Context, req *v1.RegisterPhoneAuth_Request) (rsp *v1.RegisterPhoneAuth_Reply, err error) {
	if !s.VerifyName(req.Name) {
		return nil, cerrors.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.VerifyNickname(*req.Nickname) {
		return nil, cerrors.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters (emoji, Chinese, etc.)")
	}
	if !s.VerifyPassword(req.Password) {
		return nil, cerrors.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number, and may include letters, numbers, and special symbols @#$%^&*!()_+-=[]{};:'\",.<>/?`~|\\")
	}
	code, token, err := s.authUsecase.RegisterPhone(ctx, &model.User{
		Phone:    &req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	return &v1.RegisterPhoneAuth_Reply{Code: code, CodeToken: token}, err
}

func (s *AuthService) RegisterPhoneVerify(ctx context.Context, req *v1.RegisterPhoneVerifyAuth_Request) (rsp *v1.RegisterPhoneVerifyAuth_Reply, err error) {
	err = s.authUsecase.RegisterPhoneVerify(ctx, req.CodeToken, req.Code)
	return &v1.RegisterPhoneVerifyAuth_Reply{}, err
}

func (s *AuthService) ExistEmail(ctx context.Context, req *v1.ExistEmailAuth_Request) (rsp *v1.ExistEmailAuth_Reply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, req.Email)
	return &v1.ExistEmailAuth_Reply{Exist: &exist}, err
}

func (s *AuthService) ExistPhone(ctx context.Context, req *v1.ExistPhoneAuth_Request) (rsp *v1.ExistPhoneAuth_Reply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, req.Phone)
	return &v1.ExistPhoneAuth_Reply{Exist: &exist}, err
}

func (s *AuthService) ExistUsername(ctx context.Context, req *v1.ExistUsernameAuth_Request) (rsp *v1.ExistUsernameAuth_Reply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, req.Username)
	return &v1.ExistUsernameAuth_Reply{Exist: &exist}, err
}

func (s *AuthService) LoginAccount(ctx context.Context, req *v1.LoginAccountAuth_Request) (rsp *v1.LoginAccountAuth_Reply, err error) {
	token, user, err := s.authUsecase.LoginAccount(ctx, req.Account, req.Password)
	if err != nil {
		return nil, cerrors.ErrorBadRequest("account not exist or password is incorrect").WithCause(err)
	}
	return &v1.LoginAccountAuth_Reply{
		Token: token,
		User:  assembleUserProto(ctx, user, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo),
	}, err
}

func (s *AuthService) Logout(ctx context.Context, req *v1.LogoutAuth_Request) (rsp *v1.LogoutAuth_Reply, err error) {
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.authUsecase.Logout(ctx, token)
	return &v1.LogoutAuth_Reply{}, err
}
