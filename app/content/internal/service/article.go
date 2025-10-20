package service

import (
	v1 "common/api/content/v1"
	"content/internal/biz"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ArticleService struct {
	v1.UnimplementedArticleServiceServer
	*BaseService

	articleDomain *biz.ArticleDomain
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHttp(hs *http.Server) {
	v1.RegisterArticleServiceHTTPServer(hs, s)
}

func NewArticleService(baseService *BaseService, articleDomain *biz.ArticleDomain) *ArticleService {
	return &ArticleService{
		BaseService:   baseService,
		articleDomain: articleDomain,
	}
}
