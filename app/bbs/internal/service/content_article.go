package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ContentArticleService struct {
	bbscontentv1.UnimplementedArticleServiceServer
	contentArticleUsecase *usecase.ContentArticleUsecase
}

func NewContentArticleService(contentArticleUsecase *usecase.ContentArticleUsecase) *ContentArticleService {
	return &ContentArticleService{contentArticleUsecase: contentArticleUsecase}
}

func (s *ContentArticleService) RegisterGrpc(gs *grpc.Server) {
	bbscontentv1.RegisterArticleServiceServer(gs, s)
}

func (s *ContentArticleService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterArticleServiceHTTPServer(hs, s)
}

func (s *ContentArticleService) Create(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
	return s.contentArticleUsecase.CreateArticle(ctx, req)
}

func (s *ContentArticleService) UpdateDraft(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
	return s.contentArticleUsecase.UpdateDraftArticle(ctx, req)
}

func (s *ContentArticleService) Publish(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
	return s.contentArticleUsecase.PublishArticle(ctx, req)
}

func (s *ContentArticleService) Delete(ctx context.Context, req *bbscontentv1.DeleteArticle_Request) (*bbscontentv1.DeleteArticle_Reply, error) {
	return s.contentArticleUsecase.DeleteArticle(ctx, req)
}

func (s *ContentArticleService) List(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
	return s.contentArticleUsecase.ListArticles(ctx, req)
}

func (s *ContentArticleService) Get(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
	return s.contentArticleUsecase.GetArticle(ctx, req)
}

func (s *ContentArticleService) Like(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
	return s.contentArticleUsecase.LikeArticle(ctx, req)
}

func (s *ContentArticleService) Thank(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
	return s.contentArticleUsecase.ThankArticle(ctx, req)
}

func (s *ContentArticleService) Collect(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
	return s.contentArticleUsecase.CollectArticle(ctx, req)
}

func (s *ContentArticleService) Watch(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
	return s.contentArticleUsecase.WatchArticle(ctx, req)
}

func (s *ContentArticleService) Reward(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
	return s.contentArticleUsecase.RewardArticle(ctx, req)
}

func (s *ContentArticleService) AcceptAnswer(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
	return s.contentArticleUsecase.AcceptAnswerArticle(ctx, req)
}
