package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsnotifyv1 "common/proto/gen/bbs/v1/notify"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type NotificationService struct {
	bbsnotifyv1.UnimplementedNotificationServiceServer
	notificationUsecase *usecase.NotificationUsecase
}

func NewNotificationService(
	notificationUsecase *usecase.NotificationUsecase,
) *NotificationService {
	return &NotificationService{
		notificationUsecase: notificationUsecase,
	}
}

func (s *NotificationService) RegisterGrpc(gs *grpc.Server) {
}

func (s *NotificationService) RegisterHttp(hs *http.Server) {
	bbsnotifyv1.RegisterNotificationServiceHTTPServer(hs, s)
}

func (s *NotificationService) List(ctx context.Context, req *bbsnotifyv1.ListNotifications_Req) (*bbsnotifyv1.ListNotifications_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.notificationUsecase.ListNotifications(ctx, &usecase.ListNotificationsReq{
		UserID: user.ID,
		Page:   req.GetPage(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.ListNotifications_Resp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

func (s *NotificationService) MarkRead(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Req) (*bbsnotifyv1.MarkReadNotification_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.notificationUsecase.MarkReadNotification(ctx, &usecase.MarkReadNotificationReq{
		UserID: user.ID,
		IDs:    req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.MarkReadNotification_Resp{
		Count: resp,
	}, nil
}

func (s *NotificationService) CountUnread(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Req) (*bbsnotifyv1.CountUnreadNotifications_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	count, err := s.notificationUsecase.CountUnreadNotifications(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.CountUnreadNotifications_Resp{
		Count: count,
	}, nil
}
