package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"common/pkg/client"
	commonenum "common/pkg/enum"
	"common/proto/gen/common/enums"
	"notify/internal/biz/usecase"

	"google.golang.org/protobuf/encoding/protojson"
)

type Consumer struct {
	log             *slog.Logger
	natsClient      *client.NatsClient
	consumerUsecase *usecase.ConsumerUsecase
	subjects        []commonenum.EventSubject
	ctx             context.Context
	cancel          context.CancelFunc
}

func NewConsumer(
	logger *slog.Logger,
	natsClient *client.NatsClient,
	consumerUsecase *usecase.ConsumerUsecase,
	subjects []commonenum.EventSubject,
) *Consumer {
	return &Consumer{
		log:             logger,
		natsClient:      natsClient,
		consumerUsecase: consumerUsecase,
		subjects:        subjects,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)
	queueGroup := string(commonenum.EventQueueGroupNotify)
	for _, subject := range c.subjects {
		subjectName := string(subject)
		_, err := c.natsClient.QueueSubscribe(c.ctx, subjectName, queueGroup, func(ctx context.Context, msg *client.Message) error {
			if msg == nil {
				return nil
			}
			payload := msg.Data
			var event enums.Event
			if err := protojson.Unmarshal(payload, &event); err != nil {
				c.log.Error(fmt.Sprintf("unmarshal event failed: subject=%s err=%v", msg.Subject, err))
				return err
			}
			if event.EventId == "" {
				c.log.Error(fmt.Sprintf("event id is empty, inbox status cannot be saved: subject=%s type=%v", msg.Subject, event.Type))
				return errors.New("event id is required")
			}
			eventType, ok := commonenum.EventTypeMap.ToEnum(event.Type)
			if !ok {
				err := fmt.Errorf("unknown event type: %s", event.Type.String())
				c.log.Error(fmt.Sprintf("event type invalid: event_id=%s subject=%s type=%v err=%v", event.EventId, msg.Subject, event.Type, err))
				return err
			}
			eventSubject := commonenum.EventSubject(msg.Subject)
			if _, ok := commonenum.EventSubjectMap.ToProto(eventSubject); !ok {
				err := fmt.Errorf("unknown event subject: %s", msg.Subject)
				c.log.Error(fmt.Sprintf("event subject invalid: event_id=%s subject=%s type=%v err=%v", event.EventId, msg.Subject, event.Type, err))
				return err
			}
			if expectedSubject, ok := commonenum.EventSubjectMap.ToEnum(enums.EventSubject(event.Type)); !ok || expectedSubject != eventSubject {
				err := fmt.Errorf("event subject mismatch: subject=%s type=%s", msg.Subject, event.Type.String())
				c.log.Error(fmt.Sprintf("event subject mismatch: event_id=%s subject=%s type=%v err=%v", event.EventId, msg.Subject, event.Type, err))
				return err
			}
			return c.consumerUsecase.Handle(ctx, &usecase.ConsumerHandleReq{
				Event:     &event,
				EventType: eventType,
				Subject:   eventSubject,
				Payload:   string(payload),
			})
		})
		if err != nil {
			c.log.Error(fmt.Sprintf("queue subscribe failed: subject=%s queue=%s err=%v", subjectName, queueGroup, err))
			continue
		}
	}
	return nil
}

func (c *Consumer) Stop(ctx context.Context) error {
	if c.cancel == nil {
		return nil
	}
	c.cancel()
	return nil
}
