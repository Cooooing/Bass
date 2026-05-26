package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// InboxEvent 通知服务 inbox 事件表，用于记录已接收事件并保证消费幂等。
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
		field.String("event_id").Comment("消息体中的事件幂等 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Enum("subject").NamedValues(commonenum.EventSubjectMap.EnumValues()...).Comment("收到消息的 NATS 主题"),
		field.Bytes("payload").Comment("原始事件消息体"),
		field.Enum("status").Values(commonenum.InboxEventStatusMap.EnumValues()...).Default(string(commonenum.InboxEventStatusReceived)).Comment("处理状态"),
		field.Int32("retry_count").Comment("处理重试次数").Default(0),
		field.Time("received_at").Comment("接收时间").Default(time.Now),
		field.Time("processed_at").Comment("处理完成时间").Optional().Nillable(),
	}
}

func (InboxEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (InboxEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_id").
			Unique().
			StorageKey("notify_inbox_events_event_id"),
		index.Fields("status", "id").
			StorageKey("notify_inbox_events_status_id"),
		index.Fields("event_type", "status").
			StorageKey("notify_inbox_events_event_type_status"),
	}
}
