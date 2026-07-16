package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/loginlog"
	"user/internal/enum"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.LoginLogRepo = (*LoginLogRepo)(nil)

type LoginLogRepo struct {
	db *gen.Client
}

func NewLoginLogRepo(db *gen.Client) repo.LoginLogRepo {
	return &LoginLogRepo{
		db: db,
	}
}

func (r *LoginLogRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *LoginLogRepo) Create(ctx context.Context, req *repo.LoginLogCreateReq) (*repo.LoginLogCreateResponse, error) {
	log, err := r.create(ctx, req.Log)
	if err != nil {
		return nil, err
	}
	return &repo.LoginLogCreateResponse{Log: log}, nil
}

func (r *LoginLogRepo) Get(ctx context.Context, req *repo.LoginLogGetReq) (*repo.LoginLogGetResponse, error) {
	log, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.LoginLogGetResponse{Log: log}, nil
}

func (r *LoginLogRepo) List(ctx context.Context, req *repo.LoginLogGetReq) (*repo.LoginLogListResponse, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.LoginLogListResponse{Rows: rows}, nil
}

func (r *LoginLogRepo) Map(ctx context.Context, req *repo.LoginLogGetReq) (*repo.LoginLogMapResponse, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.LoginLogMapResponse{Rows: rows}, nil
}

func (r *LoginLogRepo) Count(ctx context.Context, req *repo.LoginLogGetReq) (*repo.LoginLogCountResponse, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.LoginLogCountResponse{Count: count}, nil
}

func (r *LoginLogRepo) Page(ctx context.Context, req *repo.LoginLogPageReq) (*repo.LoginLogPageResponse, error) {
	rows, page, err := r.page(ctx, &common.PageRequest{Page: req.Page.Page, Size: req.Page.Size}, &req.Query)
	if err != nil {
		return nil, err
	}
	resp := repo.PageResponse{}
	if page != nil {
		resp = repo.PageResponse{Total: page.GetTotal(), Page: page.GetPage(), Size: page.GetSize()}
	}
	return &repo.LoginLogPageResponse{Rows: rows, Page: resp}, nil
}
func (r *LoginLogRepo) create(ctx context.Context, l *model.LoginLog) (*model.LoginLog, error) {
	tx := r.getClient(ctx)
	created, err := tx.LoginLog.Create().
		SetNillableUserID(l.UserID).
		SetLoginMethod(loginlog.LoginMethod(l.LoginMethod)).
		SetStatus(loginlog.Status(l.Status)).
		SetNillableIP(l.IP).
		SetNillableCountry(l.Country).
		SetNillableCountryCode(l.CountryCode).
		SetNillableProvince(l.Province).
		SetNillableCity(l.City).
		SetNillableIsp(l.ISP).
		SetNillableUserAgent(l.UserAgent).
		SetNillableDeviceID(l.DeviceID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.LoginLog{
		ID:          created.ID,
		UserID:      created.UserID,
		LoginMethod: enum.LoginMethod(created.LoginMethod),
		Status:      enum.LoginStatus(created.Status),
		IP:          created.IP,
		Country:     created.Country,
		CountryCode: created.CountryCode,
		Province:    created.Province,
		City:        created.City,
		ISP:         created.Isp,
		UserAgent:   created.UserAgent,
		DeviceID:    created.DeviceID,
		CreatedAt:   created.CreatedAt,
		UpdatedAt:   created.UpdatedAt,
	}, nil
}

func (r *LoginLogRepo) get(ctx context.Context, req *repo.LoginLogGetReq) (*model.LoginLog, error) {
	tx := r.getClient(ctx)
	query := tx.LoginLog.Query()
	query = r.getQuery(query, req)
	if req != nil && req.LastSuccess {
		query = query.Order(gen.Desc(loginlog.FieldCreatedAt))
	}
	l, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.LoginLog{
		ID:          l.ID,
		UserID:      l.UserID,
		LoginMethod: enum.LoginMethod(l.LoginMethod),
		Status:      enum.LoginStatus(l.Status),
		IP:          l.IP,
		Country:     l.Country,
		CountryCode: l.CountryCode,
		Province:    l.Province,
		City:        l.City,
		ISP:         l.Isp,
		UserAgent:   l.UserAgent,
		DeviceID:    l.DeviceID,
		CreatedAt:   l.CreatedAt,
		UpdatedAt:   l.UpdatedAt,
	}, nil
}

func (r *LoginLogRepo) list(ctx context.Context, req *repo.LoginLogGetReq) ([]*model.LoginLog, error) {
	tx := r.getClient(ctx)
	query := tx.LoginLog.Query()
	query = r.getQuery(query, req)
	list, err := query.Order(gen.Desc(loginlog.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.LoginLog, 0, len(list))
	for _, l := range list {
		result = append(result, &model.LoginLog{
			ID:          l.ID,
			UserID:      l.UserID,
			LoginMethod: enum.LoginMethod(l.LoginMethod),
			Status:      enum.LoginStatus(l.Status),
			IP:          l.IP,
			Country:     l.Country,
			CountryCode: l.CountryCode,
			Province:    l.Province,
			City:        l.City,
			ISP:         l.Isp,
			UserAgent:   l.UserAgent,
			DeviceID:    l.DeviceID,
			CreatedAt:   l.CreatedAt,
			UpdatedAt:   l.UpdatedAt,
		})
	}
	return result, nil
}

func (r *LoginLogRepo) mapRows(ctx context.Context, req *repo.LoginLogGetReq) (map[int64]*model.LoginLog, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.LoginLog, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *LoginLogRepo) count(ctx context.Context, req *repo.LoginLogGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.LoginLog.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *LoginLogRepo) page(ctx context.Context, page *common.PageRequest, req *repo.LoginLogGetReq) ([]*model.LoginLog, *common.PageResponse, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
	query := tx.LoginLog.Query()
	query = r.getQuery(query, req)

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Order(gen.Desc(loginlog.FieldCreatedAt)).
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	result := make([]*model.LoginLog, 0, len(list))
	for _, l := range list {
		result = append(result, &model.LoginLog{
			ID:          l.ID,
			UserID:      l.UserID,
			LoginMethod: enum.LoginMethod(l.LoginMethod),
			Status:      enum.LoginStatus(l.Status),
			IP:          l.IP,
			Country:     l.Country,
			CountryCode: l.CountryCode,
			Province:    l.Province,
			City:        l.City,
			ISP:         l.Isp,
			UserAgent:   l.UserAgent,
			DeviceID:    l.DeviceID,
			CreatedAt:   l.CreatedAt,
			UpdatedAt:   l.UpdatedAt,
		})
	}
	return result, &common.PageResponse{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *LoginLogRepo) getQuery(query *gen.LoginLogQuery, req *repo.LoginLogGetReq) *gen.LoginLogQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(loginlog.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(loginlog.IDIn(req.IDs...))
	}
	if req.UserID != nil {
		query = query.Where(loginlog.UserID(*req.UserID))
	}
	if len(req.UserIDs) > 0 {
		query = query.Where(loginlog.UserIDIn(req.UserIDs...))
	}
	if req.Status != nil {
		query = query.Where(loginlog.StatusEQ(loginlog.Status(*req.Status)))
	}
	if req.LastSuccess {
		query = query.Where(loginlog.StatusEQ(loginlog.Status(enum.LoginStatusSuccess)))
	}
	if req.IP != nil {
		query = query.Where(loginlog.IP(*req.IP))
	}
	return query
}
