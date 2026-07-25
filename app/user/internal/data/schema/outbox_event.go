package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OutboxEvent 存储本地事务提交后需要投递的事件。
type OutboxEvent struct {
	ent.Schema
}

func (OutboxEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "outbox_events",
		},
		entsql.WithComments(true),
	}
}

func (OutboxEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("全局幂等 ID，对应 common.enums.Event.event_id").NotEmpty().Unique(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Enum("subject").GoType(commonenum.EventSubject("")).Values(commonenum.EventSubjectMap.EnumValues()...).Comment("NATS 主题"),
		field.Text("payload").Comment("JSON 编码后的 common.enums.Event"),
		field.JSON("headers", map[string]string{}).Comment("消息头").Default(map[string]string{}),
		field.Enum("status").Values(commonenum.OutboxEventStatusMap.EnumValues()...).Default(string(commonenum.OutboxEventStatusPending)).Comment("投递状态"),
		field.Int32("retry_count").Comment("投递重试次数").Default(0),
		field.String("last_error").Comment("最近一次投递失败原因摘要").Optional().Nillable(),
		field.Time("published_at").Comment("发布时间").Optional().Nillable(),
	}
}

func (OutboxEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (OutboxEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "id").
			StorageKey("user_outbox_events_status_id"),
	}
}
