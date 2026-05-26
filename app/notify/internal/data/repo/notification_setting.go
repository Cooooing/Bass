package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationsetting"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ repo.NotificationSettingRepo = (*NotificationSettingRepo)(nil)

type NotificationSettingRepo struct {
	db *gen.Client
}

func NewNotificationSettingRepo(db *gen.Client) repo.NotificationSettingRepo {
	return &NotificationSettingRepo{
		db: db,
	}
}

func (r *NotificationSettingRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationSettingRepo) List(ctx context.Context, req *repo.NotificationSettingGetReq) ([]*model.NotificationSetting, error) {
	query := r.getClient(ctx).NotificationSetting.Query()
	if req != nil && req.UserID != nil {
		query = query.Where(notificationsetting.UserIDEQ(*req.UserID))
	}
	if req != nil && len(req.UserIDs) > 0 {
		query = query.Where(notificationsetting.UserIDIn(req.UserIDs...))
	}
	if req != nil && req.EventType != nil {
		dbEventType, _ := commonenum.EventTypeMap.ToEnum(*req.EventType)
		query = query.Where(notificationsetting.EventTypeEQ(notificationsetting.EventType(dbEventType)))
	}
	if req != nil && req.Channel != nil {
		dbChannel, _ := notifyenum.NotificationChannelMap.ToEnum(*req.Channel)
		query = query.Where(notificationsetting.ChannelEQ(notificationsetting.Channel(dbChannel)))
	}
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationSetting, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationSetting{
			ID:        item.ID,
			UserID:    item.UserID,
			EventType: commonenum.EventType(item.EventType),
			Channel:   notifyenum.NotificationChannel(item.Channel),
			Enable:    item.Enable,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationSettingRepo) Save(ctx context.Context, pref *model.NotificationSetting) (*model.NotificationSetting, error) {
	save, err := r.getClient(ctx).NotificationSetting.Create().
		SetUserID(pref.UserID).
		SetEventType(notificationsetting.EventType(pref.EventType)).
		SetChannel(notificationsetting.Channel(pref.Channel)).
		SetEnable(pref.Enable).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.NotificationSetting{
		ID:        save.ID,
		UserID:    save.UserID,
		EventType: commonenum.EventType(save.EventType),
		Channel:   notifyenum.NotificationChannel(save.Channel),
		Enable:    save.Enable,
		CreatedAt: save.CreatedAt,
		UpdatedAt: save.UpdatedAt,
	}, nil
}

func (r *NotificationSettingRepo) Upsert(ctx context.Context, pref *model.NotificationSetting) (*model.NotificationSetting, error) {
	client := r.getClient(ctx)
	exist, err := client.NotificationSetting.Query().
		Where(
			notificationsetting.UserIDEQ(pref.UserID),
			notificationsetting.EventTypeEQ(notificationsetting.EventType(pref.EventType)),
			notificationsetting.ChannelEQ(notificationsetting.Channel(pref.Channel)),
		).
		Only(ctx)
	if err == nil {
		update, err := client.NotificationSetting.UpdateOneID(exist.ID).
			SetEnable(pref.Enable).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.NotificationSetting{
			ID:        update.ID,
			UserID:    update.UserID,
			EventType: commonenum.EventType(update.EventType),
			Channel:   notifyenum.NotificationChannel(update.Channel),
			Enable:    update.Enable,
			CreatedAt: update.CreatedAt,
			UpdatedAt: update.UpdatedAt,
		}, nil
	}
	if !gen.IsNotFound(err) {
		return nil, err
	}
	return r.Save(ctx, pref)
}
