package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type TwoFactorAuthService struct {
	v1.UnimplementedUserTwoFactorAuthServiceServer
	twoFactorAuthUsecase *usecase.TwoFactorAuthUsecase
}

func NewTwoFactorAuthService(twoFactorAuthUsecase *usecase.TwoFactorAuthUsecase) *TwoFactorAuthService {
	return &TwoFactorAuthService{
		twoFactorAuthUsecase: twoFactorAuthUsecase,
	}
}

func (s *TwoFactorAuthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserTwoFactorAuthServiceServer(gs, s)
}

func (s *TwoFactorAuthService) RegisterHttp(hs *http.Server) {}

func (s *TwoFactorAuthService) Validate(ctx context.Context, req *v1.ValidateTwoFactorAuth_Request) (rsp *v1.ValidateTwoFactorAuth_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	validate := s.twoFactorAuthUsecase.Validate(ctx, user.TwofaSecret, req.Code)
	return &v1.ValidateTwoFactorAuth_Reply{
		Verified: validate,
	}, nil
}

func (s *TwoFactorAuthService) Enable(ctx context.Context, req *v1.EnableTwoFactorAuth_Request) (rsp *v1.EnableTwoFactorAuth_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if user.TwofaEnable {
		return nil, cerrors.ErrorBadRequest("2FA already enabled")
	}
	buf, err := s.twoFactorAuthUsecase.Enable(ctx, user.Name)
	if err != nil {
		return nil, err
	}

	return &v1.EnableTwoFactorAuth_Reply{
		Data:        buf,
		ContentType: "image/png",
	}, nil
}

func (s *TwoFactorAuthService) Disable(ctx context.Context, req *v1.DisableTwoFactorAuth_Request) (rsp *v1.DisableTwoFactorAuth_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if !user.TwofaEnable {
		return nil, cerrors.ErrorBadRequest("2FA already disabled")
	}
	err = s.twoFactorAuthUsecase.Disable(ctx, user.Name, user.TwofaSecret, req.Code)
	return &v1.DisableTwoFactorAuth_Reply{}, err
}

func (s *TwoFactorAuthService) Confirm(ctx context.Context, req *v1.ConfirmTwoFactorAuth_Request) (rsp *v1.ConfirmTwoFactorAuth_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err = s.twoFactorAuthUsecase.Confirm(ctx, user.Name, req.Code)
	return &v1.ConfirmTwoFactorAuth_Reply{}, err
}
