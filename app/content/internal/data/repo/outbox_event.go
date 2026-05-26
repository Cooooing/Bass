package repo

import (
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/outboxevent"
	"context"
	"fmt"

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
	if req.Event.EventId == "" {
		req.Event.EventId = uuid.NewString()
	}
	req.Event.Type = eventType
	if req.Event.Timestamp == nil {
		req.Event.Timestamp = timestamppb.Now()
	}
	payloadBytes, err := proto.Marshal(req.Event)
	if err != nil {
		return err
	}
	client := r.db
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		client = tx
	}
	return client.OutboxEvent.Create().
		SetEventID(req.Event.EventId).
		SetEventType(outboxevent.EventType(req.EventType)).
		SetSubject(outboxevent.Subject(req.Subject)).
		SetPayload(payloadBytes).
		Exec(ctx)
}
