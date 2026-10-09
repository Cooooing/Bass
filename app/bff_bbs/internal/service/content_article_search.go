package service

import (
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/server"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bff_bbs/v1/content"
	"context"

	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentArticleSearchService struct {
	bbscontentv1.UnimplementedArticleSearchServiceServer
	contentArticleUsecase *usecase.ContentArticleUsecase
	articleService        *ContentArticleService
}

func NewContentArticleSearchService(
	contentArticleUsecase *usecase.ContentArticleUsecase,
	articleService *ContentArticleService,
) *ContentArticleSearchService {
	return &ContentArticleSearchService{
		contentArticleUsecase: contentArticleUsecase,
		articleService:        articleService,
	}
}

func (s *ContentArticleSearchService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentArticleSearchService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterArticleSearchServiceHTTPServer(hs, s)
}

func (s *ContentArticleSearchService) Search(
	ctx context.Context,
	req *bbscontentv1.SearchArticles_Req,
) (*bbscontentv1.SearchArticles_Resp, error) {
	var userID int64
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		userID = user.ID
	}
	result, err := s.contentArticleUsecase.SearchArticles(ctx, &usecase.SearchArticlesReq{
		UserID:   userID,
		ClientIP: server.ClientIP(ctx),
		Query:    req.GetQuery(),
		Cursor:   req.Cursor,
		Size:     req.GetSize(),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.ArticleListItem, 0, len(result.Rows))
	for _, row := range result.Rows {
		rows = append(rows, s.articleService.articleListItem(row))
	}
	return &bbscontentv1.SearchArticles_Resp{Rows: rows, NextCursor: result.NextCursor}, nil
}
