package service

import (
	"common/api/gen/common"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	"common/pkg/util"

	"context"
	"user/internal/biz/doamin"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthenticationService struct {
	v1.UnimplementedUserAuthenticationServiceServer
	*BaseService
	*VerifyService
	authenticationDomain *doamin.AuthenticationDomain
	userRepo             repo.UserRepo
}

func NewAuthenticationService(baseService *BaseService, verifyService *VerifyService, authenticationDomain *doamin.AuthenticationDomain, userRepo repo.UserRepo) *AuthenticationService {
	return &AuthenticationService{
		BaseService:          baseService,
		VerifyService:        verifyService,
		authenticationDomain: authenticationDomain,
		userRepo:             userRepo,
	}
}

func (s *AuthenticationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserAuthenticationServiceServer(gs, s)
}

func (s *AuthenticationService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserAuthenticationServiceHTTPServer(hs, s)
}

func (s *AuthenticationService) RegisterEmail(ctx context.Context, req *v1.RegisterEmailAuth_Request) (rsp *v1.RegisterEmailAuth_Reply, err error) {
	if !s.VerifyName(req.Name) {
		return nil, common.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.VerifyNickname(*req.Nickname) {
		return nil, common.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters (emoji, Chinese, etc.)")
	}
	if !s.VerifyPassword(req.Password) {
		return nil, common.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number, and may include letters, numbers, and special symbols @#$%^&*!()_+-=[]{};:'\",.<>/?`~|\\")
	}
	code, token, err := s.authenticationDomain.RegisterEmail(ctx, &model.User{User: &gen.User{
		Email:    new(req.Email),
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	}})
	return &v1.RegisterEmailAuth_Reply{Code: code, CodeToken: token}, err
}

func (s *AuthenticationService) RegisterEmailVerify(ctx context.Context, req *v1.RegisterEmailVerifyAuth_Request) (rsp *v1.RegisterEmailVerifyAuth_Reply, err error) {
	err = s.authenticationDomain.RegisterEmailVerify(ctx, req.CodeToken, req.Code)
	return &v1.RegisterEmailVerifyAuth_Reply{}, err
}

func (s *AuthenticationService) RegisterPhone(ctx context.Context, req *v1.RegisterPhoneAuth_Request) (rsp *v1.RegisterPhoneAuth_Reply, err error) {
	if !s.VerifyName(req.Name) {
		return nil, common.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.VerifyNickname(*req.Nickname) {
		return nil, common.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters (emoji, Chinese, etc.)")
	}
	if !s.VerifyPassword(req.Password) {
		return nil, common.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number, and may include letters, numbers, and special symbols @#$%^&*!()_+-=[]{};:'\",.<>/?`~|\\")
	}
	code, token, err := s.authenticationDomain.RegisterPhone(ctx, &model.User{User: &gen.User{
		Phone:    new(req.Phone),
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	}})
	return &v1.RegisterPhoneAuth_Reply{Code: code, CodeToken: token}, err
}

func (s *AuthenticationService) RegisterPhoneVerify(ctx context.Context, req *v1.RegisterPhoneVerifyAuth_Request) (rsp *v1.RegisterPhoneVerifyAuth_Reply, err error) {
	err = s.authenticationDomain.RegisterPhoneVerify(ctx, req.CodeToken, req.Code)
	return &v1.RegisterPhoneVerifyAuth_Reply{}, err
}

func (s *AuthenticationService) ExistEmail(ctx context.Context, req *v1.ExistEmailAuth_Request) (rsp *v1.ExistEmailAuth_Reply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, s.Db, req.Email)
	return &v1.ExistEmailAuth_Reply{Exist: &exist}, err
}

func (s *AuthenticationService) ExistPhone(ctx context.Context, req *v1.ExistPhoneAuth_Request) (rsp *v1.ExistPhoneAuth_Reply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, s.Db, req.Phone)
	return &v1.ExistPhoneAuth_Reply{Exist: &exist}, err
}

func (s *AuthenticationService) ExistUsername(ctx context.Context, req *v1.ExistUsernameAuth_Request) (rsp *v1.ExistUsernameAuth_Reply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, s.Db, req.Username)
	return &v1.ExistUsernameAuth_Reply{Exist: &exist}, err
}

func (s *AuthenticationService) LoginAccount(ctx context.Context, req *v1.LoginAccountAuth_Request) (rsp *v1.LoginAccountAuth_Reply, err error) {
	token, user, err := s.authenticationDomain.LoginAccount(ctx, req.Account, req.Password)
	if err != nil {
		return nil, common.ErrorBadRequest("account not exist or password is incorrect").WithCause(err)
	}
	return &v1.LoginAccountAuth_Reply{
		Token: token,
		User:  user.ConvertToRpc(),
	}, err
}

func (s *AuthenticationService) Logout(ctx context.Context, req *v1.LogoutAuth_Request) (rsp *v1.LogoutAuth_Reply, err error) {
	token, ok := util.GetContextValue[string](ctx, constant.CtxToken)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	err = s.authenticationDomain.Logout(ctx, token)
	return &v1.LogoutAuth_Reply{}, err
}
