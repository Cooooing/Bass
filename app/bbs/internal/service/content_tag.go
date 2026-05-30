package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ContentTagService struct {
	bbscontentv1.UnimplementedTagServiceServer
	contentTagUsecase *usecase.ContentTagUsecase
}

func NewContentTagService(contentTagUsecase *usecase.ContentTagUsecase) *ContentTagService {
	return &ContentTagService{contentTagUsecase: contentTagUsecase}
}

func (s *ContentTagService) RegisterGrpc(gs *grpc.Server) {
	bbscontentv1.RegisterTagServiceServer(gs, s)
}

func (s *ContentTagService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterTagServiceHTTPServer(hs, s)
}

func (s *ContentTagService) List(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	return s.contentTagUsecase.ListTags(ctx, req)
}
