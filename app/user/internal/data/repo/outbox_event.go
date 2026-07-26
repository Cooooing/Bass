package repo

import (
	commonenum "common/pkg/enum"
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"common/proto/gen/common"
	"context"
	"errors"
	"fmt"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/outboxevent"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ repo.OutboxEventRepo = (*OutboxEventRepo)(nil)

type OutboxEventRepo struct {
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

func (r *OutboxEventRepo) Save(ctx context.Context, req *repo.OutboxEventSave) (*model.OutboxEvent, error) {
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
	if req.Event.EventId == "" {
		req.Event.EventId = uuid.NewString()
	}
	if req.Event.Timestamp == nil {
		req.Event.Timestamp = timestamppb.Now()
	}
	payloadBytes, err := protojson.MarshalOptions{
		UseProtoNames: true,
	}.Marshal(req.Event)
	if err != nil {
		return nil, err
	}
	payload := string(payloadBytes)
	headers := map[string]string{}
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(headers))

	create := r.getClient(ctx).OutboxEvent.Create().
		SetEventID(req.Event.EventId).
		SetEventType(outboxevent.EventType(eventType)).
		SetSubject(subject).
		SetPayload(payload)
	if len(headers) > 0 {
		create.SetHeaders(headers)
	}
	event, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.OutboxEvent{
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

func (r *OutboxEventRepo) Get(ctx context.Context, req *repo.OutboxEventGetReq) (*model.OutboxEvent, error) {
	query := r.getQuery(r.getClient(ctx).OutboxEvent.Query(), req)
	event, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.OutboxEvent{
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

func (r *OutboxEventRepo) List(ctx context.Context, req *repo.OutboxEventGetReq) ([]*model.OutboxEvent, error) {
	events, err := r.getQuery(r.getClient(ctx).OutboxEvent.Query(), req).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.OutboxEvent, 0, len(events))
	for _, event := range events {
		result = append(result, &model.OutboxEvent{
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

func (r *OutboxEventRepo) Map(ctx context.Context, req *repo.OutboxEventGetReq) (map[int64]*model.OutboxEvent, error) {
	events, err := r.getQuery(r.getClient(ctx).OutboxEvent.Query(), req).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.OutboxEvent, len(events))
	for _, event := range events {
		result[event.ID] = &model.OutboxEvent{
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
		}
	}
	return result, nil
}

func (r *OutboxEventRepo) Count(ctx context.Context, req *repo.OutboxEventGetReq) (int, error) {
	return r.getQuery(r.getClient(ctx).OutboxEvent.Query(), req).Count(ctx)
}

func (r *OutboxEventRepo) Page(ctx context.Context, req *repo.OutboxEventPageReq) (*repo.OutboxEventPageResp, error) {
	page := server.PageValid(&common.PageReq{
		Page: req.Page.Page,
		Size: req.Page.Size,
	})
	query := r.getQuery(r.getClient(ctx).OutboxEvent.Query(), &req.Query)
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
	rows := make([]*model.OutboxEvent, 0, len(events))
	for _, event := range events {
		rows = append(rows, &model.OutboxEvent{
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
		Rows: rows,
		Page: repo.PageResp{
			Total: uint32(total),
			Page:  page.GetPage(),
			Size:  page.GetSize(),
		},
	}, nil
}

func (r *OutboxEventRepo) ClaimOneForPublish(ctx context.Context, req *repo.OutboxEventClaimOneForPublishReq) (*model.OutboxEvent, error) {
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
	return &model.OutboxEvent{
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

func (r *OutboxEventRepo) ClaimForPublish(ctx context.Context, req *repo.OutboxEventClaimForPublishReq) ([]*model.OutboxEvent, error) {
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
	claimed := make([]*model.OutboxEvent, 0, len(events))
	for _, event := range events {
		ids = append(ids, event.ID)
		claimed = append(claimed, &model.OutboxEvent{
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
	maxRetry := req.MaxRetry
	lastError := req.LastError
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
