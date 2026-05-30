package data

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	notifyv1 "common/api/gen/notify/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.NotificationRepo = (*NotificationRepo)(nil)

type NotificationRepo struct {
	notifyClient *rpc.NotifyClient
}

func NewNotificationRepo(notifyClient *rpc.NotifyClient) repo.NotificationRepo {
	return &NotificationRepo{notifyClient: notifyClient}
}

func (r *NotificationRepo) ListNotifications(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.notifyClient.StationMessage.List(ctx, &notifyv1.ListStationMessages_Request{
		UserId: userID,
		Page:   req.GetPage(),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbsnotifyv1.Notification, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &bbsnotifyv1.Notification{
			Id:         item.GetId(),
			EventId:    item.GetEventId(),
			ReceiverId: item.GetReceiverId(),
			EventType:  item.GetEventType(),
			Title:      item.GetTitle(),
			Content:    item.GetContent(),
			ReadAt:     formatProtoTime(item.GetReadAt()),
			CreatedAt:  formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:  formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsnotifyv1.ListNotifications_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *NotificationRepo) MarkReadNotification(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.notifyClient.StationMessage.MarkRead(ctx, &notifyv1.MarkReadStationMessage_Request{
		UserId: userID,
		Ids:    req.GetIds(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.MarkReadNotification_Reply{Count: reply.GetCount()}, nil
}

func (r *NotificationRepo) CountUnreadNotifications(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.notifyClient.StationMessage.CountUnread(ctx, &notifyv1.CountUnreadStationMessages_Request{UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.CountUnreadNotifications_Reply{Count: reply.GetCount()}, nil
}
