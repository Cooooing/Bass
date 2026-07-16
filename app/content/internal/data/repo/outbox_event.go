package repo

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	commonenums "common/proto/gen/common/enums"
	"content/internal/biz/base"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/outboxevent"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ repo.OutboxEventRepo = (*OutboxEventRepo)(nil)

type OutboxEventRepo struct {
	db *gen.Client
}

func NewOutboxEventRepo(db *gen.Client) repo.OutboxEventRepo {
	return &OutboxEventRepo{db: db}
}

func (r *OutboxEventRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
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

func (r *OutboxEventRepo) Save(ctx context.Context, req *repo.OutboxEventSave) (*repo.OutboxEventSaveResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("outbox event save request is nil")
	}
	if req.Event == nil {
		return nil, fmt.Errorf("outbox event is nil")
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(req.Event.Type)
	if !ok {
		return nil, fmt.Errorf("unknown event type: %s", req.Event.Type)
	}
	subject, ok := commonenum.EventSubjectMap.ToEnum(req.Event.Subject)
	if !ok {
		return nil, fmt.Errorf("unknown event subject: %s", req.Event.Subject)
	}
	r.normalizeEvent(req.Event)
	payloadBytes, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(req.Event)
	if err != nil {
		return nil, err
	}
	payload := string(payloadBytes)
	headers := map[string]string{}
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(headers))

	if err := r.withTxClient(ctx, func(c *gen.Client) error {
		create := c.OutboxEvent.Create().
			SetEventID(req.Event.EventId).
			SetEventType(outboxevent.EventType(eventType)).
			SetSubject(subject).
			SetPayload(payload)
		if len(headers) > 0 {
			create.SetHeaders(headers)
		}
		return create.Exec(ctx)
	}); err != nil {
		return nil, err
	}
	return &repo.OutboxEventSaveResponse{}, nil
}

func (r *OutboxEventRepo) Get(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEventGetResponse, error) {
	query := r.getClient(ctx).OutboxEvent.Query()
	query = r.getQuery(query, req)
	event, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &repo.OutboxEventGetResponse{Event: &repo.OutboxEvent{
		ID:         event.ID,
		EventID:    event.EventID,
		EventType:  commonenum.EventType(event.EventType),
		Subject:    event.Subject,
		Payload:    event.Payload,
		Headers:    event.Headers,
		Status:     commonenum.OutboxEventStatus(event.Status),
		RetryCount: event.RetryCount,
		LastError:  event.LastError,
		UpdatedAt:  event.UpdatedAt,
	}}, nil
}

func (r *OutboxEventRepo) List(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEventListResponse, error) {
	query := r.getClient(ctx).OutboxEvent.Query()
	query = r.getQuery(query, req)
	events, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*repo.OutboxEvent, 0, len(events))
	for _, event := range events {
		result = append(result, &repo.OutboxEvent{
			ID:         event.ID,
			EventID:    event.EventID,
			EventType:  commonenum.EventType(event.EventType),
			Subject:    event.Subject,
			Payload:    event.Payload,
			Headers:    event.Headers,
			Status:     commonenum.OutboxEventStatus(event.Status),
			RetryCount: event.RetryCount,
			LastError:  event.LastError,
			UpdatedAt:  event.UpdatedAt,
		})
	}
	return &repo.OutboxEventListResponse{Rows: result}, nil
}

func (r *OutboxEventRepo) Map(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEventMapResponse, error) {
	listResponse, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.OutboxEventMapResponse{Rows: lo.SliceToMap(listResponse.Rows, func(item *repo.OutboxEvent) (int64, *repo.OutboxEvent) {
		return item.ID, item
	})}, nil
}

func (r *OutboxEventRepo) Count(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEventCountResponse, error) {
	query := r.getClient(ctx).OutboxEvent.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.OutboxEventCountResponse{Count: count}, nil
}

func (r *OutboxEventRepo) Page(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEventPageResponse, error) {
	page := normalizePage(req.Page)
	query := r.getClient(ctx).OutboxEvent.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	events, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*repo.OutboxEvent, 0, len(events))
	for _, event := range events {
		result = append(result, &repo.OutboxEvent{
			ID:         event.ID,
			EventID:    event.EventID,
			EventType:  commonenum.EventType(event.EventType),
			Subject:    event.Subject,
			Payload:    event.Payload,
			Headers:    event.Headers,
			Status:     commonenum.OutboxEventStatus(event.Status),
			RetryCount: event.RetryCount,
			LastError:  event.LastError,
			UpdatedAt:  event.UpdatedAt,
		})
	}
	return &repo.OutboxEventPageResponse{
		Rows: result,
		Page: &base.PageResponse{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *OutboxEventRepo) ClaimForPublish(ctx context.Context, req *repo.OutboxEventClaimForPublishReq) (*repo.OutboxEventClaimForPublishResponse, error) {
	limit := req.Limit
	staleBefore := req.StaleBefore
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
				ID:         event.ID,
				EventID:    event.EventID,
				EventType:  commonenum.EventType(event.EventType),
				Subject:    event.Subject,
				Payload:    event.Payload,
				Headers:    event.Headers,
				Status:     commonenum.OutboxEventStatus(event.Status),
				RetryCount: event.RetryCount,
				LastError:  event.LastError,
				UpdatedAt:  event.UpdatedAt,
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
	return &repo.OutboxEventClaimForPublishResponse{Rows: result}, nil
}

func (r *OutboxEventRepo) MarkPublished(ctx context.Context, req *repo.OutboxEventMarkPublishedReq) (*repo.OutboxEventMarkPublishedResponse, error) {
	id := req.ID
	publishedAt := req.PublishedAt
	if err := r.withTxClient(ctx, func(c *gen.Client) error {
		return c.OutboxEvent.UpdateOneID(id).
			SetStatus(outboxevent.Status(commonenum.OutboxEventStatusPublished)).
			SetPublishedAt(publishedAt).
			ClearLastError().
			Exec(ctx)
	}); err != nil {
		return nil, err
	}
	return &repo.OutboxEventMarkPublishedResponse{}, nil
}

func (r *OutboxEventRepo) MarkFailed(ctx context.Context, req *repo.OutboxEventMarkFailedReq) (*repo.OutboxEventMarkFailedResponse, error) {
	id := req.ID
	lastError := req.LastError
	maxRetry := req.MaxRetry
	if err := r.withTxClient(ctx, func(c *gen.Client) error {
		if maxRetry <= 0 {
			maxRetry = 1
		}
		if len(lastError) > 255 {
			lastError = lastError[:255]
		}
		event, err := c.OutboxEvent.Get(ctx, id)
		if err != nil {
			return err
		}
		retryCount := event.RetryCount + 1
		status := commonenum.OutboxEventStatusFailed
		if retryCount >= maxRetry {
			status = commonenum.OutboxEventStatusDead
		}
		return c.OutboxEvent.UpdateOneID(id).
			SetStatus(outboxevent.Status(status)).
			SetRetryCount(retryCount).
			SetLastError(lastError).
			Exec(ctx)
	}); err != nil {
		return nil, err
	}
	return &repo.OutboxEventMarkFailedResponse{}, nil
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

func (p *NatsEventClient) Publish(ctx context.Context, req *repo.EventClientPublishReq) (*repo.EventClientPublishResponse, error) {
	msg := req.Message
	if msg == nil {
		return nil, fmt.Errorf("event message is nil")
	}
	if err := p.natsClient.Publish(ctx, msg.Subject, &client.Message{
		Subject: msg.Subject,
		Data:    []byte(msg.Payload),
		Header:  msg.Headers,
	}); err != nil {
		return nil, err
	}
	return &repo.EventClientPublishResponse{}, nil
}

func (r *OutboxEventRepo) normalizeEvent(event *commonenums.Event) {
	if event.EventId == "" {
		event.EventId = uuid.NewString()
	}
	if event.Timestamp == nil {
		event.Timestamp = timestamppb.Now()
	}
}

func (r *OutboxEventRepo) getQuery(query *gen.OutboxEventQuery, req *repo.OutboxEventGetReq) *gen.OutboxEventQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(outboxevent.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(outboxevent.IDIn(req.IDs...))
	}
	if req.EventID != nil {
		query = query.Where(outboxevent.EventID(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(outboxevent.EventIDIn(req.EventIDs...))
	}
	if req.Subject != nil {
		subject, ok := commonenum.EventSubjectMap.ToEnum(*req.Subject)
		if ok {
			query = query.Where(outboxevent.SubjectEQ(subject))
		}
	}
	if req.Status != nil {
		query = query.Where(outboxevent.StatusEQ(outboxevent.Status(*req.Status)))
	}
	return query
}
