package data

import (
	"common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type ArticleActionRecordRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewArticleActionRecordRepo(baseRepo *BaseRepo, client *gen.Client) repo.ArticleActionRecordRepo {
	return &ArticleActionRecordRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (a ArticleActionRecordRepo) Save(ctx context.Context, client *gen.Client, record *model.ArticleActionRecord) (*model.ArticleActionRecord, error) {
	// TODO implement me
	panic("implement me")
}

func (a ArticleActionRecordRepo) Delete(ctx context.Context, client *gen.Client, articleId int, userId int, action v1.ArticleAction) error {
	// TODO implement me
	panic("implement me")
}
