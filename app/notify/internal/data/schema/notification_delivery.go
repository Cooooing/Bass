package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	notifyenum "notify/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationDelivery 是外部通知投递任务表。
type NotificationDelivery struct {
	ent.Schema
}

func (NotificationDelivery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_delivery"},
	}
}

func (NotificationDelivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("事件幂等 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Int64("receiver_id").Comment("接收者 ID").Optional().Nillable(),
		field.Enum("channel").Values(notifyenum.NotificationChannelMap.EnumValues()...).Comment("投递渠道"),
		field.String("target").Comment("投递目标").NotEmpty(),
		field.String("title").Comment("标题").Default(""),
		field.String("content").Comment("内容"),
		field.Enum("status").Values(notifyenum.NotificationDeliveryStatusMap.EnumValues()...).Default(string(notifyenum.NotificationDeliveryStatusPending)).Comment("投递状态"),
		field.Int32("retry_count").Comment("投递重试次数").Default(0),
		field.Time("sent_at").Comment("投递完成时间").Optional().Nillable(),
	}
}

func (NotificationDelivery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationDelivery) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_id", "channel", "target").Unique(),
		index.Fields("status", "id"),
		index.Fields("receiver_id", "status"),
	}
}
