package service

import (
	"context"

	contentv1 "common/proto/gen/content/v1"

	"content/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// TCCArticleThankService 适配文章感谢的 TCC 分支协议。
type TCCArticleThankService struct {
	contentv1.UnimplementedTCCArticleThankServiceServer
	usecase *usecase.TCCArticleThankUsecase
}

func NewTCCArticleThankService(usecase *usecase.TCCArticleThankUsecase) *TCCArticleThankService {
	return &TCCArticleThankService{usecase: usecase}
}

func (s *TCCArticleThankService) RegisterGrpc(server *grpc.Server) {
	contentv1.RegisterTCCArticleThankServiceServer(server, s)
}

func (*TCCArticleThankService) RegisterHTTP(*http.Server) {}

func (s *TCCArticleThankService) Try(ctx context.Context, req *contentv1.TCCArticleThank_Req) (*contentv1.TCCArticleThank_Resp, error) {
	if err := s.usecase.Try(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticleThank_Resp{}, nil
}

func (s *TCCArticleThankService) Confirm(ctx context.Context, req *contentv1.TCCArticleThank_Req) (*contentv1.TCCArticleThank_Resp, error) {
	if err := s.usecase.Confirm(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticleThank_Resp{}, nil
}

func (s *TCCArticleThankService) Cancel(ctx context.Context, req *contentv1.TCCArticleThank_Req) (*contentv1.TCCArticleThank_Resp, error) {
	if err := s.usecase.Cancel(ctx, req); err != nil {
		return nil, err
	}
	return &contentv1.TCCArticleThank_Resp{}, nil
}
