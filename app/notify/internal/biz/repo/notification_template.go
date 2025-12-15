package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/notify/v1"
	"common/pkg/cutil/collections/dict"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/ent/gen"
)

type NotificationTemplateRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.NotificationTemplate) (*model.NotificationTemplate, error)
	Update(ctx context.Context, tx *gen.Client, u *model.NotificationTemplate) (*model.NotificationTemplate, error)

	SaveCache(ctx context.Context, records dict.Map[string, *model.NotificationTemplate]) error
	GetCache(ctx context.Context, notificationType *v1.NotificationType, channels []*v1.NotificationChannel) (dict.Map[string, *model.NotificationTemplate], error)

	GetOne(ctx context.Context, tx *gen.Client, req *NotificationTemplateGetReq) (*model.NotificationTemplate, error)
	GetList(ctx context.Context, tx *gen.Client, req *NotificationTemplateGetReq) ([]*model.NotificationTemplate, error)
	GetMap(ctx context.Context, tx *gen.Client, req *NotificationTemplateGetReq) (dict.Map[string, *model.NotificationTemplate], error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *NotificationTemplateGetReq) ([]*model.NotificationTemplate, *cv1.PageReply, error)
}

type NotificationTemplateGetReq struct {
	NotificationTemplateId  *int64
	NotificationTemplateIds []int64
	NotificationType        *v1.NotificationType
	Channel                 *v1.NotificationChannel
	Channels                []*v1.NotificationChannel
	Enable                  *bool
}
