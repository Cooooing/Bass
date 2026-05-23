package model

import (
	"time"
	"user/internal/enum"
)

type LoginLog struct {
	ID            int64
	UserID        *int64
	Account       string
	LoginMethod   enum.LoginMethod
	Status        enum.LoginStatus
	FailureReason *string
	IP            *string
	Country       *string
	CountryCode   *string
	Province      *string
	City          *string
	ISP           *string
	UserAgent     *string
	DeviceID      *string
	DeviceName    *string
	Platform      *string
	OS            *string
	Browser       *string
	RequestID     *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
