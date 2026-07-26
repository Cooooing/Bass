package model

import (
	commonenum "common/pkg/enum"
	"time"
)

type BanRecord struct {
	ID            int64
	UserID        int64
	OperatorID    int64
	OperatorRealm commonenum.LoginRealm
	Reason        string
	Remark        string
	StartedAt     *time.Time
	BannedUntil   *time.Time
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
