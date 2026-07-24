package model

import (
	"time"
	"user/internal/enum"
)

type BanRecord struct {
	ID            int64
	UserID        int64
	OperatorID    int64
	OperatorRealm enum.LoginRealm
	Reason        string
	Remark        string
	StartedAt     time.Time
	BannedUntil   *time.Time
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
