package model

import (
	"math"
	"time"
)

type ChatSession struct {
	ID                int64
	ReceiverID        *int64
	GroupID           *int64
	IsMuted           bool
	IsPinned          bool
	LastReadMessageID *int64
	ReadCount         uint32
	MessageCount      uint32
	LastMessageID     *int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	CreatedBy         *int64
	UpdatedBy         *int64

	Group       *ChatGroup
	UnreadCount uint32
}

func (s *ChatSession) ReadDelta(latestMessageID int64) int32 {
	lastReadMessageID := int64(0)
	if s.LastReadMessageID != nil {
		lastReadMessageID = *s.LastReadMessageID
	}
	delta := latestMessageID - lastReadMessageID
	if delta <= 0 {
		return 0
	}
	if delta >= math.MaxInt32 {
		return math.MaxInt32
	}
	return int32(delta) // #nosec G115 -- bounded by MaxInt32 above.
}
