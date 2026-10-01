package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/privacysetting"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.PrivacySettingRepo = (*PrivacySettingRepo)(nil)

type PrivacySettingRepo struct {
	db *gen.Client
}

func NewPrivacySettingRepo(
	db *gen.Client,
) repo.PrivacySettingRepo {
	return &PrivacySettingRepo{
		db: db,
	}
}

func (r *PrivacySettingRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *PrivacySettingRepo) Get(ctx context.Context, req *repo.PrivacySettingGetReq) (*model.PrivacySetting, error) {
	setting, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *PrivacySettingRepo) List(ctx context.Context, req *repo.PrivacySettingGetReq) ([]*model.PrivacySetting, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *PrivacySettingRepo) Map(ctx context.Context, req *repo.PrivacySettingGetReq) (map[int64]*model.PrivacySetting, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *PrivacySettingRepo) Count(ctx context.Context, req *repo.PrivacySettingGetReq) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PrivacySettingRepo) Page(ctx context.Context, req *repo.PrivacySettingPageReq) (*repo.PrivacySettingPageResp, error) {
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
	return &repo.PrivacySettingPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *PrivacySettingRepo) UpsertByUserID(ctx context.Context, setting *model.PrivacySetting) (*model.PrivacySetting, error) {
	setting, err := r.upsertByUserID(ctx, setting)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *PrivacySettingRepo) Update(ctx context.Context, setting *model.PrivacySetting) (*model.PrivacySetting, error) {
	setting, err := r.update(ctx, setting)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *PrivacySettingRepo) get(ctx context.Context, req *repo.PrivacySettingGetReq) (*model.PrivacySetting, error) {
	tx := r.getClient(ctx)
	query := tx.PrivacySetting.Query()
	query = r.getQuery(query, req)
	p, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.PrivacySetting{
		ID:                   p.ID,
		UserID:               p.UserID,
		PublicPoints:         new(p.PublicPoints),
		PublicFollowerList:   new(p.PublicFollowerList),
		PublicArticleList:    new(p.PublicArticleList),
		PublicCommentList:    new(p.PublicCommentList),
		PublicOnlineStatus:   new(p.PublicOnlineStatus),
		PublicLocation:       new(p.PublicLocation),
		PublicMoonbreezeList: new(p.PublicMoonbreezeList),
		PublicFollowingList:  new(p.PublicFollowingList),
	}, nil
}

func (r *PrivacySettingRepo) list(ctx context.Context, req *repo.PrivacySettingGetReq) ([]*model.PrivacySetting, error) {
	tx := r.getClient(ctx)
	query := tx.PrivacySetting.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.PrivacySetting, 0, len(list))
	for _, p := range list {
		result = append(result, &model.PrivacySetting{
			ID:                   p.ID,
			UserID:               p.UserID,
			PublicPoints:         new(p.PublicPoints),
			PublicFollowerList:   new(p.PublicFollowerList),
			PublicArticleList:    new(p.PublicArticleList),
			PublicCommentList:    new(p.PublicCommentList),
			PublicOnlineStatus:   new(p.PublicOnlineStatus),
			PublicLocation:       new(p.PublicLocation),
			PublicMoonbreezeList: new(p.PublicMoonbreezeList),
			PublicFollowingList:  new(p.PublicFollowingList),
		})
	}
	return result, nil
}

func (r *PrivacySettingRepo) mapRows(ctx context.Context, req *repo.PrivacySettingGetReq) (map[int64]*model.PrivacySetting, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.PrivacySetting, len(list))
	for _, item := range list {
		result[item.UserID] = item
	}
	return result, nil
}

func (r *PrivacySettingRepo) count(ctx context.Context, req *repo.PrivacySettingGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.PrivacySetting.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *PrivacySettingRepo) page(ctx context.Context, page *common.PageReq, req *repo.PrivacySettingGetReq) ([]*model.PrivacySetting, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.PrivacySetting.Query()
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
	result := make([]*model.PrivacySetting, 0, len(list))
	for _, p := range list {
		result = append(result, &model.PrivacySetting{
			ID:                   p.ID,
			UserID:               p.UserID,
			PublicPoints:         new(p.PublicPoints),
			PublicFollowerList:   new(p.PublicFollowerList),
			PublicArticleList:    new(p.PublicArticleList),
			PublicCommentList:    new(p.PublicCommentList),
			PublicOnlineStatus:   new(p.PublicOnlineStatus),
			PublicLocation:       new(p.PublicLocation),
			PublicMoonbreezeList: new(p.PublicMoonbreezeList),
			PublicFollowingList:  new(p.PublicFollowingList),
		})
	}
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *PrivacySettingRepo) upsertByUserID(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error) {
	existing, err := r.get(ctx, &repo.PrivacySettingGetReq{
		UserID: new(p.UserID),
	})
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		create := tx.PrivacySetting.Create().
			SetUserID(p.UserID)
		if p.PublicPoints != nil {
			create.SetPublicPoints(*p.PublicPoints)
		}
		if p.PublicFollowerList != nil {
			create.SetPublicFollowerList(*p.PublicFollowerList)
		}
		if p.PublicArticleList != nil {
			create.SetPublicArticleList(*p.PublicArticleList)
		}
		if p.PublicCommentList != nil {
			create.SetPublicCommentList(*p.PublicCommentList)
		}
		if p.PublicOnlineStatus != nil {
			create.SetPublicOnlineStatus(*p.PublicOnlineStatus)
		}
		if p.PublicLocation != nil {
			create.SetPublicLocation(*p.PublicLocation)
		}
		if p.PublicMoonbreezeList != nil {
			create.SetPublicMoonbreezeList(*p.PublicMoonbreezeList)
		}
		if p.PublicFollowingList != nil {
			create.SetPublicFollowingList(*p.PublicFollowingList)
		}
		saved, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.PrivacySetting{
			ID:                   saved.ID,
			UserID:               saved.UserID,
			PublicPoints:         new(saved.PublicPoints),
			PublicFollowerList:   new(saved.PublicFollowerList),
			PublicArticleList:    new(saved.PublicArticleList),
			PublicCommentList:    new(saved.PublicCommentList),
			PublicOnlineStatus:   new(saved.PublicOnlineStatus),
			PublicLocation:       new(saved.PublicLocation),
			PublicMoonbreezeList: new(saved.PublicMoonbreezeList),
			PublicFollowingList:  new(saved.PublicFollowingList),
		}, nil
	}
	p.ID = existing.ID
	return r.update(ctx, p)
}

func (r *PrivacySettingRepo) update(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error) {
	tx := r.getClient(ctx)
	if p.PublicPoints == nil &&
		p.PublicFollowerList == nil &&
		p.PublicArticleList == nil &&
		p.PublicCommentList == nil &&
		p.PublicOnlineStatus == nil &&
		p.PublicLocation == nil &&
		p.PublicMoonbreezeList == nil &&
		p.PublicFollowingList == nil {
		saved, err := tx.PrivacySetting.Get(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		return &model.PrivacySetting{
			ID:                   saved.ID,
			UserID:               saved.UserID,
			PublicPoints:         new(saved.PublicPoints),
			PublicFollowerList:   new(saved.PublicFollowerList),
			PublicArticleList:    new(saved.PublicArticleList),
			PublicCommentList:    new(saved.PublicCommentList),
			PublicOnlineStatus:   new(saved.PublicOnlineStatus),
			PublicLocation:       new(saved.PublicLocation),
			PublicMoonbreezeList: new(saved.PublicMoonbreezeList),
			PublicFollowingList:  new(saved.PublicFollowingList),
		}, nil
	}
	update := tx.PrivacySetting.UpdateOneID(p.ID)
	if p.PublicPoints != nil {
		update.SetPublicPoints(*p.PublicPoints)
	}
	if p.PublicFollowerList != nil {
		update.SetPublicFollowerList(*p.PublicFollowerList)
	}
	if p.PublicArticleList != nil {
		update.SetPublicArticleList(*p.PublicArticleList)
	}
	if p.PublicCommentList != nil {
		update.SetPublicCommentList(*p.PublicCommentList)
	}
	if p.PublicOnlineStatus != nil {
		update.SetPublicOnlineStatus(*p.PublicOnlineStatus)
	}
	if p.PublicLocation != nil {
		update.SetPublicLocation(*p.PublicLocation)
	}
	if p.PublicMoonbreezeList != nil {
		update.SetPublicMoonbreezeList(*p.PublicMoonbreezeList)
	}
	if p.PublicFollowingList != nil {
		update.SetPublicFollowingList(*p.PublicFollowingList)
	}
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PrivacySetting{
		ID:                   saved.ID,
		UserID:               saved.UserID,
		PublicPoints:         new(saved.PublicPoints),
		PublicFollowerList:   new(saved.PublicFollowerList),
		PublicArticleList:    new(saved.PublicArticleList),
		PublicCommentList:    new(saved.PublicCommentList),
		PublicOnlineStatus:   new(saved.PublicOnlineStatus),
		PublicLocation:       new(saved.PublicLocation),
		PublicMoonbreezeList: new(saved.PublicMoonbreezeList),
		PublicFollowingList:  new(saved.PublicFollowingList),
	}, nil
}

func (r *PrivacySettingRepo) getQuery(query *gen.PrivacySettingQuery, req *repo.PrivacySettingGetReq) *gen.PrivacySettingQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(privacysetting.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(privacysetting.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(privacysetting.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(privacysetting.UserIDIn(req.UserIDs...))
	}
	return query
}
