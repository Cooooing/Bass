package repo

import (
	bbsnotifyv1 "common/api/gen/bbs/v1/notify"
	"context"
)

type NotifyRepo interface {
	ListNotifications(ctx context.Context, req *bbsnotifyv1.ListNotifications_Request) (*bbsnotifyv1.ListNotifications_Reply, error)
	MarkReadNotification(ctx context.Context, req *bbsnotifyv1.MarkReadNotification_Request) (*bbsnotifyv1.MarkReadNotification_Reply, error)
	CountUnreadNotifications(ctx context.Context, req *bbsnotifyv1.CountUnreadNotifications_Request) (*bbsnotifyv1.CountUnreadNotifications_Reply, error)
	ListCurrentNotificationSetting(ctx context.Context, req *bbsnotifyv1.ListCurrentNotificationSetting_Request) (*bbsnotifyv1.ListCurrentNotificationSetting_Reply, error)
	UpdateCurrentNotificationSetting(ctx context.Context, req *bbsnotifyv1.UpdateCurrentNotificationSetting_Request) (*bbsnotifyv1.UpdateCurrentNotificationSetting_Reply, error)
}
