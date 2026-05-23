package model

import (
	"time"
)

type User struct {
	// ID 是实体 ID。
	ID int64 `json:"id,omitempty"`
	// Name 是用户名。
	Name string `json:"name,omitempty"`
	// Nickname 是昵称。
	Nickname string `json:"nickname,omitempty"`
	// Password 是密码哈希。
	Password string `json:"-"`
	// Email 是邮箱。
	Email string `json:"email,omitempty"`
	// Phone 是手机号。
	Phone string `json:"phone,omitempty"`
	// URL 是用户个人主页链接。
	URL string `json:"url,omitempty"`
	// AvatarURL 是头像 URL。
	AvatarURL string `json:"avatar_url,omitempty"`
	// Introduction 是个人简介。
	Introduction string `json:"introduction,omitempty"`
	// Mbti 是用户 MBTI 类型。
	Mbti string `json:"mbti,omitempty"`
	// Status 是用户状态，0 表示正常，1 表示封禁，2 表示注销。
	Status int32 `json:"status,omitempty"`
	// Role 是用户角色。
	Role string `json:"role,omitempty"`
	// FollowCount 是关注数。
	FollowCount int32 `json:"follow_count,omitempty"`
	// FollowerCount 是粉丝数。
	FollowerCount int32 `json:"follower_count,omitempty"`
	// LastLoginTime 是最近登录时间。
	LastLoginTime *time.Time `json:"last_login_time,omitempty"`
	// LastLoginIP 是最近登录 IP。
	LastLoginIP string `json:"last_login_ip,omitempty"`
	// OnlineMinutes 是在线总时长，单位为分钟。
	OnlineMinutes int32 `json:"online_minutes,omitempty"`
	// LastCheckinTime 是最近签到时间。
	LastCheckinTime *time.Time `json:"last_checkin_time,omitempty"`
	// CurrentCheckinStreak 是当前连续签到天数。
	CurrentCheckinStreak int32 `json:"current_checkin_streak,omitempty"`
	// LongestCheckinStreak 是最长连续签到天数。
	LongestCheckinStreak int32 `json:"longest_checkin_streak,omitempty"`
	// Language 是用户语言。
	Language string `json:"language,omitempty"`
	// Timezone 是时区。
	Timezone string `json:"timezone,omitempty"`
	// Theme 是桌面端主题。
	Theme string `json:"theme,omitempty"`
	// MobileTheme 是移动端主题。
	MobileTheme string `json:"mobile_theme,omitempty"`
	// EnableWebNotify 表示是否启用 Web 通知。
	EnableWebNotify bool `json:"enable_web_notify,omitempty"`
	// EnableEmailSubscribe 表示是否启用邮件订阅。
	EnableEmailSubscribe bool `json:"enable_email_subscribe,omitempty"`
	// PublicPoints 表示是否公开积分概况。
	PublicPoints bool `json:"public_points,omitempty"`
	// PublicFollowers 表示是否公开粉丝列表。
	PublicFollowers bool `json:"public_followers,omitempty"`
	// PublicArticles 表示是否公开文章列表。
	PublicArticles bool `json:"public_articles,omitempty"`
	// PublicComments 表示是否公开评论列表。
	PublicComments bool `json:"public_comments,omitempty"`
	// PublicOnlineStatus 表示是否公开在线状态。
	PublicOnlineStatus bool `json:"public_online_status,omitempty"`
	// Country 是所在国家。
	Country string `json:"country,omitempty"`
	// Province 是所在省份。
	Province string `json:"province,omitempty"`
	// City 是所在城市。
	City string `json:"city,omitempty"`
	// PublicLocation 表示是否公开地理位置。
	PublicLocation bool `json:"public_location,omitempty"`
	// TwofaEnable 表示是否启用双因素认证。
	TwofaEnable bool `json:"twofa_enable,omitempty"`
	// TwofaEnableTime 是双因素认证启用时间。
	TwofaEnableTime *time.Time `json:"twofa_enable_time,omitempty"`
	// TwofaSecret 是双因素认证密钥。
	TwofaSecret string `json:"-"`
	// CreatedAt 是创建时间。
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// UpdatedAt 是更新时间。
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
