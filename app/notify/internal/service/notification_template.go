package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/notify/v1"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"notify/internal/biz/domain"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	notificationtemplate "notify/internal/data/gen/notificationtemplate"
	notifyenum "notify/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationTemplateService struct {
	v1.UnimplementedNotifyNotificationTemplateServiceServer
	notificationTemplateDomain *domain.NotificationTemplateDomain
}

func NewNotificationTemplateService(notificationTemplateDomain *domain.NotificationTemplateDomain) *NotificationTemplateService {
	return &NotificationTemplateService{
		notificationTemplateDomain: notificationTemplateDomain,
	}
}

func (s *NotificationTemplateService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationTemplateServiceServer(gs, s)
}

func (s *NotificationTemplateService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationTemplateServiceHTTPServer(hs, s)
}

func (s *NotificationTemplateService) Page(ctx context.Context, req *v1.PageNotificationTemplate_Request) (rsp *v1.PageNotificationTemplate_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.NotificationTemplateQueryParams{})
	records, page, err := s.notificationTemplateDomain.Page(ctx, req.Page, &repo.NotificationTemplateGetReq{
		NotificationTemplateIds: req.Query.Ids,
		EventType:               req.Query.EventType,
		Channel:                 req.Query.Channel,
		Enable:                  req.Query.Enable,
	})

	return &v1.PageNotificationTemplate_Reply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(records),
	}, err
}

func (s *NotificationTemplateService) Add(ctx context.Context, req *v1.AddNotificationTemplate_Request) (rsp *v1.AddNotificationTemplate_Reply, err error) {
	if req.NotificationTemplate == nil {
		return nil, cerrors.ErrorBadRequest("notificationTemplate is required")
	}
	tpl := &gen.NotificationTemplate{
		Title:   req.NotificationTemplate.Title,
		Content: req.NotificationTemplate.Content,
		Enable:  req.NotificationTemplate.Enable,
	}
	dbEventType, ok := notifyenum.EventTypeMap.ToEnum(req.NotificationTemplate.EventType)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid eventType")
	}
	dbChannel, ok := notifyenum.NotificationChannelMap.ToEnum(req.NotificationTemplate.Channel)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid channel")
	}
	tpl.EventType = notificationtemplate.EventType(dbEventType)
	tpl.Channel = notificationtemplate.Channel(dbChannel)
	save, err := s.notificationTemplateDomain.Add(ctx, &model.NotificationTemplate{NotificationTemplate: tpl})
	if err != nil {
		return nil, err
	}
	return &v1.AddNotificationTemplate_Reply{
		NotificationTemplate: save.ConvertToRpc(),
	}, nil
}

func (s *NotificationTemplateService) Update(ctx context.Context, req *v1.UpdateNotificationTemplate_Request) (rsp *v1.UpdateNotificationTemplate_Reply, err error) {
	if req.EventType == nil || req.Channel == nil || req.Content == nil {
		return nil, cerrors.ErrorBadRequest("eventType, channel, content is required")
	}
	tpl := &gen.NotificationTemplate{
		Content: *req.Content,
		Enable:  util.DerefOrDefault(req.Enable, false),
	}
	dbEventType, ok := notifyenum.EventTypeMap.ToEnum(*req.EventType)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid eventType")
	}
	dbChannel, ok := notifyenum.NotificationChannelMap.ToEnum(*req.Channel)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid channel")
	}
	tpl.EventType = notificationtemplate.EventType(dbEventType)
	tpl.Channel = notificationtemplate.Channel(dbChannel)
	save, err := s.notificationTemplateDomain.Update(ctx, &model.NotificationTemplate{NotificationTemplate: tpl})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateNotificationTemplate_Reply{
		NotificationTemplate: save.ConvertToRpc(),
	}, nil
}
