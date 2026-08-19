package model

import (
	"economy/internal/enum"
	"time"
)

type Record struct {
	ID            int64
	TransactionNo string
	UserID        int64
	RecordType    enum.EconomyRecordType
	Direction     enum.EconomyRecordDirection
	Amount        int64
	BalanceBefore int64
	BalanceAfter  int64
	Remark        *string
	CreatedAt     *time.Time
}
