package repo

import (
	commonmodel "common/pkg/model"
	"context"
	"time"

	"bff_bbs/internal/biz/model"
)

type ArticleSave struct {
	Title         string
	Content       string
	RewardContent *string
	RewardPoints  *int32
	Type          int32
	Statement     *string
	Commentable   *bool
}

type ArticleQuery struct {
	TagID             *int64
	DomainID          *int64
	AuthorID          *int64
	Type              *int32
	Order             *int32
	PublishStatus     *int32
	PublishStatuses   []int32
	Visibility        *int32
	Visibilities      []int32
	Scheduled         *bool
	ExcludedAuthorIDs []int64
}

type CreateDraftArticleReq struct {
	UserID  int64
	Article *ArticleSave
}

type UpdateDraftArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ArticleSave
}

type PublishArticleReq struct {
	UserID      int64
	ArticleID   int64
	ScheduledAt *time.Time
	City        *string
}

type CancelPublishArticleReq struct {
	UserID    int64
	ArticleID int64
}

type DiscardDraftArticleReq struct {
	UserID    int64
	ArticleID int64
}

type ArchiveArticleReq struct {
	UserID    int64
	ArticleID int64
	Reason    *string
}

type ListArticlesReq struct {
	UserID int64
	Page   *commonmodel.PageReq
	Query  *ArticleQuery
}

type ListArticlesResp struct {
	Page *commonmodel.PageResp
	Rows []*model.ArticleListItem
}

type SearchArticlesReq struct {
	UserID            int64
	ClientIP          string
	Query             string
	Cursor            *string
	Size              uint32
	ExcludedAuthorIDs []int64
}

type SearchArticlesResp struct {
	Rows       []*model.ArticleListItem
	NextCursor *string
}

type ArticleSearchRateLimitCache interface {
	Allow(context.Context, *ArticleSearchRateLimitReq) (bool, error)
}

type ArticleSearchRateLimitReq struct {
	Key    string
	Window time.Duration
	Limit  int64
}

type ArticleViewHistoryRow struct {
	Article  *model.ArticleListItem
	ViewedAt *time.Time
}

type ArticleViewHistoryPageResp struct {
	Rows []*ArticleViewHistoryRow
	Page *commonmodel.PageResp
}

type GetArticleReq struct {
	UserID    int64
	ArticleID int64
}

type ViewArticleReq struct {
	UserID    int64
	ArticleID int64
	IP        *string
	UserAgent *string
}

type LikeArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type ThankArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type CollectArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type RewardArticleReq struct {
	UserID    int64
	ArticleID int64
	Points    int32
}

type ContentArticleClient interface {
	CreateDraftArticle(ctx context.Context, req *CreateDraftArticleReq) (int64, error)
	UpdateDraftArticle(ctx context.Context, req *UpdateDraftArticleReq) (int64, error)
	PublishArticle(ctx context.Context, req *PublishArticleReq) error
	CancelPublishArticle(ctx context.Context, req *CancelPublishArticleReq) error
	DiscardDraftArticle(ctx context.Context, req *DiscardDraftArticleReq) error
	ArchiveArticle(ctx context.Context, req *ArchiveArticleReq) error
	ListArticles(ctx context.Context, req *ListArticlesReq) (*ListArticlesResp, error)
	SearchArticles(ctx context.Context, req *SearchArticlesReq) (*SearchArticlesResp, error)
	PageViewHistory(ctx context.Context, userID int64, page *commonmodel.PageReq, excludedAuthorIDs []int64) (*ArticleViewHistoryPageResp, error)
	GetArticle(ctx context.Context, req *GetArticleReq) (*model.ArticleDetail, error)
	ViewArticle(ctx context.Context, req *ViewArticleReq) error
	LikeArticle(ctx context.Context, req *LikeArticleReq) (bool, error)
	ThankArticle(ctx context.Context, req *ThankArticleReq) (bool, error)
	CollectArticle(ctx context.Context, req *CollectArticleReq) (bool, error)
	RewardArticle(ctx context.Context, req *RewardArticleReq) error
}
