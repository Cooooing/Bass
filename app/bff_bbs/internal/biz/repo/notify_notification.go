package repo

import (
	"context"
	"time"
)

type NotificationClient interface {
	ListNotifications(ctx context.Context, req *ListNotificationsReq) (*ListNotificationsResp, error)
	MarkReadNotification(ctx context.Context, req *MarkReadNotificationReq) (int32, error)
	CountUnreadNotifications(ctx context.Context, userID int64) (int64, error)
}

type Notification struct {
	ID         int64
	EventID    string
	ReceiverID int64
	EventType  int32
	Title      string
	Content    string
	ReadAt     *time.Time
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type ListNotificationsReq struct {
	UserID int64
	Page   *PageReq
}

type ListNotificationsResp struct {
	Page *PageResp
	Rows []*Notification
}

type MarkReadNotificationReq struct {
	UserID int64
	IDs    []int64
}
