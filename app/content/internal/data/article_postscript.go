package data

import (
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type ArticlePostscriptRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewArticlePostscriptRepo(baseRepo *BaseRepo, client *gen.Client) repo.ArticlePostscriptRepo {
	return &ArticlePostscriptRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (a ArticlePostscriptRepo) AddPostscript(ctx context.Context, client *gen.Client, articleId int64, content string) error {
	return client.ArticlePostscript.Create().
		SetArticleID(articleId).
		SetContent(content).
		Exec(ctx)
}
