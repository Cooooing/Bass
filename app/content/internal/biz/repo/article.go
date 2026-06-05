package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	"content/internal/biz/model"
	"context"
)

type ArticleRepo interface {
	Save(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error)

	Update(ctx context.Context, article *model.Article, tags []*model.Tag) (*model.Article, error)
	UpdateContent(ctx context.Context, articleId int64, content string) error
	UpdateStatus(ctx context.Context, articleId int64, userId int64, status v1.ArticleStatus) error
	UpdateControlFields(ctx context.Context, articleId int64, userId int64, status v1.ArticleStatus, commentable bool, anonymous bool, listable *bool) error
	UpdateHasPostscript(ctx context.Context, articleId int64, userId int64, hasPostscript bool) error
	UpdateStat(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction, num int32) (*model.Article, error)
	UpdateAcceptAnswer(ctx context.Context, articleId int64, userId int64, commentId int64) (*model.Article, error)
	Publish(ctx context.Context, articleId int64, userId int64) (*model.Article, error)

	Delete(ctx context.Context, articleId int64, userId int64) error

	Exist(ctx context.Context, req *ArticleGetReq) (bool, error)
	Get(ctx context.Context, req *ArticleGetReq) (*model.Article, error)
	GetList(ctx context.Context, req *ArticleGetReq) ([]*model.Article, error)
	Page(ctx context.Context, page *common.PageRequest, req *ArticleGetReq) ([]*model.Article, *common.PageReply, error)
}

type ArticleGetReq struct {
	ArticleId *int64
	CreatedBy *int64
	TagId     *int64
	DomainId  *int64
	Status    *v1.ArticleStatus
	AuthorId  *int64
	Order     *v1.ArticleOrder
	Type      *v1.ArticleType
	Keyword   *string

	Listable *bool

	IsSummary   bool
	QueryUserId *int64
}

type ArticlePostscriptRepo interface {
	Save(ctx context.Context, articlePostscript *model.ArticlePostscript) (*model.ArticlePostscript, error)
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error)
	Delete(ctx context.Context, articleId int64, userId int64, action v1.ArticleAction) (int, error)

	Exist(ctx context.Context, req *ArticleActionRecordReq) (bool, error)
	Get(ctx context.Context, req *ArticleActionRecordReq) (*model.ArticleActionRecord, error)
	GetList(ctx context.Context, req *ArticleActionRecordReq) ([]*model.ArticleActionRecord, error)
	Page(ctx context.Context, page *common.PageRequest, req *ArticleActionRecordReq) ([]*model.ArticleActionRecord, *common.PageReply, error)
}

type ArticleActionRecordReq struct {
	ArticleId *int64
	UserId    *int64
	Type      *v1.ArticleAction
}
