package model

import "user/internal/enum"

type Preferences struct {
	// ID 是偏好设置记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// Language 是首选语言。
	Language *enum.Language
	// Timezone 是首选时区。
	Timezone *string
	// Theme 是桌面端主题。
	Theme *string
	// MobileTheme 是移动端主题。
	MobileTheme *string
}
