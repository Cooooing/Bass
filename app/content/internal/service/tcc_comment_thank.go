package service

import (
	contentv1 "common/proto/gen/content/v1"
	"content/internal/biz/usecase"
	"context"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type TCCCommentThankService struct {
	contentv1.UnimplementedTCCCommentThankServiceServer
	usecase *usecase.TCCCommentThankUsecase
}

func NewTCCCommentThankService(usecase *usecase.TCCCommentThankUsecase) *TCCCommentThankService {
	return &TCCCommentThankService{usecase: usecase}
}
func (s *TCCCommentThankService) RegisterGrpc(server *grpc.Server) {
	contentv1.RegisterTCCCommentThankServiceServer(server, s)
}
func (*TCCCommentThankService) RegisterHTTP(*http.Server) {}
func (s *TCCCommentThankService) Try(ctx context.Context, r *contentv1.TCCCommentThank_Req) (*contentv1.TCCCommentThank_Resp, error) {
	if err := s.usecase.Try(ctx, r); err != nil {
		return nil, err
	}
	return &contentv1.TCCCommentThank_Resp{}, nil
}
func (s *TCCCommentThankService) Confirm(ctx context.Context, r *contentv1.TCCCommentThank_Req) (*contentv1.TCCCommentThank_Resp, error) {
	if err := s.usecase.Confirm(ctx, r); err != nil {
		return nil, err
	}
	return &contentv1.TCCCommentThank_Resp{}, nil
}
func (s *TCCCommentThankService) Cancel(ctx context.Context, r *contentv1.TCCCommentThank_Req) (*contentv1.TCCCommentThank_Resp, error) {
	if err := s.usecase.Cancel(ctx, r); err != nil {
		return nil, err
	}
	return &contentv1.TCCCommentThank_Resp{}, nil
}
