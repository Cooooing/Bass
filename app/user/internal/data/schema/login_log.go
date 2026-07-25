package schema

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
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
		entsql.Annotation{
			Table: constant.TablePrefixUser.String() + "login_logs",
		},
		entsql.WithComments(true),
	}
}

func (LoginLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("账号 ID，登录失败且无法匹配账号时为空").Optional().Nillable(),
		field.String("account_input").Comment("登录输入账号原文").MaxLen(320).Default(""),
		field.Enum("login_type").Values(userenum.LoginTypeMap.EnumValues()...).Comment("登录方式"),
		field.Enum("realm").Values(commonenum.LoginRealmMap.EnumValues()...).Comment("登录域"),
		field.Enum("status").Values(userenum.LoginStatusMap.EnumValues()...).Comment("登录状态"),
		field.Enum("failure_reason").Values(userenum.LoginFailureReasonMap.EnumValues()...).Comment("失败原因").Optional().Nillable(),
		field.String("session_id").Comment("会话 ID").MaxLen(64).Default(""),
		field.String("ip").Comment("客户端 IP").Optional().Nillable(),
		field.String("country").Comment("国家").Optional().Nillable(),
		field.String("country_code").Comment("国家代码").Optional().Nillable(),
		field.String("province").Comment("省份").Optional().Nillable(),
		field.String("city").Comment("城市").Optional().Nillable(),
		field.String("isp").Comment("网络服务商").Optional().Nillable(),
		field.Text("user_agent").Comment("User-Agent").Optional().Nillable(),
		field.Enum("client_type").Values(userenum.ClientTypeMap.EnumValues()...).Comment("客户端类型").Optional().Nillable(),
		field.Enum("device_type").Values(userenum.DeviceTypeMap.EnumValues()...).Comment("设备类型").Optional().Nillable(),
		field.String("os_name").Comment("操作系统名称").MaxLen(64).Default(""),
		field.String("os_version").Comment("操作系统版本").MaxLen(64).Default(""),
		field.String("browser_name").Comment("浏览器名称").MaxLen(64).Default(""),
		field.String("browser_version").Comment("浏览器版本").MaxLen(64).Default(""),
		field.String("app_name").Comment("客户端应用名称").MaxLen(64).Default(""),
		field.String("app_version").Comment("客户端应用版本").MaxLen(64).Default(""),
	}
}

func (LoginLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (LoginLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "created_at").StorageKey("user_login_logs_user_created_idx"),
		index.Fields("status", "created_at").StorageKey("user_login_logs_status_created_idx"),
		index.Fields("ip", "created_at").StorageKey("user_login_logs_ip_created_idx"),
		index.Fields("realm", "created_at").StorageKey("user_login_logs_realm_created_idx"),
	}
}

func (LoginLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Account.Type).Ref("login_logs").Field("user_id").Unique(),
	}
}
