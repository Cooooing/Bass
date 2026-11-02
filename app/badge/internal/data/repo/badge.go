package repo

import (
	"badge/internal/biz/model"
	bizrepo "badge/internal/biz/repo"
	"badge/internal/data/gen"
	"badge/internal/data/gen/badge"
	commonmodel "common/pkg/model"
	"context"
)

// BadgeRepo adapts Ent badge persistence to the business query port.
type BadgeRepo struct {
	db *gen.Client
}

func NewBadgeRepo(db *gen.Client) bizrepo.BadgeRepo {
	return &BadgeRepo{db: db}
}

func (r *BadgeRepo) Get(ctx context.Context, req *bizrepo.BadgeGetReq) (*model.Badge, error) {
	row, err := r.query(req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *BadgeRepo) List(ctx context.Context, req *bizrepo.BadgeGetReq) ([]*model.Badge, error) {
	rows, err := r.query(req).All(ctx)
	if err != nil {
		return nil, err
	}
	return r.models(rows), nil
}

func (r *BadgeRepo) Page(ctx context.Context, req *bizrepo.BadgeGetReq) (*bizrepo.BadgePageResp, error) {
	var pageReq *commonmodel.PageReq
	if req != nil {
		pageReq = req.Page
	}

	page := commonmodel.NormalizePage(pageReq)
	query := r.query(req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Limit(page.Limit()).Offset(page.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.BadgePageResp{
		Rows: r.models(rows),
		Page: &commonmodel.PageResp{Total: int64(total), Page: page.Page, Size: page.Size},
	}, nil
}

func (r *BadgeRepo) Create(ctx context.Context, item *model.Badge) (*model.Badge, error) {
	row, err := r.db.Badge.Create().
		SetNillableCode(item.Code).
		SetName(item.Name).
		SetDescription(item.Description).
		SetRenderParams(item.RenderParams).
		SetDesignerUserID(item.DesignerUserID).
		SetStatus(badge.Status(item.Status)).
		SetNillablePublishedAt(item.PublishedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *BadgeRepo) Update(ctx context.Context, item *model.Badge) (*model.Badge, error) {
	row, err := r.db.Badge.UpdateOneID(item.ID).
		SetNillableCode(item.Code).
		SetName(item.Name).
		SetDescription(item.Description).
		SetRenderParams(item.RenderParams).
		SetStatus(badge.Status(item.Status)).
		SetNillablePublishedAt(item.PublishedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *BadgeRepo) Delete(ctx context.Context, id int64) error {
	return r.db.Badge.DeleteOneID(id).Exec(ctx)
}

func (r *BadgeRepo) query(req *bizrepo.BadgeGetReq) *gen.BadgeQuery {
	query := r.db.Badge.Query()
	if req == nil {
		return query.Order(gen.Desc(badge.FieldUpdatedAt), gen.Desc(badge.FieldID))
	}
	if req.BadgeID != nil {
		query = query.Where(badge.IDEQ(*req.BadgeID))
	}
	if req.Code != nil {
		query = query.Where(badge.CodeEQ(*req.Code))
	}
	if req.DesignerUserID != nil {
		query = query.Where(badge.DesignerUserIDEQ(*req.DesignerUserID))
	}
	if req.Status != nil {
		query = query.Where(badge.StatusEQ(badge.Status(*req.Status)))
	}
	return query.Order(gen.Desc(badge.FieldUpdatedAt), gen.Desc(badge.FieldID))
}

func (r *BadgeRepo) models(rows []*gen.Badge) []*model.Badge {
	items := make([]*model.Badge, 0, len(rows))
	for _, row := range rows {
		items = append(items, r.model(row))
	}
	return items
}

func (r *BadgeRepo) model(row *gen.Badge) *model.Badge {
	return &model.Badge{
		ID:             row.ID,
		Code:           row.Code,
		Name:           row.Name,
		Description:    row.Description,
		RenderParams:   row.RenderParams,
		DesignerUserID: row.DesignerUserID,
		Status:         string(row.Status),
		PublishedAt:    row.PublishedAt,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
