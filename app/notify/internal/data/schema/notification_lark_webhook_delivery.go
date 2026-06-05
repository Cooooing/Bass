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

// NotificationLarkWebhookDelivery 保存 Lark webhook 投递结果。
type NotificationLarkWebhookDelivery struct {
	ent.Schema
}

func (NotificationLarkWebhookDelivery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixNotify.String() + "notification_lark_webhook_delivery"},
		entsql.WithComments(true),
	}
}

func (NotificationLarkWebhookDelivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("来源事件 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("来源事件类型"),
		field.String("webhook_id").Comment("Lark webhook 配置标识").NotEmpty(),
		field.String("request_body").Comment("请求体"),
		field.Enum("status").Values(notifyenum.NotificationChannelStatusMap.EnumValues()...).Default(string(notifyenum.NotificationChannelStatusProcessing)).Comment("通道状态"),
		field.Int32("attempt_count").Comment("发送尝试次数").Default(0),
		field.Time("last_attempt_at").Comment("最近一次发送尝试时间").Optional().Nillable(),
		field.Int("http_status").Comment("HTTP 状态码").Optional().Nillable(),
		field.String("response_body").Comment("Lark 响应体").Optional().Nillable(),
		field.Time("sent_at").Comment("发送成功时间").Optional().Nillable(),
	}
}

func (NotificationLarkWebhookDelivery) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (NotificationLarkWebhookDelivery) Indexes() []ent.Index {
	return []ent.Index{index.Fields("event_id", "webhook_id").Unique()}
}
