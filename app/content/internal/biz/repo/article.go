package repo

import (
	cv1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type ArticleRepo interface {
	Save(ctx context.Context, tx *gen.Client, article *model.Article) (*model.Article, error)
	UpdateContent(ctx context.Context, tx *gen.Client, articleId int64, content string) error
	UpdateStatus(ctx context.Context, tx *gen.Client, articleId int64, status cv1.ArticleStatus) error
	UpdateHasPostscript(ctx context.Context, tx *gen.Client, articleId int64, hasPostscript bool) error
	UpdateStat(ctx context.Context, tx *gen.Client, articleId int64, action cv1.ArticleAction, num int32) error
	Publish(ctx context.Context, tx *gen.Client, articleId int64) error

	Delete(ctx context.Context, tx *gen.Client, articleId int64) error

	Exist(ctx context.Context, tx *gen.Client, articleId int64, status cv1.ArticleStatus) (bool, error)
	GetById(ctx context.Context, tx *gen.Client, articleId int64) (*model.Article, error)
	GetList(ctx context.Context, tx *gen.Client, req *ArticleGetReq) ([]*model.Article, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ArticleGetReq) ([]*model.Article, *cv1.PageReply, error)
}

type ArticleGetReq struct {
	TagId    *int64
	DomainId *int64
	Status   *cv1.ArticleStatus
	AuthorId *int64
	Order    *cv1.ArticleOrder
	Type     *cv1.ArticleType
	Keyword  *string
}

type ArticlePostscriptRepo interface {
	AddPostscript(ctx context.Context, tx *gen.Client, articleId int64, content string) error
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error)
	Delete(ctx context.Context, tx *gen.Client, articleId int64, userId int64, action cv1.ArticleAction) error
}
