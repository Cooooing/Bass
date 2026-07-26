package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type TotpService struct {
	bbsuserv1.UnimplementedTotpServiceServer
	totpUsecase *usecase.TotpUsecase
}

func NewTotpService(
	totpUsecase *usecase.TotpUsecase,
) *TotpService {
	return &TotpService{
		totpUsecase: totpUsecase,
	}
}

func (s *TotpService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterTotpServiceHTTPServer(hs, s)
}

func (s *TotpService) RegisterGrpc(gs *grpc.Server) {
}

func (s *TotpService) BeginEnable(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Req) (*bbsuserv1.BeginEnableTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.totpUsecase.BeginEnableTotp(ctx, &usecase.BeginEnableTotpReq{
		UserID:      user.ID,
		AccountName: user.Name,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTotp_Resp{
		Url:    resp.URL,
		QrCode: resp.QRCode,
	}, nil
}

func (s *TotpService) ConfirmEnable(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Req) (*bbsuserv1.ConfirmEnableTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.totpUsecase.ConfirmEnableTotp(ctx, &usecase.ConfirmEnableTotpReq{
		UserID: user.ID,
		Code:   req.GetCode(),
	})
	return &bbsuserv1.ConfirmEnableTotp_Resp{}, err
}

func (s *TotpService) Disable(ctx context.Context, req *bbsuserv1.DisableTotp_Req) (*bbsuserv1.DisableTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.totpUsecase.DisableTotp(ctx, &usecase.DisableTotpReq{
		UserID: user.ID,
		Code:   req.GetCode(),
	})
	return &bbsuserv1.DisableTotp_Resp{}, err
}

func (s *TotpService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Req) (*bbsuserv1.GetCurrentTotp_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	totp, err := s.totpUsecase.GetCurrentTotp(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentTotp_Resp{
		Totp: totp,
	}, nil
}
