package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	notifyv1 "common/proto/gen/notify/v1"
	"context"
)

var _ repo.NotificationClient = (*NotificationClient)(nil)

type NotificationClient struct {
	notifyClient *rpc.NotifyClient
}

func NewNotificationClient(notifyClient *rpc.NotifyClient) repo.NotificationClient {
	return &NotificationClient{notifyClient: notifyClient}
}

func (r *NotificationClient) ListNotifications(ctx context.Context, req *repo.ListNotificationsReq) (*repo.ListNotificationsResponse, error) {
	listReq := &notifyv1.ListStationMessages_Request{UserId: req.UserID}
	if req.Page != nil {
		listReq.Page = &common.PageRequest{Page: req.Page.Page, Size: req.Page.Size}
	}
	reply, err := r.notifyClient.StationMessage.List(ctx, listReq)
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.Notification, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &repo.Notification{
			ID:         item.GetId(),
			EventID:    item.GetEventId(),
			ReceiverID: item.GetReceiverId(),
			EventType:  int32(item.GetEventType()),
			Title:      item.GetTitle(),
			Content:    item.GetContent(),
			ReadAt:     formatProtoTime(item.GetReadAt()),
			CreatedAt:  formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:  formatProtoTime(item.GetUpdatedAt()),
		})
	}
	var page *repo.PageResponse
	if reply.GetPage() != nil {
		page = &repo.PageResponse{Page: reply.GetPage().GetPage(), Size: reply.GetPage().GetSize(), Total: reply.GetPage().GetTotal()}
	}
	return &repo.ListNotificationsResponse{Page: page, Rows: rows}, nil
}

func (r *NotificationClient) MarkReadNotification(ctx context.Context, req *repo.MarkReadNotificationReq) (*repo.MarkReadNotificationResponse, error) {
	reply, err := r.notifyClient.StationMessage.MarkRead(ctx, &notifyv1.MarkReadStationMessage_Request{
		UserId: req.UserID,
		Ids:    req.IDs,
	})
	if err != nil {
		return nil, err
	}
	return &repo.MarkReadNotificationResponse{Count: reply.GetCount()}, nil
}

func (r *NotificationClient) CountUnreadNotifications(ctx context.Context, req *repo.CountUnreadNotificationsReq) (*repo.CountUnreadNotificationsResponse, error) {
	reply, err := r.notifyClient.StationMessage.CountUnread(ctx, &notifyv1.CountUnreadStationMessages_Request{UserId: req.UserID})
	if err != nil {
		return nil, err
	}
	return &repo.CountUnreadNotificationsResponse{Count: reply.GetCount()}, nil
}
