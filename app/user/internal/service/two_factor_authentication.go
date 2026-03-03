package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"fmt"
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

func (s *TwoFactorAuthenticationService) Validate(ctx context.Context, req *v1.UserTwoFactorAuthenticationValidateRequest) (rsp *v1.UserTwoFactorAuthenticationValidateReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	validate := s.twoFactorAuthenticationDomain.Validate(ctx, user.TwofaSecret, req.Code)
	return &v1.UserTwoFactorAuthenticationValidateReply{
		Verified: validate,
	}, nil
}

func (s *TwoFactorAuthenticationService) Enable(ctx context.Context, req *v1.UserTwoFactorAuthenticationEnableRequest) (rsp *v1.UserTwoFactorAuthenticationEnableReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	if user.TwofaEnable {
		return nil, cv1.ErrorBadRequest("2FA already enabled")
	}
	buf, err := s.twoFactorAuthenticationDomain.Enable(ctx, user.Name)
	if w, ok := http.ResponseWriterFromServerContext(ctx); ok {
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Content-Length", fmt.Sprint(len(buf)))
		_, err := w.Write(buf)
		return nil, err
	}

	return &v1.UserTwoFactorAuthenticationEnableReply{
		Data:        buf,
		ContentType: "image/png",
	}, nil
}

func (s *TwoFactorAuthenticationService) Disable(ctx context.Context, req *v1.UserTwoFactorAuthenticationDisableRequest) (rsp *v1.UserTwoFactorAuthenticationDisableReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	if !user.TwofaEnable {
		return nil, cv1.ErrorBadRequest("2FA already disabled")
	}
	err = s.twoFactorAuthenticationDomain.Disable(ctx, user.Name, user.TwofaSecret, req.Code)
	return &v1.UserTwoFactorAuthenticationDisableReply{}, err
}

func (s *TwoFactorAuthenticationService) Confirm(ctx context.Context, req *v1.UserTwoFactorAuthenticationConfirmRequest) (rsp *v1.UserTwoFactorAuthenticationConfirmReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	err = s.twoFactorAuthenticationDomain.Confirm(ctx, user.Name, req.Code)
	return &v1.UserTwoFactorAuthenticationConfirmReply{}, err
}
