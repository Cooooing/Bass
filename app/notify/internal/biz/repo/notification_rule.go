package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

type NotificationRuleRepo interface {
	Get(ctx context.Context, req *NotificationRuleGetReq) (*NotificationRuleGetResponse, error)
	List(ctx context.Context, req *NotificationRuleListReq) (*NotificationRuleListResponse, error)
	Map(ctx context.Context, req *NotificationRuleMapReq) (*NotificationRuleMapResponse, error)
	Count(ctx context.Context, req *NotificationRuleCountReq) (*NotificationRuleCountResponse, error)
	Page(ctx context.Context, req *NotificationRulePageReq) (*NotificationRulePageResponse, error)
}

type NotificationRuleGetReq struct {
	Query *NotificationRuleQuery
}

type NotificationRuleGetResponse struct {
	Item *model.NotificationRule
}

type NotificationRuleListReq struct {
	Query *NotificationRuleQuery
}

type NotificationRuleListResponse struct {
	Rows []*model.NotificationRule
}

type NotificationRuleMapReq struct {
	Query *NotificationRuleQuery
}

type NotificationRuleMapResponse struct {
	Rows map[int64]*model.NotificationRule
}

type NotificationRuleCountReq struct {
	Query *NotificationRuleQuery
}

type NotificationRuleCountResponse struct {
	Count int
}

type NotificationRulePageReq struct {
	Query *NotificationRuleQuery
}

type NotificationRulePageResponse struct {
	Rows []*model.NotificationRule
	Page *base.PageResponse
}

type NotificationRuleQuery struct {
	Page      *base.PageRequest
	ID        *int64
	IDs       []int64
	EventType *commonenum.EventType
	Channel   *notifyenum.NotificationChannel
	Language  *notifyenum.Language
	Enabled   *bool
}
