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

// NotificationTencentSMSDelivery 保存腾讯云短信投递结果。
type NotificationTencentSMSDelivery struct {
	ent.Schema
}

func (NotificationTencentSMSDelivery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_tencent_sms_delivery",
		},
		entsql.WithComments(true),
	}
}

func (NotificationTencentSMSDelivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("event_id").Comment("来源事件 ID").NotEmpty(),
		field.Enum("event_type").Values(commonenum.EventTypeMap.EnumValues()...).Comment("来源事件类型"),
		field.Int64("receiver_id").Comment("接收者用户 ID").Optional().Nillable(),
		field.String("phone").Comment("E.164 格式手机号").NotEmpty(),
		field.String("sms_sdk_app_id").Comment("腾讯云短信应用 ID"),
		field.String("sign_name").Comment("短信签名"),
		field.String("provider_template_id").Comment("腾讯云模板 ID"),
		field.JSON("template_params", []string{}).Comment("有序模板参数"),
		field.Enum("status").Values(notifyenum.NotificationChannelStatusMap.EnumValues()...).Default(string(notifyenum.NotificationChannelStatusProcessing)).Comment("通道状态"),
		field.Int32("attempt_count").Comment("发送尝试次数").Default(0),
		field.Time("last_attempt_at").Comment("最近一次发送尝试时间").Optional().Nillable(),
		field.String("provider_request_id").Comment("腾讯云请求 ID").Optional().Nillable(),
		field.String("provider_code").Comment("腾讯云返回码").Optional().Nillable(),
		field.String("provider_message").Comment("腾讯云返回消息").Optional().Nillable(),
		field.Time("sent_at").Comment("发送成功时间").Optional().Nillable(),
	}
}

func (NotificationTencentSMSDelivery) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (NotificationTencentSMSDelivery) Indexes() []ent.Index {
	return []ent.Index{index.Fields("event_id", "phone").Unique()}
}
