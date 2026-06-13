package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type ContentArticleUsecase struct {
	contentArticleClient repo.ContentArticleClient
}

func NewContentArticleUsecase(contentArticleClient repo.ContentArticleClient) *ContentArticleUsecase {
	return &ContentArticleUsecase{contentArticleClient: contentArticleClient}
}

func (u *ContentArticleUsecase) CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error) {
	article := req.GetArticle()
	if article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	switch article.GetType() {
	case bbscontentv1.ArticleType_ARTICLE_TYPE_NORMAL, bbscontentv1.ArticleType_ARTICLE_TYPE_QA:
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	if article.GetType() != bbscontentv1.ArticleType_ARTICLE_TYPE_QA && article.BountyPoints != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	return u.contentArticleClient.CreateArticle(ctx, req)
}

func (u *ContentArticleUsecase) UpdateArticle(ctx context.Context, req *bbscontentv1.UpdateArticle_Request) (*bbscontentv1.UpdateArticle_Reply, error) {
	article := req.GetArticle()
	if article == nil || req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	switch article.GetType() {
	case bbscontentv1.ArticleType_ARTICLE_TYPE_NORMAL, bbscontentv1.ArticleType_ARTICLE_TYPE_QA:
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	if article.GetType() != bbscontentv1.ArticleType_ARTICLE_TYPE_QA && article.BountyPoints != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	return u.contentArticleClient.UpdateArticle(ctx, req)
}

func (u *ContentArticleUsecase) UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error) {
	article := req.GetArticle()
	update := &bbscontentv1.UpdateArticle_Request{ArticleId: req.GetArticleId()}
	if article != nil {
		update.Article = &bbscontentv1.UpdateArticle_Request_Article{
			Title:         article.GetTitle(),
			Content:       article.GetContent(),
			RewardContent: article.RewardContent,
			RewardPoints:  article.RewardPoints,
			Type:          bbscontentv1.ArticleType(article.GetType()),
			BountyPoints:  article.BountyPoints,
			Statement:     article.Statement,
			Commentable:   article.Commentable,
			Anonymous:     article.Anonymous,
			TagIds:        article.GetTagIds(),
		}
	}
	reply, err := u.UpdateArticle(ctx, update)
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.UpdateDraftArticle_Reply{Article: reply.GetArticle()}, nil
}

func (u *ContentArticleUsecase) PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error) {
	switch req.GetVisibility() {
	case bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_UNSPECIFIED:
		req.Visibility = bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC
	case bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC, bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_PRIVATE:
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
	}
	return u.contentArticleClient.PublishArticle(ctx, req)
}

func (u *ContentArticleUsecase) DiscardDraftArticle(ctx context.Context, req *bbscontentv1.DiscardDraftArticle_Request) (*bbscontentv1.DiscardDraftArticle_Reply, error) {
	return u.contentArticleClient.DiscardDraftArticle(ctx, req)
}

func (u *ContentArticleUsecase) ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error) {
	return u.contentArticleClient.ListArticles(ctx, req)
}

func (u *ContentArticleUsecase) GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error) {
	return u.contentArticleClient.GetArticle(ctx, req)
}

func (u *ContentArticleUsecase) ViewArticle(ctx context.Context, articleID int64) error {
	return u.contentArticleClient.ViewArticle(ctx, articleID)
}

func (u *ContentArticleUsecase) LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error) {
	return u.contentArticleClient.LikeArticle(ctx, req)
}

func (u *ContentArticleUsecase) ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error) {
	return u.contentArticleClient.ThankArticle(ctx, req)
}

func (u *ContentArticleUsecase) CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error) {
	return u.contentArticleClient.CollectArticle(ctx, req)
}

func (u *ContentArticleUsecase) WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error) {
	return u.contentArticleClient.WatchArticle(ctx, req)
}

func (u *ContentArticleUsecase) RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error) {
	return u.contentArticleClient.RewardArticle(ctx, req)
}

func (u *ContentArticleUsecase) AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error) {
	return u.contentArticleClient.AcceptAnswerArticle(ctx, req)
}
