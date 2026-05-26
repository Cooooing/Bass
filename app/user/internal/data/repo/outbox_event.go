package repo

import (
	commonenums "common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	"context"
	"fmt"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/outboxevent"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
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

func (r *OutboxEventRepo) Save(ctx context.Context, req *repo.OutboxEventSave) error {
	if req == nil {
		return fmt.Errorf("outbox event save request is nil")
	}
	if req.Event == nil {
		return fmt.Errorf("outbox event is nil")
	}
	eventType, ok := commonenum.EventTypeMap.ToProto(req.EventType)
	if !ok {
		return fmt.Errorf("unknown event type: %s", req.EventType)
	}
	if _, ok = commonenum.EventSubjectMap.ToProto(req.Subject); !ok {
		return fmt.Errorf("unknown event subject: %s", req.Subject)
	}

	normalizeEvent(req.Event, eventType)

	payloadBytes, err := proto.Marshal(req.Event)
	if err != nil {
		return err
	}

	return r.getClient(ctx).OutboxEvent.Create().
		SetEventID(req.Event.EventId).
		SetEventType(outboxevent.EventType(req.EventType)).
		SetSubject(outboxevent.Subject(req.Subject)).
		SetPayload(payloadBytes).
		Exec(ctx)
}

func normalizeEvent(event *commonenums.Event, eventType commonenums.EventType) {
	if event.EventId == "" {
		event.EventId = uuid.NewString()
	}
	event.Type = eventType
	if event.Timestamp == nil {
		event.Timestamp = timestamppb.Now()
	}
}
