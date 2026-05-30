package model

import "time"

type Totp struct {
	// ID 是 TOTP 认证记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// Enable 表示是否启用 TOTP 认证。
	Enable bool
	// EnableTime 是启用 TOTP 认证的时间。
	EnableTime *time.Time
	// Secret 是 TOTP 认证密钥。
	Secret string
}
