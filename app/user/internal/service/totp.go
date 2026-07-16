package service

import (
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TotpService struct {
	v1.UnimplementedTotpServiceServer
	totpUsecase *usecase.TotpUsecase
}

func NewTotpService(totpUsecase *usecase.TotpUsecase) *TotpService {
	return &TotpService{
		totpUsecase: totpUsecase,
	}
}

func (s *TotpService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterTotpServiceServer(gs, s)
}

func (s *TotpService) Validate(ctx context.Context, req *v1.ValidateTotp_Request) (*v1.ValidateTotp_Response, error) {
	res, err := s.totpUsecase.ValidateByUserID(ctx, &usecase.ValidateTotpByUserIDReq{UserID: req.GetUserId(), Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &v1.ValidateTotp_Response{Verified: res.Verified}, nil
}

func (s *TotpService) BeginEnable(ctx context.Context, req *v1.BeginEnableTotp_Request) (*v1.BeginEnableTotp_Response, error) {
	res, err := s.totpUsecase.BeginEnable(ctx, &usecase.BeginEnableTotpReq{UserID: req.GetUserId(), AccountName: req.GetAccountName()})
	if err != nil {
		return nil, err
	}
	return &v1.BeginEnableTotp_Response{Url: res.URL, QrCode: res.QRCode}, nil
}

func (s *TotpService) ConfirmEnable(ctx context.Context, req *v1.ConfirmEnableTotp_Request) (*v1.ConfirmEnableTotp_Response, error) {
	err := s.totpUsecase.ConfirmEnable(ctx, &usecase.ConfirmEnableTotpReq{UserID: req.GetUserId(), Code: req.GetCode()})
	return &v1.ConfirmEnableTotp_Response{}, err
}

func (s *TotpService) Disable(ctx context.Context, req *v1.DisableTotp_Request) (*v1.DisableTotp_Response, error) {
	err := s.totpUsecase.Disable(ctx, &usecase.DisableTotpReq{UserID: req.GetUserId(), Code: req.GetCode()})
	return &v1.DisableTotp_Response{}, err
}

func (s *TotpService) Get(ctx context.Context, req *v1.GetTotp_Request) (*v1.GetTotp_Response, error) {
	res, err := s.totpUsecase.GetByUserID(ctx, &usecase.GetTotpByUserIDReq{UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	reply := &v1.GetTotp_Response_Totp{UserId: req.GetUserId()}
	if res.Totp != nil {
		reply.Enable = res.Totp.Enable
		if res.Totp.EnableTime != nil {
			reply.EnableTime = timestamppb.New(*res.Totp.EnableTime)
		}
	}
	return &v1.GetTotp_Response{Totp: reply}, nil
}
