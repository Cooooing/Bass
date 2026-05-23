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

// OutboxEvent 内容服务 outbox 事件表，用于在本地事务内持久化待投递事件。
type OutboxEvent struct {
	ent.Schema
}

func (OutboxEvent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixContent.String() + "outbox_events"},
		entsql.WithComments(true),
	}
}

func (OutboxEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("事件幂等 ID，对应 common.enums.Event.event_id").NotEmpty().Unique(),
		field.Int32("event_type").Comment("事件类型，对应 common.enums.EventType 数值"),
		field.String("subject").Comment("NATS 主题").NotEmpty(),
		field.String("aggregate_type").Comment("业务聚合类型").Optional().Nillable(),
		field.String("aggregate_id").Comment("业务聚合 ID").Optional().Nillable(),
		field.String("producer_service").Comment("生产者服务名").NotEmpty(),
		field.Bytes("payload").Comment("protobuf 编码后的 common.enums.Event"),
		field.JSON("headers", map[string]string{}).Comment("消息头").Default(map[string]string{}),
		field.Enum("status").Values(commonenum.OutboxEventStatusMap.EnumValues()...).Default(string(commonenum.OutboxEventStatusPending)).Comment("投递状态"),
		field.Int32("retry_count").Comment("投递重试次数").Default(0),
		field.Int32("max_retry").Comment("最大投递重试次数").Default(10),
		field.Time("next_retry_at").Comment("下次重试时间").Optional().Nillable(),
		field.String("locked_by").Comment("分发器锁持有者").Optional().Nillable(),
		field.Time("locked_until").Comment("分发器锁过期时间").Optional().Nillable(),
		field.Time("published_at").Comment("发布时间").Optional().Nillable(),
		field.Text("last_error").Comment("最近一次投递错误").Optional().Nillable(),
	}
}

func (OutboxEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (OutboxEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "next_retry_at", "id").
			StorageKey("content_outbox_events_status_next_retry_at_id"),
		index.Fields("aggregate_type", "aggregate_id").
			StorageKey("content_outbox_events_aggregate_type_aggregate_id"),
	}
}
