package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/notify/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationRecordService struct {
	v1.UnimplementedNotifyNotificationRecordServiceServer
	notificationRecordUsecase *usecase.NotificationRecordUsecase
}

func NewNotificationRecordService(notificationRecordUsecase *usecase.NotificationRecordUsecase) *NotificationRecordService {
	return &NotificationRecordService{
		notificationRecordUsecase: notificationRecordUsecase,
	}
}

func (s *NotificationRecordService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterNotifyNotificationRecordServiceServer(gs, s)
}

func (s *NotificationRecordService) RegisterHttp(hs *http.Server) {
	v1.RegisterNotifyNotificationRecordServiceHTTPServer(hs, s)
}

func (s *NotificationRecordService) List(ctx context.Context, req *v1.ListNotificationRecords_Request) (rsp *v1.ListNotificationRecords_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	records, page, err := s.notificationRecordUsecase.Page(ctx, req.Page, &repo.NotificationRecordGetReq{
		ReceiverId: new(user.ID),
		Status:     new(v1.NotificationStatus_NOTIFICATION_STATUS_NORMAL),
		WithMeta:   true,
	})

	rows := make([]*v1.NotificationRecord, 0, len(records))
	for _, record := range records {
		row := &v1.NotificationRecord{
			CreatedAt:      timestamppb.New(*record.CreatedAt),
			UpdatedAt:      timestamppb.New(*record.UpdatedAt),
			Id:             record.ID,
			NotificationId: record.NotificationID,
			ReceiverId:     record.ReceiverID,
		}
		if record.ReadTime != nil {
			row.ReadTime = timestamppb.New(*record.ReadTime)
		}
		if record.NotificationMeta != nil {
			meta := record.NotificationMeta
			row.NotificationMeta = &v1.NotificationMeta{
				CreatedAt: timestamppb.New(*meta.CreatedAt),
				UpdatedAt: timestamppb.New(*meta.UpdatedAt),
				Id:        meta.ID,
				Title:     meta.Title,
				Content:   meta.Content,
				Status:    v1.NotificationStatus(v1.NotificationStatus_value[string(meta.Status)]),
			}
		}
		rows = append(rows, row)
	}

	return &v1.ListNotificationRecords_Reply{
		Page: page,
		Rows: rows,
	}, err
}

func (s *NotificationRecordService) MarkRead(ctx context.Context, req *v1.MarkReadNotificationRecord_Request) (rsp *v1.MarkReadNotificationRecord_Reply, err error) {
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

	count, err := s.notificationRecordUsecase.MarkRead(ctx, user.ID, startTime, endTime, req.NotificationRecordIds)
	return &v1.MarkReadNotificationRecord_Reply{
		Count: int32(count),
	}, err
}

func (s *NotificationRecordService) CountUnread(ctx context.Context, req *v1.CountUnreadNotificationRecords_Request) (*v1.CountUnreadNotificationRecords_Reply, error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	count, err := s.notificationRecordUsecase.CountUnread(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &v1.CountUnreadNotificationRecords_Reply{Count: int64(count)}, nil
}
