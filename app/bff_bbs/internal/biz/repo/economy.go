package repo

import (
	commonmodel "common/pkg/model"
	commonenums "common/proto/gen/common/enums"
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
	TypeLabel     string
	CreatedAt     *time.Time
}

type ListEconomyRecordsReq struct {
	UserID     int64
	Page       *commonmodel.PageReq
	Direction  *economyv1enum.EconomyRecordDirection
	RecordType *economyv1enum.EconomyRecordType
	Language   commonenums.Language
}

type ListEconomyRecordsResp struct {
	Rows []*EconomyRecord
	Page *commonmodel.PageResp
}

type TransferEconomyPointsReq struct {
	FromUserID     int64
	ToUserID       int64
	Amount         int64
	Remark         *string
	IdempotencyKey string
}

type TransferEconomyPointsResp struct {
	Balance int64
}

type EconomyClient interface {
	GetAccount(ctx context.Context, userID int64) (*EconomyAccount, error)
	ListRecords(ctx context.Context, req *ListEconomyRecordsReq) (*ListEconomyRecordsResp, error)
	TransferPoints(ctx context.Context, req *TransferEconomyPointsReq) (*TransferEconomyPointsResp, error)
}
