package repo

import (
	economyv1enum "common/proto/gen/economy/v1/enum"
	"context"
	"time"
)

type EconomyAccount struct {
	Balance      int64
	TotalIncome  int64
	TotalExpense int64
}

type EconomyRecord struct {
	ID            int64
	TransactionNo string
	RecordType    economyv1enum.EconomyRecordType
	Direction     economyv1enum.EconomyRecordDirection
	Amount        int64
	BalanceBefore int64
	BalanceAfter  int64
	Remark        *string
	CreatedAt     *time.Time
}

type ListEconomyRecordsReq struct {
	UserID     int64
	Page       *PageReq
	Direction  *economyv1enum.EconomyRecordDirection
	RecordType *economyv1enum.EconomyRecordType
}

type ListEconomyRecordsResp struct {
	Rows []*EconomyRecord
	Page *PageResp
}

type EconomyClient interface {
	GetAccount(ctx context.Context, userID int64) (*EconomyAccount, error)
	ListRecords(ctx context.Context, req *ListEconomyRecordsReq) (*ListEconomyRecordsResp, error)
}
