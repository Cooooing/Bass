package repo

import (
	commonenums "common/api/gen/common/enums"
	"common/pkg/client"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/outboxevent"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ repo.OutboxEventRepo = (*OutboxEventRepo)(nil)

type OutboxEventRepo struct {
	db *gen.Client
}

func NewOutboxEventRepo(db *gen.Client) repo.OutboxEventRepo {
	return &OutboxEventRepo{db: db}
}

func (r *OutboxEventRepo) withTxClient(ctx context.Context, fn func(c *gen.Client) error) error {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return fn(c)
	}
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()
	if err = fn(tx.Client()); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Join(err, fmt.Errorf("tx rollback failed: %w", rbErr))
		}
		return err
	}
	return tx.Commit()
}

func (r *OutboxEventRepo) Save(ctx context.Context, req *repo.OutboxEventSave) error {
	if req == nil {
		return fmt.Errorf("outbox event save request is nil")
	}
	if req.Event == nil {
		return fmt.Errorf("outbox event is nil")
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(req.Event.Type)
	if !ok {
		return fmt.Errorf("unknown event type: %s", req.Event.Type)
	}
	subject, ok := commonenum.EventSubjectMap.ToEnum(req.Event.Subject)
	if !ok {
		return fmt.Errorf("unknown event subject: %s", req.Event.Subject)
	}
	normalizeEvent(req.Event)
	payloadBytes, err := proto.Marshal(req.Event)
	if err != nil {
		return err
	}

	return r.withTxClient(ctx, func(c *gen.Client) error {
		return c.OutboxEvent.Create().
			SetEventID(req.Event.EventId).
			SetEventType(outboxevent.EventType(eventType)).
			SetSubject(subject).
			SetPayload(payloadBytes).
			Exec(ctx)
	})
}

func (r *OutboxEventRepo) ClaimForPublish(ctx context.Context, limit int, staleBefore time.Time) ([]*repo.OutboxEvent, error) {
	var result []*repo.OutboxEvent
	err := r.withTxClient(ctx, func(c *gen.Client) error {
		events, err := c.OutboxEvent.Query().
			Where(
				outboxevent.Or(
					outboxevent.StatusIn(
						outboxevent.Status(commonenum.OutboxEventStatusPending),
						outboxevent.Status(commonenum.OutboxEventStatusFailed),
					),
					outboxevent.And(
						outboxevent.StatusEQ(outboxevent.Status(commonenum.OutboxEventStatusPublishing)),
						outboxevent.UpdatedAtLT(staleBefore),
					),
				),
				func(s *sql.Selector) {
					s.ForUpdate(sql.WithLockAction(sql.SkipLocked))
				},
			).
			Order(outboxevent.ByID()).
			Limit(limit).
			All(ctx)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			return nil
		}
		ids := make([]int64, 0, len(events))
		claimed := make([]*repo.OutboxEvent, 0, len(events))
		for _, event := range events {
			ids = append(ids, event.ID)
			claimed = append(claimed, &repo.OutboxEvent{
				ID:      event.ID,
				EventID: event.EventID,
				Subject: event.Subject,
				Payload: event.Payload,
				Headers: event.Headers,
			})
		}
		if _, err = c.OutboxEvent.Update().
			Where(outboxevent.IDIn(ids...)).
			SetStatus(outboxevent.Status(commonenum.OutboxEventStatusPublishing)).
			Save(ctx); err != nil {
			return err
		}
		result = claimed
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *OutboxEventRepo) MarkPublished(ctx context.Context, id int64, publishedAt time.Time) error {
	return r.withTxClient(ctx, func(c *gen.Client) error {
		return c.OutboxEvent.UpdateOneID(id).
			SetStatus(outboxevent.Status(commonenum.OutboxEventStatusPublished)).
			SetPublishedAt(publishedAt).
			Exec(ctx)
	})
}

func (r *OutboxEventRepo) MarkFailed(ctx context.Context, id int64) error {
	return r.withTxClient(ctx, func(c *gen.Client) error {
		return c.OutboxEvent.UpdateOneID(id).
			SetStatus(outboxevent.Status(commonenum.OutboxEventStatusFailed)).
			AddRetryCount(1).
			Exec(ctx)
	})
}

var _ repo.EventClient = (*NatsEventClient)(nil)

type NatsEventClient struct {
	natsClient *client.NatsClient
}

func NewNatsEventClient(natsClient *client.NatsClient) repo.EventClient {
	return &NatsEventClient{
		natsClient: natsClient,
	}
}

func (p *NatsEventClient) Publish(ctx context.Context, msg *repo.EventClientMessage) error {
	if msg == nil {
		return fmt.Errorf("event message is nil")
	}
	return p.natsClient.Publish(ctx, msg.Subject, &client.Message{
		Subject: msg.Subject,
		Data:    msg.Payload,
		Header:  msg.Headers,
	})
}

func normalizeEvent(event *commonenums.Event) {
	if event.EventId == "" {
		event.EventId = uuid.NewString()
	}
	if event.Timestamp == nil {
		event.Timestamp = timestamppb.Now()
	}
}
