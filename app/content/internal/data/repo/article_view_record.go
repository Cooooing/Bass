package repo

import (
	commonmodel "common/pkg/model"
	"common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/article"
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

func (r *ArticleViewRecordRepo) Page(
	ctx context.Context,
	req *repo.ArticleViewRecordPageReq,
) (*repo.ArticleViewRecordPageResp, error) {
	if req == nil {
		req = &repo.ArticleViewRecordPageReq{}
	}
	page := commonmodel.NormalizePage(req.Page)
	query := r.getClient(ctx).ArticleViewRecord.Query().
		Where(articleviewrecord.UserIDEQ(req.UserID)).
		Where(articleviewrecord.HasArticleWith(
			article.DeletedAtIsNil(),
			article.PublishStatusEQ(article.PublishStatusPublished),
			article.VisibilityEQ(article.VisibilityPublic),
			article.RestrictionEQ(article.RestrictionNone),
		))
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	records, err := query.
		Order(gen.Desc(articleviewrecord.FieldViewedAt), gen.Desc(articleviewrecord.FieldID)).
		Limit(page.Limit()).
		Offset(page.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rows := make([]*model.ArticleViewRecord, 0, len(records))
	for _, record := range records {
		rows = append(rows, &model.ArticleViewRecord{
			ID:        record.ID,
			ArticleID: record.ArticleID,
			UserID:    record.UserID,
			ViewedAt:  &record.ViewedAt,
		})
	}
	return &repo.ArticleViewRecordPageResp{
		Rows: rows,
		Page: &commonmodel.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}
