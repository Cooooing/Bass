package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
	"regexp"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthService struct {
	bbsuserv1.UnimplementedAuthServiceServer
	authUsecase *usecase.AuthUsecase
	phoneRe     *regexp.Regexp
	nameRe      *regexp.Regexp
	codeRe      *regexp.Regexp
}

func NewAuthService(authUsecase *usecase.AuthUsecase) *AuthService {
	return &AuthService{
		authUsecase: authUsecase,
		phoneRe:     regexp.MustCompile(`^1[3-9]\d{9}$`),
		nameRe:      regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`),
		codeRe:      regexp.MustCompile(`^[A-Za-z0-9]{6}$`),
	}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAuthServiceHTTPServer(hs, s)
}

func (s *AuthService) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Reply, error) {
	if err := s.validateStartEmailRegistration(req); err != nil {
		return nil, err
	}
	return s.authUsecase.StartEmailRegistration(ctx, req)
}

func (s *AuthService) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Reply, error) {
	if err := s.validateVerifyEmailRegistration(req); err != nil {
		return nil, err
	}
	return s.authUsecase.VerifyEmailRegistration(ctx, req)
}

func (s *AuthService) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Reply, error) {
	if err := s.validateStartPhoneRegistration(req); err != nil {
		return nil, err
	}
	return s.authUsecase.StartPhoneRegistration(ctx, req)
}

func (s *AuthService) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Reply, error) {
	if err := s.validateVerifyPhoneRegistration(req); err != nil {
		return nil, err
	}
	return s.authUsecase.VerifyPhoneRegistration(ctx, req)
}

func (s *AuthService) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Reply, error) {
	if err := s.validateLoginByPassword(req); err != nil {
		return nil, err
	}
	return s.authUsecase.LoginByPassword(ctx, req)
}

func (s *AuthService) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	return s.authUsecase.Logout(ctx, req)
}
