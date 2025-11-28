package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type ArticleRepo interface {
	Save(ctx context.Context, tx *gen.Client, article *model.Article, tags []*model.Tag) (*model.Article, error)

	Update(ctx context.Context, tx *gen.Client, article *model.Article, tags []*model.Tag) (*model.Article, error)
	UpdateContent(ctx context.Context, tx *gen.Client, articleId int64, content string) error
	UpdateStatus(ctx context.Context, tx *gen.Client, articleId int64, status v1.ArticleStatus) error
	UpdateHasPostscript(ctx context.Context, tx *gen.Client, articleId int64, hasPostscript bool) error
	UpdateStat(ctx context.Context, tx *gen.Client, articleId int64, action v1.ArticleAction, num int32) error
	Publish(ctx context.Context, tx *gen.Client, articleId int64) error

	Delete(ctx context.Context, tx *gen.Client, articleId int64) error

	Exist(ctx context.Context, tx *gen.Client, req *ArticleGetReq) (bool, error)
	GetOne(ctx context.Context, tx *gen.Client, req *ArticleGetReq) (*model.Article, error)
	GetList(ctx context.Context, tx *gen.Client, req *ArticleGetReq) ([]*model.Article, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ArticleGetReq) ([]*model.Article, *cv1.PageReply, error)
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
}

type ArticlePostscriptRepo interface {
	Save(ctx context.Context, tx *gen.Client, articlePostscript *model.ArticlePostscript) (*model.ArticlePostscript, error)
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error)
	Delete(ctx context.Context, tx *gen.Client, articleId int64, userId int64, action v1.ArticleAction) error
}
