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

// HolidayEntry 是外部数据源给出的某一天配置，只有节日和调休日会出现。
type HolidayEntry struct {
	Date    time.Time
	Name    string
	Holiday bool
}

// Format 返回配置的 yyyy-MM-dd 日期键。
func (e *HolidayEntry) Format() string {
	return e.Date.Format("2006-01-02")
}

// HolidayConfigCache 是按年缓存的节假日配置快照。
type HolidayConfigCache struct {
	Entries   map[string]*HolidayEntry
	FetchedAt time.Time
}
