package model

import (
	"time"
	"user/internal/enum"
)

type Account struct {
	// ID 是账号 ID。
	ID int64
	// Name 是唯一登录账号名。
	Name string
	// Nickname 是可选展示昵称。
	Nickname *string
	// Password 是哈希后的密码。
	Password string
	// Email 是可选绑定邮箱。
	Email *string
	// Phone 是可选绑定手机号。
	Phone *string
	// URL 是可选个人主页地址。
	URL *string
	// AvatarURL 是可选头像地址。
	AvatarURL *string
	// Introduction 是可选个人简介。
	Introduction *string
	// Mbti 是可选 MBTI 类型。
	Mbti *string
	// Status 是账号状态。
	Status *enum.AccountStatus
	// GroupName 是账号组名。
	GroupName string
	// FollowCount 是当前账号关注的账号数量。
	FollowCount *int32
	// FollowerCount 是粉丝数量。
	FollowerCount *int32
	// BlockCount 是当前账号拉黑的账号数量。
	BlockCount *int32
	// BlockedCount 是拉黑当前账号的账号数量。
	BlockedCount *int32
	// CreatedAt 是记录创建时间。
	CreatedAt *time.Time
	// UpdatedAt 是记录更新时间。
	UpdatedAt *time.Time
}
