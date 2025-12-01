package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/user/v1"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"user/internal/biz"
	"user/internal/biz/model"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthenticationService struct {
	v1.UnimplementedUserAuthenticationServiceServer
	*BaseService
	*VerifyService
	authenticationDomain *biz.AuthenticationDomain
	userRepo             repo.UserRepo
}

func NewAuthenticationService(baseService *BaseService, verifyService *VerifyService, authenticationDomain *biz.AuthenticationDomain, userRepo repo.UserRepo) *AuthenticationService {
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

func (s *AuthenticationService) RegisterEmail(ctx context.Context, req *v1.RegisterEmailRequest) (rsp *v1.RegisterEmailReply, err error) {
	if !s.VerifyName(req.Name) {
		return nil, cv1.ErrorBadRequest("name must be 4-32 characters long, only letters, numbers, and single '-' allowed (cannot start or end with '-')")
	}
	if req.Nickname != nil && !s.VerifyNickname(*req.Nickname) {
		return nil, cv1.ErrorBadRequest("nickname must be 2-32 characters long, contain at least one non-digit character, and may include letters, numbers, '_', '-', or Unicode characters (emoji, Chinese, etc.)")
	}
	if !s.VerifyPassword(req.Password) {
		return nil, cv1.ErrorBadRequest("password must be 6-64 characters long, contain at least one letter and one number, and may include letters, numbers, and special symbols @#$%^&*!()_+-=[]{};:'\",.<>/?`~|\\")
	}
	code, token, err := s.authenticationDomain.RegisterEmail(ctx, &model.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	return &v1.RegisterEmailReply{Code: code, CodeToken: token}, err
}

func (s *AuthenticationService) RegisterEmailVerify(ctx context.Context, req *v1.RegisterEmailVerifyRequest) (rsp *v1.RegisterEmailVerifyReply, err error) {
	err = s.authenticationDomain.RegisterEmailVerify(ctx, req.CodeToken, req.Code)
	return &v1.RegisterEmailVerifyReply{}, err
}

func (s *AuthenticationService) ExistEmail(ctx context.Context, req *v1.ExistEmailRequest) (rsp *v1.ExistEmailReply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, s.db, req.Email)
	return &v1.ExistEmailReply{Exist: &exist}, err
}

func (s *AuthenticationService) ExistPhone(ctx context.Context, req *v1.ExistPhoneRequest) (rsp *v1.ExistPhoneReply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, s.db, req.Phone)
	return &v1.ExistPhoneReply{Exist: &exist}, err
}

func (s *AuthenticationService) ExistUsername(ctx context.Context, req *v1.ExistUsernameRequest) (rsp *v1.ExistUsernameReply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, s.db, req.Username)
	return &v1.ExistUsernameReply{Exist: &exist}, err
}

func (s *AuthenticationService) LoginAccount(ctx context.Context, req *v1.LoginAccountRequest) (rsp *v1.LoginAccountReply, err error) {
	token, user, err := s.authenticationDomain.LoginAccount(ctx, req.Account, req.Password)
	if err != nil {
		return nil, cv1.ErrorBadRequest("account not exist or password is incorrect").WithCause(err)
	}
	return &v1.LoginAccountReply{
		Token: token,
		User:  user.ConvertToRpc(),
	}, err
}

func (s *AuthenticationService) Logout(ctx context.Context, req *v1.LogoutRequest) (rsp *v1.LogoutReply, err error) {
	token := util.MustGetContextValue[string](ctx, constant.CtxToken)
	err = s.authenticationDomain.Logout(ctx, token)
	return &v1.LogoutReply{}, err
}
