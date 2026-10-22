package model

import "time"

type GlobalChatMute struct {
	UserID         int64
	MutedUntil     time.Time
	Reason         *string
	OperatorUserID int64
}
