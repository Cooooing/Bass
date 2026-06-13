package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ContentDomainService struct {
	bbscontentv1.UnimplementedDomainServiceServer
	contentDomainUsecase *usecase.ContentDomainUsecase
}

func NewContentDomainService(contentDomainUsecase *usecase.ContentDomainUsecase) *ContentDomainService {
	return &ContentDomainService{contentDomainUsecase: contentDomainUsecase}
}

func (s *ContentDomainService) RegisterGrpc(gs *grpc.Server) {
	bbscontentv1.RegisterDomainServiceServer(gs, s)
}

func (s *ContentDomainService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterDomainServiceHTTPServer(hs, s)
}

func (s *ContentDomainService) List(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error) {
	return s.contentDomainUsecase.ListDomains(ctx, req)
}
