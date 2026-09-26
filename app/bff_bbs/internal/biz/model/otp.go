package model

import "time"

type Totp struct {
	UserID     int64
	Enable     bool
	EnableTime *time.Time
}
