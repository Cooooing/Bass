package service

import (
	userv1 "common/proto/gen/user/v1"
	"context"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"user/internal/biz/usecase"
)

type TCCLivenessRewardService struct {
	userv1.UnimplementedTCCLivenessRewardServiceServer
	usecase *usecase.TCCLivenessRewardUsecase
}

func NewTCCLivenessRewardService(usecase *usecase.TCCLivenessRewardUsecase) *TCCLivenessRewardService {
	return &TCCLivenessRewardService{usecase: usecase}
}
func (s *TCCLivenessRewardService) RegisterGrpc(server *grpc.Server) {
	userv1.RegisterTCCLivenessRewardServiceServer(server, s)
}
func (*TCCLivenessRewardService) RegisterHTTP(*http.Server) {}
func (s *TCCLivenessRewardService) Try(ctx context.Context, req *userv1.TCCLivenessReward_Req) (*userv1.TCCLivenessReward_Resp, error) {
	return &userv1.TCCLivenessReward_Resp{}, s.usecase.Try(ctx, req)
}
func (s *TCCLivenessRewardService) Confirm(ctx context.Context, req *userv1.TCCLivenessReward_Req) (*userv1.TCCLivenessReward_Resp, error) {
	return &userv1.TCCLivenessReward_Resp{}, s.usecase.Confirm(ctx, req)
}
func (s *TCCLivenessRewardService) Cancel(ctx context.Context, req *userv1.TCCLivenessReward_Req) (*userv1.TCCLivenessReward_Resp, error) {
	return &userv1.TCCLivenessReward_Resp{}, s.usecase.Cancel(ctx, req)
}
