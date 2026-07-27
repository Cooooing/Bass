package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
)

type NotificationLarkWebhookTemplateRepo interface {
	Upsert(ctx context.Context, template *model.NotificationLarkWebhookTemplate) (*model.NotificationLarkWebhookTemplate, error)
	BulkUpsert(ctx context.Context, templates []*model.NotificationLarkWebhookTemplate) error
	Get(ctx context.Context, query *NotificationLarkWebhookTemplateQuery) (*model.NotificationLarkWebhookTemplate, error)
	List(ctx context.Context, query *NotificationLarkWebhookTemplateQuery) ([]*model.NotificationLarkWebhookTemplate, error)
	Map(ctx context.Context, query *NotificationLarkWebhookTemplateQuery) (map[int64]*model.NotificationLarkWebhookTemplate, error)
	Count(ctx context.Context, query *NotificationLarkWebhookTemplateQuery) (int, error)
	Page(ctx context.Context, query *NotificationLarkWebhookTemplateQuery) (*NotificationLarkWebhookTemplatePageResp, error)
}

type NotificationLarkWebhookTemplatePageResp struct {
	Rows []*model.NotificationLarkWebhookTemplate
	Page *base.PageResp
}

type NotificationLarkWebhookTemplateQuery struct {
	Page    *base.PageRequest
	ID      *int64
	IDs     []int64
	RuleID  *int64
	RuleIDs []int64
}
