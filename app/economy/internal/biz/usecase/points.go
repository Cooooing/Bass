package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/model"
	"economy/internal/biz/repo"
	"economy/internal/enum"

	"github.com/samber/lo"
)

type PointsUsecase struct {
	tx                     base.Tx
	accountRepo            repo.AccountRepo
	recordRepo             repo.RecordRepo
	transactionNoGenerator base.TransactionNoGenerator
}

func NewPointsUsecase(tx base.Tx, accountRepo repo.AccountRepo, recordRepo repo.RecordRepo, transactionNoGenerator base.TransactionNoGenerator) *PointsUsecase {
	return &PointsUsecase{tx: tx, accountRepo: accountRepo, recordRepo: recordRepo, transactionNoGenerator: transactionNoGenerator}
}

type EconomyAddPointsReq struct {
	UserID, Amount int64
	RecordType     enum.EconomyRecordType
	Remark         *string
}
type EconomyDeductPointsReq struct {
	UserID, Amount int64
	RecordType     enum.EconomyRecordType
	Remark         *string
}
type EconomyRecordResp struct {
	Account *model.Account
	Record  *model.Record
}
type EconomyTransferPointsReq struct {
	FromUserID, ToUserID, Amount int64
	OutRecordType, InRecordType  enum.EconomyRecordType
	Remark                       *string
}
type EconomyTransferPointsResp struct {
	FromAccount, ToAccount *model.Account
	OutRecord, InRecord    *model.Record
}
type EconomyListRecordsReq struct {
	UserID     int64
	Page       *base.PageRequest
	Direction  *enum.EconomyRecordDirection
	RecordType *enum.EconomyRecordType
}
type EconomyListRecordsResp struct {
	Rows []*model.Record
	Page *base.PageResp
}

func (u *PointsUsecase) GetAccount(ctx context.Context, userID int64) (*model.Account, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	account, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(userID)})
	if err != nil {
		return &model.Account{UserID: userID}, nil
	}
	return account, nil
}

func (u *PointsUsecase) MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.Account, error) {
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool { return userID > 0 })
	if len(ids) == 0 {
		return map[int64]*model.Account{}, nil
	}
	accounts, err := u.accountRepo.Map(ctx, &repo.AccountGetReq{UserIDs: ids})
	if err != nil {
		return nil, err
	}
	for _, userID := range ids {
		if accounts[userID] == nil {
			accounts[userID] = &model.Account{UserID: userID}
		}
	}
	return accounts, nil
}

func (u *PointsUsecase) AddPoints(ctx context.Context, req *EconomyAddPointsReq) (*EconomyRecordResp, error) {
	if req.UserID <= 0 || req.Amount <= 0 || !req.RecordType.IsIncome() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	var account *model.Account
	var record *model.Record
	err := u.tx(ctx, func(ctx context.Context) error {
		before, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.UserID)})
		if err != nil {
			before, err = u.accountRepo.Save(ctx, &model.Account{UserID: req.UserID})
			if err != nil {
				return err
			}
		}
		account, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.UserID, BalanceDelta: req.Amount, IncomeDelta: req.Amount})
		if err != nil {
			return err
		}
		transactionNo, err := u.transactionNoGenerator.NewTransactionNo()
		if err != nil {
			return err
		}
		record, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, UserID: req.UserID, RecordType: req.RecordType, Direction: enum.EconomyRecordDirectionIncome, Amount: req.Amount, BalanceBefore: before.Balance, BalanceAfter: account.Balance, Remark: req.Remark})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &EconomyRecordResp{Account: account, Record: record}, nil
}

func (u *PointsUsecase) DeductPoints(ctx context.Context, req *EconomyDeductPointsReq) (*EconomyRecordResp, error) {
	if req.UserID <= 0 || req.Amount <= 0 || !req.RecordType.IsExpense() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	var account *model.Account
	var record *model.Record
	err := u.tx(ctx, func(ctx context.Context) error {
		before, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.UserID)})
		if err != nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_INSUFFICIENT_BALANCE)
		}
		account, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.UserID, BalanceDelta: -req.Amount, ExpenseDelta: req.Amount, BalanceMin: new(req.Amount)})
		if err != nil {
			return err
		}
		transactionNo, err := u.transactionNoGenerator.NewTransactionNo()
		if err != nil {
			return err
		}
		record, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, UserID: req.UserID, RecordType: req.RecordType, Direction: enum.EconomyRecordDirectionExpense, Amount: req.Amount, BalanceBefore: before.Balance, BalanceAfter: account.Balance, Remark: req.Remark})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &EconomyRecordResp{Account: account, Record: record}, nil
}

func (u *PointsUsecase) TransferPoints(ctx context.Context, req *EconomyTransferPointsReq) (*EconomyTransferPointsResp, error) {
	if req.FromUserID <= 0 || req.ToUserID <= 0 || req.Amount <= 0 || !req.OutRecordType.IsExpense() || !req.InRecordType.IsIncome() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.FromUserID == req.ToUserID {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_TRANSFER_SELF_NOT_ALLOWED)
	}
	var fromAccount, toAccount *model.Account
	var outRecord, inRecord *model.Record
	err := u.tx(ctx, func(ctx context.Context) error {
		fromBefore, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.FromUserID)})
		if err != nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_INSUFFICIENT_BALANCE)
		}
		toBefore, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.ToUserID)})
		if err != nil {
			toBefore, err = u.accountRepo.Save(ctx, &model.Account{UserID: req.ToUserID})
			if err != nil {
				return err
			}
		}
		fromAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.FromUserID, BalanceDelta: -req.Amount, ExpenseDelta: req.Amount, BalanceMin: new(req.Amount)})
		if err != nil {
			return err
		}
		toAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.ToUserID, BalanceDelta: req.Amount, IncomeDelta: req.Amount})
		if err != nil {
			return err
		}
		transactionNo, err := u.transactionNoGenerator.NewTransactionNo()
		if err != nil {
			return err
		}
		outRecord, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, UserID: req.FromUserID, RecordType: req.OutRecordType, Direction: enum.EconomyRecordDirectionExpense, Amount: req.Amount, BalanceBefore: fromBefore.Balance, BalanceAfter: fromAccount.Balance, Remark: req.Remark})
		if err != nil {
			return err
		}
		inRecord, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, UserID: req.ToUserID, RecordType: req.InRecordType, Direction: enum.EconomyRecordDirectionIncome, Amount: req.Amount, BalanceBefore: toBefore.Balance, BalanceAfter: toAccount.Balance, Remark: req.Remark})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &EconomyTransferPointsResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}

func (u *PointsUsecase) ListRecords(ctx context.Context, req *EconomyListRecordsReq) (*EconomyListRecordsResp, error) {
	if req.UserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	page, err := u.recordRepo.Page(ctx, &repo.RecordGetReq{Page: req.Page, UserID: new(req.UserID), Direction: req.Direction, RecordType: req.RecordType})
	if err != nil {
		return nil, err
	}
	return &EconomyListRecordsResp{Rows: page.Rows, Page: page.Page}, nil
}
