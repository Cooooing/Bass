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
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_inbox_event"},
		entsql.WithComments(true),
	}
}

func (InboxEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("消息体中的事件幂等 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Enum("subject").GoType(commonenum.EventSubject("")).Comment("收到消息的 NATS 主题"),
		field.Text("payload").Comment("原始事件 JSON 消息体"),
		field.Enum("status").Values(commonenum.InboxEventStatusMap.EnumValues()...).Default(string(commonenum.InboxEventStatusProcessing)).Comment("处理状态"),
		field.Int32("attempt_count").Comment("处理尝试次数").Default(0),
		field.String("last_error").Comment("最近一次失败原因摘要").Optional().Nillable(),
		field.Time("processing_started_at").Comment("最近一次开始处理时间").Default(time.Now),
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
			StorageKey("notify_notification_inbox_event_event_id"),
	}
}
