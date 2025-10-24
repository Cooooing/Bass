package data

import (
	v1 "common/api/common/v1"
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

func (r *ArticleRepo) Save(ctx context.Context, client *gen.Client, article *model.Article) (*model.Article, error) {
	// TODO implement me
	panic("implement me")
}

func (r *ArticleRepo) UpdateHasPostscript(ctx context.Context, client *gen.Client, articleId int, hasPostscript bool) error {
	// TODO implement me
	panic("implement me")
}

func (r *ArticleRepo) UpdateStat(ctx context.Context, client *gen.Client, articleId int, action v1.ArticleAction, num int) error {
	// TODO implement me
	panic("implement me")
}

func (r *ArticleRepo) Delete(ctx context.Context, articleId int) error {
	// TODO implement me
	panic("implement me")
}

func (r *ArticleRepo) GetArticleById(ctx context.Context, client *gen.Client, id int) (*model.Article, error) {
	// TODO implement me
	panic("implement me")
}
