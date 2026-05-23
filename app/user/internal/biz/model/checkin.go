package model

import "time"

type CheckinStat struct {
	// ID 是签到统计记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// TotalOnlineMinutes 是累计在线时长。
	TotalOnlineMinutes *int32
	// CurrentStreak 是当前连续签到天数。
	CurrentStreak *int32
	// LongestStreak 是最长连续签到天数。
	LongestStreak *int32
}

type CheckinRecord struct {
	// ID 是签到记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// Date 是签到日期。
	Date *time.Time
	// OnlineMinutes 是当天在线时长。
	OnlineMinutes *int32
	// Activity 是当天活跃度。
	Activity *int32
	// Checked 表示是否完成签到目标。
	Checked bool
}
