package model

import (
	"time"

	"bff_bbs/internal/enum"
)

// HolidayInfo 是 BFF 对客户端暴露的节假日状态。
type HolidayInfo struct {
	Date          time.Time
	Type          enum.HolidayType
	DayName       string
	HolidayName   string
	DaysToHoliday int32
	DaysRemaining int32
}
