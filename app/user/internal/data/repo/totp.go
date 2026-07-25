package repo

import (
	"common/proto/gen/common"
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/totp"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.TotpRepo = (*TotpRepo)(nil)

type TotpRepo struct {
	db *gen.Client
}

func NewTotpRepo(
	db *gen.Client,
) repo.TotpRepo {
	return &TotpRepo{
		db: db,
	}
}

func (r *TotpRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *TotpRepo) Get(ctx context.Context, req *repo.TotpGetReq) (*model.Totp, error) {
	row, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (r *TotpRepo) List(ctx context.Context, req *repo.TotpGetReq) ([]*model.Totp, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *TotpRepo) Map(ctx context.Context, req *repo.TotpGetReq) (map[int64]*model.Totp, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *TotpRepo) Count(ctx context.Context, req *repo.TotpGetReq) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TotpRepo) Page(ctx context.Context, req *repo.TotpPageReq) (*repo.TotpPageResp, error) {
	rows, page, err := r.page(ctx, &common.PageReq{
		Page: req.Page.Page,
		Size: req.Page.Size,
	}, &req.Query)
	if err != nil {
		return nil, err
	}
	resp := repo.PageResp{}
	if page != nil {
		resp = repo.PageResp{
			Total: page.GetTotal(),
			Page:  page.GetPage(),
			Size:  page.GetSize(),
		}
	}
	return &repo.TotpPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *TotpRepo) UpsertEnabledByUserID(ctx context.Context, req *repo.TotpUpsertEnabledByUserIDReq) (*model.Totp, error) {
	row, err := r.upsertEnabledByUserID(ctx, req.UserID, req.Secret)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (r *TotpRepo) DisableByUserID(ctx context.Context, userID int64) (*model.Totp, error) {
	row, err := r.disableByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return row, nil
}

func (r *TotpRepo) get(ctx context.Context, req *repo.TotpGetReq) (*model.Totp, error) {
	tx := r.getClient(ctx)
	query := tx.Totp.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Totp{
		ID:         row.ID,
		UserID:     row.UserID,
		Enable:     row.Enable,
		EnableTime: row.EnableTime,
		Secret:     row.Secret,
	}, nil
}

func (r *TotpRepo) list(ctx context.Context, req *repo.TotpGetReq) ([]*model.Totp, error) {
	tx := r.getClient(ctx)
	query := tx.Totp.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Totp, 0, len(list))
	for _, row := range list {
		result = append(result, &model.Totp{
			ID:         row.ID,
			UserID:     row.UserID,
			Enable:     row.Enable,
			EnableTime: row.EnableTime,
			Secret:     row.Secret,
		})
	}
	return result, nil
}

func (r *TotpRepo) mapRows(ctx context.Context, req *repo.TotpGetReq) (map[int64]*model.Totp, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Totp, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *TotpRepo) count(ctx context.Context, req *repo.TotpGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.Totp.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *TotpRepo) page(ctx context.Context, page *common.PageReq, req *repo.TotpGetReq) ([]*model.Totp, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.Totp.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.Totp, 0, len(list))
	for _, row := range list {
		result = append(result, &model.Totp{
			ID:         row.ID,
			UserID:     row.UserID,
			Enable:     row.Enable,
			EnableTime: row.EnableTime,
			Secret:     row.Secret,
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *TotpRepo) upsertEnabledByUserID(ctx context.Context, userID int64, secret string) (*model.Totp, error) {
	tx := r.getClient(ctx)
	existing, err := r.get(ctx, &repo.TotpGetReq{
		UserID: &userID,
	})
	if err != nil {
		return nil, err
	}
	if existing != nil {
		saved, err := tx.Totp.UpdateOneID(existing.ID).
			SetEnable(true).
			SetEnableTime(time.Now()).
			SetSecret(secret).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Totp{
			ID:         saved.ID,
			UserID:     saved.UserID,
			Enable:     saved.Enable,
			EnableTime: saved.EnableTime,
			Secret:     saved.Secret,
		}, nil
	}
	saved, err := tx.Totp.Create().
		SetUserID(userID).
		SetEnable(true).
		SetEnableTime(time.Now()).
		SetSecret(secret).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Totp{
		ID:         saved.ID,
		UserID:     saved.UserID,
		Enable:     saved.Enable,
		EnableTime: saved.EnableTime,
		Secret:     saved.Secret,
	}, nil
}

func (r *TotpRepo) disableByUserID(ctx context.Context, userID int64) (*model.Totp, error) {
	tx := r.getClient(ctx)
	existing, err := r.get(ctx, &repo.TotpGetReq{
		UserID: &userID,
	})
	if err != nil {
		return nil, err
	}
	saved, err := tx.Totp.UpdateOneID(existing.ID).
		SetEnable(false).
		SetSecret("").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Totp{
		ID:         saved.ID,
		UserID:     saved.UserID,
		Enable:     saved.Enable,
		EnableTime: saved.EnableTime,
		Secret:     saved.Secret,
	}, nil
}

func (r *TotpRepo) getQuery(query *gen.TotpQuery, req *repo.TotpGetReq) *gen.TotpQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(totp.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(totp.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(totp.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(totp.UserIDIn(req.UserIDs...))
	}
	if req.Enable != nil {
		query = query.Where(totp.Enable(*req.Enable))
	}
	return query
}
