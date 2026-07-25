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

// NotificationTencentSMSTemplate 保存腾讯云短信模板字段。
type NotificationTencentSMSTemplate struct {
	ent.Schema
}

func (NotificationTencentSMSTemplate) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixNotify.String() + "notification_tencent_sms_template",
		},
		entsql.WithComments(true),
	}
}

func (NotificationTencentSMSTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("rule_id").Comment("通知规则 ID"),
		field.String("sms_sdk_app_id").Comment("腾讯云短信应用 ID"),
		field.String("sign_name").Comment("短信签名"),
		field.String("provider_template_id").Comment("腾讯云模板 ID"),
		field.JSON("param_templates", []string{}).Comment("有序模板参数模板"),
	}
}

func (NotificationTencentSMSTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}}
}

func (NotificationTencentSMSTemplate) Indexes() []ent.Index {
	return []ent.Index{index.Fields("rule_id").Unique()}
}

func (NotificationTencentSMSTemplate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rule", NotificationRule.Type).Ref("tencent_sms_template").Required().Unique().Field("rule_id"),
	}
}
