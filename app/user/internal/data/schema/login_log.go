package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	userenum "user/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// LoginLog 记录账号每次登录尝试，用于安全审计。
type LoginLog struct {
	ent.Schema
}

func (LoginLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "login_logs"},
		entsql.WithComments(true),
	}
}

func (LoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID，登录失败且无法匹配账号时为空").Optional().Nillable(),
		field.String("account").Comment("提交登录的账号").NotEmpty(),
		field.Enum("login_method").Values(userenum.LoginMethodMap.EnumValues()...).Default(string(userenum.LoginMethodPassword)).Comment("登录方式"),
		field.Enum("status").Values(userenum.LoginStatusMap.EnumValues()...).Comment("登录状态"),
		field.String("failure_reason").Comment("失败原因").Optional().Nillable(),
		field.String("ip").Comment("客户端 IP").Optional().Nillable(),
		field.String("country").Comment("国家").Optional().Nillable(),
		field.String("country_code").Comment("国家代码").Optional().Nillable(),
		field.String("province").Comment("省份").Optional().Nillable(),
		field.String("city").Comment("城市").Optional().Nillable(),
		field.String("isp").Comment("网络服务商").Optional().Nillable(),
		field.Text("user_agent").Comment("User-Agent").Optional().Nillable(),
		field.String("device_id").Comment("设备 ID").Optional().Nillable(),
		field.String("device_name").Comment("设备名称").Optional().Nillable(),
		field.String("platform").Comment("客户端平台").Optional().Nillable(),
		field.String("os").Comment("操作系统").Optional().Nillable(),
		field.String("browser").Comment("浏览器").Optional().Nillable(),
		field.String("request_id").Comment("请求 ID").Optional().Nillable(),
	}
}

func (LoginLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (LoginLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "created_at"),
		index.Fields("account", "created_at"),
		index.Fields("status", "created_at"),
		index.Fields("ip", "created_at"),
	}
}

func (LoginLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("login_logs").Field("user_id").Unique(),
	}
}
