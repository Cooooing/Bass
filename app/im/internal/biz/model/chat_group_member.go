package model

import (
	"im/internal/enum"
	"time"
)

type ChatGroupMember struct {
	ID        int64
	GroupID   int64
	UserID    int64
	Nickname  *string
	Role      enum.ChatGroupMemberRole
	MuteEndAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
	CreatedBy *int64
	UpdatedBy *int64

	Member *AccountBasic
}
