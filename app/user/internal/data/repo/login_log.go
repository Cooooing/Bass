package repo

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/loginlog"
	"user/internal/enum"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.LoginLogRepo = (*LoginLogRepo)(nil)

type LoginLogRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewLoginLogRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.LoginLogRepo {
	return &LoginLogRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *LoginLogRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func loginLogToDomain(l *gen.LoginLog) *model.LoginLog {
	return &model.LoginLog{
		ID:            l.ID,
		UserID:        l.UserID,
		Account:       l.Account,
		LoginMethod:   enum.LoginMethod(l.LoginMethod),
		Status:        enum.LoginStatus(l.Status),
		FailureReason: l.FailureReason,
		IP:            l.IP,
		Country:       l.Country,
		CountryCode:   l.CountryCode,
		Province:      l.Province,
		City:          l.City,
		ISP:           l.Isp,
		UserAgent:     l.UserAgent,
		DeviceID:      l.DeviceID,
		DeviceName:    l.DeviceName,
		Platform:      l.Platform,
		OS:            l.Os,
		Browser:       l.Browser,
		RequestID:     l.RequestID,
		CreatedAt:     l.CreatedAt,
		UpdatedAt:     l.UpdatedAt,
	}
}

func (r *LoginLogRepo) Create(ctx context.Context, l *model.LoginLog) (*model.LoginLog, error) {
	tx := r.getClient(ctx)
	created, err := tx.LoginLog.Create().
		SetNillableUserID(l.UserID).
		SetAccount(l.Account).
		SetLoginMethod(loginlog.LoginMethod(l.LoginMethod)).
		SetStatus(loginlog.Status(l.Status)).
		SetNillableFailureReason(l.FailureReason).
		SetNillableIP(l.IP).
		SetNillableCountry(l.Country).
		SetNillableCountryCode(l.CountryCode).
		SetNillableProvince(l.Province).
		SetNillableCity(l.City).
		SetNillableIsp(l.ISP).
		SetNillableUserAgent(l.UserAgent).
		SetNillableDeviceID(l.DeviceID).
		SetNillableDeviceName(l.DeviceName).
		SetNillablePlatform(l.Platform).
		SetNillableOs(l.OS).
		SetNillableBrowser(l.Browser).
		SetNillableRequestID(l.RequestID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return loginLogToDomain(created), nil
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
	return loginLogToDomain(l), nil
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
		result = append(result, loginLogToDomain(l))
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
		result = append(result, loginLogToDomain(l))
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
	if req.Account != nil {
		query = query.Where(loginlog.Account(*req.Account))
	}
	if req.Status != nil {
		query = query.Where(loginlog.StatusEQ(loginlog.Status(*req.Status)))
	}
	if req.IP != nil {
		query = query.Where(loginlog.IP(*req.IP))
	}
	return query
}
