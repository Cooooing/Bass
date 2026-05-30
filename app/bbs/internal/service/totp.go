package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type TotpService struct {
	bbsuserv1.UnimplementedTotpServiceServer
	totpUsecase *usecase.TotpUsecase
}

func NewTotpService(totpUsecase *usecase.TotpUsecase) *TotpService {
	return &TotpService{totpUsecase: totpUsecase}
}

func (s *TotpService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterTotpServiceServer(gs, s)
}

func (s *TotpService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterTotpServiceHTTPServer(hs, s)
}

func (s *TotpService) BeginEnable(ctx context.Context, req *bbsuserv1.BeginEnableTotp_Request) (*bbsuserv1.BeginEnableTotp_Reply, error) {
	return s.totpUsecase.BeginEnableTotp(ctx, req)
}

func (s *TotpService) ConfirmEnable(ctx context.Context, req *bbsuserv1.ConfirmEnableTotp_Request) (*bbsuserv1.ConfirmEnableTotp_Reply, error) {
	return s.totpUsecase.ConfirmEnableTotp(ctx, req)
}

func (s *TotpService) Disable(ctx context.Context, req *bbsuserv1.DisableTotp_Request) (*bbsuserv1.DisableTotp_Reply, error) {
	return s.totpUsecase.DisableTotp(ctx, req)
}

func (s *TotpService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentTotp_Request) (*bbsuserv1.GetCurrentTotp_Reply, error) {
	return s.totpUsecase.GetCurrentTotp(ctx, req)
}
