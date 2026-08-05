package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	economyv1 "common/proto/gen/economy/v1"
	"context"
)

var _ repo.EconomyClient = (*EconomyClient)(nil)

type EconomyClient struct {
	economyClient *rpc.EconomyClient
}

func NewEconomyClient(economyClient *rpc.EconomyClient) repo.EconomyClient {
	return &EconomyClient{economyClient: economyClient}
}

func (r *EconomyClient) GetAccount(ctx context.Context, userID int64) (*repo.EconomyAccount, error) {
	reply, err := r.economyClient.Economy.GetAccount(ctx, &economyv1.GetEconomyAccount_Req{UserId: userID})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	return &repo.EconomyAccount{Balance: account.GetBalance(), TotalIncome: account.GetTotalIncome(), TotalExpense: account.GetTotalExpense()}, nil
}

func (r *EconomyClient) ListRecords(ctx context.Context, req *repo.ListEconomyRecordsReq) (*repo.ListEconomyRecordsResp, error) {
	var page *common.PageReq
	if req.Page != nil {
		page = &common.PageReq{Page: req.Page.Page, Size: req.Page.Size}
	}
	rpcReq := &economyv1.ListEconomyRecords_Req{UserId: req.UserID, Page: page, Direction: req.Direction, RecordType: req.RecordType}
	reply, err := r.economyClient.Economy.ListRecords(ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.EconomyRecord, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		item := &repo.EconomyRecord{ID: row.GetId(), TransactionNo: row.GetTransactionNo(), RecordType: row.GetRecordType(), Direction: row.GetDirection(), Amount: row.GetAmount(), BalanceBefore: row.GetBalanceBefore(), BalanceAfter: row.GetBalanceAfter(), Remark: row.Remark}
		if row.GetCreatedAt() != nil {
			item.CreatedAt = new(row.GetCreatedAt().AsTime())
		}
		rows = append(rows, item)
	}
	var pageResp *repo.PageResp
	if reply.GetPage() != nil {
		pageResp = &repo.PageResp{Page: reply.GetPage().GetPage(), Size: reply.GetPage().GetSize(), Total: reply.GetPage().GetTotal()}
	}
	return &repo.ListEconomyRecordsResp{Rows: rows, Page: pageResp}, nil
}
