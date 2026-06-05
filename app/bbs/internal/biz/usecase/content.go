package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentUsecase struct {
	contentRepo repo.ContentRepo
}

func NewContentUsecase(contentRepo repo.ContentRepo) *ContentUsecase {
	return &ContentUsecase{contentRepo: contentRepo}
}

func (u *ContentUsecase) CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
	return u.contentRepo.CreateArticle(ctx, req)
}

func (u *ContentUsecase) UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
	return u.contentRepo.UpdateDraftArticle(ctx, req)
}

func (u *ContentUsecase) PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
	return u.contentRepo.PublishArticle(ctx, req)
}

func (u *ContentUsecase) DeleteArticle(ctx context.Context, req *bbscontentv1.DeleteArticle_Request) (*bbscontentv1.DeleteArticle_Reply, error) {
	return u.contentRepo.DeleteArticle(ctx, req)
}

func (u *ContentUsecase) ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
	return u.contentRepo.ListArticles(ctx, req)
}

func (u *ContentUsecase) GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
	return u.contentRepo.GetArticle(ctx, req)
}

func (u *ContentUsecase) AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
	return u.contentRepo.AddPostscript(ctx, req)
}

func (u *ContentUsecase) LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
	return u.contentRepo.LikeArticle(ctx, req)
}

func (u *ContentUsecase) ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
	return u.contentRepo.ThankArticle(ctx, req)
}

func (u *ContentUsecase) CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
	return u.contentRepo.CollectArticle(ctx, req)
}

func (u *ContentUsecase) WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
	return u.contentRepo.WatchArticle(ctx, req)
}

func (u *ContentUsecase) RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
	return u.contentRepo.RewardArticle(ctx, req)
}

func (u *ContentUsecase) AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
	return u.contentRepo.AcceptAnswerArticle(ctx, req)
}

func (u *ContentUsecase) CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	return u.contentRepo.CreateComment(ctx, req)
}

func (u *ContentUsecase) ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	return u.contentRepo.ListComments(ctx, req)
}

func (u *ContentUsecase) LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	return u.contentRepo.LikeComment(ctx, req)
}

func (u *ContentUsecase) ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	return u.contentRepo.ThankComment(ctx, req)
}

func (u *ContentUsecase) ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error) {
	return u.contentRepo.ListDomains(ctx, req)
}

func (u *ContentUsecase) ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	return u.contentRepo.ListTags(ctx, req)
}
