package repo

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/gen"
)

type NotificationSettingRepo interface {
	List(ctx context.Context, req *NotificationSettingGetReq) ([]*model.NotificationSetting, error)
	Save(ctx context.Context, tx *gen.Client, pref *model.NotificationSetting) (*model.NotificationSetting, error)
}

type NotificationSettingGetReq struct {
	UserID    *int64
	EventType *enums.EventType
}
