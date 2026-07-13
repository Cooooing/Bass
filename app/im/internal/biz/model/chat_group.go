package model

import (
	"time"

	userv1 "common/proto/gen/user/v1"
	"im/internal/enum"
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
	Owner     *userv1.AccountBasic
}
