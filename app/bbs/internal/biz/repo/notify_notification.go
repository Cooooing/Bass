package repo

import "context"

type NotificationClient interface {
	ListNotifications(ctx context.Context, req *ListNotificationsReq) (*ListNotificationsResponse, error)
	MarkReadNotification(ctx context.Context, req *MarkReadNotificationReq) (*MarkReadNotificationResponse, error)
	CountUnreadNotifications(ctx context.Context, req *CountUnreadNotificationsReq) (*CountUnreadNotificationsResponse, error)
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

type ListNotificationsResponse struct {
	Page *PageResponse
	Rows []*Notification
}

type MarkReadNotificationReq struct {
	UserID int64
	IDs    []int64
}

type MarkReadNotificationResponse struct {
	Count int32
}

type CountUnreadNotificationsReq struct {
	UserID int64
}

type CountUnreadNotificationsResponse struct {
	Count int64
}
