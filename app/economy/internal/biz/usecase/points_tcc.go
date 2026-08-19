package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"economy/internal/biz/model"
	"economy/internal/biz/repo"
	"economy/internal/enum"
)

type PointsTccReq struct {
	UserID, Amount int64
	RecordType     enum.EconomyRecordType
	Remark         *string
}
type PointsTccResp struct {
	Account *model.Account
	Record  *model.Record
}

func (u *PointsUsecase) Try(ctx context.Context, req *PointsTccReq) (*PointsTccResp, error) {
	if req.UserID <= 0 || req.Amount <= 0 || (!req.RecordType.IsIncome() && !req.RecordType.IsExpense()) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.RecordType.IsIncome() {
		account, err := u.GetAccount(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
		return &PointsTccResp{Account: account}, nil
	}
	var account *model.Account
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		account, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
			UserID:       req.UserID,
			FrozenDelta:  req.Amount,
			AvailableMin: new(req.Amount),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &PointsTccResp{Account: account}, nil
}

func (u *PointsUsecase) Comfirm(ctx context.Context, req *PointsTccReq) (*PointsTccResp, error) {
	if req.UserID <= 0 || req.Amount <= 0 || (!req.RecordType.IsIncome() && !req.RecordType.IsExpense()) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.RecordType.IsIncome() {
		resp, err := u.AddPoints(ctx, &EconomyAddPointsReq{
			UserID:     req.UserID,
			Amount:     req.Amount,
			RecordType: req.RecordType,
			Remark:     req.Remark,
		})
		if err != nil {
			return nil, err
		}
		return &PointsTccResp{
			Account: resp.Account,
			Record:  resp.Record,
		}, nil
	}
	var account *model.Account
	var record *model.Record
	err := u.tx(ctx, func(ctx context.Context) error {
		before, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.UserID)})
		if err != nil {
			return err
		}
		account, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
			UserID:       req.UserID,
			BalanceDelta: -req.Amount,
			FrozenDelta:  -req.Amount,
			ExpenseDelta: req.Amount,
			BalanceMin:   new(req.Amount),
			FrozenMin:    new(req.Amount),
		})
		if err != nil {
			return err
		}
		transactionNo, err := u.transactionNoGenerator.NewTransactionNo()
		if err != nil {
			return err
		}
		record, err = u.recordRepo.Save(ctx, &model.Record{
			TransactionNo: transactionNo,
			UserID:        req.UserID,
			RecordType:    req.RecordType,
			Direction:     enum.EconomyRecordDirectionExpense,
			Amount:        req.Amount,
			BalanceBefore: before.Balance,
			BalanceAfter:  account.Balance,
			Remark:        req.Remark,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &PointsTccResp{Account: account, Record: record}, nil
}

func (u *PointsUsecase) Cancel(ctx context.Context, req *PointsTccReq) (*PointsTccResp, error) {
	if req.UserID <= 0 || req.Amount <= 0 || (!req.RecordType.IsIncome() && !req.RecordType.IsExpense()) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.RecordType.IsIncome() {
		account, err := u.GetAccount(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
		return &PointsTccResp{Account: account}, nil
	}
	var account *model.Account
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		account, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
			UserID:      req.UserID,
			FrozenDelta: -req.Amount,
			FrozenMin:   new(req.Amount),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &PointsTccResp{Account: account}, nil
}
