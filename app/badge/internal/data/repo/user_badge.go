package repo

import (
	"badge/internal/biz/model"
	bizrepo "badge/internal/biz/repo"
	"badge/internal/data/gen"
	"badge/internal/data/gen/predicate"
	"badge/internal/data/gen/userbadge"
	commonmodel "common/pkg/model"
	"context"
	"time"
)

type UserBadgeRepo struct{ db *gen.Client }

func NewUserBadgeRepo(db *gen.Client) bizrepo.UserBadgeRepo { return &UserBadgeRepo{db: db} }
func (r *UserBadgeRepo) Get(ctx context.Context, req *bizrepo.UserBadgeGetReq) (*model.UserBadge, error) {
	row, err := r.query(req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *UserBadgeRepo) List(ctx context.Context, req *bizrepo.UserBadgeGetReq) ([]*model.UserBadge, error) {
	rows, err := r.query(req).All(ctx)
	if err != nil {
		return nil, err
	}
	return r.models(rows), nil
}

func (r *UserBadgeRepo) Page(ctx context.Context, req *bizrepo.UserBadgeGetReq) (*bizrepo.UserBadgePageResp, error) {
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
	return &bizrepo.UserBadgePageResp{
		Rows: r.models(rows),
		Page: &commonmodel.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}
func (r *UserBadgeRepo) Grant(ctx context.Context, items []*model.UserBadge) (_ []*model.UserBadge, err error) {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	builders := make([]*gen.UserBadgeCreate, 0, len(items))
	predicates := make([]predicate.UserBadge, 0, len(items))
	for _, item := range items {
		builders = append(builders, tx.UserBadge.Create().
			SetUserID(item.UserID).
			SetBadgeID(item.BadgeID).
			SetVariables(item.Variables).
			SetObtainedAt(item.ObtainedAt).
			SetNillableExpiresAt(item.ExpiresAt).
			SetDisplayed(item.Displayed).
			SetDisplayOrder(item.DisplayOrder))
		predicates = append(predicates, userbadge.And(
			userbadge.UserIDEQ(item.UserID),
			userbadge.BadgeIDEQ(item.BadgeID),
		))
	}

	err = tx.UserBadge.CreateBulk(builders...).
		OnConflictColumns(userbadge.FieldUserID, userbadge.FieldBadgeID).
		Update(func(upsert *gen.UserBadgeUpsert) {
			upsert.UpdateVariables()
			upsert.UpdateObtainedAt()
			upsert.UpdateExpiresAt()
		}).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	rows, err := tx.UserBadge.Query().Where(userbadge.Or(predicates...)).All(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return r.models(rows), nil
}
func (r *UserBadgeRepo) Revoke(ctx context.Context, items []*model.UserBadge) (int, error) {
	predicates := make([]predicate.UserBadge, 0, len(items))
	for _, item := range items {
		predicates = append(predicates, userbadge.And(
			userbadge.UserIDEQ(item.UserID),
			userbadge.BadgeIDEQ(item.BadgeID),
		))
	}
	return r.db.UserBadge.Delete().Where(userbadge.Or(predicates...)).Exec(ctx)
}
func (r *UserBadgeRepo) UpdatePresentation(ctx context.Context, userID, badgeID int64, displayed bool, displayOrder int32) (*model.UserBadge, error) {
	updated, err := r.db.UserBadge.Update().
		Where(userbadge.UserIDEQ(userID), userbadge.BadgeIDEQ(badgeID)).
		SetDisplayed(displayed).
		SetDisplayOrder(displayOrder).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	if updated == 0 {
		return nil, nil
	}
	return r.Get(ctx, &bizrepo.UserBadgeGetReq{
		UserID:  &userID,
		BadgeID: &badgeID,
	})
}

func (r *UserBadgeRepo) query(req *bizrepo.UserBadgeGetReq) *gen.UserBadgeQuery {
	query := r.db.UserBadge.Query()
	if req == nil {
		return query.Order(gen.Asc(userbadge.FieldDisplayOrder), gen.Desc(userbadge.FieldID))
	}
	if req.UserID != nil {
		query = query.Where(userbadge.UserIDEQ(*req.UserID))
	}
	if req.BadgeID != nil {
		query = query.Where(userbadge.BadgeIDEQ(*req.BadgeID))
	}
	if req.Displayed != nil {
		query = query.Where(userbadge.DisplayedEQ(*req.Displayed))
	}
	if req.ActiveOnly {
		query = query.Where(userbadge.Or(
			userbadge.ExpiresAtIsNil(),
			userbadge.ExpiresAtGT(time.Now()),
		))
	}
	return query.Order(gen.Asc(userbadge.FieldDisplayOrder), gen.Desc(userbadge.FieldID))
}
func (*UserBadgeRepo) model(row *gen.UserBadge) *model.UserBadge {
	return &model.UserBadge{
		ID:           row.ID,
		UserID:       row.UserID,
		BadgeID:      row.BadgeID,
		Variables:    row.Variables,
		ObtainedAt:   row.ObtainedAt,
		ExpiresAt:    row.ExpiresAt,
		Displayed:    row.Displayed,
		DisplayOrder: row.DisplayOrder,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}

func (r *UserBadgeRepo) models(rows []*gen.UserBadge) []*model.UserBadge {
	items := make([]*model.UserBadge, 0, len(rows))
	for _, row := range rows {
		items = append(items, r.model(row))
	}
	return items
}
