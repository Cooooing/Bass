package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

type NotificationRuleRepo interface {
	Upsert(ctx context.Context, rule *model.NotificationRule) (*model.NotificationRule, error)
	BulkUpsert(ctx context.Context, rules []*model.NotificationRule) error
	Get(ctx context.Context, query *NotificationRuleQuery) (*model.NotificationRule, error)
	List(ctx context.Context, query *NotificationRuleQuery) ([]*model.NotificationRule, error)
	Map(ctx context.Context, query *NotificationRuleQuery) (map[int64]*model.NotificationRule, error)
	Count(ctx context.Context, query *NotificationRuleQuery) (int, error)
	Page(ctx context.Context, query *NotificationRuleQuery) (*NotificationRulePageResp, error)
}

type NotificationRulePageResp struct {
	Rows []*model.NotificationRule
	Page *base.PageResp
}

type NotificationRuleQuery struct {
	Page       *base.PageRequest
	ID         *int64
	IDs        []int64
	EventType  *commonenum.EventType
	EventTypes []commonenum.EventType
	Channel    *notifyenum.NotificationChannel
	Channels   []notifyenum.NotificationChannel
	Language   *notifyenum.Language
	Languages  []notifyenum.Language
	Enabled    *bool
}
