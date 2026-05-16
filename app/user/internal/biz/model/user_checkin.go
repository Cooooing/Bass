package model

import "time"

type UserCheckinStat struct {
	// 用户签到聚合统计ID
	ID int64
	// 用户ID
	UserID int64
	// 累计在线总时长（分钟）
	TotalOnlineMinutes *int32
	// 当前连续签到天数
	CurrentStreak *int32
	// 最长连续签到天数
	LongestStreak *int32
}

type UserCheckinRecord struct {
	// 用户每日签到记录ID
	ID int64
	// 用户ID
	UserID int64
	// 签到日期
	Date *time.Time
	// 当日在线时长（分钟）
	OnlineMinutes *int32
	// 当日活跃度
	Activity *int32
	// 是否达标签到（由活跃度判断）
	Checked bool
}
