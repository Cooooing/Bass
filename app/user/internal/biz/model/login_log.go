package model

import (
	"time"
	"user/internal/enum"
)

type LoginLog struct {
	ID          int64
	UserID      *int64
	LoginMethod enum.LoginMethod
	Status      enum.LoginStatus
	IP          *string
	Country     *string
	CountryCode *string
	Province    *string
	City        *string
	ISP         *string
	UserAgent   *string
	DeviceID    *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
