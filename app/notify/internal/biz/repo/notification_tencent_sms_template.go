package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
)

type NotificationTencentSMSTemplateRepo interface {
	Upsert(ctx context.Context, template *model.NotificationTencentSMSTemplate) (*model.NotificationTencentSMSTemplate, error)
	BulkUpsert(ctx context.Context, templates []*model.NotificationTencentSMSTemplate) error
	Get(ctx context.Context, query *NotificationTencentSMSTemplateQuery) (*model.NotificationTencentSMSTemplate, error)
	List(ctx context.Context, query *NotificationTencentSMSTemplateQuery) ([]*model.NotificationTencentSMSTemplate, error)
	Map(ctx context.Context, query *NotificationTencentSMSTemplateQuery) (map[int64]*model.NotificationTencentSMSTemplate, error)
	Count(ctx context.Context, query *NotificationTencentSMSTemplateQuery) (int, error)
	Page(ctx context.Context, query *NotificationTencentSMSTemplateQuery) (*NotificationTencentSMSTemplatePageResp, error)
}

type NotificationTencentSMSTemplatePageResp struct {
	Rows []*model.NotificationTencentSMSTemplate
	Page *base.PageResp
}

type NotificationTencentSMSTemplateQuery struct {
	Page    *base.PageRequest
	ID      *int64
	IDs     []int64
	RuleID  *int64
	RuleIDs []int64
}
