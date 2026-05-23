package repo

import (
	commonClient "common/pkg/client"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationsetting"
	notifyenum "notify/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.NotificationSettingRepo = (*NotificationSettingRepo)(nil)

type NotificationSettingRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
}

func NewNotificationSettingRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
) repo.NotificationSettingRepo {
	return &NotificationSettingRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
	}
}

func (r *NotificationSettingRepo) List(ctx context.Context, req *repo.NotificationSettingGetReq) ([]*model.NotificationSetting, error) {
	query := r.db.NotificationSetting.Query()
	if req != nil && req.UserID != nil {
		query = query.Where(notificationsetting.UserIDEQ(*req.UserID))
	}
	if req != nil && req.EventType != nil {
		dbEventType, _ := notifyenum.EventTypeMap.ToEnum(*req.EventType)
		query = query.Where(notificationsetting.EventTypeEQ(notificationsetting.EventType(dbEventType)))
	}
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationSetting, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationSetting{NotificationSetting: item})
	}
	return result, nil
}

func (r *NotificationSettingRepo) Save(ctx context.Context, tx *gen.Client, pref *model.NotificationSetting) (*model.NotificationSetting, error) {
	save, err := tx.NotificationSetting.Create().
		SetUserID(pref.UserID).
		SetEventType(pref.EventType).
		SetChannel(pref.Channel).
		SetEnable(pref.Enable).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationSetting{NotificationSetting: save}, nil
}
