package model

import (
	"time"
)

type User struct {
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// 用户名
	Name string `json:"name,omitempty"`
	// 昵称
	Nickname string `json:"nickname,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 用户个人主页链接
	URL string `json:"url,omitempty"`
	// 头像URL
	AvatarURL string `json:"avatar_url,omitempty"`
	// 个人简介
	Introduction string `json:"introduction,omitempty"`
	// 用户 MBTI 类型
	Mbti string `json:"mbti,omitempty"`
	// 用户状态：0-正常，1-封禁，2-注销
	Status int32 `json:"status,omitempty"`
	// 用户角色
	Role string `json:"role,omitempty"`
	// 关注数
	FollowCount int32 `json:"follow_count,omitempty"`
	// 粉丝数
	FollowerCount int32 `json:"follower_count,omitempty"`
	// 最近登录时间
	LastLoginTime *time.Time `json:"last_login_time,omitempty"`
	// 最近登录IP
	LastLoginIP string `json:"last_login_ip,omitempty"`
	// 在线总时长（分钟）
	OnlineMinutes int32 `json:"online_minutes,omitempty"`
	// 最近签到时间
	LastCheckinTime *time.Time `json:"last_checkin_time,omitempty"`
	// 当前连续签到天数
	CurrentCheckinStreak int32 `json:"current_checkin_streak,omitempty"`
	// 最长连续签到天数
	LongestCheckinStreak int32 `json:"longest_checkin_streak,omitempty"`
	// 用户语言
	Language string `json:"language,omitempty"`
	// 时区
	Timezone string `json:"timezone,omitempty"`
	// 皮肤主题
	Theme string `json:"theme,omitempty"`
	// 移动端皮肤主题
	MobileTheme string `json:"mobile_theme,omitempty"`
	// 启用Web通知
	EnableWebNotify bool `json:"enable_web_notify,omitempty"`
	// 启用邮件订阅
	EnableEmailSubscribe bool `json:"enable_email_subscribe,omitempty"`
	// 是否公开积分榜
	PublicPoints bool `json:"public_points,omitempty"`
	// 是否公开粉丝列表
	PublicFollowers bool `json:"public_followers,omitempty"`
	// 是否公开帖子列表
	PublicArticles bool `json:"public_articles,omitempty"`
	// 是否公开评论列表
	PublicComments bool `json:"public_comments,omitempty"`
	// 是否公开在线状态
	PublicOnlineStatus bool `json:"public_online_status,omitempty"`
	// 所在国家
	Country string `json:"country,omitempty"`
	// 所在省份
	Province string `json:"province,omitempty"`
	// 所在城市
	City string `json:"city,omitempty"`
	// 是否公开地理位置
	PublicLocation bool `json:"public_location,omitempty"`
	// 二步验证Secret
	TwofaSecret string `json:"twofa_secret,omitempty"`
	// 创建时间
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
