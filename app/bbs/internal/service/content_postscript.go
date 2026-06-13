package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ContentPostscriptService struct {
	bbscontentv1.UnimplementedPostscriptServiceServer
	contentPostscriptUsecase *usecase.ContentPostscriptUsecase
}

func NewContentPostscriptService(contentPostscriptUsecase *usecase.ContentPostscriptUsecase) *ContentPostscriptService {
	return &ContentPostscriptService{contentPostscriptUsecase: contentPostscriptUsecase}
}

func (s *ContentPostscriptService) RegisterGrpc(gs *grpc.Server) {
	bbscontentv1.RegisterPostscriptServiceServer(gs, s)
}

func (s *ContentPostscriptService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterPostscriptServiceHTTPServer(hs, s)
}

func (s *ContentPostscriptService) Add(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
	return s.contentPostscriptUsecase.AddPostscript(ctx, req)
}
