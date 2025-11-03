package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
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

	Exist(ctx context.Context, tx *gen.Client, id int64, status cv1.ArticleStatus) (bool, error)
	GetArticleById(ctx context.Context, tx *gen.Client, id int64) (*model.Article, error)

	GetOne(ctx context.Context, tx *gen.Client, articleId int64) (*v1.GetArticleOneReply, error)
	GetList(ctx context.Context, tx *gen.Client, req *v1.GetArticleRequest) (*v1.GetArticleReply, error)
}

type ArticlePostscriptRepo interface {
	AddPostscript(ctx context.Context, tx *gen.Client, articleId int64, content string) error
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, tx *gen.Client, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error)
	Delete(ctx context.Context, tx *gen.Client, articleId int64, userId int64, action cv1.ArticleAction) error
}
