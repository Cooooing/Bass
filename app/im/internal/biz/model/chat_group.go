package model

import (
	"im/internal/enum"
	"time"
)

type ChatGroup struct {
	ID            int64
	Name          string
	Avatar        *string
	Introduction  *string
	OwnerID       int64
	Status        enum.ChatGroupStatus
	MemberCount   uint32
	MessageCount  uint32
	LastMessageID *int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	CreatedBy     *int64
	UpdatedBy     *int64

	WithOwner bool
	Owner     *AccountBasic
}
