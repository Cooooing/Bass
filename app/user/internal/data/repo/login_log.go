package repo

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/loginlog"
	"user/internal/enum"

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

func (r *LoginLogRepo) Create(ctx context.Context, l *model.LoginLog) (*model.LoginLog, error) {
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

func (r *LoginLogRepo) FindLastSuccessByUserID(ctx context.Context, userID int64) (*model.LoginLog, error) {
	tx := r.getClient(ctx)
	l, err := tx.LoginLog.Query().
		Where(loginlog.UserID(userID)).
		Where(loginlog.StatusEQ(loginlog.Status(enum.LoginStatusSuccess))).
		Order(gen.Desc(loginlog.FieldCreatedAt)).
		First(ctx)
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

func (r *LoginLogRepo) List(ctx context.Context, req *repo.LoginLogGetReq) ([]*model.LoginLog, error) {
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

func (r *LoginLogRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.LoginLogGetReq) ([]*model.LoginLog, *common.PageReply, error) {
	tx := r.getClient(ctx)
	page = constant.PageValid(page)
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
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *LoginLogRepo) getQuery(query *gen.LoginLogQuery, req *repo.LoginLogGetReq) *gen.LoginLogQuery {
	if req == nil {
		return query
	}
	if req.UserID != nil {
		query = query.Where(loginlog.UserID(*req.UserID))
	}
	if req.Status != nil {
		query = query.Where(loginlog.StatusEQ(loginlog.Status(*req.Status)))
	}
	if req.IP != nil {
		query = query.Where(loginlog.IP(*req.IP))
	}
	return query
}
