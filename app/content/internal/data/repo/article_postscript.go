package repo

import (
	"content/internal/biz/base"
	"context"

	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/articlepostscript"
	"content/internal/enum"

	"github.com/samber/lo"
)

var _ repo.ArticlePostscriptRepo = (*ArticlePostscriptRepo)(nil)

type ArticlePostscriptRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewArticlePostscriptRepo(
	db *gen.Client,
) repo.ArticlePostscriptRepo {
	return &ArticlePostscriptRepo{
		db: db,
	}
}

func (r *ArticlePostscriptRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ArticlePostscriptRepo) Save(ctx context.Context, articlePostscript *model.ArticlePostscript) (*model.ArticlePostscript, error) {
	postscript := articlePostscript
	save, err := r.getClient(ctx).ArticlePostscript.Create().
		SetArticleID(postscript.ArticleID).
		SetContent(postscript.Content).
		SetRestriction(articlepostscript.Restriction(postscript.Restriction)).
		SetNillableCreatedBy(postscript.CreatedBy).
		SetNillableUpdatedBy(postscript.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.ArticlePostscript{
		ID:          save.ID,
		ArticleID:   save.ArticleID,
		Content:     save.Content,
		Restriction: enum.ContentRestriction(save.Restriction),
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
		CreatedBy:   save.CreatedBy,
		UpdatedBy:   save.UpdatedBy,
	}, nil
}

func (r *ArticlePostscriptRepo) Get(ctx context.Context, req *repo.ArticlePostscriptGetReq) (*model.ArticlePostscript, error) {
	query := r.getClient(ctx).ArticlePostscript.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.ArticlePostscript{
		ID:          row.ID,
		ArticleID:   row.ArticleID,
		Content:     row.Content,
		Restriction: enum.ContentRestriction(row.Restriction),
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		CreatedBy:   row.CreatedBy,
		UpdatedBy:   row.UpdatedBy,
	}, nil
}

func (r *ArticlePostscriptRepo) List(ctx context.Context, req *repo.ArticlePostscriptGetReq) ([]*model.ArticlePostscript, error) {
	query := r.getClient(ctx).ArticlePostscript.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ArticlePostscript, 0, len(list))
	for _, row := range list {
		result = append(result, &model.ArticlePostscript{
			ID:          row.ID,
			ArticleID:   row.ArticleID,
			Content:     row.Content,
			Restriction: enum.ContentRestriction(row.Restriction),
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
			CreatedBy:   row.CreatedBy,
			UpdatedBy:   row.UpdatedBy,
		})
	}
	return result, nil
}

func (r *ArticlePostscriptRepo) Map(ctx context.Context, req *repo.ArticlePostscriptGetReq) (map[int64]*model.
	ArticlePostscript, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(listResp, func(item *model.ArticlePostscript) (int64, *model.ArticlePostscript) {
		return item.ID, item
	}), nil
}

func (r *ArticlePostscriptRepo) Count(ctx context.Context, req *repo.ArticlePostscriptGetReq) (int, error) {
	query := r.getClient(ctx).ArticlePostscript.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *ArticlePostscriptRepo) Page(ctx context.Context, req *repo.ArticlePostscriptGetReq) (*repo.ArticlePostscriptPageResp, error) {
	page := r.normalizePage(req.Page)
	query := r.getClient(ctx).ArticlePostscript.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ArticlePostscript, 0, len(list))
	for _, row := range list {
		result = append(result, &model.ArticlePostscript{
			ID:          row.ID,
			ArticleID:   row.ArticleID,
			Content:     row.Content,
			Restriction: enum.ContentRestriction(row.Restriction),
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
			CreatedBy:   row.CreatedBy,
			UpdatedBy:   row.UpdatedBy,
		})
	}
	return &repo.ArticlePostscriptPageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *ArticlePostscriptRepo) getQuery(query *gen.ArticlePostscriptQuery, req *repo.ArticlePostscriptGetReq) *gen.ArticlePostscriptQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(articlepostscript.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(articlepostscript.IDIn(req.IDs...))
	}
	if req.ArticleID != nil {
		query = query.Where(articlepostscript.ArticleIDEQ(*req.ArticleID))
	}
	if len(req.ArticleIDs) > 0 {
		query = query.Where(articlepostscript.ArticleIDIn(req.ArticleIDs...))
	}
	if req.CreatedBy != nil {
		query = query.Where(articlepostscript.CreatedByEQ(*req.CreatedBy))
	}
	if req.Restriction != nil {
		query = query.Where(articlepostscript.RestrictionEQ(articlepostscript.Restriction(*req.Restriction)))
	}
	return query
}
