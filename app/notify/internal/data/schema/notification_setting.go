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

type NotificationSetting struct {
	ent.Schema
}

func (NotificationSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_setting"},
	}
}

func (NotificationSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户ID"),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Enum("channel").Values(notifyenum.NotificationChannelMap.EnumValues()...).Comment("通知渠道"),
		field.Bool("enable").Comment("是否启用").Default(true),
	}
}

func (NotificationSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationSetting) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "event_type", "channel").Unique(),
	}
}
