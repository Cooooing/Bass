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

func NewTotpService(totpUsecase *usecase.TotpUsecase) *TotpService {
	return &TotpService{
		totpUsecase: totpUsecase,
	}
}

func (s *TotpService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterTotpServiceServer(gs, s)
}

func (s *TotpService) RegisterHttp(hs *http.Server) {}

func (s *TotpService) Validate(ctx context.Context, req *v1.ValidateTotp_Request) (*v1.ValidateTotp_Reply, error) {
	verified, err := s.totpUsecase.ValidateByUserID(ctx, req.GetUserId(), req.GetCode())
	if err != nil {
		return nil, err
	}
	return &v1.ValidateTotp_Reply{Verified: verified}, nil
}

func (s *TotpService) BeginEnable(ctx context.Context, req *v1.BeginEnableTotp_Request) (*v1.BeginEnableTotp_Reply, error) {
	url, qrCode, err := s.totpUsecase.BeginEnable(ctx, req.GetUserId(), req.GetAccountName())
	if err != nil {
		return nil, err
	}
	return &v1.BeginEnableTotp_Reply{Url: url, QrCode: qrCode}, nil
}

func (s *TotpService) ConfirmEnable(ctx context.Context, req *v1.ConfirmEnableTotp_Request) (*v1.ConfirmEnableTotp_Reply, error) {
	err := s.totpUsecase.ConfirmEnable(ctx, req.GetUserId(), req.GetCode())
	return &v1.ConfirmEnableTotp_Reply{}, err
}

func (s *TotpService) Disable(ctx context.Context, req *v1.DisableTotp_Request) (*v1.DisableTotp_Reply, error) {
	err := s.totpUsecase.Disable(ctx, req.GetUserId(), req.GetCode())
	return &v1.DisableTotp_Reply{}, err
}

func (s *TotpService) Get(ctx context.Context, req *v1.GetTotp_Request) (*v1.GetTotp_Reply, error) {
	totpSetting, err := s.totpUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.Totp{UserId: req.GetUserId()}
	if totpSetting != nil {
		reply.Enable = totpSetting.Enable
		if totpSetting.EnableTime != nil {
			reply.EnableTime = timestamppb.New(*totpSetting.EnableTime)
		}
	}
	return &v1.GetTotp_Reply{Totp: reply}, nil
}
