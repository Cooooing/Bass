package data

import (
	"bbs/internal/biz/repo"
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	notifyv1 "common/api/gen/notify/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.NotifyRepo = (*NotifyRepo)(nil)

type NotifyRepo struct {
	notifyClient *rpc.NotifyClient
}

func NewNotifyRepo(notifyClient *rpc.NotifyClient) repo.NotifyRepo {
	return &NotifyRepo{notifyClient: notifyClient}
}

func (r *NotifyRepo) ListNotifications(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error) {
	reply, err := r.notifyClient.Record.List(forwardAuth(ctx), &notifyv1.ListNotificationRecords_Request{
		Page:  req.GetPage(),
		Query: &notifyv1.NotificationRecordQueryParams{},
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbsnotifyv1.Notification, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &bbsnotifyv1.Notification{
			Id:             item.GetId(),
			NotificationId: item.GetNotificationId(),
			ReceiverId:     item.GetReceiverId(),
			ReadTime:       formatProtoTime(item.GetReadTime()),
			CreatedAt:      formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:      formatProtoTime(item.GetUpdatedAt()),
		}
		if meta := item.GetNotificationMeta(); meta != nil {
			row.Title = meta.GetTitle()
			row.Content = meta.GetContent()
		}
		rows = append(rows, row)
	}
	return &bbsnotifyv1.ListNotifications_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *NotifyRepo) MarkReadNotification(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error) {
	reply, err := r.notifyClient.Record.MarkRead(forwardAuth(ctx), &notifyv1.MarkReadNotificationRecord_Request{
		NotificationRecordIds: req.GetNotificationRecordIds(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.MarkReadNotification_Reply{Count: reply.GetCount()}, nil
}

func (r *NotifyRepo) CountUnreadNotifications(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error) {
	reply, err := r.notifyClient.Record.CountUnread(forwardAuth(ctx), &notifyv1.CountUnreadNotificationRecords_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsnotifyv1.CountUnreadNotifications_Reply{Count: reply.GetCount()}, nil
}

func (r *NotifyRepo) ListCurrentNotificationSetting(ctx context.Context, req *bbsnotifyv1.ListCurrentNotificationSetting_Request) (*bbsnotifyv1.ListCurrentNotificationSetting_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbsnotifyv1.NotificationSettingQuery{}
	}
	notifyQuery := &notifyv1.NotificationSettingQueryParams{EventType: query.EventType}
	if query.Channel != nil {
		notifyQuery.Channel = new(notifyv1.NotificationChannel(*query.Channel))
	}
	reply, err := r.notifyClient.Setting.ListCurrent(forwardAuth(ctx), &notifyv1.ListCurrentNotificationSetting_Request{Query: notifyQuery})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbsnotifyv1.NotificationSetting, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		rows = append(rows, &bbsnotifyv1.NotificationSetting{
			Id:        item.GetId(),
			UserId:    item.GetUserId(),
			EventType: item.GetEventType(),
			Channel:   bbsnotifyv1.NotificationChannel(item.GetChannel()),
			Enable:    item.GetEnable(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsnotifyv1.ListCurrentNotificationSetting_Reply{Rows: rows}, nil
}

func (r *NotifyRepo) UpdateCurrentNotificationSetting(ctx context.Context, req *bbsnotifyv1.UpdateCurrentNotificationSetting_Request) (*bbsnotifyv1.UpdateCurrentNotificationSetting_Reply, error) {
	reply, err := r.notifyClient.Setting.UpdateCurrent(forwardAuth(ctx), &notifyv1.UpdateCurrentNotificationSetting_Request{
		EventType: req.GetEventType(),
		Channel:   notifyv1.NotificationChannel(req.GetChannel()),
		Enable:    req.GetEnable(),
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetNotificationSetting()
	return &bbsnotifyv1.UpdateCurrentNotificationSetting_Reply{NotificationSetting: &bbsnotifyv1.NotificationSetting{
		Id:        item.GetId(),
		UserId:    item.GetUserId(),
		EventType: item.GetEventType(),
		Channel:   bbsnotifyv1.NotificationChannel(item.GetChannel()),
		Enable:    item.GetEnable(),
		CreatedAt: formatProtoTime(item.GetCreatedAt()),
		UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}
