package service

import (
	"context"

	contentv1 "common/proto/gen/content/v1"

	"content/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// TCCArticlePublishService 适配文章发布 TCC 分支协议。
type TCCArticlePublishService struct {
	contentv1.UnimplementedTCCArticlePublishServiceServer
	usecase *usecase.TCCArticlePublishUsecase
}

func NewTCCArticlePublishService(usecase *usecase.TCCArticlePublishUsecase) *TCCArticlePublishService {
	return &TCCArticlePublishService{usecase: usecase}
}

func (s *TCCArticlePublishService) RegisterGrpc(server *grpc.Server) {
	contentv1.RegisterTCCArticlePublishServiceServer(server, s)
}

func (*TCCArticlePublishService) RegisterHTTP(*http.Server) {}

func (s *TCCArticlePublishService) Try(ctx context.Context, req *contentv1.TCCArticlePublish_Req) (*contentv1.TCCArticlePublish_Resp, error) {
	if err := s.usecase.Try(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticlePublish_Resp{}, nil
}

func (s *TCCArticlePublishService) Confirm(ctx context.Context, req *contentv1.TCCArticlePublish_Req) (*contentv1.TCCArticlePublish_Resp, error) {
	if err := s.usecase.Confirm(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticlePublish_Resp{}, nil
}

func (s *TCCArticlePublishService) Cancel(ctx context.Context, req *contentv1.TCCArticlePublish_Req) (*contentv1.TCCArticlePublish_Resp, error) {
	if err := s.usecase.Cancel(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticlePublish_Resp{}, nil
}
