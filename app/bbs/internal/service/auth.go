package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AuthService struct {
	bbsuserv1.UnimplementedAuthServiceServer
	userClient *rpc.UserClient
}

func NewAuthService(userClient *rpc.UserClient) *AuthService {
	return &AuthService{userClient: userClient}
}

func (s *AuthService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterAuthServiceServer(gs, s)
}

func (s *AuthService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAuthServiceHTTPServer(hs, s)
}

func (s *AuthService) RegisterEmail(ctx context.Context, req *bbsuserv1.RegisterEmail_Request) (*bbsuserv1.RegisterEmail_Reply, error) {
	reply, err := s.userClient.Auth.RegisterEmail(ctx, &userv1.RegisterEmail_Request{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.RegisterEmail_Reply{CodeToken: reply.GetCodeToken()}, nil
}

func (s *AuthService) VerifyEmailRegister(ctx context.Context, req *bbsuserv1.VerifyEmailRegister_Request) (*bbsuserv1.VerifyEmailRegister_Reply, error) {
	_, err := s.userClient.Auth.VerifyEmailRegister(ctx, &userv1.VerifyEmailRegister_Request{
		Code:      req.GetCode(),
		CodeToken: req.GetCodeToken(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyEmailRegister_Reply{}, nil
}

func (s *AuthService) RegisterPhone(ctx context.Context, req *bbsuserv1.RegisterPhone_Request) (*bbsuserv1.RegisterPhone_Reply, error) {
	reply, err := s.userClient.Auth.RegisterPhone(ctx, &userv1.RegisterPhone_Request{
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.RegisterPhone_Reply{CodeToken: reply.GetCodeToken()}, nil
}

func (s *AuthService) VerifyPhoneRegister(ctx context.Context, req *bbsuserv1.VerifyPhoneRegister_Request) (*bbsuserv1.VerifyPhoneRegister_Reply, error) {
	_, err := s.userClient.Auth.VerifyPhoneRegister(ctx, &userv1.VerifyPhoneRegister_Request{
		Code:      req.GetCode(),
		CodeToken: req.GetCodeToken(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyPhoneRegister_Reply{}, nil
}

func (s *AuthService) LoginPassword(ctx context.Context, req *bbsuserv1.LoginPassword_Request) (*bbsuserv1.LoginPassword_Reply, error) {
	reply, err := s.userClient.Auth.LoginPassword(ctx, &userv1.LoginPassword_Request{
		Account:  req.GetAccount(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.LoginPassword_Reply{
		Token:   reply.GetToken(),
		Account: toBFFAccount(reply.GetAccount()),
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	_, err := s.userClient.Auth.Logout(forwardAuth(ctx), &userv1.Logout_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.Logout_Reply{}, nil
}
