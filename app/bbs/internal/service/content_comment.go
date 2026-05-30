package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ContentCommentService struct {
	bbscontentv1.UnimplementedCommentServiceServer
	contentCommentUsecase *usecase.ContentCommentUsecase
}

func NewContentCommentService(contentCommentUsecase *usecase.ContentCommentUsecase) *ContentCommentService {
	return &ContentCommentService{contentCommentUsecase: contentCommentUsecase}
}

func (s *ContentCommentService) RegisterGrpc(gs *grpc.Server) {
	bbscontentv1.RegisterCommentServiceServer(gs, s)
}

func (s *ContentCommentService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterCommentServiceHTTPServer(hs, s)
}

func (s *ContentCommentService) Create(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	return s.contentCommentUsecase.CreateComment(ctx, req)
}

func (s *ContentCommentService) List(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	return s.contentCommentUsecase.ListComments(ctx, req)
}

func (s *ContentCommentService) Like(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	return s.contentCommentUsecase.LikeComment(ctx, req)
}

func (s *ContentCommentService) Thank(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	return s.contentCommentUsecase.ThankComment(ctx, req)
}
