package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type ContentArticleUsecase struct {
	contentArticleClient repo.ContentArticleClient
}

func NewContentArticleUsecase(contentArticleClient repo.ContentArticleClient) *ContentArticleUsecase {
	return &ContentArticleUsecase{contentArticleClient: contentArticleClient}
}

type ContentArticleViewerActionState = repo.ArticleViewerActionState

type ContentArticlePostscript = repo.ArticlePostscript

type ContentArticleListItem = repo.ArticleListItem

type ContentArticleDetail = repo.ArticleDetail

type ContentAccountProfile = repo.AccountProfile

type ContentPageResponse = repo.PageResponse
type ContentArticleSave struct {
	Title         string
	Content       string
	RewardContent *string
	RewardPoints  *int32
	Type          bbscontentv1.ArticleType
	BountyPoints  *int32
	Statement     *string
	Commentable   *bool
	Anonymous     *bool
	TagIDs        []int64
}

type CreateArticleReq struct {
	UserID  int64
	Article *ContentArticleSave
}

type CreateArticleResponse struct {
	Article *ContentArticleDetail
}

func (u *ContentArticleUsecase) CreateArticle(ctx context.Context, req *CreateArticleReq) (*CreateArticleResponse, error) {
	if req == nil || req.Article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	article := req.Article
	if err := validateArticleType(article.Type, article.BountyPoints); err != nil {
		return nil, err
	}
	response, err := u.contentArticleClient.CreateArticle(ctx, &repo.CreateArticleReq{
		UserID: req.UserID,
		Article: &repo.ArticleSave{
			Title:         article.Title,
			Content:       article.Content,
			RewardContent: article.RewardContent,
			RewardPoints:  article.RewardPoints,
			Type:          int32(article.Type),
			BountyPoints:  article.BountyPoints,
			Statement:     article.Statement,
			Commentable:   article.Commentable,
			Anonymous:     article.Anonymous,
			TagIDs:        article.TagIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	return &CreateArticleResponse{Article: response.Article}, nil
}

type UpdateArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ContentArticleSave
}

type UpdateArticleResponse struct {
	Article *ContentArticleDetail
}

func (u *ContentArticleUsecase) UpdateArticle(ctx context.Context, req *UpdateArticleReq) (*UpdateArticleResponse, error) {
	if req == nil || req.Article == nil || req.ArticleID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	article := req.Article
	if err := validateArticleType(article.Type, article.BountyPoints); err != nil {
		return nil, err
	}
	response, err := u.contentArticleClient.UpdateArticle(ctx, &repo.UpdateArticleReq{
		UserID:    req.UserID,
		ArticleID: req.ArticleID,
		Article: &repo.ArticleSave{
			Title:         article.Title,
			Content:       article.Content,
			RewardContent: article.RewardContent,
			RewardPoints:  article.RewardPoints,
			Type:          int32(article.Type),
			BountyPoints:  article.BountyPoints,
			Statement:     article.Statement,
			Commentable:   article.Commentable,
			Anonymous:     article.Anonymous,
			TagIDs:        article.TagIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	return &UpdateArticleResponse{Article: response.Article}, nil
}

type UpdateDraftArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ContentArticleSave
}

type UpdateDraftArticleResponse struct {
	Article *ContentArticleDetail
}

func (u *ContentArticleUsecase) UpdateDraftArticle(ctx context.Context, req *UpdateDraftArticleReq) (*UpdateDraftArticleResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	reply, err := u.UpdateArticle(ctx, &UpdateArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Article: req.Article})
	if err != nil {
		return nil, err
	}
	return &UpdateDraftArticleResponse{Article: reply.Article}, nil
}

type PublishArticleReq struct {
	UserID     int64
	ArticleID  int64
	Visibility bbscontentv1.ArticleVisibility
}

func (u *ContentArticleUsecase) PublishArticle(ctx context.Context, req *PublishArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	visibility := req.Visibility
	switch visibility {
	case bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_UNSPECIFIED:
		visibility = bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC
	case bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC, bbscontentv1.ArticleVisibility_ARTICLE_VISIBILITY_PRIVATE:
	default:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
	}
	_, err := u.contentArticleClient.PublishArticle(ctx, &repo.PublishArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Visibility: int32(visibility)})
	return err
}

type DiscardDraftArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) DiscardDraftArticle(ctx context.Context, req *DiscardDraftArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := u.contentArticleClient.DiscardDraftArticle(ctx, &repo.DiscardDraftArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
	return err
}

type ListArticlesReq struct {
	UserID int64
	Page   *common.PageRequest
	Query  *bbscontentv1.ListArticles_Request_ArticleQuery
}

type ListArticlesResponse struct {
	Page *ContentPageResponse
	Rows []*ContentArticleListItem
}

func (u *ContentArticleUsecase) ListArticles(ctx context.Context, req *ListArticlesReq) (*ListArticlesResponse, error) {
	if req == nil {
		req = &ListArticlesReq{}
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	var query *repo.ArticleQuery
	if req.Query != nil {
		query = &repo.ArticleQuery{
			TagID:    req.Query.TagId,
			DomainID: req.Query.DomainId,
			Keyword:  req.Query.Keyword,
			AuthorID: req.Query.AuthorId,
		}
		if req.Query.Type != nil {
			value := int32(*req.Query.Type)
			query.Type = &value
		}
		if req.Query.Order != nil {
			value := int32(*req.Query.Order)
			query.Order = &value
		}
		if req.Query.PublishStatus != nil {
			value := int32(*req.Query.PublishStatus)
			query.PublishStatus = &value
		}
		for _, item := range req.Query.PublishStatuses {
			query.PublishStatuses = append(query.PublishStatuses, int32(item))
		}
		if req.Query.Visibility != nil {
			value := int32(*req.Query.Visibility)
			query.Visibility = &value
		}
		for _, item := range req.Query.Visibilities {
			query.Visibilities = append(query.Visibilities, int32(item))
		}
	}
	response, err := u.contentArticleClient.ListArticles(ctx, &repo.ListArticlesReq{UserID: req.UserID, Page: page, Query: query})
	if err != nil {
		return nil, err
	}
	return &ListArticlesResponse{Page: response.Page, Rows: response.Rows}, nil
}

type GetArticleReq struct {
	UserID    int64
	ArticleID int64
}

type GetArticleResponse struct {
	Article *ContentArticleDetail
}

func (u *ContentArticleUsecase) GetArticle(ctx context.Context, req *GetArticleReq) (*GetArticleResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	response, err := u.contentArticleClient.GetArticle(ctx, &repo.GetArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
	if err != nil {
		return nil, err
	}
	return &GetArticleResponse{Article: response.Article}, nil
}

type ViewArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) ViewArticle(ctx context.Context, req *ViewArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := u.contentArticleClient.ViewArticle(ctx, &repo.ViewArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
	return err
}

type LikeArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type LikeArticleResponse struct {
	Liked bool
}

func (u *ContentArticleUsecase) LikeArticle(ctx context.Context, req *LikeArticleReq) (*LikeArticleResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	response, err := u.contentArticleClient.LikeArticle(ctx, &repo.LikeArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
	if err != nil {
		return nil, err
	}
	return &LikeArticleResponse{Liked: response.Liked}, nil
}

type ThankArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type ThankArticleResponse struct {
	Thanked bool
}

func (u *ContentArticleUsecase) ThankArticle(ctx context.Context, req *ThankArticleReq) (*ThankArticleResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	response, err := u.contentArticleClient.ThankArticle(ctx, &repo.ThankArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
	if err != nil {
		return nil, err
	}
	return &ThankArticleResponse{Thanked: response.Thanked}, nil
}

type CollectArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type CollectArticleResponse struct {
	Collected bool
}

func (u *ContentArticleUsecase) CollectArticle(ctx context.Context, req *CollectArticleReq) (*CollectArticleResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	response, err := u.contentArticleClient.CollectArticle(ctx, &repo.CollectArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
	if err != nil {
		return nil, err
	}
	return &CollectArticleResponse{Collected: response.Collected}, nil
}

type WatchArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type WatchArticleResponse struct {
	Watched bool
}

func (u *ContentArticleUsecase) WatchArticle(ctx context.Context, req *WatchArticleReq) (*WatchArticleResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	response, err := u.contentArticleClient.WatchArticle(ctx, &repo.WatchArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
	if err != nil {
		return nil, err
	}
	return &WatchArticleResponse{Watched: response.Watched}, nil
}

type RewardArticleReq struct {
	UserID    int64
	ArticleID int64
	Points    int32
}

func (u *ContentArticleUsecase) RewardArticle(ctx context.Context, req *RewardArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := u.contentArticleClient.RewardArticle(ctx, &repo.RewardArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Points: req.Points})
	return err
}

type AcceptAnswerArticleReq struct {
	UserID    int64
	ArticleID int64
	CommentID int64
}

func (u *ContentArticleUsecase) AcceptAnswerArticle(ctx context.Context, req *AcceptAnswerArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := u.contentArticleClient.AcceptAnswerArticle(ctx, &repo.AcceptAnswerArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, CommentID: req.CommentID})
	return err
}

func validateArticleType(articleType bbscontentv1.ArticleType, bountyPoints *int32) error {
	switch articleType {
	case bbscontentv1.ArticleType_ARTICLE_TYPE_NORMAL, bbscontentv1.ArticleType_ARTICLE_TYPE_QA:
	default:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	if articleType != bbscontentv1.ArticleType_ARTICLE_TYPE_QA && bountyPoints != nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	return nil
}
