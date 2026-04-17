package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"user/internal/biz/doamin"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type TwoFactorAuthenticationService struct {
	v1.UnimplementedUserTwoFactorAuthenticationServiceServer
	*BaseService
	twoFactorAuthenticationDomain *doamin.TwoFactorAuthenticationDomain
}

func NewTwoFactorAuthenticationService(baseService *BaseService, userRelationDomain *doamin.TwoFactorAuthenticationDomain) *TwoFactorAuthenticationService {
	return &TwoFactorAuthenticationService{
		BaseService:                   baseService,
		twoFactorAuthenticationDomain: userRelationDomain,
	}
}

func (s *TwoFactorAuthenticationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserTwoFactorAuthenticationServiceServer(gs, s)
}

func (s *TwoFactorAuthenticationService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserTwoFactorAuthenticationServiceHTTPServer(hs, s)
}

func (s *TwoFactorAuthenticationService) Validate(ctx context.Context, req *v1.ValidateTwoFactorAuthentication_Request) (rsp *v1.ValidateTwoFactorAuthentication_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	validate := s.twoFactorAuthenticationDomain.Validate(ctx, user.TwofaSecret, req.Code)
	return &v1.ValidateTwoFactorAuthentication_Reply{
		Verified: validate,
	}, nil
}

func (s *TwoFactorAuthenticationService) Enable(ctx context.Context, req *v1.EnableTwoFactorAuthentication_Request) (rsp *v1.EnableTwoFactorAuthentication_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if user.TwofaEnable {
		return nil, cerrors.ErrorBadRequest("2FA already enabled")
	}
	buf, err := s.twoFactorAuthenticationDomain.Enable(ctx, user.Name)
	if err != nil {
		return nil, err
	}

	return &v1.EnableTwoFactorAuthentication_Reply{
		Data:        buf,
		ContentType: "image/png",
	}, nil
}

func (s *TwoFactorAuthenticationService) Disable(ctx context.Context, req *v1.DisableTwoFactorAuthentication_Request) (rsp *v1.DisableTwoFactorAuthentication_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if !user.TwofaEnable {
		return nil, cerrors.ErrorBadRequest("2FA already disabled")
	}
	err = s.twoFactorAuthenticationDomain.Disable(ctx, user.Name, user.TwofaSecret, req.Code)
	return &v1.DisableTwoFactorAuthentication_Reply{}, err
}

func (s *TwoFactorAuthenticationService) Confirm(ctx context.Context, req *v1.ConfirmTwoFactorAuthentication_Request) (rsp *v1.ConfirmTwoFactorAuthentication_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.twoFactorAuthenticationDomain.Confirm(ctx, user.Name, req.Code)
	return &v1.ConfirmTwoFactorAuthentication_Reply{}, err
}
