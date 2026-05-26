package repo

import (
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
)

type NotificationSettingRepo interface {
	List(ctx context.Context, req *NotificationSettingGetReq) ([]*model.NotificationSetting, error)
	Save(ctx context.Context, pref *model.NotificationSetting) (*model.NotificationSetting, error)
	Upsert(ctx context.Context, pref *model.NotificationSetting) (*model.NotificationSetting, error)
}

type NotificationSettingGetReq struct {
	UserID    *int64
	UserIDs   []int64
	EventType *enums.EventType
	Channel   *v1.NotificationChannel
}
