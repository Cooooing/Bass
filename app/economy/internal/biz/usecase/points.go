package usecase

import (
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"

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
	IdempotencyKey *string
}
type EconomyDeductPointsReq struct {
	UserID, Amount int64
	RecordType     enum.EconomyRecordType
	Remark         *string
	IdempotencyKey *string
}
type EconomyFreezePointsReq struct {
	UserID, Amount int64
}
type EconomyCommitFrozenPointsReq struct {
	UserID, Amount int64
	RecordType     enum.EconomyRecordType
	Remark         *string
	IdempotencyKey *string
}
type EconomyRecordResp struct {
	Account *model.Account
	Record  *model.Record
}
type EconomyTransferPointsReq struct {
	FromUserID, ToUserID, Amount int64
	OutRecordType, InRecordType  enum.EconomyRecordType
	Remark                       *string
	IdempotencyKey               *string
}
type EconomyTransferPointsResp struct {
	FromAccount, ToAccount *model.Account
	OutRecord, InRecord    *model.Record
}
type EconomyCommitFrozenTransferReq struct {
	FromUserID, ToUserID, Amount int64
	OutRecordType, InRecordType  enum.EconomyRecordType
	Remark                       *string
	IdempotencyKey               *string
}
type EconomyListRecordsReq struct {
	UserID     int64
	Page       *commonmodel.PageReq
	Direction  *enum.EconomyRecordDirection
	RecordType *enum.EconomyRecordType
}
type EconomyListRecordsResp struct {
	Rows []*model.Record
	Page *commonmodel.PageResp
}

func (u *PointsUsecase) GetAccount(ctx context.Context, userID int64) (*model.Account, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	account, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(userID)})
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (u *PointsUsecase) EnsureAccount(ctx context.Context, userID int64) (*model.Account, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var account *model.Account
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		account, err = u.accountRepo.Ensure(ctx, userID)
		return err
	})
	if err != nil {
		return nil, err
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
		before, err := u.accountRepo.Ensure(ctx, req.UserID)
		if err != nil {
			return err
		}
		account, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.UserID, BalanceDelta: req.Amount, IncomeDelta: req.Amount})
		if err != nil {
			return err
		}
		transactionNo, err := u.transactionNoGenerator.NewTransactionNo()
		if err != nil {
			return err
		}
		record, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, IdempotencyKey: req.IdempotencyKey, UserID: req.UserID, RecordType: req.RecordType, Direction: enum.EconomyRecordDirectionIncome, Amount: req.Amount, BalanceBefore: before.Balance, BalanceAfter: account.Balance, Remark: req.Remark})
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
		record, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, IdempotencyKey: req.IdempotencyKey, UserID: req.UserID, RecordType: req.RecordType, Direction: enum.EconomyRecordDirectionExpense, Amount: req.Amount, BalanceBefore: before.Balance, BalanceAfter: account.Balance, Remark: req.Remark})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &EconomyRecordResp{Account: account, Record: record}, nil
}

func (u *PointsUsecase) FreezePoints(ctx context.Context, req *EconomyFreezePointsReq) (*model.Account, error) {
	if req == nil || req.UserID <= 0 || req.Amount <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
		UserID:       req.UserID,
		FrozenDelta:  req.Amount,
		AvailableMin: new(req.Amount),
	})
}

func (u *PointsUsecase) UnfreezePoints(ctx context.Context, req *EconomyFreezePointsReq) (*model.Account, error) {
	if req == nil || req.UserID <= 0 || req.Amount <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
		UserID:      req.UserID,
		FrozenDelta: -req.Amount,
		FrozenMin:   new(req.Amount),
	})
}

func (u *PointsUsecase) CommitFrozenPoints(ctx context.Context, req *EconomyCommitFrozenPointsReq) (*EconomyRecordResp, error) {
	if req == nil || req.UserID <= 0 || req.Amount <= 0 || !req.RecordType.IsExpense() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
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
			TransactionNo:  transactionNo,
			IdempotencyKey: req.IdempotencyKey,
			UserID:         req.UserID,
			RecordType:     req.RecordType,
			Direction:      enum.EconomyRecordDirectionExpense,
			Amount:         req.Amount,
			BalanceBefore:  before.Balance,
			BalanceAfter:   account.Balance,
			Remark:         req.Remark,
		})
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
	existing, err := u.existingTransfer(ctx, req)
	if err != nil || existing != nil {
		return existing, err
	}
	var fromAccount, toAccount *model.Account
	var outRecord, inRecord *model.Record
	err = u.tx(ctx, func(ctx context.Context) error {
		fromBefore, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.FromUserID)})
		if err != nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_INSUFFICIENT_BALANCE)
		}
		toBefore, err := u.accountRepo.Ensure(ctx, req.ToUserID)
		if err != nil {
			return err
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
		outRecord, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, IdempotencyKey: req.IdempotencyKey, UserID: req.FromUserID, RecordType: req.OutRecordType, Direction: enum.EconomyRecordDirectionExpense, Amount: req.Amount, BalanceBefore: fromBefore.Balance, BalanceAfter: fromAccount.Balance, Remark: req.Remark})
		if err != nil {
			return err
		}
		inRecord, err = u.recordRepo.Save(ctx, &model.Record{TransactionNo: transactionNo, IdempotencyKey: req.IdempotencyKey, UserID: req.ToUserID, RecordType: req.InRecordType, Direction: enum.EconomyRecordDirectionIncome, Amount: req.Amount, BalanceBefore: toBefore.Balance, BalanceAfter: toAccount.Balance, Remark: req.Remark})
		return err
	})
	if err != nil {
		existing, existingErr := u.existingTransfer(ctx, req)
		if existingErr == nil && existing != nil {
			return existing, nil
		}
		return nil, err
	}
	return &EconomyTransferPointsResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
}

func (u *PointsUsecase) existingTransfer(ctx context.Context, req *EconomyTransferPointsReq) (*EconomyTransferPointsResp, error) {
	if req.IdempotencyKey == nil || strings.TrimSpace(*req.IdempotencyKey) == "" {
		return nil, nil
	}
	outRecord, err := u.recordRepo.Get(ctx, &repo.RecordGetReq{UserID: &req.FromUserID, IdempotencyKey: req.IdempotencyKey})
	if err != nil {
		if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_NOT_FOUND {
			return nil, nil
		}
		return nil, err
	}
	inRecord, err := u.recordRepo.Get(ctx, &repo.RecordGetReq{TransactionNo: &outRecord.TransactionNo, UserID: &req.ToUserID})
	if err != nil || outRecord.RecordType != req.OutRecordType || inRecord.RecordType != req.InRecordType || outRecord.Amount != req.Amount || inRecord.Amount != req.Amount || outRecord.Direction != enum.EconomyRecordDirectionExpense || inRecord.Direction != enum.EconomyRecordDirectionIncome {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_IDEMPOTENCY_CONFLICT)
	}
	return &EconomyTransferPointsResp{
		FromAccount: &model.Account{UserID: req.FromUserID, Balance: outRecord.BalanceAfter},
		ToAccount:   &model.Account{UserID: req.ToUserID, Balance: inRecord.BalanceAfter},
		OutRecord:   outRecord,
		InRecord:    inRecord,
	}, nil
}

func (u *PointsUsecase) CommitFrozenTransfer(ctx context.Context, req *EconomyCommitFrozenTransferReq) (*EconomyTransferPointsResp, error) {
	if req == nil || req.FromUserID <= 0 || req.ToUserID <= 0 || req.Amount <= 0 || !req.OutRecordType.IsExpense() || !req.InRecordType.IsIncome() {
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
		toBefore, err := u.accountRepo.Ensure(ctx, req.ToUserID)
		if err != nil {
			return err
		}
		fromAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
			UserID:       req.FromUserID,
			BalanceDelta: -req.Amount,
			FrozenDelta:  -req.Amount,
			ExpenseDelta: req.Amount,
			BalanceMin:   new(req.Amount),
			FrozenMin:    new(req.Amount),
		})
		if err != nil {
			return err
		}
		toAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
			UserID:       req.ToUserID,
			BalanceDelta: req.Amount,
			IncomeDelta:  req.Amount,
		})
		if err != nil {
			return err
		}
		transactionNo, err := u.transactionNoGenerator.NewTransactionNo()
		if err != nil {
			return err
		}
		outRecord, err = u.recordRepo.Save(ctx, &model.Record{
			TransactionNo:  transactionNo,
			IdempotencyKey: req.IdempotencyKey,
			UserID:         req.FromUserID,
			RecordType:     req.OutRecordType,
			Direction:      enum.EconomyRecordDirectionExpense,
			Amount:         req.Amount,
			BalanceBefore:  fromBefore.Balance,
			BalanceAfter:   fromAccount.Balance,
			Remark:         req.Remark,
		})
		if err != nil {
			return err
		}
		inRecord, err = u.recordRepo.Save(ctx, &model.Record{
			TransactionNo:  transactionNo,
			IdempotencyKey: req.IdempotencyKey,
			UserID:         req.ToUserID,
			RecordType:     req.InRecordType,
			Direction:      enum.EconomyRecordDirectionIncome,
			Amount:         req.Amount,
			BalanceBefore:  toBefore.Balance,
			BalanceAfter:   toAccount.Balance,
			Remark:         req.Remark,
		})
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
