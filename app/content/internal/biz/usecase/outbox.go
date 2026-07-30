package usecase

import (
	"common/pkg/constant"
	"content/internal/biz/base"
	"content/internal/biz/repo"
	"context"
	"log/slog"
	"time"
)

type OutboxUsecase struct {
	logger         *slog.Logger
	tx             base.Tx
	outboxRepo     repo.OutboxEventRepo
	eventClient    repo.EventClient
	defaultLimit   int
	defaultTimeout time.Duration
	defaultRetry   int32
}

func NewOutboxUsecase(
	logger *slog.Logger,
	tx base.Tx,
	outboxRepo repo.OutboxEventRepo,
	eventClient repo.EventClient,
) *OutboxUsecase {
	return &OutboxUsecase{
		logger:         logger,
		tx:             tx,
		outboxRepo:     outboxRepo,
		eventClient:    eventClient,
		defaultLimit:   1000,
		defaultTimeout: 1 * time.Minute,
		defaultRetry:   3,
	}
}

type PublishOutboxEventReq struct {
	ID             int64
	PublishTimeout time.Duration
	MaxRetry       int32
}

type PublishOutboxEventResp struct {
	Published bool
	Skipped   bool
}

func (u *OutboxUsecase) Publish(ctx context.Context, req *PublishOutboxEventReq) (*PublishOutboxEventResp, error) {
	publishTimeout := u.defaultTimeout
	maxRetry := u.defaultRetry
	if req.PublishTimeout > 0 {
		publishTimeout = req.PublishTimeout
	}
	if req.MaxRetry > 0 {
		maxRetry = req.MaxRetry
	}
	var event *repo.OutboxEvent
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		event, err = u.outboxRepo.ClaimOneForPublish(ctx, &repo.OutboxEventClaimOneForPublishReq{
			ID:          req.ID,
			StaleBefore: new(time.Now().Add(-publishTimeout)),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	if event == nil {
		return &PublishOutboxEventResp{
			Skipped: true,
		}, nil
	}
	published, err := u.publishClaimedEvent(ctx, event, maxRetry)
	if err != nil {
		return nil, err
	}
	return &PublishOutboxEventResp{
		Published: published,
	}, nil
}

type PublishOutboxEventsReq struct {
	Limit          int
	PublishTimeout time.Duration
	MaxRetry       int32
}

type PublishOutboxEventsResp struct {
	Claimed   int32
	Published int32
	Failed    int32
	Skipped   int32
}

func (u *OutboxUsecase) PublishBatch(ctx context.Context, req *PublishOutboxEventsReq) (*PublishOutboxEventsResp, error) {
	limit := u.defaultLimit
	publishTimeout := u.defaultTimeout
	maxRetry := u.defaultRetry
	if req != nil && req.Limit > 0 {
		limit = req.Limit
	}
	if req != nil && req.PublishTimeout > 0 {
		publishTimeout = req.PublishTimeout
	}
	if req != nil && req.MaxRetry > 0 {
		maxRetry = req.MaxRetry
	}
	var events []*repo.OutboxEvent
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		events, err = u.outboxRepo.ClaimForPublish(ctx, &repo.OutboxEventClaimForPublishReq{
			Limit:       limit,
			StaleBefore: new(time.Now().Add(-publishTimeout)),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	resp := &PublishOutboxEventsResp{
		Claimed: int32(len(events)),
	}
	for _, event := range events {
		if event == nil {
			resp.Skipped++
			continue
		}
		published, err := u.publishClaimedEvent(ctx, event, maxRetry)
		if err != nil {
			resp.Failed++
			u.logger.ErrorContext(ctx, "publish outbox event failed", constant.LogFieldEventID, event.EventID, constant.LogFieldErr, err)
			continue
		}
		if published {
			resp.Published++
		} else {
			resp.Failed++
		}
	}
	return resp, nil
}

func (u *OutboxUsecase) publishClaimedEvent(ctx context.Context, event *repo.OutboxEvent, maxRetry int32) (bool, error) {
	err := u.eventClient.Publish(ctx, &repo.EventClientMessage{
		Subject: event.Subject.String(),
		Payload: event.Payload,
		Headers: event.Headers,
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "publish outbox event failed", constant.LogFieldEventID, event.EventID, constant.LogFieldErr, err)
		if ctx.Err() != nil {
			return false, ctx.Err()
		}
		if markErr := u.outboxRepo.MarkFailed(ctx, &repo.OutboxEventMarkFailedReq{
			ID:        event.ID,
			LastError: err.Error(),
			MaxRetry:  maxRetry,
		}); markErr != nil {
			return false, markErr
		}
		return false, nil
	}
	if err = u.outboxRepo.MarkPublished(ctx, &repo.OutboxEventMarkPublishedReq{
		ID:          event.ID,
		PublishedAt: new(time.Now()),
	}); err != nil {
		return false, err
	}
	return true, nil
}
