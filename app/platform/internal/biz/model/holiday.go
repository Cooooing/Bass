package model

import (
	"time"

	"platform/internal/enum"
)

// HolidayInfo 是某一天的节假日状态快照。
type HolidayInfo struct {
	Date          time.Time
	Type          enum.HolidayType
	DayName       string
	HolidayName   string
	DaysToHoliday int32
	DaysRemaining int32
}
