package repo

import (
	"context"
	"time"

	utilent "common/pkg/util/ent"

	"content/internal/biz/repo"
	"content/internal/data/gen"
	articleqaent "content/internal/data/gen/articleqa"
)

type ArticleQARepo struct {
	db *gen.Client
}

func NewArticleQARepo(db *gen.Client) repo.ArticleQARepo {
	return &ArticleQARepo{db: db}
}

func (r *ArticleQARepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticleQARepo) GetAcceptedCommentID(ctx context.Context, articleID int64) (*int64, error) {
	row, err := r.client(ctx).ArticleQA.Query().Where(articleqaent.ArticleIDEQ(articleID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return new(row.AcceptedCommentID), nil
}

func (r *ArticleQARepo) Accept(ctx context.Context, articleID, commentID int64, acceptedAt time.Time) error {
	_, err := r.client(ctx).ArticleQA.Create().SetArticleID(articleID).SetAcceptedCommentID(commentID).SetAcceptedAt(acceptedAt).Save(ctx)
	return err
}
