package data

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type ArticlePostscriptRepo struct {
	*BaseData
	client *gen.Client
}

func NewArticlePostscriptRepo(BaseData *BaseData, client *gen.Client) repo.ArticlePostscriptRepo {
	return &ArticlePostscriptRepo{
		BaseData: BaseData,
		client:   client,
	}
}

func (a ArticlePostscriptRepo) Save(ctx context.Context, client *gen.Client, articlePostscript *model.ArticlePostscript) (*model.ArticlePostscript, error) {
	save, err := client.ArticlePostscript.Create().
		SetArticleID(articlePostscript.ArticleID).
		SetContent(articlePostscript.Content).
		Save(ctx)
	return &model.ArticlePostscript{ArticlePostscript: save}, err
}
