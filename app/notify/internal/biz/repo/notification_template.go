package repo

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"context"
	"notify/internal/biz/model"
)

type NotificationTemplateRepo interface {
	Save(ctx context.Context, u *model.NotificationTemplate) (*model.NotificationTemplate, error)
	Update(ctx context.Context, u *model.NotificationTemplate) (*model.NotificationTemplate, error)

	// GetTemplates 按事件类型和语言获取所有渠道模板，内部带缓存
	GetTemplates(ctx context.Context, eventType enums.EventType, language string) ([]*model.NotificationTemplate, error)

	Get(ctx context.Context, req *NotificationTemplateGetReq) (*model.NotificationTemplate, error)
	GetList(ctx context.Context, req *NotificationTemplateGetReq) ([]*model.NotificationTemplate, error)
	GetMap(ctx context.Context, req *NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error)
	GetPage(ctx context.Context, page *commonv1.PageRequest, req *NotificationTemplateGetReq) ([]*model.NotificationTemplate, *commonv1.PageReply, error)
}

type NotificationTemplateGetReq struct {
	NotificationTemplateId  *int64
	NotificationTemplateIds []int64
	EventType               *enums.EventType
	Channel                 *v1.NotificationChannel
	Channels                []*v1.NotificationChannel
	Language                *string
	Enable                  *bool
}
