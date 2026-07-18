package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentPostscriptService struct {
	bbscontentv1.UnimplementedPostscriptServiceServer
	contentPostscriptUsecase *usecase.ContentPostscriptUsecase
}

func NewContentPostscriptService(contentPostscriptUsecase *usecase.ContentPostscriptUsecase) *ContentPostscriptService {
	return &ContentPostscriptService{contentPostscriptUsecase: contentPostscriptUsecase}
}

func (s *ContentPostscriptService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentPostscriptService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterPostscriptServiceHTTPServer(hs, s)
}

func (s *ContentPostscriptService) Add(ctx context.Context, req *bbscontentv1.AddPostscript_Req) (*bbscontentv1.AddPostscript_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentPostscriptUsecase.AddPostscript(ctx, &usecase.AddPostscriptReq{UserID: userID, ArticleID: req.GetArticleId(), Content: req.GetContent()})
	if err != nil {
		return nil, err
	}
	var postscript *bbscontentv1.AddPostscript_Resp_ArticlePostscript
	if resp != nil {
		postscript = &bbscontentv1.AddPostscript_Resp_ArticlePostscript{Id: resp.ID, ArticleId: resp.ArticleID, Content: resp.Content, ContentRender: resp.ContentRender, Restriction: bbscontentv1.ContentRestriction(resp.Restriction), CreatedBy: resp.CreatedBy, UpdatedBy: resp.UpdatedBy, CreatedAt: resp.CreatedAt, UpdatedAt: resp.UpdatedAt}
	}
	return &bbscontentv1.AddPostscript_Resp{Postscript: postscript}, nil
}
