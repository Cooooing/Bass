package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

type NotificationRuleRepo interface {
	ListEnabled(ctx context.Context, eventType commonenum.EventType, language notifyenum.Language) ([]*model.NotificationRule, error)
}
