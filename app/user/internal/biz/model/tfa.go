package model

import "time"

type TFA struct {
	// ID 是双因素认证记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// Enable 表示是否启用双因素认证。
	Enable bool
	// EnableTime 是启用双因素认证的时间。
	EnableTime *time.Time
	// Secret 是双因素认证密钥。
	Secret string
}
