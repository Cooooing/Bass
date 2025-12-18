package schema

import (
	"common/pkg"
	"common/pkg/constant"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// NotificationTemplate 通知记录表
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
		field.Int32("notification_type").Comment("通知类型"),
		field.Int32("channel").Comment("通知渠道"),
		field.String("title").Comment("标题").Default(""), // 暂时仅用于邮件主题
		field.String("content").Comment("模板内容"),
		field.JSON("processors", []string{}).Comment("处理器链（有序执行）"),
		field.Bool("enable").Comment("是否启用").Default(false),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
}

func (NotificationTemplate) Indexes() []ent.Index {
	return []ent.Index{}
}

func (NotificationTemplate) Edges() []ent.Edge {
	return []ent.Edge{}
}
