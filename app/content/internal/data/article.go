package data

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type ArticleRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewArticleRepo(baseRepo *BaseRepo, client *gen.Client) repo.ArticleRepo {
	return &ArticleRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (r *ArticleRepo) Save(ctx context.Context, a *model.Article) (*model.Article, error) {
	return nil, nil
}

func (r *ArticleRepo) GetArticleById(ctx context.Context, id int) (*model.Article, error) {
	return nil, nil
}
