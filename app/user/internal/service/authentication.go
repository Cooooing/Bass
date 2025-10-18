package service

import (
	v1 "common/api/user/v1"
	"context"
	"user/internal/biz"
	"user/internal/biz/model"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthenticationService struct {
	v1.UnimplementedAuthenticationServer
	*BaseService
	authenticationDomain *biz.AuthenticationDomain
	userRepo             repo.UserRepo
}

func NewAuthenticationService(baseService *BaseService, authenticationDomain *biz.AuthenticationDomain, userRepo repo.UserRepo) *AuthenticationService {
	return &AuthenticationService{
		BaseService:          baseService,
		authenticationDomain: authenticationDomain,
		userRepo:             userRepo,
	}
}

func (s *AuthenticationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAuthenticationServer(gs, s)
}

func (s *AuthenticationService) RegisterHttp(hs *http.Server) {
	v1.RegisterAuthenticationHTTPServer(hs, s)
}

func (s *AuthenticationService) ExistEmail(ctx context.Context, req *v1.ExistEmailRequest) (rsp *v1.ExistEmailReply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, req.Email)
	return &v1.ExistEmailReply{Exist: &exist}, err
}

func (s *AuthenticationService) ExistPhone(ctx context.Context, req *v1.ExistPhoneRequest) (rsp *v1.ExistPhoneReply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, req.Phone)
	return &v1.ExistPhoneReply{Exist: &exist}, err
}

func (s *AuthenticationService) ExistUsername(ctx context.Context, req *v1.ExistUsernameRequest) (rsp *v1.ExistUsernameReply, err error) {
	exist, err := s.userRepo.ConstantAccount(ctx, req.Username)
	return &v1.ExistUsernameReply{Exist: &exist}, err
}

func (s *AuthenticationService) LoginAccount(ctx context.Context, req *v1.LoginAccountRequest) (rsp *v1.LoginAccountReply, err error) {
	token, err := s.authenticationDomain.LoginAccount(ctx, req.Account, req.Password)
	return &v1.LoginAccountReply{Token: token}, err
}

func (s *AuthenticationService) RegisterEmail(ctx context.Context, req *v1.RegisterEmailRequest) (rsp *v1.RegisterEmailReply, err error) {
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
