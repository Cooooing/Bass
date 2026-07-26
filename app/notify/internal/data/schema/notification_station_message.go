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

// NotificationStationMessage 保存站内信消息。
type NotificationStationMessage struct {
	ent.Schema
}

func (NotificationStationMessage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_station_message",
		},
		entsql.WithComments(true),
	}
}

func (NotificationStationMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("来源事件 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("来源事件类型"),
		field.Int64("receiver_id").Comment("接收者用户 ID"),
		field.String("title").Comment("标题"),
		field.String("content").Comment("内容"),
		field.Enum("status").Values(notifyenum.NotificationChannelStatusMap.EnumValues()...).Default(string(notifyenum.NotificationChannelStatusSucceeded)).Comment("通道状态"),
		field.Time("read_at").Comment("已读时间").Optional().Nillable(),
	}
}

func (NotificationStationMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationStationMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_id", "receiver_id").Unique(),
		index.Fields("receiver_id", "created_at"),
		index.Fields("receiver_id", "read_at"),
	}
}
