package model

import "time"

type User struct {
	// 用户ID
	ID int64
	// 用户名
	Name string
	// 昵称
	Nickname *string
	// 密码
	Password string
	// 邮箱
	Email *string
	// 手机号
	Phone *string
	// 用户个人主页链接
	URL *string
	// 头像URL
	AvatarURL *string
	// 个人简介
	Introduction *string
	// 用户 MBTI 类型
	Mbti *string
	// 用户状态：1-正常，2-封禁，3-注销
	Status *int32
	// 用户组名称
	GroupName string
	// 关注数
	FollowCount *int32
	// 粉丝数
	FollowerCount *int32
	// 屏蔽数
	BlockCount *int32
	// 被屏蔽数
	BlockedCount *int32
	// 最近登录时间
	LastLoginTime *time.Time
	// 最近登录IP
	LastLoginIP *string
	// 创建时间
	CreatedAt *time.Time
	// 更新时间
	UpdatedAt *time.Time
}
