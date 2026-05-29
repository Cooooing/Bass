package usecase

import (
	"context"
	"time"

	base "content/internal/biz/base"
	"content/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	outboxMinPollInterval = time.Second
	outboxMaxPollInterval = 8 * time.Second
	outboxPublishLimit    = 1000
	outboxPublishTimeout  = 5 * time.Minute
)

type OutboxPublisher struct {
	log         *log.Helper
	tx          base.Tx
	outboxRepo  repo.OutboxEventRepo
	eventClient repo.EventClient
	cancel      context.CancelFunc
}

func NewOutboxPublisher(
	logger log.Logger,
	tx base.Tx,
	outboxRepo repo.OutboxEventRepo,
	eventClient repo.EventClient,
) *OutboxPublisher {
	return &OutboxPublisher{
		log:         log.NewHelper(logger),
		tx:          tx,
		outboxRepo:  outboxRepo,
		eventClient: eventClient,
	}
}

func (p *OutboxPublisher) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	go func() {
		pollInterval := outboxMinPollInterval
		for {
			published, err := p.publishBatch(runCtx)
			if err != nil {
				p.log.Errorf("publish outbox events failed: %v", err)
			}
			waitInterval := pollInterval
			if published {
				pollInterval = outboxMinPollInterval
				waitInterval = outboxMinPollInterval
			} else if pollInterval < outboxMaxPollInterval {
				pollInterval *= 2
				if pollInterval > outboxMaxPollInterval {
					pollInterval = outboxMaxPollInterval
				}
			}
			select {
			case <-runCtx.Done():
				return
			case <-time.After(waitInterval):
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

func (p *OutboxPublisher) publishBatch(ctx context.Context) (bool, error) {
	var events []*repo.OutboxEvent
	err := p.tx(ctx, func(ctx context.Context) error {
		claimed, err := p.outboxRepo.ClaimForPublish(ctx, outboxPublishLimit, time.Now().Add(-outboxPublishTimeout))
		if err != nil {
			return err
		}
		events = claimed
		return nil
	})
	if err != nil {
		return false, err
	}
	if len(events) == 0 {
		return false, nil
	}
	published := false
	var batchErr error
	for _, event := range events {
		err = p.eventClient.Publish(ctx, &repo.EventClientMessage{
			Subject: string(event.Subject),
			Payload: event.Payload,
			Headers: event.Headers,
		})
		if err != nil {
			p.log.Errorf("publish outbox event failed: event_id=%s err=%v", event.EventID, err)
			if ctx.Err() != nil {
				return published, ctx.Err()
			}
			if markErr := p.outboxRepo.MarkFailed(ctx, event.ID); markErr != nil {
				p.log.Errorf("mark outbox event failed: event_id=%s err=%v", event.EventID, markErr)
				if batchErr == nil {
					batchErr = markErr
				}
			}
			continue
		}
		if err = p.outboxRepo.MarkPublished(ctx, event.ID, time.Now()); err != nil {
			if ctx.Err() != nil {
				return published, ctx.Err()
			}
			p.log.Errorf("mark outbox event published failed: event_id=%s err=%v", event.EventID, err)
			if batchErr == nil {
				batchErr = err
			}
			continue
		}
		published = true
	}
	return published, batchErr
}
