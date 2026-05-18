package repo

import (
	"common/api/gen/common/enums"
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

func (r *NotificationSettingRepo) GetByUser(ctx context.Context, userID int64) ([]*model.NotificationSetting, error) {
	list, err := r.db.NotificationSetting.Query().
		Where(notificationsetting.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationSetting, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationSetting{NotificationSetting: item})
	}
	return result, nil
}

func (r *NotificationSettingRepo) GetByUserAndEvent(ctx context.Context, userID int64, eventType enums.EventType) ([]*model.NotificationSetting, error) {
	dbEventType, _ := notifyenum.EventTypeMap.ToEnum(eventType)
	list, err := r.db.NotificationSetting.Query().
		Where(
			notificationsetting.UserIDEQ(userID),
			notificationsetting.EventTypeEQ(notificationsetting.EventType(dbEventType)),
		).
		All(ctx)
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
