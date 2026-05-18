package repo

import (
	"common/api/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/gen"
)

type NotificationSettingRepo interface {
	GetByUser(ctx context.Context, userID int64) ([]*model.NotificationSetting, error)
	GetByUserAndEvent(ctx context.Context, userID int64, eventType enums.EventType) ([]*model.NotificationSetting, error)
	Save(ctx context.Context, tx *gen.Client, pref *model.NotificationSetting) (*model.NotificationSetting, error)
}
