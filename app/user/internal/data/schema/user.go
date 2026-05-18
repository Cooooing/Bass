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

// User 用户实体定义
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixUser.String() + "users"},
		entsql.WithComments(true),
	}
}

// Fields 定义表字段
func (User) Fields() []ent.Field {
	fields := []ent.Field{
		// --- 基础信息 ---
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("用户名").NotEmpty(),
		field.String("nickname").Comment("昵称").Optional().Nillable(),
		field.String("password").Comment("密码").NotEmpty().Sensitive(),
		field.String("email").Comment("邮箱").Optional().Nillable(),
		field.String("phone").Comment("手机号").Optional().Nillable(),
		field.String("url").Comment("用户个人主页链接").Optional().Nillable(),
		field.String("avatar_url").Comment("头像URL").Optional().Nillable(),
		field.String("introduction").Comment("个人简介").Optional().Nillable(),
		field.String("mbti").Comment("用户 MBTI 类型").Optional().Nillable(),

		// --- 状态 ---
		field.Enum("status").Values(userenum.UserStatusMap.EnumValues()...).Default(string(userenum.UserStatusNormal)).Nillable().Comment("用户状态"),
		field.String("group_name").Comment("用户组名称").Optional(),

		// --- 社交信息 ---
		field.Int32("follow_count").Comment("关注数").Default(0).Nillable(),
		field.Int32("follower_count").Comment("粉丝数").Default(0).Nillable(),
		field.Int32("block_count").Comment("屏蔽数").Default(0).Nillable(),
		field.Int32("blocked_count").Comment("被屏蔽数").Default(0).Nillable(),

		// --- 登录信息 ---
		field.Time("last_login_time").Comment("最近登录时间").Optional().Nillable(),
		field.String("last_login_ip").Comment("最近登录IP").Optional().Nillable(),
	}
	return fields
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique(),
		index.Fields("email").Unique(),
		index.Fields("phone").Unique(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// 关联用户关系发起者 一对多
		edge.To("relations_as_actor", UserRelation.Type),
		// 关联用户关系目标 一对多
		edge.To("relations_as_target", UserRelation.Type),
		// 用户偏好设置 一对一
		edge.To("preferences", UserPreferences.Type),
		// 用户隐私设置 一对一
		edge.To("privacy", UserPrivacy.Type),
		// 用户地理信息 一对一
		edge.To("location", UserLocation.Type),
		// 用户二步验证 一对一
		edge.To("tfa", UserTFA.Type),
		// 用户签到记录 一对多
		edge.To("checkin_records", UserCheckinRecord.Type),
		// 用户签到聚合统计 一对一
		edge.To("checkin_stat", UserCheckinStat.Type),
	}
}
