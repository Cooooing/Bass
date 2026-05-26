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

// NotificationTemplate 通知模板表
type NotificationTemplate struct {
	ent.Schema
}

func (NotificationTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_template"},
	}
}

// Fields 定义表字段
func (NotificationTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Enum("channel").Values(notifyenum.NotificationChannelMap.EnumValues()...).Comment("通知渠道"),
		field.Enum("language").Values(notifyenum.LanguageMap.EnumValues()...).Default(string(notifyenum.LanguageZhCN)).Comment("语言"),
		field.String("title").Comment("标题").Default(""),
		field.String("content").Comment("模板内容"),
		field.Bool("enable").Comment("是否启用").Default(false),
	}
}

func (NotificationTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationTemplate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type", "channel", "language", "enable"),
		index.Fields("event_type", "channel", "language").Unique(),
	}
}

func (NotificationTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}
