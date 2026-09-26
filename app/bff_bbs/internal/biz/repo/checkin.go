package repo

import (
	"context"
	"time"
)

type CheckinClient interface {
	CheckIn(ctx context.Context, userID int64) (*Checkin, error)
	GetOverview(ctx context.Context, userID int64, month string) (*CheckinOverview, error)
}

type Checkin struct {
	RecordID      int64
	Date          *time.Time
	CurrentStreak int32
	LongestStreak int32
}

type CheckinOverview struct {
	Records       []*CheckinRecord
	CurrentStreak int32
	LongestStreak int32
}

type CheckinRecord struct {
	ID      int64
	Date    *time.Time
	Checked bool
}
