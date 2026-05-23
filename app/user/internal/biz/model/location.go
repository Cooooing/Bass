package model

type Location struct {
	// ID 是位置记录 ID。
	ID int64
	// UserID 是归属账号 ID。
	UserID int64
	// Country 是可选国家名称。
	Country *string
	// Province 是可选省份名称。
	Province *string
	// City 是可选城市名称。
	City *string
}
