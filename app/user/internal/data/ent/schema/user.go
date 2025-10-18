package schema

import (
	"common/pkg"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User 实体定义
type User struct {
	ent.Schema
}

// Fields 定义表字段
func (User) Fields() []ent.Field {
	fields := []ent.Field{
		// --- 基础信息 ---
		field.String("name").Comment("用户名").NotEmpty(),
		field.String("nickname").Comment("昵称").Optional(),
		field.String("password").Comment("密码").NotEmpty(),
		field.String("email").Comment("邮箱").Optional(),
		field.String("phone").Comment("手机号").Optional(),
		field.String("url").Comment("用户个人主页链接").Optional(),
		field.String("avatar_url").Comment("头像URL").Optional(),
		field.String("introduction").Comment("个人简介").Optional(),
		field.String("mbti").Comment("用户 MBTI 类型").Optional(),

		// --- 状态 ---
		field.Int("status").Comment("用户状态：0-正常，1-封禁，2-注销").Default(0),
		field.String("role").Comment("用户角色").Default("user"),

		// --- 社交信息 ---
		field.Int("follow_count").Comment("关注数").Default(0),
		field.Int("follower_count").Comment("粉丝数").Default(0),

		// --- 登录信息 ---
		field.Time("last_login_time").Comment("最近登录时间").Optional().Nillable(),
		field.String("last_login_ip").Comment("最近登录IP").Optional(),

		// --- 行为统计 ---
		field.Int("online_minutes").Comment("在线总时长（分钟）").Default(0),
		field.Time("last_checkin_time").Comment("最近签到时间").Optional().Nillable(),
		field.Int("current_checkin_streak").Comment("当前连续签到天数").Default(0),
		field.Int("longest_checkin_streak").Comment("最长连续签到天数").Default(0),

		// --- 用户偏好设置 ---
		field.String("language").Comment("用户语言").Default("zh-CN"),
		field.String("timezone").Comment("时区").Default("Asia/Shanghai"),
		field.String("theme").Comment("皮肤主题").Default("default"),
		field.String("mobile_theme").Comment("移动端皮肤主题").Default("default"),
		field.Bool("enable_web_notify").Comment("启用Web通知").Default(true),
		field.Bool("enable_email_subscribe").Comment("启用邮件订阅").Default(true),

		// --- 隐私设置 ---
		field.Bool("public_points").Comment("是否公开积分榜").Default(true),
		field.Bool("public_followers").Comment("是否公开粉丝列表").Default(true),
		field.Bool("public_articles").Comment("是否公开帖子列表").Default(true),
		field.Bool("public_comments").Comment("是否公开评论列表").Default(true),
		field.Bool("public_online_status").Comment("是否公开在线状态").Default(true),

		// --- 地理信息 ---
		field.String("country").Comment("所在国家").Optional(),
		field.String("province").Comment("所在省份").Optional(),
		field.String("city").Comment("所在城市").Optional(),
		field.Bool("public_location").Comment("是否公开地理位置").Default(true),

		// --- 其他 ---
		field.String("twofa_secret").Comment("二步验证Secret").Optional(),
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
