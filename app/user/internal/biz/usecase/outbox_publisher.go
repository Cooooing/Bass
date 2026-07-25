package usecase

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"context"
	"time"
	"user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"

	"log/slog"
)

const (
	outboxMinPollInterval = time.Second
	outboxMaxPollInterval = 8 * time.Second
	outboxPublishLimit    = 1000
	outboxPublishTimeout  = 5 * time.Minute
)

type OutboxPublisher struct {
	logger      *slog.Logger
	conf        *config.Bootstrap
	tx          base.Tx
	outboxRepo  repo.OutboxEventRepo
	eventClient repo.EventClient
	redisLock   *commonClient.RedisLock
	cancel      context.CancelFunc
}

func NewOutboxPublisher(
	logger *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	outboxRepo repo.OutboxEventRepo,
	eventClient repo.EventClient,
	redisLock *commonClient.RedisLock,
) *OutboxPublisher {
	return &OutboxPublisher{
		logger:      logger,
		conf:        conf,
		tx:          tx,
		outboxRepo:  outboxRepo,
		eventClient: eventClient,
		redisLock:   redisLock,
	}
}

func (p *OutboxPublisher) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	p.cancel = cancel
	go func() {
		pollInterval := p.minPollInterval()
		for {
			published, err := p.publishBatch(runCtx)
			if err != nil {
				p.logger.ErrorContext(runCtx, "publish outbox events failed", constant.LogFieldErr, err)
			}
			waitInterval := pollInterval
			if published {
				pollInterval = p.minPollInterval()
				waitInterval = p.minPollInterval()
			} else if pollInterval < p.maxPollInterval() {
				pollInterval *= 2
				if pollInterval > p.maxPollInterval() {
					pollInterval = p.maxPollInterval()
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
	lock, acquired, err := p.redisLock.TryAcquire(ctx, constant.GetKeyOutboxPublisherLock(p.serviceName()), p.pollLockTTL())
	if err != nil {
		p.logger.WarnContext(ctx, "acquire outbox publisher lock failed", constant.LogFieldErr, err)
		return false, nil
	}
	if !acquired {
		return false, nil
	}
	defer func() {
		if err := lock.Release(ctx); err != nil {
			p.logger.WarnContext(ctx, "release outbox publisher lock failed", constant.LogFieldErr, err)
		}
	}()
	var events []*model.OutboxEvent
	err = p.tx(ctx, func(ctx context.Context) error {
		claimed, err := p.outboxRepo.ClaimForPublish(ctx, &repo.OutboxEventClaimForPublishReq{
			Limit:       p.publishLimit(),
			StaleBefore: time.Now().Add(-p.publishTimeout()),
		})
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
			p.logger.ErrorContext(ctx, "publish outbox event failed", constant.LogFieldEventID, event.EventID, constant.LogFieldErr, err)
			if ctx.Err() != nil {
				return published, ctx.Err()
			}
			if markErr := p.outboxRepo.MarkFailed(ctx, &repo.OutboxEventMarkFailedReq{
				ID:        event.ID,
				LastError: err.Error(),
				MaxRetry:  p.maxRetry(),
			}); markErr != nil {
				p.logger.ErrorContext(ctx, "mark outbox event failed", constant.LogFieldEventID, event.EventID, constant.LogFieldErr, markErr)
				if batchErr == nil {
					batchErr = markErr
				}
			}
			continue
		}
		if err = p.outboxRepo.MarkPublished(ctx, &repo.OutboxEventMarkPublishedReq{
			ID:          event.ID,
			PublishedAt: time.Now(),
		}); err != nil {
			if ctx.Err() != nil {
				return published, ctx.Err()
			}
			p.logger.ErrorContext(ctx, "mark outbox event published failed", constant.LogFieldEventID, event.EventID, constant.LogFieldErr, err)
			if batchErr == nil {
				batchErr = err
			}
			continue
		}
		published = true
	}
	return published, batchErr
}

func (p *OutboxPublisher) serviceName() string {
	if p.conf != nil && p.conf.GetServer() != nil && p.conf.GetServer().GetName() != "" {
		return p.conf.GetServer().GetName()
	}
	return "user"
}

func (p *OutboxPublisher) publishLimit() int {
	if p.conf != nil && p.conf.GetEvent() != nil && p.conf.GetEvent().GetOutbox() != nil && p.conf.GetEvent().GetOutbox().GetPublishLimit() > 0 {
		return int(p.conf.GetEvent().GetOutbox().GetPublishLimit())
	}
	return outboxPublishLimit
}

func (p *OutboxPublisher) maxRetry() int32 {
	if p.conf != nil && p.conf.GetEvent() != nil && p.conf.GetEvent().GetOutbox() != nil && p.conf.GetEvent().GetOutbox().GetMaxRetry() > 0 {
		return p.conf.GetEvent().GetOutbox().GetMaxRetry()
	}
	return 10
}

func (p *OutboxPublisher) publishTimeout() time.Duration {
	if p.conf != nil && p.conf.GetEvent() != nil && p.conf.GetEvent().GetOutbox() != nil && p.conf.GetEvent().GetOutbox().GetPublishTimeout() != nil && p.conf.GetEvent().GetOutbox().GetPublishTimeout().AsDuration() > 0 {
		return p.conf.GetEvent().GetOutbox().GetPublishTimeout().AsDuration()
	}
	return outboxPublishTimeout
}

func (p *OutboxPublisher) pollLockTTL() time.Duration {
	if p.conf != nil && p.conf.GetEvent() != nil && p.conf.GetEvent().GetOutbox() != nil && p.conf.GetEvent().GetOutbox().GetPollLockTtl() != nil && p.conf.GetEvent().GetOutbox().GetPollLockTtl().AsDuration() > 0 {
		return p.conf.GetEvent().GetOutbox().GetPollLockTtl().AsDuration()
	}
	return time.Minute
}

func (p *OutboxPublisher) minPollInterval() time.Duration {
	if p.conf != nil && p.conf.GetEvent() != nil && p.conf.GetEvent().GetOutbox() != nil && p.conf.GetEvent().GetOutbox().GetMinPollInterval() != nil && p.conf.GetEvent().GetOutbox().GetMinPollInterval().AsDuration() > 0 {
		return p.conf.GetEvent().GetOutbox().GetMinPollInterval().AsDuration()
	}
	return outboxMinPollInterval
}

func (p *OutboxPublisher) maxPollInterval() time.Duration {
	if p.conf != nil && p.conf.GetEvent() != nil && p.conf.GetEvent().GetOutbox() != nil && p.conf.GetEvent().GetOutbox().GetMaxPollInterval() != nil && p.conf.GetEvent().GetOutbox().GetMaxPollInterval().AsDuration() > 0 {
		return p.conf.GetEvent().GetOutbox().GetMaxPollInterval().AsDuration()
	}
	return outboxMaxPollInterval
}
