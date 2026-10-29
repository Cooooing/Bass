package service

import (
	userv1 "common/proto/gen/user/v1"
	"context"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"user/internal/biz/usecase"
)

type TCCCheckinService struct {
	userv1.UnimplementedTCCCheckinServiceServer
	usecase *usecase.TCCCheckinUsecase
}

func NewTCCCheckinService(usecase *usecase.TCCCheckinUsecase) *TCCCheckinService {
	return &TCCCheckinService{usecase: usecase}
}
func (s *TCCCheckinService) RegisterGrpc(server *grpc.Server) {
	userv1.RegisterTCCCheckinServiceServer(server, s)
}
func (*TCCCheckinService) RegisterHTTP(*http.Server) {}
func (s *TCCCheckinService) Try(ctx context.Context, req *userv1.TCCCheckin_Req) (*userv1.TCCCheckin_Resp, error) {
	return &userv1.TCCCheckin_Resp{}, s.usecase.Try(ctx, req)
}
func (s *TCCCheckinService) Confirm(ctx context.Context, req *userv1.TCCCheckin_Req) (*userv1.TCCCheckin_Resp, error) {
	return &userv1.TCCCheckin_Resp{}, s.usecase.Confirm(ctx, req)
}
func (s *TCCCheckinService) Cancel(ctx context.Context, req *userv1.TCCCheckin_Req) (*userv1.TCCCheckin_Resp, error) {
	return &userv1.TCCCheckin_Resp{}, s.usecase.Cancel(ctx, req)
}
