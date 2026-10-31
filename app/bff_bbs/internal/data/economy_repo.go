package data

import (
	"common/pkg/client/rpc"
	commonmodel "common/pkg/model"
	economyv1 "common/proto/gen/economy/v1"
	economyv1enum "common/proto/gen/economy/v1/enum"
	"context"

	"bff_bbs/internal/biz/repo"
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
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{Page: req.Page.Page, Size: req.Page.Size}
	}
	rpcReq := &economyv1.ListEconomyRecords_Req{UserId: req.UserID, Page: page.ToProto(), Direction: req.Direction, RecordType: req.RecordType, Language: req.Language}
	reply, err := r.economyClient.Economy.ListRecords(ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.EconomyRecord, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		item := &repo.EconomyRecord{ID: row.GetId(), TransactionNo: row.GetTransactionNo(), RecordType: row.GetRecordType(), Direction: row.GetDirection(), Amount: row.GetAmount(), BalanceBefore: row.GetBalanceBefore(), BalanceAfter: row.GetBalanceAfter(), Remark: row.Remark, TypeLabel: row.GetTypeLabel()}
		if row.GetCreatedAt() != nil {
			item.CreatedAt = new(row.GetCreatedAt().AsTime())
		}
		rows = append(rows, item)
	}
	var pageResp *commonmodel.PageResp
	if reply.GetPage() != nil {
		pageResp = commonmodel.PageRespFromProto(reply.GetPage())
	}
	return &repo.ListEconomyRecordsResp{Rows: rows, Page: pageResp}, nil
}

func (r *EconomyClient) TransferPoints(ctx context.Context, req *repo.TransferEconomyPointsReq) (*repo.TransferEconomyPointsResp, error) {
	reply, err := r.economyClient.Economy.TransferPoints(ctx, &economyv1.TransferEconomyPoints_Req{
		FromUserId:     req.FromUserID,
		ToUserId:       req.ToUserID,
		Amount:         req.Amount,
		OutRecordType:  economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ACCOUNT_TRANSFER_OUT,
		InRecordType:   economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ACCOUNT_TRANSFER_IN,
		Remark:         req.Remark,
		IdempotencyKey: &req.IdempotencyKey,
	})
	if err != nil {
		return nil, err
	}
	return &repo.TransferEconomyPointsResp{Balance: reply.GetFromAccount().GetBalance()}, nil
}
