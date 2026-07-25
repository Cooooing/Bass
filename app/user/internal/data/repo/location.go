package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/location"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.LocationRepo = (*LocationRepo)(nil)

type LocationRepo struct {
	db *gen.Client
}

func NewLocationRepo(
	db *gen.Client,
) repo.LocationRepo {
	return &LocationRepo{
		db: db,
	}
}

func (r *LocationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *LocationRepo) Get(ctx context.Context, req *repo.LocationGetReq) (*model.Location, error) {
	location, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (r *LocationRepo) List(ctx context.Context, req *repo.LocationGetReq) ([]*model.Location, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *LocationRepo) Map(ctx context.Context, req *repo.LocationGetReq) (map[int64]*model.Location, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *LocationRepo) Count(ctx context.Context, req *repo.LocationGetReq) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *LocationRepo) Page(ctx context.Context, req *repo.LocationPageReq) (*repo.LocationPageResp, error) {
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
	return &repo.LocationPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *LocationRepo) UpsertByUserID(ctx context.Context, location *model.Location) (*model.Location, error) {
	location, err := r.upsertByUserID(ctx, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (r *LocationRepo) Update(ctx context.Context, location *model.Location) (*model.Location, error) {
	location, err := r.update(ctx, location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (r *LocationRepo) get(ctx context.Context, req *repo.LocationGetReq) (*model.Location, error) {
	tx := r.getClient(ctx)
	query := tx.Location.Query()
	query = r.getQuery(query, req)
	l, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Location{
		ID:       l.ID,
		UserID:   l.UserID,
		Country:  l.Country,
		Province: l.Province,
		City:     l.City,
	}, nil
}

func (r *LocationRepo) list(ctx context.Context, req *repo.LocationGetReq) ([]*model.Location, error) {
	tx := r.getClient(ctx)
	query := tx.Location.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Location, 0, len(list))
	for _, l := range list {
		result = append(result, &model.Location{
			ID:       l.ID,
			UserID:   l.UserID,
			Country:  l.Country,
			Province: l.Province,
			City:     l.City,
		})
	}
	return result, nil
}

func (r *LocationRepo) mapRows(ctx context.Context, req *repo.LocationGetReq) (map[int64]*model.Location, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Location, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *LocationRepo) count(ctx context.Context, req *repo.LocationGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.Location.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *LocationRepo) page(ctx context.Context, page *common.PageReq, req *repo.LocationGetReq) ([]*model.Location, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.Location.Query()
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
	result := make([]*model.Location, 0, len(list))
	for _, l := range list {
		result = append(result, &model.Location{
			ID:       l.ID,
			UserID:   l.UserID,
			Country:  l.Country,
			Province: l.Province,
			City:     l.City,
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *LocationRepo) upsertByUserID(ctx context.Context, l *model.Location) (*model.Location, error) {
	existing, err := r.get(ctx, &repo.LocationGetReq{
		UserID: new(l.UserID),
	})
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		saved, err := tx.Location.Create().
			SetUserID(l.UserID).
			SetNillableCountry(l.Country).
			SetNillableProvince(l.Province).
			SetNillableCity(l.City).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Location{
			ID:       saved.ID,
			UserID:   saved.UserID,
			Country:  saved.Country,
			Province: saved.Province,
			City:     saved.City,
		}, nil
	}
	l.ID = existing.ID
	return r.update(ctx, l)
}

func (r *LocationRepo) update(ctx context.Context, l *model.Location) (*model.Location, error) {
	tx := r.getClient(ctx)
	saved, err := tx.Location.UpdateOneID(l.ID).
		SetNillableCountry(l.Country).
		SetNillableProvince(l.Province).
		SetNillableCity(l.City).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Location{
		ID:       saved.ID,
		UserID:   saved.UserID,
		Country:  saved.Country,
		Province: saved.Province,
		City:     saved.City,
	}, nil
}

func (r *LocationRepo) getQuery(query *gen.LocationQuery, req *repo.LocationGetReq) *gen.LocationQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(location.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(location.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(location.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(location.UserIDIn(req.UserIDs...))
	}
	return query
}
