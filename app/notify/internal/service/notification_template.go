package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/notify/v1"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"context"
	"notify/internal/biz/domain"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationTemplateService struct {
	v1.UnimplementedNotifyNotificationTemplateServiceServer
	*BaseService
	notificationTemplateDomain *domain.NotificationTemplateDomain
}

func NewNotificationTemplateService(baseService *BaseService, notificationTemplateDomain *domain.NotificationTemplateDomain) *NotificationTemplateService {
	return &NotificationTemplateService{
		BaseService:                baseService,
		notificationTemplateDomain: notificationTemplateDomain,
	}
}

func (s *NotificationTemplateService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationTemplateServiceServer(gs, s)
}

func (s *NotificationTemplateService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationTemplateServiceHTTPServer(hs, s)
}

func (s *NotificationTemplateService) Page(ctx context.Context, req *v1.NotificationTemplatePageRequest) (rsp *v1.NotificationTemplatePageReply, err error) {
	req.Query = base.OrDefault(req.Query, &v1.NotificationTemplateQueryParams{})
	records, page, err := s.notificationTemplateDomain.Page(ctx, req.Page, &repo.NotificationTemplateGetReq{
		NotificationTemplateIds: req.Query.Ids,
		NotificationType:        req.Query.NotificationType,
		Channel:                 req.Query.Channel,
		Enable:                  req.Query.Enable,
	})

	return &v1.NotificationTemplatePageReply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(records),
	}, err
}

func (s *NotificationTemplateService) Add(ctx context.Context, req *v1.NotificationTemplateAddRequest) (rsp *v1.NotificationTemplateAddReply, err error) {
	if req.NotificationTemplate == nil {
		return nil, cv1.ErrorBadRequest("notificationTemplate is required")
	}
	if _, ok := v1.NotificationType_name[int32(req.NotificationTemplate.NotificationType)]; !ok {
		return nil, cv1.ErrorBadRequest("invalid notificationType")
	}
	if _, ok := v1.NotificationChannel_name[int32(req.NotificationTemplate.Channel)]; !ok {
		return nil, cv1.ErrorBadRequest("invalid channel")
	}
	tpl, err := s.notificationTemplateDomain.Add(ctx, &model.NotificationTemplate{NotificationTemplate: &gen.NotificationTemplate{
		NotificationType: int32(req.NotificationTemplate.NotificationType),
		Channel:          int32(req.NotificationTemplate.Channel),
		Content:          req.NotificationTemplate.Content,
		Processors:       req.NotificationTemplate.Processors,
		Enable:           req.NotificationTemplate.Enable,
	}})
	if err != nil {
		return nil, err
	}
	return &v1.NotificationTemplateAddReply{
		NotificationTemplate: tpl.ConvertToRpc(),
	}, nil
}

func (s *NotificationTemplateService) Update(ctx context.Context, req *v1.NotificationTemplateUpdateRequest) (rsp *v1.NotificationTemplateUpdateReply, err error) {
	if req.NotificationType == nil || req.Channel == nil || req.Content == nil {
		return nil, cv1.ErrorBadRequest("notificationType, channel, content is required")
	}
	if _, ok := v1.NotificationType_name[int32(*req.NotificationType)]; !ok {
		return nil, cv1.ErrorBadRequest("invalid notificationType")
	}
	if _, ok := v1.NotificationChannel_name[int32(*req.Channel)]; !ok {
		return nil, cv1.ErrorBadRequest("invalid channel")
	}
	tpl, err := s.notificationTemplateDomain.Update(ctx, &model.NotificationTemplate{NotificationTemplate: &gen.NotificationTemplate{
		NotificationType: int32(*req.NotificationType),
		Channel:          int32(*req.Channel),
		Content:          *req.Content,
		Processors:       req.Processors,
		Enable:           base.DerefOrDefault(req.Enable, false),
	}})
	if err != nil {
		return nil, err
	}
	return &v1.NotificationTemplateUpdateReply{
		NotificationTemplate: tpl.ConvertToRpc(),
	}, nil
}
