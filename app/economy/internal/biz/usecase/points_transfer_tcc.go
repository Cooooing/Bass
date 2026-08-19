package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"economy/internal/biz/model"
	"economy/internal/biz/repo"
	"economy/internal/enum"
)

type PointsTransferTccReq struct {
	FromUserID, ToUserID, Amount int64
	OutRecordType, InRecordType  enum.EconomyRecordType
	Remark                       *string
}
type PointsTransferTccResp struct {
	FromAccount, ToAccount *model.Account
	OutRecord, InRecord    *model.Record
}

func (u *PointsUsecase) TransferTry(ctx context.Context, req *PointsTransferTccReq) (*PointsTransferTccResp, error) {
	if req.FromUserID <= 0 || req.ToUserID <= 0 || req.Amount <= 0 || !req.OutRecordType.IsExpense() || !req.InRecordType.IsIncome() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.FromUserID == req.ToUserID {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_TRANSFER_SELF_NOT_ALLOWED)
	}
	var fromAccount *model.Account
	var toAccount *model.Account
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		fromAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.FromUserID, FrozenDelta: req.Amount, AvailableMin: new(req.Amount)})
		if err != nil {
			return err
		}
		toAccount, err = u.GetAccount(ctx, req.ToUserID)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &PointsTransferTccResp{FromAccount: fromAccount, ToAccount: toAccount}, nil
}

func (u *PointsUsecase) TransferComfirm(ctx context.Context, req *PointsTransferTccReq) (*PointsTransferTccResp, error) {
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
			return err
		}
		toBefore, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.ToUserID)})
		if err != nil {
			toBefore, err = u.accountRepo.Save(ctx, &model.Account{UserID: req.ToUserID})
			if err != nil {
				return err
			}
		}
		fromAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.FromUserID, BalanceDelta: -req.Amount, FrozenDelta: -req.Amount, ExpenseDelta: req.Amount, BalanceMin: new(req.Amount), FrozenMin: new(req.Amount)})
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
	return &PointsTransferTccResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}

func (u *PointsUsecase) TransferCancel(ctx context.Context, req *PointsTransferTccReq) (*PointsTransferTccResp, error) {
	if req.FromUserID <= 0 || req.ToUserID <= 0 || req.Amount <= 0 || !req.OutRecordType.IsExpense() || !req.InRecordType.IsIncome() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.FromUserID == req.ToUserID {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_TRANSFER_SELF_NOT_ALLOWED)
	}
	var fromAccount *model.Account
	var toAccount *model.Account
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		fromAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.FromUserID, FrozenDelta: -req.Amount, FrozenMin: new(req.Amount)})
		if err != nil {
			return err
		}
		toAccount, err = u.GetAccount(ctx, req.ToUserID)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &PointsTransferTccResp{FromAccount: fromAccount, ToAccount: toAccount}, nil
}
