package usecase

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/proto/gen/bbs/v1/notify"
	"common/proto/gen/common"
	commonenums "common/proto/gen/common/enums"
	"context"
)

type NotificationUsecase struct {
	notificationClient repo.NotificationClient
}

func NewNotificationUsecase(notificationClient repo.NotificationClient) *NotificationUsecase {
	return &NotificationUsecase{notificationClient: notificationClient}
}

type ListNotificationsReq struct {
	UserID int64
	Page   *common.PageRequest
}

type ListNotificationsResponse struct {
	Page *common.PageResponse
	Rows []*bbsnotifyv1.ListNotifications_Response_Notification
}

func (u *NotificationUsecase) ListNotifications(ctx context.Context, req *ListNotificationsReq) (*ListNotificationsResponse, error) {
	if req == nil {
		req = &ListNotificationsReq{}
	}
	var pageReq *repo.PageReq
	if req.Page != nil {
		pageReq = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	reply, err := u.notificationClient.ListNotifications(ctx, &repo.ListNotificationsReq{UserID: req.UserID, Page: pageReq})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if reply.Page != nil {
		page = &common.PageResponse{Page: reply.Page.Page, Size: reply.Page.Size, Total: reply.Page.Total}
	}
	rows := make([]*bbsnotifyv1.ListNotifications_Response_Notification, 0, len(reply.Rows))
	for _, item := range reply.Rows {
		rows = append(rows, &bbsnotifyv1.ListNotifications_Response_Notification{
			Id:         item.ID,
			EventId:    item.EventID,
			ReceiverId: item.ReceiverID,
			EventType:  commonenums.EventType(item.EventType),
			Title:      item.Title,
			Content:    item.Content,
			ReadAt:     item.ReadAt,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}
	return &ListNotificationsResponse{Page: page, Rows: rows}, nil
}

type MarkReadNotificationReq struct {
	UserID int64
	IDs    []int64
}

type MarkReadNotificationResponse struct {
	Count int32
}

func (u *NotificationUsecase) MarkReadNotification(ctx context.Context, req *MarkReadNotificationReq) (*MarkReadNotificationResponse, error) {
	reply, err := u.notificationClient.MarkReadNotification(ctx, &repo.MarkReadNotificationReq{UserID: req.UserID, IDs: req.IDs})
	if err != nil {
		return nil, err
	}
	return &MarkReadNotificationResponse{Count: reply.Count}, nil
}

type CountUnreadNotificationsReq struct {
	UserID int64
}

type CountUnreadNotificationsResponse struct {
	Count int64
}

func (u *NotificationUsecase) CountUnreadNotifications(ctx context.Context, req *CountUnreadNotificationsReq) (*CountUnreadNotificationsResponse, error) {
	reply, err := u.notificationClient.CountUnreadNotifications(ctx, &repo.CountUnreadNotificationsReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	return &CountUnreadNotificationsResponse{Count: reply.Count}, nil
}
