package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
)

type NotificationStationTemplateRepo interface {
	Upsert(ctx context.Context, template *model.NotificationStationTemplate) (*model.NotificationStationTemplate, error)
	BulkUpsert(ctx context.Context, templates []*model.NotificationStationTemplate) error
	Get(ctx context.Context, query *NotificationStationTemplateQuery) (*model.NotificationStationTemplate, error)
	List(ctx context.Context, query *NotificationStationTemplateQuery) ([]*model.NotificationStationTemplate, error)
	Map(ctx context.Context, query *NotificationStationTemplateQuery) (map[int64]*model.NotificationStationTemplate, error)
	Count(ctx context.Context, query *NotificationStationTemplateQuery) (int, error)
	Page(ctx context.Context, query *NotificationStationTemplateQuery) (*NotificationStationTemplatePageResp, error)
}

type NotificationStationTemplatePageResp struct {
	Rows []*model.NotificationStationTemplate
	Page *base.PageResp
}

type NotificationStationTemplateQuery struct {
	Page    *base.PageRequest
	ID      *int64
	IDs     []int64
	RuleID  *int64
	RuleIDs []int64
}
