package model

import (
	commonenum "common/pkg/enum"
	"time"
	"user/internal/enum"
)

type LoginLog struct {
	ID             int64
	UserID         *int64
	AccountInput   string
	LoginType      enum.LoginType
	Realm          commonenum.LoginRealm
	Status         enum.LoginStatus
	FailureReason  *enum.LoginFailureReason
	SessionID      string
	IP             *string
	Country        *string
	CountryCode    *string
	Province       *string
	City           *string
	ISP            *string
	UserAgent      *string
	ClientType     *enum.ClientType
	DeviceType     *enum.DeviceType
	OSName         string
	OSVersion      string
	BrowserName    string
	BrowserVersion string
	AppName        string
	AppVersion     string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
