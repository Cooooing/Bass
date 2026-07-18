package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
	"time"
)

type ArticleRepo interface {
	Save(ctx context.Context, article *model.Article) (*model.Article, error)
	Update(ctx context.Context, article *model.Article) (*model.Article, error)
	UpdatePublishStatus(ctx context.Context, req *ArticleUpdatePublishStatusReq) error
	UpdateVisibility(ctx context.Context, req *ArticleUpdateVisibilityReq) error
	UpdateRestriction(ctx context.Context, req *ArticleUpdateRestrictionReq) error
	DiscardDraft(ctx context.Context, articleID int64) error
	UpdateHasPostscript(ctx context.Context, req *ArticleUpdateHasPostscriptReq) error
	AddStats(ctx context.Context, req *ArticleAddStatsReq) error
	UpdateAcceptedAnswerID(ctx context.Context, req *ArticleUpdateAcceptedAnswerIDReq) (*model.Article, error)
	ReplaceTags(ctx context.Context, req *ArticleReplaceTagsReq) error
	Exist(ctx context.Context, req *ArticleGetReq) (bool, error)
	Get(ctx context.Context, req *ArticleGetReq) (*model.Article, error)
	List(ctx context.Context, req *ArticleGetReq) ([]*model.Article, error)
	Map(ctx context.Context, req *ArticleGetReq) (map[int64]*model.Article, error)
	Count(ctx context.Context, req *ArticleGetReq) (int, error)
	Page(ctx context.Context, req *ArticleGetReq) (*ArticlePageResp, error)
}

type ArticleUpdatePublishStatusReq struct {
	ArticleID     int64
	PublishStatus enum.ArticlePublishStatus
	Visibility    enum.ArticleVisibility
	PublishedAt   *time.Time
	UpdatedBy     int64
}

type ArticleUpdateVisibilityReq struct {
	ArticleID  int64
	Visibility enum.ArticleVisibility
	UpdatedBy  int64
}

type ArticleUpdateRestrictionReq struct {
	ArticleID   int64
	Restriction enum.ContentRestriction
	UpdatedBy   int64
}

type ArticleUpdateHasPostscriptReq struct {
	ArticleID     int64
	HasPostscript bool
	UpdatedBy     int64
}

type ArticleAddStatsReq struct {
	ArticleID int64
	Stats     ArticleStatUpdate
	UpdatedBy *int64
}

type ArticleUpdateAcceptedAnswerIDReq struct {
	ArticleID int64
	CommentID int64
	UpdatedBy int64
}

type ArticleReplaceTagsReq struct {
	ArticleID int64
	TagIDs    []int64
}

type ArticleStatUpdate struct {
	ViewCount    int32
	ThankCount   int32
	LikeCount    int32
	CollectCount int32
	WatchCount   int32
	ReplyCount   int32
}

type ArticlePageResp struct {
	Rows []*model.Article
	Page *base.PageResp
}

type ArticleGetReq struct {
	Page            *base.PageRequest
	ArticleId       *int64
	ArticleIds      []int64
	CreatedBy       *int64
	TagId           *int64
	DomainId        *int64
	PublishStatus   *enum.ArticlePublishStatus
	PublishStatuses []enum.ArticlePublishStatus
	Visibility      *enum.ArticleVisibility
	Visibilities    []enum.ArticleVisibility
	Restriction     *enum.ContentRestriction
	Restrictions    []enum.ContentRestriction
	AuthorId        *int64
	Order           *enum.ArticleOrder
	Type            *enum.ArticleType
	Keyword         *string
}
type ArticlePostscriptRepo interface {
	Save(ctx context.Context, articlePostscript *model.ArticlePostscript) (*model.ArticlePostscript, error)
	Get(ctx context.Context, req *ArticlePostscriptGetReq) (*model.ArticlePostscript, error)
	List(ctx context.Context, req *ArticlePostscriptGetReq) ([]*model.ArticlePostscript, error)
	Map(ctx context.Context, req *ArticlePostscriptGetReq) (map[int64]*model.ArticlePostscript, error)
	Count(ctx context.Context, req *ArticlePostscriptGetReq) (int, error)
	Page(ctx context.Context, req *ArticlePostscriptGetReq) (*ArticlePostscriptPageResp, error)
}

type ArticlePostscriptPageResp struct {
	Rows []*model.ArticlePostscript
	Page *base.PageResp
}

type ArticlePostscriptGetReq struct {
	Page        *base.PageRequest
	ID          *int64
	IDs         []int64
	ArticleID   *int64
	ArticleIDs  []int64
	CreatedBy   *int64
	Restriction *enum.ContentRestriction
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, record *model.ArticleActionRecord) (bool, error)
	Delete(ctx context.Context, req *ArticleActionRecordDeleteReq) (int, error)
	Exist(ctx context.Context, req *ArticleActionRecordReq) (bool, error)
	Get(ctx context.Context, req *ArticleActionRecordReq) (*model.ArticleActionRecord, error)
	List(ctx context.Context, req *ArticleActionRecordReq) ([]*model.ArticleActionRecord, error)
	Map(ctx context.Context, req *ArticleActionRecordReq) (map[int64]*model.ArticleActionRecord, error)
	Count(ctx context.Context, req *ArticleActionRecordReq) (int, error)
	Page(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordPageResp, error)
}

type ArticleActionRecordDeleteReq struct {
	ArticleID int64
	UserID    int64
	Action    enum.ArticleAction
}

type ArticleActionRecordPageResp struct {
	Rows []*model.ArticleActionRecord
	Page *base.PageResp
}

type ArticleActionRecordReq struct {
	Page       *base.PageRequest
	ID         *int64
	IDs        []int64
	ArticleId  *int64
	ArticleIds []int64
	UserId     *int64
	UserIds    []int64
	Type       *enum.ArticleAction
	Types      []enum.ArticleAction
}
