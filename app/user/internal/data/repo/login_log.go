package repo

import (
	commonenum "common/pkg/enum"
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

func NewLoginLogRepo(
	db *gen.Client,
) repo.LoginLogRepo {
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

func (r *LoginLogRepo) Create(ctx context.Context, log *model.LoginLog) (*model.LoginLog, error) {
	return r.create(ctx, log)
}

func (r *LoginLogRepo) Get(ctx context.Context, req *repo.LoginLogGetReq) (*model.LoginLog, error) {
	return r.get(ctx, req)
}

func (r *LoginLogRepo) List(ctx context.Context, req *repo.LoginLogGetReq) ([]*model.LoginLog, error) {
	return r.list(ctx, req)
}

func (r *LoginLogRepo) Map(ctx context.Context, req *repo.LoginLogGetReq) (map[int64]*model.LoginLog, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.LoginLog, len(rows))
	for _, item := range rows {
		result[item.ID] = item
	}
	return result, nil
}

func (r *LoginLogRepo) Count(ctx context.Context, req *repo.LoginLogGetReq) (int, error) {
	query := r.getClient(ctx).LoginLog.Query()
	return r.getQuery(query, req).Count(ctx)
}

func (r *LoginLogRepo) Page(ctx context.Context, req *repo.LoginLogPageReq) (*repo.LoginLogPageResp, error) {
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
	return &repo.LoginLogPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *LoginLogRepo) create(ctx context.Context, l *model.LoginLog) (*model.LoginLog, error) {
	create := r.getClient(ctx).LoginLog.Create().
		SetNillableUserID(l.UserID).
		SetAccountInput(l.AccountInput).
		SetLoginType(loginlog.LoginType(l.LoginType)).
		SetRealm(loginlog.Realm(l.Realm)).
		SetStatus(loginlog.Status(l.Status)).
		SetSessionID(l.SessionID).
		SetNillableIP(l.IP).
		SetNillableCountry(l.Country).
		SetNillableCountryCode(l.CountryCode).
		SetNillableProvince(l.Province).
		SetNillableCity(l.City).
		SetNillableIsp(l.ISP).
		SetNillableUserAgent(l.UserAgent).
		SetOsName(l.OSName).
		SetOsVersion(l.OSVersion).
		SetBrowserName(l.BrowserName).
		SetBrowserVersion(l.BrowserVersion).
		SetAppName(l.AppName).
		SetAppVersion(l.AppVersion)
	if l.FailureReason != nil {
		create.SetFailureReason(loginlog.FailureReason(*l.FailureReason))
	}
	if l.ClientType != nil {
		create.SetClientType(loginlog.ClientType(*l.ClientType))
	}
	if l.DeviceType != nil {
		create.SetDeviceType(loginlog.DeviceType(*l.DeviceType))
	}
	created, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.loginLog(created), nil
}

func (r *LoginLogRepo) get(ctx context.Context, req *repo.LoginLogGetReq) (*model.LoginLog, error) {
	query := r.getClient(ctx).LoginLog.Query()
	query = r.getQuery(query, req)
	if req != nil && req.LastSuccess {
		query = query.Order(gen.Desc(loginlog.FieldCreatedAt))
	}
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.loginLog(row), nil
}

func (r *LoginLogRepo) list(ctx context.Context, req *repo.LoginLogGetReq) ([]*model.LoginLog, error) {
	query := r.getClient(ctx).LoginLog.Query()
	rows, err := r.getQuery(query, req).Order(gen.Desc(loginlog.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.LoginLog, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.loginLog(row))
	}
	return result, nil
}

func (r *LoginLogRepo) page(ctx context.Context, page *common.PageReq, req *repo.LoginLogGetReq) ([]*model.LoginLog, *common.PageResp, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).LoginLog.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	rows, err := query.Order(gen.Desc(loginlog.FieldCreatedAt)).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.LoginLog, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.loginLog(row))
	}
	return result, &common.PageResp{
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

func (r *LoginLogRepo) loginLog(row *gen.LoginLog) *model.LoginLog {
	if row == nil {
		return nil
	}
	result := &model.LoginLog{
		ID:             row.ID,
		UserID:         row.UserID,
		AccountInput:   row.AccountInput,
		LoginType:      enum.LoginType(row.LoginType),
		Realm:          commonenum.LoginRealm(row.Realm),
		Status:         enum.LoginStatus(row.Status),
		SessionID:      row.SessionID,
		IP:             row.IP,
		Country:        row.Country,
		CountryCode:    row.CountryCode,
		Province:       row.Province,
		City:           row.City,
		ISP:            row.Isp,
		UserAgent:      row.UserAgent,
		OSName:         row.OsName,
		OSVersion:      row.OsVersion,
		BrowserName:    row.BrowserName,
		BrowserVersion: row.BrowserVersion,
		AppName:        row.AppName,
		AppVersion:     row.AppVersion,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
	if row.FailureReason != nil {
		result.FailureReason = new(enum.LoginFailureReason(*row.FailureReason))
	}
	if row.ClientType != nil {
		result.ClientType = new(enum.ClientType(*row.ClientType))
	}
	if row.DeviceType != nil {
		result.DeviceType = new(enum.DeviceType(*row.DeviceType))
	}
	return result
}
