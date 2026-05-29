package repo

import (
	"context"

	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/articlepostscript"
	"content/internal/enum"
)

var _ repo.ArticlePostscriptRepo = (*ArticlePostscriptRepo)(nil)

type ArticlePostscriptRepo struct {
	db *gen.Client
}

func NewArticlePostscriptRepo(db *gen.Client) repo.ArticlePostscriptRepo {
	return &ArticlePostscriptRepo{db: db}
}

func (r *ArticlePostscriptRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticlePostscriptRepo) Save(ctx context.Context, postscript *model.ArticlePostscript) (*model.ArticlePostscript, error) {
	save, err := r.getClient(ctx).ArticlePostscript.Create().
		SetArticleID(postscript.ArticleID).
		SetContent(postscript.Content).
		SetNillableCreatedBy(postscript.CreatedBy).
		SetNillableUpdatedBy(postscript.UpdatedBy).
		SetStatus(articlepostscript.Status(postscript.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ArticlePostscript{
		ID:        save.ID,
		ArticleID: save.ArticleID,
		Content:   save.Content,
		Status:    enum.ArticlePostscriptStatus(save.Status),
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
		CreatedBy: save.CreatedBy,
		UpdatedBy: save.UpdatedBy,
	}, nil
}
