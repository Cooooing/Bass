package repo

import "context"

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
	ReadAt     string
	CreatedAt  string
	UpdatedAt  string
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
