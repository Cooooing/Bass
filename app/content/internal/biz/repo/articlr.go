package repo

import (
	"content/internal/biz/model"
	"context"
)

type ArticleRepo interface {
	Save(ctx context.Context, article *model.Article) (*model.Article, error)

	GetArticleById(ctx context.Context, id int) (*model.Article, error)
}
