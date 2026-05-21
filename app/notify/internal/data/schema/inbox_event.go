package schema

import (
	"common/pkg/constant"
	commonevent "common/pkg/event"
	utilent "common/pkg/util/ent"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// InboxEvent records events consumed by the notify service for idempotency.
type InboxEvent struct {
	ent.Schema
}

func (InboxEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "inbox_events"},
		entsql.WithComments(true),
	}
}

func (InboxEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("event id from message body").NotEmpty(),
		field.Int32("event_type").Comment("common enums.EventType value"),
		field.String("subject").Comment("received NATS subject").NotEmpty(),
		field.String("producer_service").Comment("producer service name").Optional().Nillable(),
		field.String("consumer_service").Comment("consumer service name").NotEmpty(),
		field.String("consumer_group").Comment("consumer group name").NotEmpty(),
		field.String("payload_hash").Comment("payload hash for duplicate diagnostics").Optional().Nillable(),
		field.Enum("status").Values(commonevent.InboxStatusValues()...).Default(string(commonevent.InboxStatusReceived)).Comment("processing status"),
		field.Int32("retry_count").Comment("processing retry count").Default(0),
		field.Int32("max_retry").Comment("maximum processing retries").Default(10),
		field.String("locked_by").Comment("consumer lock owner").Optional().Nillable(),
		field.Time("locked_until").Comment("consumer lock expiration").Optional().Nillable(),
		field.Time("received_at").Comment("received time").Default(time.Now),
		field.Time("processed_at").Comment("processed time").Optional().Nillable(),
		field.Text("last_error").Comment("last processing error").Optional().Nillable(),
	}
}

func (InboxEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (InboxEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("consumer_group", "event_id").Unique(),
		index.Fields("status", "id"),
		index.Fields("event_type", "status"),
	}
}
