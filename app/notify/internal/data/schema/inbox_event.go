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
		field.Int32("event_type").Comment("事件类型，对应 common.enums.EventType 数值"),
		field.String("subject").Comment("收到消息的 NATS 主题").NotEmpty(),
		field.String("producer_service").Comment("生产者服务名").Optional().Nillable(),
		field.String("consumer_service").Comment("消费者服务名").NotEmpty(),
		field.String("consumer_group").Comment("消费者分组名").NotEmpty(),
		field.String("payload_hash").Comment("消息体哈希，用于排查重复消息内容是否一致").Optional().Nillable(),
		field.Enum("status").Values(commonenum.InboxEventStatusMap.EnumValues()...).Default(string(commonenum.InboxEventStatusReceived)).Comment("处理状态"),
		field.Int32("retry_count").Comment("处理重试次数").Default(0),
		field.Int32("max_retry").Comment("最大处理重试次数").Default(10),
		field.String("locked_by").Comment("消费者锁持有者").Optional().Nillable(),
		field.Time("locked_until").Comment("消费者锁过期时间").Optional().Nillable(),
		field.Time("received_at").Comment("接收时间").Default(time.Now),
		field.Time("processed_at").Comment("处理完成时间").Optional().Nillable(),
		field.Text("last_error").Comment("最近一次处理错误").Optional().Nillable(),
	}
}

func (InboxEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (InboxEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("consumer_group", "event_id").
			Unique().
			StorageKey("notify_inbox_events_consumer_group_event_id"),
		index.Fields("status", "id").
			StorageKey("notify_inbox_events_status_id"),
		index.Fields("event_type", "status").
			StorageKey("notify_inbox_events_event_type_status"),
	}
}
