package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	utilent "common/pkg/util/ent"
	notifyenum "notify/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// NotificationRule 保存事件类型、通道和语言维度的通知启用规则。
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
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("事件类型"),
		field.Enum("channel").Values(notifyenum.NotificationChannelMap.EnumValues()...).Comment("通知通道"),
		field.Enum("language").Values(notifyenum.LanguageMap.EnumValues()...).Default(string(notifyenum.LanguageZhCN)).Comment("语言"),
		field.Bool("enabled").Comment("是否启用").Default(false),
	}
}

func (NotificationRule) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (NotificationRule) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("event_type", "channel", "language").Unique(),
		index.Fields("event_type", "language", "enabled"),
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
