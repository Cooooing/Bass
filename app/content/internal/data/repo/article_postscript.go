package repo

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	basedata "content/internal/data/base"
	"content/internal/data/ent/gen"
	"context"
)

type ArticlePostscriptRepo struct {
	*basedata.BaseData
	client *gen.Client
}

func NewArticlePostscriptRepo(BaseData *basedata.BaseData, client *gen.Client) repo.ArticlePostscriptRepo {
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
