package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationEmailTemplate 保存邮件模板字段。
type NotificationEmailTemplate struct {
	ent.Schema
}

func (NotificationEmailTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_email_template",
		},
		entsql.WithComments(true),
	}
}

func (NotificationEmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("rule_id").Comment("通知规则 ID"),
		field.String("subject_template").Comment("主题模板"),
		field.String("body_template").Comment("正文模板"),
		field.String("content_type").Comment("正文类型"),
	}
}

func (NotificationEmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationEmailTemplate) Indexes() []ent.Index {
	return []ent.Index{index.Fields("rule_id").Unique()}
}

func (NotificationEmailTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rule", NotificationRule.Type).Ref("email_template").Required().Unique().Field("rule_id"),
	}
}
