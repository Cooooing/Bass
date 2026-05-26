package service

import (
	"common/api/gen/common/errors"
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	commonModel "common/pkg/model"
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

type NotificationSettingService struct {
	v1.UnimplementedNotifyNotificationSettingServiceServer
	notificationSettingUsecase *usecase.NotificationSettingUsecase
}

func NewNotificationSettingService(notificationSettingUsecase *usecase.NotificationSettingUsecase) *NotificationSettingService {
	return &NotificationSettingService{notificationSettingUsecase: notificationSettingUsecase}
}

func (s *NotificationSettingService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationSettingServiceServer(gs, s)
}

func (s *NotificationSettingService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationSettingServiceHTTPServer(hs, s)
}

func (s *NotificationSettingService) ListCurrent(ctx context.Context, req *v1.ListCurrentNotificationSetting_Request) (*v1.ListCurrentNotificationSetting_Reply, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, errors.ErrorUnauthorized("user not login")
	}
	req = util.OrDefault(req, &v1.ListCurrentNotificationSetting_Request{})
	req.Query = util.OrDefault(req.Query, &v1.NotificationSettingQueryParams{})
	getReq := &repo.NotificationSettingGetReq{
		UserID:    &user.ID,
		EventType: req.Query.EventType,
	}
	if req.Query.Channel != nil {
		if _, ok := notifyenum.NotificationChannelMap.ToEnum(*req.Query.Channel); !ok {
			return nil, errors.ErrorBadRequest("invalid notification channel")
		}
		getReq.Channel = req.Query.Channel
	}
	rows, err := s.notificationSettingUsecase.List(ctx, getReq)
	if err != nil {
		return nil, err
	}
	reply := make([]*v1.NotificationSetting, 0, len(rows))
	for _, row := range rows {
		item := &v1.NotificationSetting{
			Id:        row.ID,
			UserId:    row.UserID,
			EventType: commonenum.EventTypeMap.MustToProto(commonenum.EventType(row.EventType)),
			Channel:   notifyenum.NotificationChannelMap.MustToProto(notifyenum.NotificationChannel(row.Channel)),
			Enable:    row.Enable,
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		reply = append(reply, item)
	}
	return &v1.ListCurrentNotificationSetting_Reply{Rows: reply}, nil
}

func (s *NotificationSettingService) UpdateCurrent(ctx context.Context, req *v1.UpdateCurrentNotificationSetting_Request) (*v1.UpdateCurrentNotificationSetting_Reply, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, errors.ErrorUnauthorized("user not login")
	}
	dbEventType, ok := commonenum.EventTypeMap.ToEnum(req.EventType)
	if !ok {
		return nil, errors.ErrorBadRequest("invalid event type")
	}
	dbChannel, ok := notifyenum.NotificationChannelMap.ToEnum(req.Channel)
	if !ok {
		return nil, errors.ErrorBadRequest("invalid notification channel")
	}
	save, err := s.notificationSettingUsecase.Upsert(ctx, &model.NotificationSetting{
		UserID:    user.ID,
		EventType: dbEventType,
		Channel:   dbChannel,
		Enable:    req.Enable,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.NotificationSetting{
		Id:        save.ID,
		UserId:    save.UserID,
		EventType: commonenum.EventTypeMap.MustToProto(commonenum.EventType(save.EventType)),
		Channel:   notifyenum.NotificationChannelMap.MustToProto(notifyenum.NotificationChannel(save.Channel)),
		Enable:    save.Enable,
	}
	if save.CreatedAt != nil {
		reply.CreatedAt = timestamppb.New(*save.CreatedAt)
	}
	if save.UpdatedAt != nil {
		reply.UpdatedAt = timestamppb.New(*save.UpdatedAt)
	}
	return &v1.UpdateCurrentNotificationSetting_Reply{NotificationSetting: reply}, nil
}
