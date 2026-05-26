package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/notify/v1"
	commonenum "common/pkg/enum"
	"common/pkg/util"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	notifyenum "notify/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationTemplateService struct {
	v1.UnimplementedNotifyNotificationTemplateServiceServer
	notificationTemplateUsecase *usecase.NotificationTemplateUsecase
}

func NewNotificationTemplateService(notificationTemplateUsecase *usecase.NotificationTemplateUsecase) *NotificationTemplateService {
	return &NotificationTemplateService{
		notificationTemplateUsecase: notificationTemplateUsecase,
	}
}

func (s *NotificationTemplateService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationTemplateServiceServer(gs, s)
}

func (s *NotificationTemplateService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationTemplateServiceHTTPServer(hs, s)
}

func (s *NotificationTemplateService) List(ctx context.Context, req *v1.ListNotificationTemplates_Request) (rsp *v1.ListNotificationTemplates_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.NotificationTemplateQueryParams{})
	records, page, err := s.notificationTemplateUsecase.Page(ctx, req.Page, &repo.NotificationTemplateGetReq{
		NotificationTemplateIds: req.Query.Ids,
		EventType:               req.Query.EventType,
		Channel:                 req.Query.Channel,
		Enable:                  req.Query.Enable,
	})

	rows := make([]*v1.NotificationTemplate, 0, len(records))
	for _, record := range records {
		rows = append(rows, &v1.NotificationTemplate{
			CreatedAt: timestamppb.New(*record.CreatedAt),
			UpdatedAt: timestamppb.New(*record.UpdatedAt),
			Id:        record.ID,
			EventType: commonenum.EventTypeMap.MustToProto(commonenum.EventType(record.EventType)),
			Channel:   v1.NotificationChannel(v1.NotificationChannel_value[string(record.Channel)]),
			Title:     record.Title,
			Content:   record.Content,
			Enable:    record.Enable,
		})
	}

	return &v1.ListNotificationTemplates_Reply{
		Page: page,
		Rows: rows,
	}, err
}

func (s *NotificationTemplateService) Create(ctx context.Context, req *v1.CreateNotificationTemplate_Request) (rsp *v1.CreateNotificationTemplate_Reply, err error) {
	if req.NotificationTemplate == nil {
		return nil, cerrors.ErrorBadRequest("notificationTemplate is required")
	}
	dbEventType, ok := commonenum.EventTypeMap.ToEnum(req.NotificationTemplate.EventType)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid eventType")
	}
	dbChannel, ok := notifyenum.NotificationChannelMap.ToEnum(req.NotificationTemplate.Channel)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid channel")
	}
	save, err := s.notificationTemplateUsecase.Add(ctx, &model.NotificationTemplate{
		EventType: dbEventType,
		Channel:   dbChannel,
		Language:  notifyenum.LanguageZhCN,
		Title:     req.NotificationTemplate.Title,
		Content:   req.NotificationTemplate.Content,
		Enable:    req.NotificationTemplate.Enable,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateNotificationTemplate_Reply{
		NotificationTemplate: &v1.NotificationTemplate{
			CreatedAt: timestamppb.New(*save.CreatedAt),
			UpdatedAt: timestamppb.New(*save.UpdatedAt),
			Id:        save.ID,
			EventType: commonenum.EventTypeMap.MustToProto(commonenum.EventType(save.EventType)),
			Channel:   v1.NotificationChannel(v1.NotificationChannel_value[string(save.Channel)]),
			Title:     save.Title,
			Content:   save.Content,
			Enable:    save.Enable,
		},
	}, nil
}

func (s *NotificationTemplateService) Update(ctx context.Context, req *v1.UpdateNotificationTemplate_Request) (rsp *v1.UpdateNotificationTemplate_Reply, err error) {
	if req.EventType == nil || req.Channel == nil || req.Content == nil {
		return nil, cerrors.ErrorBadRequest("eventType, channel, content is required")
	}
	dbEventType, ok := commonenum.EventTypeMap.ToEnum(*req.EventType)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid eventType")
	}
	dbChannel, ok := notifyenum.NotificationChannelMap.ToEnum(*req.Channel)
	if !ok {
		return nil, cerrors.ErrorBadRequest("invalid channel")
	}
	save, err := s.notificationTemplateUsecase.Update(ctx, &model.NotificationTemplate{
		EventType: dbEventType,
		Channel:   dbChannel,
		Language:  notifyenum.LanguageZhCN,
		Content:   *req.Content,
		Enable:    util.DerefOrDefault(req.Enable, false),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateNotificationTemplate_Reply{
		NotificationTemplate: &v1.NotificationTemplate{
			CreatedAt: timestamppb.New(*save.CreatedAt),
			UpdatedAt: timestamppb.New(*save.UpdatedAt),
			Id:        save.ID,
			EventType: commonenum.EventTypeMap.MustToProto(commonenum.EventType(save.EventType)),
			Channel:   v1.NotificationChannel(v1.NotificationChannel_value[string(save.Channel)]),
			Title:     save.Title,
			Content:   save.Content,
			Enable:    save.Enable,
		},
	}, nil
}
