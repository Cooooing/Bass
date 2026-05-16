package model

import "time"

type UserTFA struct {
	// 用户二步验证ID
	ID int64
	// 用户ID
	UserID int64
	// 是否开启二步验证
	Enable bool
	// 二步验证启用时间
	EnableTime *time.Time
	// 二步验证Secret
	Secret string
}
