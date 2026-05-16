package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"common/pkg/enum"
	"common/pkg/util"

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
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("event_type").Comment("事件类型").
			GoType(enum.EventType("")),
		field.Enum("channel").Comment("通知渠道").
			GoType(enum.NotificationChannel("")),
		field.String("title").Comment("标题").Default(""),
		field.String("content").Comment("模板内容"),
		field.Bool("enable").Comment("是否启用").Default(false),
	}
	return fields
}

func (NotificationTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationTemplate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type", "channel", "enable"),
	}
}

func (NotificationTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}
