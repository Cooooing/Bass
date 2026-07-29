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

// NotificationEmailDelivery 保存邮件投递结果。
type NotificationEmailDelivery struct {
	ent.Schema
}

func (NotificationEmailDelivery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_email_delivery",
		},
		entsql.WithComments(true),
	}
}

func (NotificationEmailDelivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("来源事件 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("来源事件类型"),
		field.Int64("receiver_id").Comment("接收者用户 ID").Optional().Nillable(),
		field.String("to_email").Comment("标准化收件邮箱").NotEmpty(),
		field.String("subject").Comment("邮件主题"),
		field.String("body").Comment("邮件正文"),
		field.String("content_type").Comment("正文类型"),
		field.Enum("status").Values(notifyenum.NotificationChannelStatusMap.EnumValues()...).Default(notifyenum.NotificationChannelStatusProcessing.String()).Comment("通道状态"),
		field.Int32("attempt_count").Comment("发送尝试次数").Default(0),
		field.Time("last_attempt_at").Comment("最近一次发送尝试时间").Optional().Nillable(),
		field.String("provider_message_id").Comment("服务商消息 ID").Optional().Nillable(),
		field.String("provider_resp").Comment("服务商响应摘要").Optional().Nillable(),
		field.Time("sent_at").Comment("发送成功时间").Optional().Nillable(),
	}
}

func (NotificationEmailDelivery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (NotificationEmailDelivery) Indexes() []ent.Index {
	return []ent.Index{index.Fields("event_id", "to_email").Unique()}
}
