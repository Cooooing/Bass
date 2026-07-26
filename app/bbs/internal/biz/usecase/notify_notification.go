package usecase

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/proto/gen/bbs/v1/notify"
	"common/proto/gen/common"
	commonenums "common/proto/gen/common/enums"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationUsecase struct {
	notificationClient repo.NotificationClient
}

func NewNotificationUsecase(
	notificationClient repo.NotificationClient,
) *NotificationUsecase {
	return &NotificationUsecase{
		notificationClient: notificationClient,
	}
}

type ListNotificationsReq struct {
	UserID int64
	Page   *common.PageReq
}

type ListNotificationsResp struct {
	Page *common.PageResp
	Rows []*bbsnotifyv1.ListNotifications_Resp_Notification
}

func (u *NotificationUsecase) ListNotifications(ctx context.Context, req *ListNotificationsReq) (*ListNotificationsResp, error) {
	if req == nil {
		req = &ListNotificationsReq{}
	}
	var pageReq *repo.PageReq
	if req.Page != nil {
		pageReq = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	reply, err := u.notificationClient.ListNotifications(ctx, &repo.ListNotificationsReq{
		UserID: req.UserID,
		Page:   pageReq,
	})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if reply.Page != nil {
		page = &common.PageResp{
			Page:  reply.Page.Page,
			Size:  reply.Page.Size,
			Total: reply.Page.Total,
		}
	}
	rows := make([]*bbsnotifyv1.ListNotifications_Resp_Notification, 0, len(reply.Rows))
	for _, item := range reply.Rows {
		row := &bbsnotifyv1.ListNotifications_Resp_Notification{
			Id:         item.ID,
			EventId:    item.EventID,
			ReceiverId: item.ReceiverID,
			EventType:  commonenums.EventType(item.EventType),
			Title:      item.Title,
			Content:    item.Content,
		}
		if item.ReadAt != nil {
			row.ReadAt = timestamppb.New(*item.ReadAt)
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		rows = append(rows, row)
	}
	return &ListNotificationsResp{
		Page: page,
		Rows: rows,
	}, nil
}

type MarkReadNotificationReq struct {
	UserID int64
	IDs    []int64
}

func (u *NotificationUsecase) MarkReadNotification(ctx context.Context, req *MarkReadNotificationReq) (int32, error) {
	reply, err := u.notificationClient.MarkReadNotification(ctx, &repo.MarkReadNotificationReq{
		UserID: req.UserID,
		IDs:    req.IDs,
	})
	if err != nil {
		return 0, err
	}
	return reply, nil
}

func (u *NotificationUsecase) CountUnreadNotifications(ctx context.Context, userID int64) (int64, error) {
	reply, err := u.notificationClient.CountUnreadNotifications(ctx, userID)
	if err != nil {
		return 0, err
	}
	return reply, nil
}
