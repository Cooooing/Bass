package repo

import (
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

func (r *InboxEventRepo) SaveReceived(ctx context.Context, req *bizrepo.InboxEventSave) (*model.InboxEvent, error) {
	if req == nil {
		return nil, fmt.Errorf("inbox event save request is nil")
	}
	if req.EventID == "" {
		return nil, fmt.Errorf("event id is required")
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(req.EventType)
	if !ok {
		return nil, fmt.Errorf("unknown event type: %s", req.EventType.String())
	}
	if _, ok := commonenum.EventSubjectMap.ToProto(req.Subject); !ok {
		return nil, fmt.Errorf("unknown event subject: %s", req.Subject)
	}
	client := r.getClient(ctx)
	save, err := client.InboxEvent.Create().
		SetEventID(req.EventID).
		SetEventType(inboxevent.EventType(eventType)).
		SetSubject(inboxevent.Subject(req.Subject)).
		SetPayload(req.Payload).
		Save(ctx)
	if err == nil {
		return &model.InboxEvent{
			ID:          save.ID,
			EventID:     save.EventID,
			EventType:   commonenum.EventType(save.EventType),
			Subject:     commonenum.EventSubject(save.Subject),
			Payload:     save.Payload,
			Status:      commonenum.InboxEventStatus(save.Status),
			RetryCount:  save.RetryCount,
			ReceivedAt:  save.ReceivedAt,
			ProcessedAt: save.ProcessedAt,
			CreatedAt:   save.CreatedAt,
			UpdatedAt:   save.UpdatedAt,
		}, nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}
	exist, getErr := client.InboxEvent.Query().
		Where(inboxevent.EventIDEQ(req.EventID)).
		Only(ctx)
	if getErr != nil {
		return nil, err
	}
	return &model.InboxEvent{
		ID:          exist.ID,
		EventID:     exist.EventID,
		EventType:   commonenum.EventType(exist.EventType),
		Subject:     commonenum.EventSubject(exist.Subject),
		Payload:     exist.Payload,
		Status:      commonenum.InboxEventStatus(exist.Status),
		RetryCount:  exist.RetryCount,
		ReceivedAt:  exist.ReceivedAt,
		ProcessedAt: exist.ProcessedAt,
		CreatedAt:   exist.CreatedAt,
		UpdatedAt:   exist.UpdatedAt,
	}, nil
}

func (r *InboxEventRepo) MarkProcessing(ctx context.Context, eventID string) (bool, error) {
	count, err := r.getClient(ctx).InboxEvent.Update().
		Where(
			inboxevent.EventIDEQ(eventID),
			inboxevent.StatusIn(
				inboxevent.Status(commonenum.InboxEventStatusReceived),
				inboxevent.Status(commonenum.InboxEventStatusFailed),
			),
		).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessing)).
		AddRetryCount(1).
		Save(ctx)
	return count > 0, err
}

func (r *InboxEventRepo) MarkProcessed(ctx context.Context, eventID string) error {
	now := time.Now()
	return r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(eventID)).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessed)).
		SetProcessedAt(now).
		Exec(ctx)
}

func (r *InboxEventRepo) MarkFailed(ctx context.Context, eventID string) error {
	return r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(eventID)).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusFailed)).
		Exec(ctx)
}
