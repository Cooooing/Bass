package repo

import (
	"common/proto/gen/common"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
	"time"
)

type ArticleRepo interface {
	Save(ctx context.Context, article *model.Article) (*model.Article, error)

	Update(ctx context.Context, article *model.Article) (*model.Article, error)
	UpdatePublishStatus(ctx context.Context, articleId int64, publishStatus enum.ArticlePublishStatus, visibility enum.ArticleVisibility, publishedAt *time.Time, updatedBy int64) error
	UpdateVisibility(ctx context.Context, articleId int64, visibility enum.ArticleVisibility, updatedBy int64) error
	UpdateRestriction(ctx context.Context, articleId int64, restriction enum.ContentRestriction, updatedBy int64) error
	DiscardDraft(ctx context.Context, articleId int64) error
	UpdateHasPostscript(ctx context.Context, articleId int64, hasPostscript bool, updatedBy int64) error
	AddStats(ctx context.Context, articleId int64, stats ArticleStatUpdate, updatedBy *int64) error
	UpdateAcceptedAnswerID(ctx context.Context, articleId int64, commentId int64, updatedBy int64) (*model.Article, error)
	ReplaceTags(ctx context.Context, articleId int64, tagIds []int64) error

	Exist(ctx context.Context, req *ArticleGetReq) (bool, error)
	Get(ctx context.Context, req *ArticleGetReq) (*model.Article, error)
	List(ctx context.Context, req *ArticleGetReq) ([]*model.Article, error)
	Map(ctx context.Context, req *ArticleGetReq) (map[int64]*model.Article, error)
	Count(ctx context.Context, req *ArticleGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ArticleGetReq) ([]*model.Article, *common.PageReply, error)
}

type ArticleStatUpdate struct {
	ViewCount    int32
	ThankCount   int32
	LikeCount    int32
	CollectCount int32
	WatchCount   int32
	ReplyCount   int32
}

type ArticleGetReq struct {
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
	Page(ctx context.Context, page *common.PageRequest, req *ArticlePostscriptGetReq) ([]*model.ArticlePostscript, *common.PageReply, error)
}

type ArticlePostscriptGetReq struct {
	ID          *int64
	IDs         []int64
	ArticleID   *int64
	ArticleIDs  []int64
	CreatedBy   *int64
	Restriction *enum.ContentRestriction
}

type ArticleActionRecordRepo interface {
	// Save 保存行为记录，返回值表示是否实际写入了新记录；唯一约束冲突表示记录已存在。
	Save(ctx context.Context, record *model.ArticleActionRecord) (bool, error)
	Delete(ctx context.Context, articleId int64, userId int64, action enum.ArticleAction) (int, error)

	Exist(ctx context.Context, req *ArticleActionRecordReq) (bool, error)
	Get(ctx context.Context, req *ArticleActionRecordReq) (*model.ArticleActionRecord, error)
	List(ctx context.Context, req *ArticleActionRecordReq) ([]*model.ArticleActionRecord, error)
	Map(ctx context.Context, req *ArticleActionRecordReq) (map[int64]*model.ArticleActionRecord, error)
	Count(ctx context.Context, req *ArticleActionRecordReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ArticleActionRecordReq) ([]*model.ArticleActionRecord, *common.PageReply, error)
}

type ArticleActionRecordReq struct {
	ID         *int64
	IDs        []int64
	ArticleId  *int64
	ArticleIds []int64
	UserId     *int64
	UserIds    []int64
	Type       *enum.ArticleAction
	Types      []enum.ArticleAction
}
