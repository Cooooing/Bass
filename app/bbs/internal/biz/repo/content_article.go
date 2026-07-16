package repo

import "context"

type ArticleViewerActionState struct {
	Liked     bool
	Thanked   bool
	Collected bool
	Watched   bool
}

type ArticlePostscript struct {
	ID            int64
	ArticleID     int64
	Content       string
	ContentRender string
	Restriction   int32
	CreatedBy     *int64
	UpdatedBy     *int64
	CreatedAt     string
	UpdatedAt     string
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
	BountyPoints      *int32
	AcceptedAnswerID  *int64
	Statement         *string
	Commentable       bool
	Anonymous         bool
	ViewCount         int32
	ThankCount        int32
	LikeCount         int32
	CollectCount      int32
	WatchCount        int32
	ReplyCount        int32
	CoverImageURL     *string
	ViewerActionState *ArticleViewerActionState
	LastReplyAt       string
	LastReplyUser     *AccountProfile
	AuthorUser        *AccountProfile
	CreatedBy         *int64
	UpdatedBy         *int64
	CreatedAt         string
	UpdatedAt         string
	PublishedAt       string
	EditedAt          string
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
	BountyPoints        *int32
	AcceptedAnswerID    *int64
	Statement           *string
	Commentable         bool
	Anonymous           bool
	ViewCount           int32
	ThankCount          int32
	LikeCount           int32
	CollectCount        int32
	WatchCount          int32
	ReplyCount          int32
	CoverImageURL       *string
	ViewerActionState   *ArticleViewerActionState
	LastReplyAt         string
	LastReplyUser       *AccountProfile
	Postscripts         []*ArticlePostscript
	AuthorUser          *AccountProfile
	CreatedBy           *int64
	UpdatedBy           *int64
	CreatedAt           string
	UpdatedAt           string
	PublishedAt         string
	EditedAt            string
}

type ArticleSave struct {
	Title         string
	Content       string
	RewardContent *string
	RewardPoints  *int32
	Type          int32
	BountyPoints  *int32
	Statement     *string
	Commentable   *bool
	Anonymous     *bool
	TagIDs        []int64
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

type CreateArticleReq struct {
	UserID  int64
	Article *ArticleSave
}

type CreateArticleResponse struct {
	Article *ArticleDetail
}

type UpdateArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ArticleSave
}

type UpdateArticleResponse struct {
	Article *ArticleDetail
}

type UpdateDraftArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ArticleSave
}

type UpdateDraftArticleResponse struct {
	Article *ArticleDetail
}

type PublishArticleReq struct {
	UserID     int64
	ArticleID  int64
	Visibility int32
}

type PublishArticleResponse struct{}

type DiscardDraftArticleReq struct {
	UserID    int64
	ArticleID int64
}

type DiscardDraftArticleResponse struct{}

type ListArticlesReq struct {
	UserID int64
	Page   *PageReq
	Query  *ArticleQuery
}

type ListArticlesResponse struct {
	Page *PageResponse
	Rows []*ArticleListItem
}

type GetArticleReq struct {
	UserID    int64
	ArticleID int64
}

type GetArticleResponse struct {
	Article *ArticleDetail
}

type ViewArticleReq struct {
	UserID    int64
	ArticleID int64
}

type ViewArticleResponse struct{}

type LikeArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type LikeArticleResponse struct {
	Liked bool
}

type ThankArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type ThankArticleResponse struct {
	Thanked bool
}

type CollectArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type CollectArticleResponse struct {
	Collected bool
}

type WatchArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type WatchArticleResponse struct {
	Watched bool
}

type RewardArticleReq struct {
	UserID    int64
	ArticleID int64
	Points    int32
}

type RewardArticleResponse struct{}

type AcceptAnswerArticleReq struct {
	UserID    int64
	ArticleID int64
	CommentID int64
}

type AcceptAnswerArticleResponse struct{}

type ContentArticleClient interface {
	CreateArticle(ctx context.Context, req *CreateArticleReq) (*CreateArticleResponse, error)
	UpdateArticle(ctx context.Context, req *UpdateArticleReq) (*UpdateArticleResponse, error)
	UpdateDraftArticle(ctx context.Context, req *UpdateDraftArticleReq) (*UpdateDraftArticleResponse, error)
	PublishArticle(ctx context.Context, req *PublishArticleReq) (*PublishArticleResponse, error)
	DiscardDraftArticle(ctx context.Context, req *DiscardDraftArticleReq) (*DiscardDraftArticleResponse, error)
	ListArticles(ctx context.Context, req *ListArticlesReq) (*ListArticlesResponse, error)
	GetArticle(ctx context.Context, req *GetArticleReq) (*GetArticleResponse, error)
	ViewArticle(ctx context.Context, req *ViewArticleReq) (*ViewArticleResponse, error)
	LikeArticle(ctx context.Context, req *LikeArticleReq) (*LikeArticleResponse, error)
	ThankArticle(ctx context.Context, req *ThankArticleReq) (*ThankArticleResponse, error)
	CollectArticle(ctx context.Context, req *CollectArticleReq) (*CollectArticleResponse, error)
	WatchArticle(ctx context.Context, req *WatchArticleReq) (*WatchArticleResponse, error)
	RewardArticle(ctx context.Context, req *RewardArticleReq) (*RewardArticleResponse, error)
	AcceptAnswerArticle(ctx context.Context, req *AcceptAnswerArticleReq) (*AcceptAnswerArticleResponse, error)
}
