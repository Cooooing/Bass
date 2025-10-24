package repo

import (
	v1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type ArticleRepo interface {
	Save(ctx context.Context, client *gen.Client, article *model.Article) (*model.Article, error)
	UpdateHasPostscript(ctx context.Context, client *gen.Client, articleId int, hasPostscript bool) error
	UpdateStat(ctx context.Context, client *gen.Client, articleId int, action v1.ArticleAction, num int) error

	Delete(ctx context.Context, articleId int) error

	GetArticleById(ctx context.Context, client *gen.Client, id int) (*model.Article, error)
}

type ArticlePostscriptRepo interface {
	AddPostscript(ctx context.Context, client *gen.Client, articleId int, content string) error
}

type ArticleActionRecordRepo interface {
	Save(ctx context.Context, client *gen.Client, record *model.ActionRecord) (*model.ActionRecord, error)
	Delete(ctx context.Context, client *gen.Client, articleId int, userId int, action v1.ArticleAction) error
}
