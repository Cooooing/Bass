package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/notify/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"notify/internal/biz"
	"notify/internal/biz/repo"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type NotificationRecordService struct {
	v1.UnimplementedNotifyNotificationRecordServiceServer
	*BaseService
	notificationRecordDomain *biz.NotificationRecordDomain
}

func NewNotificationRecordService(baseService *BaseService, notificationRecordDomain *biz.NotificationRecordDomain) *NotificationRecordService {
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

func (s *NotificationRecordService) Page(ctx context.Context, req *v1.NotificationRecordPageRequest) (rsp *v1.NotificationRecordPageReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	records, page, err := s.notificationRecordDomain.Page(ctx, req.Page, &repo.NotificationRecordGetReq{
		NotificationType: req.Query.NotificationType,
		ReceiverId:       base.Ptr(user.ID),
		Status:           base.Ptr(v1.NotificationStatus_NOTIFICATION_STATUS_NORMAL),
		WithMeta:         true,
	})

	return &v1.NotificationRecordPageReply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(records),
	}, err
}

func (s *NotificationRecordService) Read(ctx context.Context, req *v1.NotificationRecordReadRequest) (rsp *v1.NotificationRecordReadReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}

	var (
		startTime *time.Time
		endTime   *time.Time
	)
	if req.ReadTimeRange != nil {
		if req.ReadTimeRange.Start != nil {
			startTime = base.Ptr(req.ReadTimeRange.Start.AsTime())
		}
		if req.ReadTimeRange.End != nil {
			endTime = base.Ptr(req.ReadTimeRange.End.AsTime())
		}
	}

	count, err := s.notificationRecordDomain.Read(ctx, user.ID, startTime, endTime, req.NotificationRecordIds)
	return &v1.NotificationRecordReadReply{
		Count: int32(count),
	}, err
}
