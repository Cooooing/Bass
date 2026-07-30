package repo

import (
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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ repo.OutboxEventRepo = (*OutboxEventRepo)(nil)

type OutboxEventRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewOutboxEventRepo(
	db *gen.Client,
) repo.OutboxEventRepo {
	return &OutboxEventRepo{
		db: db,
	}
}

func (r *OutboxEventRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *OutboxEventRepo) Save(ctx context.Context, event *commonenums.Event) (*repo.OutboxEvent, error) {
	if event == nil {
		return nil, fmt.Errorf("outbox event is nil")
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(event.Type)
	if !ok {
		return nil, fmt.Errorf("unknown event type: %s", event.Type)
	}
	subject, ok := commonenum.EventSubjectMap.ToEnum(event.Subject)
	if !ok {
		return nil, fmt.Errorf("unknown event subject: %s", event.Subject)
	}
	if event.EventId == "" {
		event.EventId = uuid.NewString()
	}
	if event.Timestamp == nil {
		event.Timestamp = timestamppb.Now()
	}
	payloadBytes, err := protojson.MarshalOptions{
		UseProtoNames: true,
	}.Marshal(event)
	if err != nil {
		return nil, err
	}
	payload := string(payloadBytes)
	headers := map[string]string{}
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(headers))

	create := r.getClient(ctx).OutboxEvent.Create().
		SetEventID(event.EventId).
		SetEventType(outboxevent.EventType(eventType)).
		SetSubject(subject).
		SetPayload(payload)
	if len(headers) > 0 {
		create.SetHeaders(headers)
	}
	saved, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.OutboxEvent{
		ID:         saved.ID,
		EventID:    saved.EventID,
		EventType:  commonenum.EventType(saved.EventType),
		Subject:    saved.Subject,
		Payload:    saved.Payload,
		Headers:    saved.Headers,
		Status:     commonenum.OutboxEventStatus(saved.Status),
		RetryCount: saved.RetryCount,
		LastError:  saved.LastError,
		UpdatedAt:  saved.UpdatedAt,
	}, nil
}

func (r *OutboxEventRepo) Get(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEvent, error) {
	query := r.getClient(ctx).OutboxEvent.Query()
	query = r.getQuery(query, req)
	event, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &repo.OutboxEvent{
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
	}, nil
}

func (r *OutboxEventRepo) List(ctx context.Context, req *repo.OutboxEventGetReq) ([]*repo.OutboxEvent, error) {
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
	return result, nil
}

func (r *OutboxEventRepo) Map(ctx context.Context, req *repo.OutboxEventGetReq) (map[int64]*repo.OutboxEvent, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*repo.OutboxEvent, len(listResp))
	for _, item := range listResp {
		result[item.ID] = item
	}
	return result, nil
}

func (r *OutboxEventRepo) Count(ctx context.Context, req *repo.OutboxEventGetReq) (int, error) {
	query := r.getClient(ctx).OutboxEvent.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *OutboxEventRepo) Page(ctx context.Context, req *repo.OutboxEventGetReq) (*repo.OutboxEventPageResp, error) {
	page := r.normalizePage(req.Page)
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
	return &repo.OutboxEventPageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *OutboxEventRepo) ClaimOneForPublish(ctx context.Context, req *repo.OutboxEventClaimOneForPublishReq) (*repo.OutboxEvent, error) {
	if req == nil || req.ID == 0 || req.StaleBefore == nil {
		return nil, errors.New("outbox event id and stale_before are required")
	}
	c := r.getClient(ctx)
	event, err := c.OutboxEvent.Query().
		Where(
			outboxevent.ID(req.ID),
			outboxevent.Or(
				outboxevent.StatusIn(
					outboxevent.Status(commonenum.OutboxEventStatusPending),
					outboxevent.Status(commonenum.OutboxEventStatusFailed),
				),
				outboxevent.And(
					outboxevent.StatusEQ(outboxevent.Status(commonenum.OutboxEventStatusPublishing)),
					outboxevent.UpdatedAtLT(*req.StaleBefore),
				),
			),
			func(s *sql.Selector) {
				s.ForUpdate(sql.WithLockAction(sql.SkipLocked))
			},
		).
		Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if _, err = c.OutboxEvent.Update().
		Where(outboxevent.ID(event.ID)).
		SetStatus(outboxevent.Status(commonenum.OutboxEventStatusPublishing)).
		Save(ctx); err != nil {
		return nil, err
	}
	return &repo.OutboxEvent{
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
	}, nil
}

func (r *OutboxEventRepo) ClaimForPublish(ctx context.Context, req *repo.OutboxEventClaimForPublishReq) ([]*repo.OutboxEvent, error) {
	if req == nil || req.StaleBefore == nil {
		return nil, errors.New("outbox event stale_before is required")
	}
	c := r.getClient(ctx)
	events, err := c.OutboxEvent.Query().
		Where(
			outboxevent.Or(
				outboxevent.StatusIn(
					outboxevent.Status(commonenum.OutboxEventStatusPending),
					outboxevent.Status(commonenum.OutboxEventStatusFailed),
				),
				outboxevent.And(
					outboxevent.StatusEQ(outboxevent.Status(commonenum.OutboxEventStatusPublishing)),
					outboxevent.UpdatedAtLT(*req.StaleBefore),
				),
			),
			func(s *sql.Selector) {
				s.ForUpdate(sql.WithLockAction(sql.SkipLocked))
			},
		).
		Order(outboxevent.ByID()).
		Limit(req.Limit).
		All(ctx)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, nil
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
		return nil, err
	}
	return claimed, nil
}

func (r *OutboxEventRepo) MarkPublished(ctx context.Context, req *repo.OutboxEventMarkPublishedReq) error {
	if req == nil || req.PublishedAt == nil {
		return errors.New("outbox event published_at is required")
	}
	return r.getClient(ctx).OutboxEvent.UpdateOneID(req.ID).
		SetStatus(outboxevent.Status(commonenum.OutboxEventStatusPublished)).
		SetPublishedAt(*req.PublishedAt).
		ClearLastError().
		Exec(ctx)
}

func (r *OutboxEventRepo) MarkFailed(ctx context.Context, req *repo.OutboxEventMarkFailedReq) error {
	if req == nil {
		return errors.New("outbox event mark failed request is required")
	}
	lastError := req.LastError
	maxRetry := req.MaxRetry
	if maxRetry <= 0 {
		maxRetry = 1
	}
	if len(lastError) > 255 {
		lastError = lastError[:255]
	}
	c := r.getClient(ctx)
	event, err := c.OutboxEvent.Get(ctx, req.ID)
	if err != nil {
		return err
	}
	retryCount := event.RetryCount + 1
	status := commonenum.OutboxEventStatusFailed
	if retryCount >= maxRetry {
		status = commonenum.OutboxEventStatusDead
	}
	return c.OutboxEvent.UpdateOneID(req.ID).
		SetStatus(outboxevent.Status(status)).
		SetRetryCount(retryCount).
		SetLastError(lastError).
		Exec(ctx)
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
