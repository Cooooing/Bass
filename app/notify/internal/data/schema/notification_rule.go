package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	notifyenum "notify/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationRule 保存通知类型、通道和语言维度的启用规则。
type NotificationRule struct {
	ent.Schema
}

func (NotificationRule) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_rule",
		},
		entsql.WithComments(true),
	}
}

func (NotificationRule) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Enum("notification_type").Values(notifyenum.NotificationTypeMap.EnumValues()...).Comment("通知类型"),
		field.Enum("channel").Values(notifyenum.NotificationChannelMap.EnumValues()...).Comment("通知通道"),
		field.Enum("language").Values(notifyenum.LanguageMap.EnumValues()...).Default(notifyenum.LanguageZhCN.String()).Comment("语言"),
		field.Bool("enabled").Comment("是否启用").Default(false),
	}
}

func (NotificationRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (NotificationRule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("notification_type", "channel", "language").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("notification_type", "language", "enabled").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}

func (NotificationRule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("station_template", NotificationStationTemplate.Type).Unique(),
		edge.To("email_template", NotificationEmailTemplate.Type).Unique(),
		edge.To("tencent_sms_template", NotificationTencentSMSTemplate.Type).Unique(),
		edge.To("lark_webhook_template", NotificationLarkWebhookTemplate.Type).Unique(),
	}
}
