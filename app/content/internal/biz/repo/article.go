package repo

import (
	v1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type ArticleRepo interface {
	Save(ctx context.Context, client *gen.Client, article *model.Article) (*model.Article, error)
	UpdateContent(ctx context.Context, client *gen.Client, articleId int64, content string) error
	UpdateStatus(ctx context.Context, client *gen.Client, articleId int64, status v1.ArticleStatus) error
	UpdateHasPostscript(ctx context.Context, client *gen.Client, articleId int64, hasPostscript bool) error
	UpdateStat(ctx context.Context, client *gen.Client, articleId int64, action v1.ArticleAction, num int32) error

	Delete(ctx context.Context, client *gen.Client, articleId int64) error

	Exist(ctx context.Context, client *gen.Client, id int64, status v1.ArticleStatus) (bool, error)
	GetArticleById(ctx context.Context, client *gen.Client, id int64) (*model.Article, error)

	Publish(ctx context.Context, client *gen.Client, articleId int64) error
}

type ArticlePostscriptRepo interface {
	AddPostscript(ctx context.Context, client *gen.Client, articleId int64, content string) error
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, client *gen.Client, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error)
	Delete(ctx context.Context, client *gen.Client, articleId int64, userId int64, action v1.ArticleAction) error
}
