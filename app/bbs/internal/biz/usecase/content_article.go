package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentArticleUsecase struct {
	contentArticleRepo repo.ContentArticleRepo
}

func NewContentArticleUsecase(contentArticleRepo repo.ContentArticleRepo) *ContentArticleUsecase {
	return &ContentArticleUsecase{contentArticleRepo: contentArticleRepo}
}

func (u *ContentArticleUsecase) CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
	return u.contentArticleRepo.CreateArticle(ctx, req)
}

func (u *ContentArticleUsecase) UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
	return u.contentArticleRepo.UpdateDraftArticle(ctx, req)
}

func (u *ContentArticleUsecase) PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
	return u.contentArticleRepo.PublishArticle(ctx, req)
}

func (u *ContentArticleUsecase) DeleteArticle(ctx context.Context, req *bbscontentv1.DeleteArticle_Request) (*bbscontentv1.DeleteArticle_Reply, error) {
	return u.contentArticleRepo.DeleteArticle(ctx, req)
}

func (u *ContentArticleUsecase) ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
	return u.contentArticleRepo.ListArticles(ctx, req)
}

func (u *ContentArticleUsecase) GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
	return u.contentArticleRepo.GetArticle(ctx, req)
}

func (u *ContentArticleUsecase) LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
	return u.contentArticleRepo.LikeArticle(ctx, req)
}

func (u *ContentArticleUsecase) ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
	return u.contentArticleRepo.ThankArticle(ctx, req)
}

func (u *ContentArticleUsecase) CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
	return u.contentArticleRepo.CollectArticle(ctx, req)
}

func (u *ContentArticleUsecase) WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
	return u.contentArticleRepo.WatchArticle(ctx, req)
}

func (u *ContentArticleUsecase) RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
	return u.contentArticleRepo.RewardArticle(ctx, req)
}

func (u *ContentArticleUsecase) AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
	return u.contentArticleRepo.AcceptAnswerArticle(ctx, req)
}
