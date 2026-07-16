package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
	"time"
)

type ArticleRepo interface {
	Save(ctx context.Context, req *ArticleSaveReq) (*ArticleSaveResponse, error)
	Update(ctx context.Context, req *ArticleUpdateReq) (*ArticleUpdateResponse, error)
	UpdatePublishStatus(ctx context.Context, req *ArticleUpdatePublishStatusReq) (*ArticleUpdatePublishStatusResponse, error)
	UpdateVisibility(ctx context.Context, req *ArticleUpdateVisibilityReq) (*ArticleUpdateVisibilityResponse, error)
	UpdateRestriction(ctx context.Context, req *ArticleUpdateRestrictionReq) (*ArticleUpdateRestrictionResponse, error)
	DiscardDraft(ctx context.Context, req *ArticleDiscardDraftReq) (*ArticleDiscardDraftResponse, error)
	UpdateHasPostscript(ctx context.Context, req *ArticleUpdateHasPostscriptReq) (*ArticleUpdateHasPostscriptResponse, error)
	AddStats(ctx context.Context, req *ArticleAddStatsReq) (*ArticleAddStatsResponse, error)
	UpdateAcceptedAnswerID(ctx context.Context, req *ArticleUpdateAcceptedAnswerIDReq) (*ArticleUpdateAcceptedAnswerIDResponse, error)
	ReplaceTags(ctx context.Context, req *ArticleReplaceTagsReq) (*ArticleReplaceTagsResponse, error)
	Exist(ctx context.Context, req *ArticleGetReq) (*ArticleExistResponse, error)
	Get(ctx context.Context, req *ArticleGetReq) (*ArticleGetResponse, error)
	List(ctx context.Context, req *ArticleGetReq) (*ArticleListResponse, error)
	Map(ctx context.Context, req *ArticleGetReq) (*ArticleMapResponse, error)
	Count(ctx context.Context, req *ArticleGetReq) (*ArticleCountResponse, error)
	Page(ctx context.Context, req *ArticleGetReq) (*ArticlePageResponse, error)
}

type ArticleSaveReq struct {
	Article *model.Article
}

type ArticleSaveResponse struct {
	Article *model.Article
}

type ArticleUpdateReq struct {
	Article *model.Article
}

type ArticleUpdateResponse struct {
	Article *model.Article
}

type ArticleUpdatePublishStatusReq struct {
	ArticleID     int64
	PublishStatus enum.ArticlePublishStatus
	Visibility    enum.ArticleVisibility
	PublishedAt   *time.Time
	UpdatedBy     int64
}

type ArticleUpdatePublishStatusResponse struct{}

type ArticleUpdateVisibilityReq struct {
	ArticleID  int64
	Visibility enum.ArticleVisibility
	UpdatedBy  int64
}

type ArticleUpdateVisibilityResponse struct{}

type ArticleUpdateRestrictionReq struct {
	ArticleID   int64
	Restriction enum.ContentRestriction
	UpdatedBy   int64
}

type ArticleUpdateRestrictionResponse struct{}

type ArticleDiscardDraftReq struct {
	ArticleID int64
}

type ArticleDiscardDraftResponse struct{}

type ArticleUpdateHasPostscriptReq struct {
	ArticleID     int64
	HasPostscript bool
	UpdatedBy     int64
}

type ArticleUpdateHasPostscriptResponse struct{}

type ArticleAddStatsReq struct {
	ArticleID int64
	Stats     ArticleStatUpdate
	UpdatedBy *int64
}

type ArticleAddStatsResponse struct{}

type ArticleUpdateAcceptedAnswerIDReq struct {
	ArticleID int64
	CommentID int64
	UpdatedBy int64
}

type ArticleUpdateAcceptedAnswerIDResponse struct {
	Article *model.Article
}

type ArticleReplaceTagsReq struct {
	ArticleID int64
	TagIDs    []int64
}

type ArticleReplaceTagsResponse struct{}

type ArticleStatUpdate struct {
	ViewCount    int32
	ThankCount   int32
	LikeCount    int32
	CollectCount int32
	WatchCount   int32
	ReplyCount   int32
}

type ArticleExistResponse struct {
	Exist bool
}

type ArticleGetResponse struct {
	Article *model.Article
}

type ArticleListResponse struct {
	Rows []*model.Article
}

type ArticleMapResponse struct {
	Rows map[int64]*model.Article
}

type ArticleCountResponse struct {
	Count int
}

type ArticlePageResponse struct {
	Rows []*model.Article
	Page *base.PageResponse
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
	Save(ctx context.Context, req *ArticlePostscriptSaveReq) (*ArticlePostscriptSaveResponse, error)
	Get(ctx context.Context, req *ArticlePostscriptGetReq) (*ArticlePostscriptGetResponse, error)
	List(ctx context.Context, req *ArticlePostscriptGetReq) (*ArticlePostscriptListResponse, error)
	Map(ctx context.Context, req *ArticlePostscriptGetReq) (*ArticlePostscriptMapResponse, error)
	Count(ctx context.Context, req *ArticlePostscriptGetReq) (*ArticlePostscriptCountResponse, error)
	Page(ctx context.Context, req *ArticlePostscriptGetReq) (*ArticlePostscriptPageResponse, error)
}

type ArticlePostscriptSaveReq struct {
	ArticlePostscript *model.ArticlePostscript
}

type ArticlePostscriptSaveResponse struct {
	ArticlePostscript *model.ArticlePostscript
}

type ArticlePostscriptGetResponse struct {
	ArticlePostscript *model.ArticlePostscript
}

type ArticlePostscriptListResponse struct {
	Rows []*model.ArticlePostscript
}

type ArticlePostscriptMapResponse struct {
	Rows map[int64]*model.ArticlePostscript
}

type ArticlePostscriptCountResponse struct {
	Count int
}

type ArticlePostscriptPageResponse struct {
	Rows []*model.ArticlePostscript
	Page *base.PageResponse
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
	Save(ctx context.Context, req *ArticleActionRecordSaveReq) (*ArticleActionRecordSaveResponse, error)
	Delete(ctx context.Context, req *ArticleActionRecordDeleteReq) (*ArticleActionRecordDeleteResponse, error)
	Exist(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordExistResponse, error)
	Get(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordGetResponse, error)
	List(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordListResponse, error)
	Map(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordMapResponse, error)
	Count(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordCountResponse, error)
	Page(ctx context.Context, req *ArticleActionRecordReq) (*ArticleActionRecordPageResponse, error)
}

type ArticleActionRecordSaveReq struct {
	Record *model.ArticleActionRecord
}

type ArticleActionRecordSaveResponse struct {
	Created bool
}

type ArticleActionRecordDeleteReq struct {
	ArticleID int64
	UserID    int64
	Action    enum.ArticleAction
}

type ArticleActionRecordDeleteResponse struct {
	Deleted int
}

type ArticleActionRecordExistResponse struct {
	Exist bool
}

type ArticleActionRecordGetResponse struct {
	Record *model.ArticleActionRecord
}

type ArticleActionRecordListResponse struct {
	Rows []*model.ArticleActionRecord
}

type ArticleActionRecordMapResponse struct {
	Rows map[int64]*model.ArticleActionRecord
}

type ArticleActionRecordCountResponse struct {
	Count int
}

type ArticleActionRecordPageResponse struct {
	Rows []*model.ArticleActionRecord
	Page *base.PageResponse
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
