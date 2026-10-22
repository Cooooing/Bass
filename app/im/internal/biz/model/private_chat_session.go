package model

import "time"

type PrivateChatSession struct {
	ID                int64
	UserID            int64
	PeerUserID        int64
	LastMessageID     *int64
	LastMessageAt     *time.Time
	LastMessage       *PrivateChatMessage
	LastReadMessageID *int64
	UnreadCount       uint32
	Muted             bool
	Pinned            bool
}
