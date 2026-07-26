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

// NotificationLarkWebhookTemplate 保存 Lark webhook 模板字段。
type NotificationLarkWebhookTemplate struct {
	ent.Schema
}

func (NotificationLarkWebhookTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_lark_webhook_template",
		},
		entsql.WithComments(true),
	}
}

func (NotificationLarkWebhookTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("rule_id").Comment("通知规则 ID"),
		field.String("webhook_id").Comment("Lark webhook 配置标识"),
		field.String("token").Sensitive().Comment("Lark webhook token"),
		field.String("secret").Optional().Nillable().Sensitive().Comment("Lark webhook 签名密钥"),
		field.String("msg_type").Default("text").Comment("Lark 消息类型"),
		field.String("content_template").Comment("Lark content 对象模板"),
	}
}

func (NotificationLarkWebhookTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationLarkWebhookTemplate) Indexes() []ent.Index {
	return []ent.Index{index.Fields("rule_id").Unique()}
}

func (NotificationLarkWebhookTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rule", NotificationRule.Type).Ref("lark_webhook_template").Required().Unique().Field("rule_id"),
	}
}
