package model

type UserPreferences struct {
	// 用户偏好设置ID
	ID int64
	// 用户ID
	UserID int64
	// 用户语言
	Language *string
	// 时区
	Timezone *string
	// 桌面皮肤主题
	Theme *string
	// 移动端皮肤主题
	MobileTheme *string
	// 启用Web通知
	EnableWebNotify *bool
	// 启用邮件订阅
	EnableEmailSubscribe *bool
}
