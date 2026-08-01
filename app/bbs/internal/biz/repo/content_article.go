package repo

import (
	"context"
	"time"
)

type ArticleViewerActionState struct {
	Liked     bool
	Thanked   bool
	Collected bool
	Rewarded  bool
}

type ArticlePostscript struct {
	ID            int64
	ArticleID     int64
	Content       string
	ContentRender string
	Restriction   int32
	CreatedBy     *int64
	UpdatedBy     *int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

type ArticleListItem struct {
	ID                int64
	Title             string
	Content           string
	ContentRender     string
	HasPostscript     bool
	HasReward         bool
	PublishStatus     int32
	Visibility        int32
	Restriction       int32
	Type              int32
	Statement         *string
	Commentable       bool
	ViewCount         int32
	ThankCount        int32
	LikeCount         int32
	CollectCount      int32
	RewardCount       int32
	ReplyCount        int32
	CoverImageURL     *string
	ViewerActionState *ArticleViewerActionState
	LastReplyAt       *time.Time
	LastReplyUser     *AccountProfile
	AuthorUser        *AccountProfile
	CreatedBy         *int64
	UpdatedBy         *int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	PublishedAt       *time.Time
	EditedAt          *time.Time
}

type ArticleDetail struct {
	ID                  int64
	Title               string
	Content             string
	ContentRender       string
	HasPostscript       bool
	HasReward           bool
	RewardContent       *string
	RewardContentRender *string
	RewardPoints        *int32
	PublishStatus       int32
	Visibility          int32
	Restriction         int32
	Type                int32
	Statement           *string
	Commentable         bool
	ViewCount           int32
	ThankCount          int32
	LikeCount           int32
	CollectCount        int32
	RewardCount         int32
	ReplyCount          int32
	CoverImageURL       *string
	ViewerActionState   *ArticleViewerActionState
	LastReplyAt         *time.Time
	LastReplyUser       *AccountProfile
	Postscripts         []*ArticlePostscript
	AuthorUser          *AccountProfile
	CreatedBy           *int64
	UpdatedBy           *int64
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	PublishedAt         *time.Time
	EditedAt            *time.Time
}

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
	TagID           *int64
	DomainID        *int64
	Keyword         *string
	AuthorID        *int64
	Type            *int32
	Order           *int32
	PublishStatus   *int32
	PublishStatuses []int32
	Visibility      *int32
	Visibilities    []int32
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
	UserID     int64
	ArticleID  int64
	Visibility int32
}

type SchedulePublishArticleReq struct {
	UserID    int64
	ArticleID int64
	PublishAt time.Time
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
	Page   *PageReq
	Query  *ArticleQuery
}

type ListArticlesResp struct {
	Page *PageResp
	Rows []*ArticleListItem
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
	CreateDraftArticle(ctx context.Context, req *CreateDraftArticleReq) (*ArticleDetail, error)
	UpdateDraftArticle(ctx context.Context, req *UpdateDraftArticleReq) (*ArticleDetail, error)
	PublishArticle(ctx context.Context, req *PublishArticleReq) error
	SchedulePublishArticle(ctx context.Context, req *SchedulePublishArticleReq) error
	CancelPublishArticle(ctx context.Context, req *CancelPublishArticleReq) error
	DiscardDraftArticle(ctx context.Context, req *DiscardDraftArticleReq) error
	ArchiveArticle(ctx context.Context, req *ArchiveArticleReq) error
	ListArticles(ctx context.Context, req *ListArticlesReq) (*ListArticlesResp, error)
	GetArticle(ctx context.Context, req *GetArticleReq) (*ArticleDetail, error)
	ViewArticle(ctx context.Context, req *ViewArticleReq) error
	LikeArticle(ctx context.Context, req *LikeArticleReq) (bool, error)
	ThankArticle(ctx context.Context, req *ThankArticleReq) (bool, error)
	CollectArticle(ctx context.Context, req *CollectArticleReq) (bool, error)
	RewardArticle(ctx context.Context, req *RewardArticleReq) error
}
