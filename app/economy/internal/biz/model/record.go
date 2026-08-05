package model

import (
	"economy/internal/enum"
	"time"
)

type Record struct {
	ID             int64
	TransactionNo  string
	UserID         int64
	RecordType     enum.EconomyRecordType
	Direction      enum.EconomyRecordDirection
	Amount         int64
	BalanceBefore  int64
	BalanceAfter   int64
	SourceID       *string
	IdempotencyKey string
	Remark         *string
	CreatedAt      *time.Time
}

func (r *Record) SameBusiness(userID int64, recordType enum.EconomyRecordType, amount int64, sourceID *string) bool {
	if r == nil || r.UserID != userID || r.RecordType != recordType || r.Amount != amount {
		return false
	}
	if r.SourceID == nil && sourceID == nil {
		return true
	}
	if r.SourceID == nil || sourceID == nil {
		return false
	}
	return *r.SourceID == *sourceID
}
