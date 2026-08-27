package service

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"strings"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	authUsecase *usecase.AuthUsecase
}

func NewAuthService(
	authUsecase *usecase.AuthUsecase,
) *AuthService {
	return &AuthService{
		authUsecase: authUsecase,
	}
}

func (s *AuthService) RegisterGrpc(*grpc.Server) {
}

func (s *AuthService) RegisterHttp(hs *http.Server) {
	v1.RegisterAuthServiceHTTPServer(hs, s)
}

func (s *AuthService) Register(ctx context.Context, req *v1.RegisterAccount_Req) (*v1.RegisterAccount_Resp, error) {
	email := strings.ToLower(strings.TrimSpace(req.GetEmail()))
	return &v1.RegisterAccount_Resp{}, s.authUsecase.Register(ctx, &usecase.RegisterReq{
		Password: req.GetPassword(),
		Email:    email,
	})
}

func (s *AuthService) Login(ctx context.Context, req *v1.LoginAccount_Req) (*v1.LoginAccount_Resp, error) {
	email := strings.ToLower(strings.TrimSpace(req.GetEmail()))
	token, err := s.authUsecase.Login(ctx, &usecase.LoginReq{
		Email:    email,
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.LoginAccount_Resp{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		UserId:       token.UserID,
		Name:         token.Name,
		Nickname:     token.Nickname,
	}
	if token.AccessTokenExpiresAt != nil {
		reply.AccessTokenExpiresAt = timestamppb.New(*token.AccessTokenExpiresAt)
	}
	if token.RefreshTokenExpiresAt != nil {
		reply.RefreshTokenExpiresAt = timestamppb.New(*token.RefreshTokenExpiresAt)
	}
	if token.SessionExpiresAt != nil {
		reply.SessionExpiresAt = timestamppb.New(*token.SessionExpiresAt)
	}
	return reply, nil
}
