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

// Account 定义账号核心表，登录历史和可选扩展资料拆分到独立业务表。
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "accounts"},
		entsql.WithComments(true),
	}
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("账号名").NotEmpty(),
		field.String("nickname").Comment("展示昵称").Optional().Nillable(),
		field.String("password").Comment("密码哈希").NotEmpty().Sensitive(),
		field.String("email").Comment("邮箱地址").Optional().Nillable(),
		field.String("phone").Comment("手机号").Optional().Nillable(),
		field.String("url").Comment("个人主页 URL").Optional().Nillable(),
		field.String("avatar_url").Comment("头像 URL").Optional().Nillable(),
		field.String("introduction").Comment("个人简介").Optional().Nillable(),
		field.Enum("mbti").Values(userenum.MBTIMap.EnumValues()...).Comment("MBTI 类型").Optional().Nillable(),
		field.Enum("status").Values(userenum.AccountStatusMap.EnumValues()...).Default(string(userenum.AccountStatusNormal)).Nillable().Comment("账号状态"),
		field.Int32("follow_count").Comment("关注数").Default(0).Nillable(),
		field.Int32("follower_count").Comment("粉丝数").Default(0).Nillable(),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("email").Unique(),
		index.Fields("phone").Unique(),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("relations_as_actor", Relation.Type),
		edge.To("relations_as_target", Relation.Type),
		edge.To("preferences", Preferences.Type),
		edge.To("privacy_settings", PrivacySetting.Type),
		edge.To("location", Location.Type),
		edge.To("tfa", TFA.Type),
		edge.To("checkin_records", CheckinRecord.Type),
		edge.To("checkin_stat", CheckinStat.Type),
		edge.To("login_logs", LoginLog.Type),
	}
}
