package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	"common/proto/gen/common/enums"
	"context"
	"fmt"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/inboxevent"
	"time"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ bizrepo.InboxEventRepo = (*InboxEventRepo)(nil)

type InboxEventRepo struct {
	db *gen.Client
}

func NewInboxEventRepo(db *gen.Client) bizrepo.InboxEventRepo {
	return &InboxEventRepo{db: db}
}

func (r *InboxEventRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *InboxEventRepo) SaveProcessing(ctx context.Context, req *bizrepo.InboxEventSave, now time.Time) (*model.InboxEvent, bool, error) {
	if req == nil {
		return nil, false, fmt.Errorf("inbox event save request is nil")
	}
	if req.EventID == "" {
		return nil, false, fmt.Errorf("event id is required")
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(enums.EventType(req.EventType))
	if !ok {
		return nil, false, fmt.Errorf("unknown event type: %s", req.EventType.String())
	}
	if _, ok := commonenum.EventSubjectMap.ToProto(req.Subject); !ok {
		return nil, false, fmt.Errorf("unknown event subject: %s", req.Subject)
	}

	client := r.getClient(ctx)
	save, err := client.InboxEvent.Create().
		SetEventID(req.EventID).
		SetEventType(inboxevent.EventType(eventType)).
		SetSubject(req.Subject).
		SetPayload(req.Payload).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessing)).
		SetAttemptCount(1).
		SetProcessingStartedAt(now).
		Save(ctx)
	if err == nil {
		return &model.InboxEvent{
			ID:                  save.ID,
			EventID:             save.EventID,
			EventType:           commonenum.EventType(save.EventType),
			Subject:             save.Subject,
			Payload:             save.Payload,
			Status:              commonenum.InboxEventStatus(save.Status),
			AttemptCount:        save.AttemptCount,
			LastError:           save.LastError,
			ProcessingStartedAt: save.ProcessingStartedAt,
			ProcessedAt:         save.ProcessedAt,
			CreatedAt:           save.CreatedAt,
			UpdatedAt:           save.UpdatedAt,
		}, true, nil
	}
	if !gen.IsConstraintError(err) {
		return nil, false, err
	}

	exist, getErr := client.InboxEvent.Query().
		Where(inboxevent.EventIDEQ(req.EventID)).
		Only(ctx)
	if getErr != nil {
		return nil, false, getErr
	}
	return &model.InboxEvent{
		ID:                  exist.ID,
		EventID:             exist.EventID,
		EventType:           commonenum.EventType(exist.EventType),
		Subject:             exist.Subject,
		Payload:             exist.Payload,
		Status:              commonenum.InboxEventStatus(exist.Status),
		AttemptCount:        exist.AttemptCount,
		LastError:           exist.LastError,
		ProcessingStartedAt: exist.ProcessingStartedAt,
		ProcessedAt:         exist.ProcessedAt,
		CreatedAt:           exist.CreatedAt,
		UpdatedAt:           exist.UpdatedAt,
	}, false, nil
}

func (r *InboxEventRepo) Get(ctx context.Context, req *bizrepo.InboxEventQuery) (*model.InboxEvent, error) {
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.InboxEvent{
		ID:                  row.ID,
		EventID:             row.EventID,
		EventType:           commonenum.EventType(row.EventType),
		Subject:             row.Subject,
		Payload:             row.Payload,
		Status:              commonenum.InboxEventStatus(row.Status),
		AttemptCount:        row.AttemptCount,
		LastError:           row.LastError,
		ProcessingStartedAt: row.ProcessingStartedAt,
		ProcessedAt:         row.ProcessedAt,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
	}, nil
}

func (r *InboxEventRepo) List(ctx context.Context, req *bizrepo.InboxEventQuery) ([]*model.InboxEvent, error) {
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.InboxEvent, 0, len(list))
	for _, row := range list {
		result = append(result, &model.InboxEvent{
			ID:                  row.ID,
			EventID:             row.EventID,
			EventType:           commonenum.EventType(row.EventType),
			Subject:             row.Subject,
			Payload:             row.Payload,
			Status:              commonenum.InboxEventStatus(row.Status),
			AttemptCount:        row.AttemptCount,
			LastError:           row.LastError,
			ProcessingStartedAt: row.ProcessingStartedAt,
			ProcessedAt:         row.ProcessedAt,
			CreatedAt:           row.CreatedAt,
			UpdatedAt:           row.UpdatedAt,
		})
	}
	return result, nil
}

func (r *InboxEventRepo) Map(ctx context.Context, req *bizrepo.InboxEventQuery) (map[int64]*model.InboxEvent, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.InboxEvent, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *InboxEventRepo) Count(ctx context.Context, req *bizrepo.InboxEventQuery) (int, error) {
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *InboxEventRepo) Page(ctx context.Context, page *common.PageRequest, req *bizrepo.InboxEventQuery) ([]*model.InboxEvent, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.InboxEvent, 0, len(list))
	for _, row := range list {
		result = append(result, &model.InboxEvent{
			ID:                  row.ID,
			EventID:             row.EventID,
			EventType:           commonenum.EventType(row.EventType),
			Subject:             row.Subject,
			Payload:             row.Payload,
			Status:              commonenum.InboxEventStatus(row.Status),
			AttemptCount:        row.AttemptCount,
			LastError:           row.LastError,
			ProcessingStartedAt: row.ProcessingStartedAt,
			ProcessedAt:         row.ProcessedAt,
			CreatedAt:           row.CreatedAt,
			UpdatedAt:           row.UpdatedAt,
		})
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *InboxEventRepo) ClaimRetry(ctx context.Context, eventID string, now time.Time, processingTimeout time.Duration, maxRetry int32) (bool, error) {
	if maxRetry <= 0 {
		maxRetry = 1
	}
	count, err := r.getClient(ctx).InboxEvent.Update().
		Where(
			inboxevent.EventIDEQ(eventID),
			inboxevent.AttemptCountLT(maxRetry),
			inboxevent.Or(
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusFailed)),
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusReceived)),
				inboxevent.And(
					inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusProcessing)),
					inboxevent.ProcessingStartedAtLTE(now.Add(-processingTimeout)),
				),
			),
		).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessing)).
		SetProcessingStartedAt(now).
		AddAttemptCount(1).
		ClearLastError().
		ClearProcessedAt().
		Save(ctx)
	if err != nil || count > 0 {
		return count > 0, err
	}
	_, err = r.getClient(ctx).InboxEvent.Update().
		Where(
			inboxevent.EventIDEQ(eventID),
			inboxevent.AttemptCountGTE(maxRetry),
			inboxevent.Or(
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusFailed)),
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusReceived)),
				inboxevent.And(
					inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusProcessing)),
					inboxevent.ProcessingStartedAtLTE(now.Add(-processingTimeout)),
				),
			),
		).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusDead)).
		Save(ctx)
	return false, err
}

func (r *InboxEventRepo) MarkProcessed(ctx context.Context, eventID string, now time.Time) error {
	return r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(eventID)).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessed)).
		SetProcessedAt(now).
		ClearLastError().
		Exec(ctx)
}

func (r *InboxEventRepo) MarkFailed(ctx context.Context, eventID string, lastError string, maxRetry int32) error {
	if maxRetry <= 0 {
		maxRetry = 1
	}
	if len(lastError) > 255 {
		lastError = lastError[:255]
	}
	event, err := r.getClient(ctx).InboxEvent.Query().
		Where(inboxevent.EventIDEQ(eventID)).
		Only(ctx)
	if err != nil {
		return err
	}
	status := commonenum.InboxEventStatusFailed
	if event.AttemptCount >= maxRetry {
		status = commonenum.InboxEventStatusDead
	}
	return r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(eventID)).
		SetStatus(inboxevent.Status(status)).
		SetLastError(lastError).
		Exec(ctx)
}

func (r *InboxEventRepo) getQuery(query *gen.InboxEventQuery, req *bizrepo.InboxEventQuery) *gen.InboxEventQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(inboxevent.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(inboxevent.IDIn(req.IDs...))
	}
	if req.EventID != nil {
		query = query.Where(inboxevent.EventIDEQ(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(inboxevent.EventIDIn(req.EventIDs...))
	}
	if req.EventType != nil {
		query = query.Where(inboxevent.EventTypeEQ(inboxevent.EventType(*req.EventType)))
	}
	if req.Subject != nil {
		query = query.Where(inboxevent.SubjectEQ(*req.Subject))
	}
	if req.Status != nil {
		query = query.Where(inboxevent.StatusEQ(inboxevent.Status(*req.Status)))
	}
	return query
}
