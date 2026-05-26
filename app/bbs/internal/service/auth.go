package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthService struct {
	bbsuserv1.UnimplementedAuthServiceServer
	userUsecase *usecase.UserUsecase
}

func NewAuthService(userUsecase *usecase.UserUsecase) *AuthService {
	return &AuthService{userUsecase: userUsecase}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAuthServiceHTTPServer(hs, s)
}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Reply, error) {
	return s.userUsecase.StartEmailRegistration(ctx, req)
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Reply, error) {
	return s.userUsecase.VerifyEmailRegistration(ctx, req)
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Reply, error) {
	return s.userUsecase.StartPhoneRegistration(ctx, req)
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Reply, error) {
	return s.userUsecase.VerifyPhoneRegistration(ctx, req)
}

func (s *AuthService) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Reply, error) {
	return s.userUsecase.LoginByPassword(ctx, req)
}

func (s *AuthService) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	return s.userUsecase.Logout(ctx, req)
}
