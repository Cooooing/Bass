package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type ContentPostscriptService struct {
	bbscontentv1.UnimplementedPostscriptServiceServer
	contentPostscriptUsecase *usecase.ContentPostscriptUsecase
}

func NewContentPostscriptService(contentPostscriptUsecase *usecase.ContentPostscriptUsecase) *ContentPostscriptService {
	return &ContentPostscriptService{contentPostscriptUsecase: contentPostscriptUsecase}
}
func (s *ContentPostscriptService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterPostscriptServiceHTTPServer(hs, s)
}
func (s *ContentPostscriptService) Add(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.contentPostscriptUsecase.AddPostscript(ctx, &usecase.AddPostscriptReq{UserID: userID, ArticleID: req.GetArticleId(), Content: req.GetContent()})
	if err != nil {
		return nil, err
	}
	var postscript *bbscontentv1.AddPostscript_Response_ArticlePostscript
	if response.Postscript != nil {
		postscript = &bbscontentv1.AddPostscript_Response_ArticlePostscript{Id: response.Postscript.ID, ArticleId: response.Postscript.ArticleID, Content: response.Postscript.Content, ContentRender: response.Postscript.ContentRender, Restriction: bbscontentv1.ContentRestriction(response.Postscript.Restriction), CreatedBy: response.Postscript.CreatedBy, UpdatedBy: response.Postscript.UpdatedBy, CreatedAt: response.Postscript.CreatedAt, UpdatedAt: response.Postscript.UpdatedAt}
	}
	return &bbscontentv1.AddPostscript_Response{Postscript: postscript}, nil
}
