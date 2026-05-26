package usecase

import (
	commonenum "common/pkg/enum"
	"context"
	"errors"
	notifychannel "notify/internal/biz/channel"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func TestDispatchPendingSendsDelivery(t *testing.T) {
	delivery := &model.NotificationDelivery{
		ID:        1,
		EventID:   "event-1",
		EventType: commonenum.EventTypeUserRegister,
		Channel:   notifyenum.NotificationChannelEmail,
		Target:    "user@example.com",
		Title:     "title",
		Content:   "content",
		Status:    notifyenum.NotificationDeliveryStatusPending,
	}
	repo := &fakeDeliveryRepo{deliveries: []*model.NotificationDelivery{delivery}}
	channel := &fakeDeliveryChannel{}
	usecase := NewDeliveryUsecase(log.DefaultLogger, &fakeChannelClient{channel: channel}, repo)

	if err := usecase.DispatchPending(context.Background(), 10); err != nil {
		t.Fatalf("dispatch pending: %v", err)
	}
	if !repo.sendingMarked {
		t.Fatal("sending status was not marked")
	}
	if !repo.sentMarked {
		t.Fatal("sent status was not marked")
	}
	if channel.req == nil || channel.req.Target != "user@example.com" {
		t.Fatalf("send target = %v, want user@example.com", channel.req)
	}
}

func TestDispatchPendingMarksFailedWhenChannelMissing(t *testing.T) {
	repo := &fakeDeliveryRepo{deliveries: []*model.NotificationDelivery{
		{
			ID:        1,
			EventID:   "event-1",
			EventType: commonenum.EventTypeUserRegister,
			Channel:   notifyenum.NotificationChannelEmail,
			Target:    "user@example.com",
			Status:    notifyenum.NotificationDeliveryStatusPending,
		},
	}}
	usecase := NewDeliveryUsecase(log.DefaultLogger, &fakeChannelClient{err: errors.New("missing channel")}, repo)

	if err := usecase.DispatchPending(context.Background(), 10); err != nil {
		t.Fatalf("dispatch pending: %v", err)
	}
	if !repo.failedMarked {
		t.Fatal("failed status was not marked")
	}
	if repo.sentMarked {
		t.Fatal("sent status should not be marked")
	}
}

type fakeDeliveryRepo struct {
	deliveries    []*model.NotificationDelivery
	sendingMarked bool
	sentMarked    bool
	failedMarked  bool
}

func (r *fakeDeliveryRepo) Saves(context.Context, []*model.NotificationDelivery) ([]*model.NotificationDelivery, error) {
	panic("not implemented")
}

func (r *fakeDeliveryRepo) ListPending(context.Context, int) ([]*model.NotificationDelivery, error) {
	return r.deliveries, nil
}

func (r *fakeDeliveryRepo) MarkSending(context.Context, int64) (bool, error) {
	r.sendingMarked = true
	return true, nil
}

func (r *fakeDeliveryRepo) MarkSent(context.Context, int64, time.Time) error {
	r.sentMarked = true
	return nil
}

func (r *fakeDeliveryRepo) MarkFailed(context.Context, int64) error {
	r.failedMarked = true
	return nil
}

type fakeChannelClient struct {
	channel notifychannel.Channel
	err     error
}

func (c *fakeChannelClient) GetChannel(notifyenum.NotificationChannel) (notifychannel.Channel, error) {
	if c.err != nil {
		return nil, c.err
	}
	return c.channel, nil
}

type fakeDeliveryChannel struct {
	req *notifychannel.SendReq
}

func (c *fakeDeliveryChannel) Send(_ context.Context, req *notifychannel.SendReq) error {
	c.req = req
	return nil
}
