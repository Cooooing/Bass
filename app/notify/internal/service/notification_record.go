package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"notify/internal/biz/domain"
	"notify/internal/biz/repo"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationRecordService struct {
	v1.UnimplementedNotifyNotificationRecordServiceServer
	*BaseService
	notificationRecordDomain *domain.NotificationRecordDomain
}

func NewNotificationRecordService(baseService *BaseService, notificationRecordDomain *domain.NotificationRecordDomain) *NotificationRecordService {
	return &NotificationRecordService{
		BaseService:              baseService,
		notificationRecordDomain: notificationRecordDomain,
	}
}

func (s *NotificationRecordService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationRecordServiceServer(gs, s)
}

func (s *NotificationRecordService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationRecordServiceHTTPServer(hs, s)
}

func (s *NotificationRecordService) Page(ctx context.Context, req *v1.PageNotificationRecord_Request) (rsp *v1.PageNotificationRecord_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	records, page, err := s.notificationRecordDomain.Page(ctx, req.Page, &repo.NotificationRecordGetReq{
		ReceiverId: new(user.ID),
		Status:     new(v1.NotificationStatus_NOTIFICATION_STATUS_NORMAL),
		WithMeta:   true,
	})

	return &v1.PageNotificationRecord_Reply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(records),
	}, err
}

func (s *NotificationRecordService) Read(ctx context.Context, req *v1.ReadNotificationRecord_Request) (rsp *v1.ReadNotificationRecord_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}

	var (
		startTime *time.Time
		endTime   *time.Time
	)
	if req.ReadTimeRange != nil {
		if req.ReadTimeRange.Start != nil {
			startTime = new(req.ReadTimeRange.Start.AsTime())
		}
		if req.ReadTimeRange.End != nil {
			endTime = new(req.ReadTimeRange.End.AsTime())
		}
	}

	count, err := s.notificationRecordDomain.Read(ctx, user.ID, startTime, endTime, req.NotificationRecordIds)
	return &v1.ReadNotificationRecord_Reply{
		Count: int32(count),
	}, err
}
