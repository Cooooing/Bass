package usecase

import (
	"context"
	notifychannel "notify/internal/biz/channel"
	"notify/internal/biz/repo"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

const defaultDeliveryLimit = 50

type DeliveryUsecase struct {
	log                      *log.Helper
	channelClient            notifychannel.Client
	notificationDeliveryRepo repo.NotificationDeliveryRepo
}

func NewDeliveryUsecase(
	logger log.Logger,
	channelClient notifychannel.Client,
	notificationDeliveryRepo repo.NotificationDeliveryRepo,
) *DeliveryUsecase {
	return &DeliveryUsecase{
		log:                      log.NewHelper(logger),
		channelClient:            channelClient,
		notificationDeliveryRepo: notificationDeliveryRepo,
	}
}

func (u *DeliveryUsecase) DispatchPending(ctx context.Context, limit int) error {
	if limit <= 0 {
		limit = defaultDeliveryLimit
	}
	deliveries, err := u.notificationDeliveryRepo.ListPending(ctx, limit)
	if err != nil {
		return err
	}
	for _, delivery := range deliveries {
		if delivery == nil {
			continue
		}
		claimed, err := u.notificationDeliveryRepo.MarkSending(ctx, delivery.ID)
		if err != nil {
			u.log.Errorf("mark delivery sending failed: id=%d err=%v", delivery.ID, err)
			continue
		}
		if !claimed {
			continue
		}
		channel, err := u.channelClient.GetChannel(delivery.Channel)
		if err == nil {
			err = channel.Send(ctx, &notifychannel.SendReq{
				EventID:    delivery.EventID,
				EventType:  delivery.EventType,
				ReceiverID: delivery.ReceiverID,
				Channel:    delivery.Channel,
				Target:     delivery.Target,
				Title:      delivery.Title,
				Content:    delivery.Content,
			})
		}
		if err != nil {
			if markErr := u.notificationDeliveryRepo.MarkFailed(ctx, delivery.ID); markErr != nil {
				u.log.Errorf("mark delivery failed status failed: id=%d err=%v", delivery.ID, markErr)
			}
			u.log.Errorf("send delivery failed: id=%d channel=%s target=%s err=%v", delivery.ID, delivery.Channel, delivery.Target, err)
			continue
		}
		if err := u.notificationDeliveryRepo.MarkSent(ctx, delivery.ID, time.Now()); err != nil {
			u.log.Errorf("mark delivery sent failed: id=%d err=%v", delivery.ID, err)
		}
	}
	return nil
}
