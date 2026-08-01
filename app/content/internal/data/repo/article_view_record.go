package repo

import (
	"common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/articleviewrecord"
	"context"

	"entgo.io/ent/dialect/sql"
)

var _ repo.ArticleViewRecordRepo = (*ArticleViewRecordRepo)(nil)

type ArticleViewRecordRepo struct {
	db *gen.Client
}

func NewArticleViewRecordRepo(
	db *gen.Client,
) repo.ArticleViewRecordRepo {
	return &ArticleViewRecordRepo{
		db: db,
	}
}

func (r *ArticleViewRecordRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := ent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticleViewRecordRepo) Save(ctx context.Context, record *model.ArticleViewRecord) error {
	create := r.getClient(ctx).ArticleViewRecord.Create().
		SetArticleID(record.ArticleID).
		SetUserID(record.UserID).
		SetNillableIP(record.IP).
		SetNillableUserAgent(record.UserAgent).
		SetNillableBrowserFingerprint(record.BrowserFingerprint)
	if record.ViewedAt != nil {
		create.SetViewedAt(*record.ViewedAt)
	}
	return create.OnConflict(
		sql.ConflictColumns(articleviewrecord.FieldArticleID, articleviewrecord.FieldUserID),
	).UpdateNewValues().Exec(ctx)
}
