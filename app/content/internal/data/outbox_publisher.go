package data

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	"content/internal/data/gen"
	"content/internal/data/gen/outboxevent"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type OutboxPublisher struct {
	log        *log.Helper
	db         *gen.Client
	natsClient *client.NatsClient
	cancel     context.CancelFunc
}

func NewOutboxPublisher(logger log.Logger, db *gen.Client, natsClient *client.NatsClient) *OutboxPublisher {
	return &OutboxPublisher{
		log:        log.NewHelper(logger),
		db:         db,
		natsClient: natsClient,
	}
}

func (p *OutboxPublisher) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			events, err := p.db.OutboxEvent.Query().
				Where(outboxevent.StatusIn(
					outboxevent.Status(commonenum.OutboxEventStatusPending),
					outboxevent.Status(commonenum.OutboxEventStatusFailed),
				)).
				Order(outboxevent.ByID()).
				Limit(50).
				All(runCtx)
			if err != nil {
				p.log.Errorf("query outbox events failed: %v", err)
			}
			for _, event := range events {
				if err := p.natsClient.Publish(runCtx, string(event.Subject), &client.Message{
					Subject: string(event.Subject),
					Data:    event.Payload,
					Header:  event.Headers,
				}); err != nil {
					p.log.Errorf("publish outbox event failed: event_id=%s err=%v", event.EventID, err)
					_ = p.db.OutboxEvent.UpdateOneID(event.ID).
						SetStatus(outboxevent.Status(commonenum.OutboxEventStatusFailed)).
						AddRetryCount(1).
						Exec(runCtx)
					continue
				}
				now := time.Now()
				_ = p.db.OutboxEvent.UpdateOneID(event.ID).
					SetStatus(outboxevent.Status(commonenum.OutboxEventStatusPublished)).
					SetPublishedAt(now).
					Exec(runCtx)
			}
			select {
			case <-runCtx.Done():
				return
			case <-ticker.C:
			}
		}
	}()
	return nil
}

func (p *OutboxPublisher) Stop(_ context.Context) error {
	if p.cancel != nil {
		p.cancel()
	}
	return nil
}
