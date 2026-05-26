package repo

import (
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationdelivery"
	notifyenum "notify/internal/enum"
	"time"

	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
)

var _ repo.NotificationDeliveryRepo = (*NotificationDeliveryRepo)(nil)

type NotificationDeliveryRepo struct {
	db *gen.Client
}

func NewNotificationDeliveryRepo(db *gen.Client) repo.NotificationDeliveryRepo {
	return &NotificationDeliveryRepo{db: db}
}

func (r *NotificationDeliveryRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationDeliveryRepo) Saves(ctx context.Context, deliveries []*model.NotificationDelivery) ([]*model.NotificationDelivery, error) {
	client := r.getClient(ctx)
	result := make([]*model.NotificationDelivery, 0, len(deliveries))
	for _, delivery := range deliveries {
		if delivery == nil {
			continue
		}
		create := client.NotificationDelivery.Create().
			SetEventID(delivery.EventID).
			SetEventType(notificationdelivery.EventType(delivery.EventType)).
			SetNillableReceiverID(delivery.ReceiverID).
			SetChannel(notificationdelivery.Channel(delivery.Channel)).
			SetTarget(delivery.Target).
			SetTitle(delivery.Title).
			SetContent(delivery.Content).
			SetStatus(notificationdelivery.Status(delivery.Status)).
			SetRetryCount(delivery.RetryCount).
			SetNillableSentAt(delivery.SentAt)
		save, err := create.Save(ctx)
		if err != nil {
			if gen.IsConstraintError(err) {
				continue
			}
			return nil, err
		}
		result = append(result, &model.NotificationDelivery{
			ID:         save.ID,
			EventID:    save.EventID,
			EventType:  commonenum.EventType(save.EventType),
			ReceiverID: save.ReceiverID,
			Channel:    notifyenum.NotificationChannel(save.Channel),
			Target:     save.Target,
			Title:      save.Title,
			Content:    save.Content,
			Status:     notifyenum.NotificationDeliveryStatus(save.Status),
			RetryCount: save.RetryCount,
			SentAt:     save.SentAt,
			CreatedAt:  save.CreatedAt,
			UpdatedAt:  save.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationDeliveryRepo) ListPending(ctx context.Context, limit int) ([]*model.NotificationDelivery, error) {
	if limit <= 0 {
		limit = 50
	}
	list, err := r.getClient(ctx).NotificationDelivery.Query().
		Where(notificationdelivery.StatusIn(
			notificationdelivery.Status(notifyenum.NotificationDeliveryStatusPending),
			notificationdelivery.Status(notifyenum.NotificationDeliveryStatusFailed),
		)).
		Order(notificationdelivery.ByID()).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationDelivery, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationDelivery{
			ID:         item.ID,
			EventID:    item.EventID,
			EventType:  commonenum.EventType(item.EventType),
			ReceiverID: item.ReceiverID,
			Channel:    notifyenum.NotificationChannel(item.Channel),
			Target:     item.Target,
			Title:      item.Title,
			Content:    item.Content,
			Status:     notifyenum.NotificationDeliveryStatus(item.Status),
			RetryCount: item.RetryCount,
			SentAt:     item.SentAt,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationDeliveryRepo) MarkSending(ctx context.Context, id int64) (bool, error) {
	count, err := r.getClient(ctx).NotificationDelivery.Update().
		Where(
			notificationdelivery.IDEQ(id),
			notificationdelivery.StatusIn(
				notificationdelivery.Status(notifyenum.NotificationDeliveryStatusPending),
				notificationdelivery.Status(notifyenum.NotificationDeliveryStatusFailed),
			),
		).
		SetStatus(notificationdelivery.Status(notifyenum.NotificationDeliveryStatusSending)).
		Save(ctx)
	return count > 0, err
}

func (r *NotificationDeliveryRepo) MarkSent(ctx context.Context, id int64, sentAt time.Time) error {
	return r.getClient(ctx).NotificationDelivery.Update().
		Where(notificationdelivery.IDEQ(id)).
		SetStatus(notificationdelivery.Status(notifyenum.NotificationDeliveryStatusSent)).
		SetSentAt(sentAt).
		Exec(ctx)
}

func (r *NotificationDeliveryRepo) MarkFailed(ctx context.Context, id int64) error {
	return r.getClient(ctx).NotificationDelivery.Update().
		Where(notificationdelivery.IDEQ(id)).
		SetStatus(notificationdelivery.Status(notifyenum.NotificationDeliveryStatusFailed)).
		AddRetryCount(1).
		Exec(ctx)
}
