package schema

import (
	"common/pkg/constant"
	commonevent "common/pkg/event"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OutboxEvent stores durable events produced by the user service.
type OutboxEvent struct {
	ent.Schema
}

func (OutboxEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "outbox_events"},
		entsql.WithComments(true),
	}
}

func (OutboxEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("event id for idempotency").NotEmpty().Unique(),
		field.Int32("event_type").Comment("common enums.EventType value"),
		field.String("subject").Comment("NATS subject").NotEmpty(),
		field.String("aggregate_type").Comment("business aggregate type").Optional().Nillable(),
		field.String("aggregate_id").Comment("business aggregate id").Optional().Nillable(),
		field.String("producer_service").Comment("producer service name").NotEmpty(),
		field.Bytes("payload").Comment("protobuf encoded common enums.Event"),
		field.JSON("headers", map[string]string{}).Comment("message headers").Default(map[string]string{}),
		field.Enum("status").Values(commonevent.OutboxStatusValues()...).Default(string(commonevent.OutboxStatusPending)).Comment("dispatch status"),
		field.Int32("retry_count").Comment("dispatch retry count").Default(0),
		field.Int32("max_retry").Comment("maximum dispatch retries").Default(10),
		field.Time("next_retry_at").Comment("next retry time").Optional().Nillable(),
		field.String("locked_by").Comment("dispatcher lock owner").Optional().Nillable(),
		field.Time("locked_until").Comment("dispatcher lock expiration").Optional().Nillable(),
		field.Time("published_at").Comment("published time").Optional().Nillable(),
		field.Text("last_error").Comment("last dispatch error").Optional().Nillable(),
	}
}

func (OutboxEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (OutboxEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "next_retry_at", "id"),
		index.Fields("aggregate_type", "aggregate_id"),
	}
}
