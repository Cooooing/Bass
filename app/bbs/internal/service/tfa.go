package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type TfaService struct {
	bbsuserv1.UnimplementedTfaServiceServer
	userUsecase *usecase.UserUsecase
}

func NewTfaService(userUsecase *usecase.UserUsecase) *TfaService {
	return &TfaService{userUsecase: userUsecase}
}

func (s *TfaService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterTfaServiceServer(gs, s)
}

func (s *TfaService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterTfaServiceHTTPServer(hs, s)
}

func (s *TfaService) Validate(ctx context.Context, req *bbsuserv1.ValidateTfa_Request) (*bbsuserv1.ValidateTfa_Reply, error) {
	return s.userUsecase.ValidateTfa(ctx, req)
}

func (s *TfaService) BeginEnable(ctx context.Context, req *bbsuserv1.BeginEnableTfa_Request) (*bbsuserv1.BeginEnableTfa_Reply, error) {
	return s.userUsecase.BeginEnableTfa(ctx, req)
}

func (s *TfaService) ConfirmEnable(ctx context.Context, req *bbsuserv1.ConfirmEnableTfa_Request) (*bbsuserv1.ConfirmEnableTfa_Reply, error) {
	return s.userUsecase.ConfirmEnableTfa(ctx, req)
}

func (s *TfaService) Disable(ctx context.Context, req *bbsuserv1.DisableTfa_Request) (*bbsuserv1.DisableTfa_Reply, error) {
	return s.userUsecase.DisableTfa(ctx, req)
}

func (s *TfaService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentTfa_Request) (*bbsuserv1.GetCurrentTfa_Reply, error) {
	return s.userUsecase.GetCurrentTfa(ctx, req)
}
