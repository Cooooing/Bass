package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"net/mail"
	"strings"
	"unicode/utf8"

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
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || !strings.Contains(email, "@") || utf8.RuneCountInString(email) > 254 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	password := req.GetPassword()
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range password {
		if r < '!' || r > '~' {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return &v1.RegisterAccount_Resp{}, s.authUsecase.Register(ctx, &usecase.RegisterReq{
		Password: password,
		Email:    email,
	})
}

func (s *AuthService) Login(ctx context.Context, req *v1.LoginAccount_Req) (*v1.LoginAccount_Resp, error) {
	email := strings.ToLower(strings.TrimSpace(req.GetEmail()))
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || !strings.Contains(email, "@") || utf8.RuneCountInString(email) > 254 || strings.TrimSpace(req.GetPassword()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
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
