package repo

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"fmt"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/inboxevent"
	"time"

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

func (r *InboxEventRepo) ClaimRetry(ctx context.Context, eventID string, now time.Time, processingTimeout time.Duration) (bool, error) {
	count, err := r.getClient(ctx).InboxEvent.Update().
		Where(
			inboxevent.EventIDEQ(eventID),
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
	return count > 0, err
}

func (r *InboxEventRepo) MarkProcessed(ctx context.Context, eventID string, now time.Time) error {
	return r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(eventID)).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessed)).
		SetProcessedAt(now).
		ClearLastError().
		Exec(ctx)
}

func (r *InboxEventRepo) MarkFailed(ctx context.Context, eventID string, lastError string) error {
	return r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(eventID)).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusFailed)).
		SetLastError(lastError).
		Exec(ctx)
}
