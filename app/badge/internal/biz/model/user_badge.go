package model

import "time"

// UserBadge is a user's owned instance of a badge.
type UserBadge struct {
	ID           int64
	UserID       int64
	BadgeID      int64
	Variables    []byte
	ObtainedAt   time.Time
	ExpiresAt    *time.Time
	Displayed    bool
	DisplayOrder int32
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
