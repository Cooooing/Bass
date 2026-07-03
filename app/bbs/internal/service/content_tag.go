package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentTagService struct {
	bbscontentv1.UnimplementedTagServiceServer
	contentTagUsecase *usecase.ContentTagUsecase
}

func NewContentTagService(contentTagUsecase *usecase.ContentTagUsecase) *ContentTagService {
	return &ContentTagService{contentTagUsecase: contentTagUsecase}
}

func (s *ContentTagService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterTagServiceHTTPServer(hs, s)
}

func (s *ContentTagService) Create(ctx context.Context, req *bbscontentv1.CreateTag_Request) (*bbscontentv1.CreateTag_Reply, error) {
	return s.contentTagUsecase.CreateTag(ctx, req)
}

func (s *ContentTagService) Update(ctx context.Context, req *bbscontentv1.UpdateTag_Request) (*bbscontentv1.UpdateTag_Reply, error) {
	return s.contentTagUsecase.UpdateTag(ctx, req)
}

func (s *ContentTagService) List(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	return s.contentTagUsecase.ListTags(ctx, req)
}
