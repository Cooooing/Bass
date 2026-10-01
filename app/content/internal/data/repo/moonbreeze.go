package repo

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/moonbreeze"
	"context"

	utilent "common/pkg/util/ent"
)

var _ repo.MoonbreezeRepo = (*MoonbreezeRepo)(nil)

type MoonbreezeRepo struct {
	db *gen.Client
}

func NewMoonbreezeRepo(db *gen.Client) repo.MoonbreezeRepo {
	return &MoonbreezeRepo{db: db}
}

func (r *MoonbreezeRepo) client(ctx context.Context) *gen.Client {
	if client, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return client
	}
	return r.db
}

func (r *MoonbreezeRepo) Create(ctx context.Context, row *model.Moonbreeze) (*model.Moonbreeze, error) {
	saved, err := r.client(ctx).Moonbreeze.Create().
		SetContent(row.Content).
		SetAuthorID(row.AuthorID).
		SetNillableCity(row.City).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.moonbreezeModel(saved), nil
}

func (r *MoonbreezeRepo) Page(ctx context.Context, req *repo.MoonbreezePageReq) (*repo.MoonbreezePageResp, error) {
	if req == nil {
		req = &repo.MoonbreezePageReq{}
	}
	size := req.Size
	if size <= 0 || size > 51 {
		size = 20
	}
	query := r.client(ctx).Moonbreeze.Query()
	if len(req.AuthorIDs) > 0 {
		query = query.Where(moonbreeze.AuthorIDIn(req.AuthorIDs...))
	}
	if req.BeforeAt != nil && req.BeforeID != nil {
		query = query.Where(moonbreeze.Or(
			moonbreeze.CreatedAtLT(*req.BeforeAt),
			moonbreeze.And(moonbreeze.CreatedAtEQ(*req.BeforeAt), moonbreeze.IDLT(*req.BeforeID)),
		))
	}
	rows, err := query.
		Order(gen.Desc(moonbreeze.FieldCreatedAt), gen.Desc(moonbreeze.FieldID)).
		Limit(size).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Moonbreeze, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.moonbreezeModel(row))
	}
	return &repo.MoonbreezePageResp{Rows: result}, nil
}

func (*MoonbreezeRepo) moonbreezeModel(row *gen.Moonbreeze) *model.Moonbreeze {
	if row == nil {
		return nil
	}
	return &model.Moonbreeze{
		ID:        row.ID,
		Content:   row.Content,
		AuthorID:  row.AuthorID,
		City:      row.City,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
