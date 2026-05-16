package model

type UserLocation struct {
	// 用户地理信息ID
	ID int64
	// 用户ID
	UserID int64
	// 所在国家
	Country *string
	// 所在省份
	Province *string
	// 所在城市
	City *string
}
