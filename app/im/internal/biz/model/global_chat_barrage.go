package model

import (
	"im/internal/enum"
	"time"
)

type GlobalChatBarrage struct {
	ID              int64
	UserID          int64
	ClientRequestID string
	Content         string
	PointsCost      int64
	Status          enum.GlobalChatBarrageStatus
	CreatedAt       *time.Time
}
