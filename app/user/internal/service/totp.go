package service

import (
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TotpService struct {
	v1.UnimplementedTotpServiceServer
	totpUsecase *usecase.TotpUsecase
}

func NewTotpService(
	totpUsecase *usecase.TotpUsecase,
) *TotpService {
	return &TotpService{
		totpUsecase: totpUsecase,
	}
}

func (s *TotpService) RegisterGrpc(
	gs *grpc.Server,
) {
	v1.RegisterTotpServiceServer(gs, s)
}

func (s *TotpService) RegisterHttp(
	hs *http.Server,
) {
}

func (s *TotpService) Validate(
	ctx context.Context,
	req *v1.ValidateTotp_Req,
) (*v1.ValidateTotp_Resp, error) {
	res, err := s.totpUsecase.ValidateByUserID(ctx, &usecase.ValidateTotpByUserIDReq{
		UserID: req.GetUserId(),
		Code:   req.GetCode(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.ValidateTotp_Resp{
		Verified: res,
	}, nil
}

func (s *TotpService) BeginEnable(
	ctx context.Context,
	req *v1.BeginEnableTotp_Req,
) (*v1.BeginEnableTotp_Resp, error) {
	res, err := s.totpUsecase.BeginEnable(ctx, &usecase.BeginEnableTotpReq{
		UserID:      req.GetUserId(),
		AccountName: req.GetAccountName(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.BeginEnableTotp_Resp{
		Url:    res.URL,
		QrCode: res.QRCode,
	}, nil
}

func (s *TotpService) ConfirmEnable(
	ctx context.Context,
	req *v1.ConfirmEnableTotp_Req,
) (*v1.ConfirmEnableTotp_Resp, error) {
	err := s.totpUsecase.ConfirmEnable(ctx, &usecase.ConfirmEnableTotpReq{
		UserID: req.GetUserId(),
		Code:   req.GetCode(),
	})
	return &v1.ConfirmEnableTotp_Resp{}, err
}

func (s *TotpService) Disable(
	ctx context.Context,
	req *v1.DisableTotp_Req,
) (*v1.DisableTotp_Resp, error) {
	err := s.totpUsecase.Disable(ctx, &usecase.DisableTotpReq{
		UserID: req.GetUserId(),
		Code:   req.GetCode(),
	})
	return &v1.DisableTotp_Resp{}, err
}

func (s *TotpService) Get(
	ctx context.Context,
	req *v1.GetTotp_Req,
) (*v1.GetTotp_Resp, error) {
	res, err := s.totpUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.GetTotp_Resp_Totp{
		UserId: req.GetUserId(),
	}
	if res != nil {
		reply.Enable = res.Enable
		if res.EnableTime != nil {
			reply.EnableTime = timestamppb.New(*res.EnableTime)
		}
	}
	return &v1.GetTotp_Resp{
		Totp: reply,
	}, nil
}
