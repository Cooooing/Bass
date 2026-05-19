package repo

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationmeta"
	notifyenum "notify/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.NotificationMetaRepo = (*NotificationMetaRepo)(nil)

type NotificationMetaRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
}

func NewNotificationMetaRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
) repo.NotificationMetaRepo {
	return &NotificationMetaRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
	}
}

func (r *NotificationMetaRepo) Save(ctx context.Context, tx *gen.Client, u *model.NotificationMeta) (*model.NotificationMeta, error) {
	save, err := tx.NotificationMeta.Create().
		SetUUID(u.UUID).
		SetEventType(u.EventType).
		SetMeta(u.Meta).
		SetTitle(u.Title).
		SetContent(u.Content).
		SetIsGlobal(u.IsGlobal).
		SetStatus(u.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationMeta{NotificationMeta: save}, nil
}

func (r *NotificationMetaRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.NotificationMetaGetReq) (*model.NotificationMeta, error) {
	query := tx.NotificationMeta.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cerrors.ErrorBadRequest("notification meta is not found")
	}
	return &model.NotificationMeta{NotificationMeta: t}, err
}

func (r *NotificationMetaRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, error) {
	var (
		records []*model.NotificationMeta
		err     error
	)
	query := tx.NotificationMeta.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		records = append(records, &model.NotificationMeta{NotificationMeta: item})
	}
	return records, nil
}

func (r *NotificationMetaRepo) GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *repo.NotificationMetaGetReq) ([]*model.NotificationMeta, *common.PageReply, error) {
	var (
		notificationMetas []*model.NotificationMeta
		err               error
	)
	page = constant.PageValid(page)
	query := tx.NotificationMeta.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		notificationMetas = append(notificationMetas, &model.NotificationMeta{NotificationMeta: item})
	}
	return notificationMetas, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *NotificationMetaRepo) getQuery(query *gen.NotificationMetaQuery, req *repo.NotificationMetaGetReq) *gen.NotificationMetaQuery {
	if req.NotificationMetaId != nil {
		query = query.Where(notificationmeta.IDEQ(*req.NotificationMetaId))
	}
	if len(req.NotificationMeraIds) > 0 {
		query = query.Where(notificationmeta.IDIn(req.NotificationMeraIds...))
	}
	if req.UUID != nil {
		query = query.Where(notificationmeta.UUIDEQ(*req.UUID))
	}
	if len(req.UUIDs) > 0 {
		query = query.Where(notificationmeta.UUIDIn(req.UUIDs...))
	}
	if req.EventType != nil {
		dbEventType, _ := notifyenum.EventTypeMap.ToEnum(*req.EventType)
		query = query.Where(notificationmeta.EventTypeEQ(notificationmeta.EventType(dbEventType)))
	}
	if req.Status != nil {
		dbStatus, _ := notifyenum.NotificationStatusMap.ToEnum(*req.Status)
		query = query.Where(notificationmeta.StatusEQ(notificationmeta.Status(dbStatus)))
	}
	return query
}
