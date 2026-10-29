package model

import (
	"time"

	"user/internal/enum"
)

type CheckinStat struct {
	// ID 是签到统计记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
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
	Date          *time.Time
	Status        enum.CheckinStatus
	CheckinPoints *int64
	StreakPoints  *int64
	CheckedAt     *time.Time
}
