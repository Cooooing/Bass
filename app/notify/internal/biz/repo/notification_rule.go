package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	"context"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

type NotificationRuleRepo interface {
	Get(ctx context.Context, req *NotificationRuleQuery) (*model.NotificationRule, error)
	List(ctx context.Context, req *NotificationRuleQuery) ([]*model.NotificationRule, error)
	Map(ctx context.Context, req *NotificationRuleQuery) (map[int64]*model.NotificationRule, error)
	Count(ctx context.Context, req *NotificationRuleQuery) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *NotificationRuleQuery) ([]*model.NotificationRule, *common.PageReply, error)
}

type NotificationRuleQuery struct {
	ID        *int64
	IDs       []int64
	EventType *commonenum.EventType
	Channel   *notifyenum.NotificationChannel
	Language  *notifyenum.Language
	Enabled   *bool
}
