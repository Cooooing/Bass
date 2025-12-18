package schema

import (
	"common/pkg"
	"common/pkg/constant"

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
		field.String("email").Comment("邮箱").Optional(),
		field.String("phone").Comment("手机号").Optional().Nillable(),
		field.String("url").Comment("用户个人主页链接").Optional().Nillable(),
		field.String("avatar_url").Comment("头像URL").Optional().Nillable(),
		field.String("introduction").Comment("个人简介").Optional().Nillable(),
		field.String("mbti").Comment("用户 MBTI 类型").Optional().Nillable(),

		// --- 状态 ---
		field.Int32("status").Comment("用户状态：0-正常，1-封禁，2-注销").Default(0).Nillable(),
		field.String("group_name").Comment("用户组名称").Optional(),

		// --- 社交信息 ---
		field.Int32("follow_count").Comment("关注数").Default(0).Nillable(),
		field.Int32("follower_count").Comment("粉丝数").Default(0).Nillable(),
		field.Int32("block_count").Comment("屏蔽数").Default(0).Nillable(),
		field.Int32("blocked_count").Comment("被屏蔽数").Default(0).Nillable(),

		// --- 登录信息 ---
		field.Time("last_login_time").Comment("最近登录时间").Optional().Nillable(),
		field.String("last_login_ip").Comment("最近登录IP").Optional().Nillable(),

		// --- 行为统计 ---
		field.Int32("online_minutes").Comment("在线总时长（分钟）").Default(0).Nillable(),
		field.Time("last_checkin_time").Comment("最近签到时间").Optional().Nillable(),
		field.Int32("current_checkin_streak").Comment("当前连续签到天数").Default(0).Nillable(),
		field.Int32("longest_checkin_streak").Comment("最长连续签到天数").Default(0).Nillable(),

		// --- 用户偏好设置 ---
		field.String("language").Comment("用户语言").Default("zh-CN").Nillable(),
		field.String("timezone").Comment("时区").Default("Asia/Shanghai").Nillable(),
		field.String("theme").Comment("皮肤主题").Default("default").Nillable(),
		field.String("mobile_theme").Comment("移动端皮肤主题").Default("default").Nillable(),
		field.Bool("enable_web_notify").Comment("启用Web通知").Default(true).Nillable(),
		field.Bool("enable_email_subscribe").Comment("启用邮件订阅").Default(true).Nillable(),

		// --- 隐私设置 ---
		field.Bool("public_points").Comment("是否公开积分榜").Default(true).Nillable(),
		field.Bool("public_followers").Comment("是否公开粉丝列表").Default(true).Nillable(),
		field.Bool("public_articles").Comment("是否公开帖子列表").Default(true).Nillable(),
		field.Bool("public_comments").Comment("是否公开评论列表").Default(true).Nillable(),
		field.Bool("public_online_status").Comment("是否公开在线状态").Default(true).Nillable(),

		// --- 地理信息 ---
		field.String("country").Comment("所在国家").Optional().Nillable(),
		field.String("province").Comment("所在省份").Optional().Nillable(),
		field.String("city").Comment("所在城市").Optional().Nillable(),
		field.Bool("public_location").Comment("是否公开地理位置").Default(true).Nillable(),

		// --- 其他 ---
		field.Bool("twofa_enable").Comment("是否开启二步验证").Default(false),
		field.Time("twofa_enable_time").Comment("二步验证启用时间").Optional().Nillable(),
		field.String("twofa_secret").Comment("二步验证Secret").Default(""),
	}
	fields = append(fields, pkg.TimeAuditFields()...)
	return fields
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
	}
}
