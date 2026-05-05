package schema

import (
	"common/pkg/constant"
	"common/pkg/util"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationTemplate 通知记录表
type NotificationTemplate struct {
	ent.Schema
}

func (NotificationTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixMsgCenter.String() + "notification_template"},
	}
}

// Fields 定义表字段
func (NotificationTemplate) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("notification_type").Comment("通知类型"),
		field.String("channel").Comment("通知渠道"),
		field.String("title").Comment("标题").Default(""),
		field.String("content").Comment("模板内容"),
		field.Bool("enable").Comment("是否启用").Default(false),
	}
	fields = append(fields, util.TimeAuditFields()...)
	return fields
}

func (NotificationTemplate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("notification_type", "channel", "enable"),
	}
}

func (NotificationTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}
