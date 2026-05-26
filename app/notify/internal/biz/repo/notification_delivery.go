package repo

import (
	"context"
	"notify/internal/biz/model"
	"time"
)

type NotificationDeliveryRepo interface {
	Saves(ctx context.Context, deliveries []*model.NotificationDelivery) ([]*model.NotificationDelivery, error)
	ListPending(ctx context.Context, limit int) ([]*model.NotificationDelivery, error)
	MarkSending(ctx context.Context, id int64) (bool, error)
	MarkSent(ctx context.Context, id int64, sentAt time.Time) error
	MarkFailed(ctx context.Context, id int64) error
}
