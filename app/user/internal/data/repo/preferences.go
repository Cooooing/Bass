package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/preferences"
	"user/internal/enum"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.PreferencesRepo = (*PreferencesRepo)(nil)

type PreferencesRepo struct {
	db *gen.Client
}

func NewPreferencesRepo(
	db *gen.Client,
) repo.PreferencesRepo {
	return &PreferencesRepo{
		db: db,
	}
}

func (r *PreferencesRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *PreferencesRepo) Get(ctx context.Context, req *repo.PreferencesGetReq) (*model.Preferences, error) {
	preferences, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return preferences, nil
}

func (r *PreferencesRepo) List(ctx context.Context, req *repo.PreferencesGetReq) ([]*model.Preferences, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *PreferencesRepo) Map(ctx context.Context, req *repo.PreferencesGetReq) (map[int64]*model.Preferences, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *PreferencesRepo) Count(ctx context.Context, req *repo.PreferencesGetReq) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PreferencesRepo) Page(ctx context.Context, req *repo.PreferencesPageReq) (*repo.PreferencesPageResp, error) {
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
	return &repo.PreferencesPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *PreferencesRepo) UpsertByUserID(ctx context.Context, preferences *model.Preferences) (*model.Preferences, error) {
	preferences, err := r.upsertByUserID(ctx, preferences)
	if err != nil {
		return nil, err
	}
	return preferences, nil
}

func (r *PreferencesRepo) Update(ctx context.Context, preferences *model.Preferences) (*model.Preferences, error) {
	preferences, err := r.update(ctx, preferences)
	if err != nil {
		return nil, err
	}
	return preferences, nil
}

func (r *PreferencesRepo) get(ctx context.Context, req *repo.PreferencesGetReq) (*model.Preferences, error) {
	tx := r.getClient(ctx)
	query := tx.Preferences.Query()
	query = r.getQuery(query, req)
	p, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Preferences{
		ID:          p.ID,
		UserID:      p.UserID,
		Language:    new(enum.Language(p.Language)),
		Timezone:    new(p.Timezone),
		Theme:       new(p.Theme),
		MobileTheme: new(p.MobileTheme),
	}, nil
}

func (r *PreferencesRepo) list(ctx context.Context, req *repo.PreferencesGetReq) ([]*model.Preferences, error) {
	tx := r.getClient(ctx)
	query := tx.Preferences.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Preferences, 0, len(list))
	for _, p := range list {
		result = append(result, &model.Preferences{
			ID:          p.ID,
			UserID:      p.UserID,
			Language:    new(enum.Language(p.Language)),
			Timezone:    new(p.Timezone),
			Theme:       new(p.Theme),
			MobileTheme: new(p.MobileTheme),
		})
	}
	return result, nil
}

func (r *PreferencesRepo) mapRows(ctx context.Context, req *repo.PreferencesGetReq) (map[int64]*model.Preferences, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Preferences, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *PreferencesRepo) count(ctx context.Context, req *repo.PreferencesGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.Preferences.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *PreferencesRepo) page(ctx context.Context, page *common.PageReq, req *repo.PreferencesGetReq) ([]*model.Preferences, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.Preferences.Query()
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
	result := make([]*model.Preferences, 0, len(list))
	for _, p := range list {
		result = append(result, &model.Preferences{
			ID:          p.ID,
			UserID:      p.UserID,
			Language:    new(enum.Language(p.Language)),
			Timezone:    new(p.Timezone),
			Theme:       new(p.Theme),
			MobileTheme: new(p.MobileTheme),
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *PreferencesRepo) upsertByUserID(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	existing, err := r.get(ctx, &repo.PreferencesGetReq{
		UserID: new(p.UserID),
	})
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		create := tx.Preferences.Create().
			SetUserID(p.UserID)
		if p.Language != nil {
			create.SetLanguage(preferences.Language(*p.Language))
		}
		if p.Timezone != nil {
			create.SetTimezone(*p.Timezone)
		}
		if p.Theme != nil {
			create.SetTheme(*p.Theme)
		}
		if p.MobileTheme != nil {
			create.SetMobileTheme(*p.MobileTheme)
		}
		saved, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Preferences{
			ID:          saved.ID,
			UserID:      saved.UserID,
			Language:    new(enum.Language(saved.Language)),
			Timezone:    new(saved.Timezone),
			Theme:       new(saved.Theme),
			MobileTheme: new(saved.MobileTheme),
		}, nil
	}
	p.ID = existing.ID
	return r.update(ctx, p)
}

func (r *PreferencesRepo) update(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	tx := r.getClient(ctx)
	if p.Language == nil && p.Timezone == nil && p.Theme == nil && p.MobileTheme == nil {
		saved, err := tx.Preferences.Get(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		return &model.Preferences{
			ID:          saved.ID,
			UserID:      saved.UserID,
			Language:    new(enum.Language(saved.Language)),
			Timezone:    new(saved.Timezone),
			Theme:       new(saved.Theme),
			MobileTheme: new(saved.MobileTheme),
		}, nil
	}
	update := tx.Preferences.UpdateOneID(p.ID)
	if p.Language != nil {
		update.SetLanguage(preferences.Language(*p.Language))
	}
	if p.Timezone != nil {
		update.SetTimezone(*p.Timezone)
	}
	if p.Theme != nil {
		update.SetTheme(*p.Theme)
	}
	if p.MobileTheme != nil {
		update.SetMobileTheme(*p.MobileTheme)
	}
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Preferences{
		ID:          saved.ID,
		UserID:      saved.UserID,
		Language:    new(enum.Language(saved.Language)),
		Timezone:    new(saved.Timezone),
		Theme:       new(saved.Theme),
		MobileTheme: new(saved.MobileTheme),
	}, nil
}

func (r *PreferencesRepo) getQuery(query *gen.PreferencesQuery, req *repo.PreferencesGetReq) *gen.PreferencesQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(preferences.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(preferences.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(preferences.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(preferences.UserIDIn(req.UserIDs...))
	}
	return query
}
