package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
)

type NotificationEmailTemplateRepo interface {
	Upsert(ctx context.Context, template *model.NotificationEmailTemplate) (*model.NotificationEmailTemplate, error)
	BulkUpsert(ctx context.Context, templates []*model.NotificationEmailTemplate) error
	Get(ctx context.Context, query *NotificationEmailTemplateQuery) (*model.NotificationEmailTemplate, error)
	List(ctx context.Context, query *NotificationEmailTemplateQuery) ([]*model.NotificationEmailTemplate, error)
	Map(ctx context.Context, query *NotificationEmailTemplateQuery) (map[int64]*model.NotificationEmailTemplate, error)
	Count(ctx context.Context, query *NotificationEmailTemplateQuery) (int, error)
	Page(ctx context.Context, query *NotificationEmailTemplateQuery) (*NotificationEmailTemplatePageResp, error)
}

type NotificationEmailTemplatePageResp struct {
	Rows []*model.NotificationEmailTemplate
	Page *base.PageResp
}

type NotificationEmailTemplateQuery struct {
	Page    *base.PageRequest
	ID      *int64
	IDs     []int64
	RuleID  *int64
	RuleIDs []int64
}
