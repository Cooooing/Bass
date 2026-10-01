package repo

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/breezemoon"
	"context"

	utilent "common/pkg/util/ent"
)

var _ repo.BreezemoonRepo = (*BreezemoonRepo)(nil)

type BreezemoonRepo struct {
	db *gen.Client
}

func NewBreezemoonRepo(db *gen.Client) repo.BreezemoonRepo {
	return &BreezemoonRepo{db: db}
}

func (r *BreezemoonRepo) client(ctx context.Context) *gen.Client {
	if client, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return client
	}
	return r.db
}

func (r *BreezemoonRepo) Create(ctx context.Context, row *model.Breezemoon) (*model.Breezemoon, error) {
	saved, err := r.client(ctx).Breezemoon.Create().
		SetContent(row.Content).
		SetAuthorID(row.AuthorID).
		SetNillableCity(row.City).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.breezemoonModel(saved), nil
}

func (r *BreezemoonRepo) Page(ctx context.Context, req *repo.BreezemoonPageReq) (*repo.BreezemoonPageResp, error) {
	if req == nil {
		req = &repo.BreezemoonPageReq{}
	}
	size := req.Size
	if size <= 0 || size > 51 {
		size = 20
	}
	query := r.client(ctx).Breezemoon.Query()
	if len(req.AuthorIDs) > 0 {
		query = query.Where(breezemoon.AuthorIDIn(req.AuthorIDs...))
	}
	if req.BeforeAt != nil && req.BeforeID != nil {
		query = query.Where(breezemoon.Or(
			breezemoon.CreatedAtLT(*req.BeforeAt),
			breezemoon.And(breezemoon.CreatedAtEQ(*req.BeforeAt), breezemoon.IDLT(*req.BeforeID)),
		))
	}
	rows, err := query.
		Order(gen.Desc(breezemoon.FieldCreatedAt), gen.Desc(breezemoon.FieldID)).
		Limit(size).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Breezemoon, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.breezemoonModel(row))
	}
	return &repo.BreezemoonPageResp{Rows: result}, nil
}

func (*BreezemoonRepo) breezemoonModel(row *gen.Breezemoon) *model.Breezemoon {
	if row == nil {
		return nil
	}
	return &model.Breezemoon{
		ID:        row.ID,
		Content:   row.Content,
		AuthorID:  row.AuthorID,
		City:      row.City,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
