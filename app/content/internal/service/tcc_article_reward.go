package service

import (
	"context"

	contentv1 "common/proto/gen/content/v1"

	"content/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// TCCArticleRewardService 适配文章打赏的 TCC 分支协议。
type TCCArticleRewardService struct {
	contentv1.UnimplementedTCCArticleRewardServiceServer
	usecase *usecase.TCCArticleRewardUsecase
}

func NewTCCArticleRewardService(usecase *usecase.TCCArticleRewardUsecase) *TCCArticleRewardService {
	return &TCCArticleRewardService{usecase: usecase}
}

func (s *TCCArticleRewardService) RegisterGrpc(server *grpc.Server) {
	contentv1.RegisterTCCArticleRewardServiceServer(server, s)
}

func (*TCCArticleRewardService) RegisterHTTP(*http.Server) {}

func (s *TCCArticleRewardService) Try(ctx context.Context, req *contentv1.TCCArticleReward_Req) (*contentv1.TCCArticleReward_Resp, error) {
	if err := s.usecase.Try(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticleReward_Resp{}, nil
}

func (s *TCCArticleRewardService) Confirm(ctx context.Context, req *contentv1.TCCArticleReward_Req) (*contentv1.TCCArticleReward_Resp, error) {
	if err := s.usecase.Confirm(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticleReward_Resp{}, nil
}

func (s *TCCArticleRewardService) Cancel(ctx context.Context, req *contentv1.TCCArticleReward_Req) (*contentv1.TCCArticleReward_Resp, error) {
	if err := s.usecase.Cancel(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticleReward_Resp{}, nil
}
