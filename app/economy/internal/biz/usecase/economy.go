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

type EconomyUsecase struct {
	tx          base.Tx
	accountRepo repo.AccountRepo
	recordRepo  repo.RecordRepo
}

func NewEconomyUsecase(
	tx base.Tx,
	accountRepo repo.AccountRepo,
	recordRepo repo.RecordRepo,
) *EconomyUsecase {
	return &EconomyUsecase{
		tx:          tx,
		accountRepo: accountRepo,
		recordRepo:  recordRepo,
	}
}

func (u *EconomyUsecase) GetAccount(ctx context.Context, userID int64) (*model.Account, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	account, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(userID)})
	if err != nil {
		return &model.Account{UserID: userID}, nil
	}
	return account, nil
}

func (u *EconomyUsecase) MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.Account, error) {
	ids := lo.Filter(lo.Uniq(userIDs), func(userID int64, _ int) bool {
		return userID > 0
	})
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

type EconomyAddPointsReq struct {
	UserID         int64
	Amount         int64
	RecordType     enum.EconomyRecordType
	SourceID       *string
	IdempotencyKey string
	Remark         *string
}

type EconomyRecordResp struct {
	Account *model.Account
	Record  *model.Record
}

func (u *EconomyUsecase) AddPoints(ctx context.Context, req *EconomyAddPointsReq) (*EconomyRecordResp, error) {
	if req.Amount <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_AMOUNT_INVALID)
	}
	if req.UserID <= 0 || req.IdempotencyKey == "" || !req.RecordType.IsIncome() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return u.changePoints(ctx, &economyChangePointsReq{
		UserID:         req.UserID,
		Amount:         req.Amount,
		RecordType:     req.RecordType,
		SourceID:       req.SourceID,
		IdempotencyKey: req.IdempotencyKey,
		Remark:         req.Remark,
	})
}

type EconomyDeductPointsReq struct {
	UserID         int64
	Amount         int64
	RecordType     enum.EconomyRecordType
	SourceID       *string
	IdempotencyKey string
	Remark         *string
}

func (u *EconomyUsecase) DeductPoints(ctx context.Context, req *EconomyDeductPointsReq) (*EconomyRecordResp, error) {
	if req.Amount <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_AMOUNT_INVALID)
	}
	if req.UserID <= 0 || req.IdempotencyKey == "" || !req.RecordType.IsExpense() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return u.changePoints(ctx, &economyChangePointsReq{
		UserID:         req.UserID,
		Amount:         req.Amount,
		RecordType:     req.RecordType,
		SourceID:       req.SourceID,
		IdempotencyKey: req.IdempotencyKey,
		Remark:         req.Remark,
	})
}

type economyChangePointsReq struct {
	UserID         int64
	Amount         int64
	RecordType     enum.EconomyRecordType
	SourceID       *string
	IdempotencyKey string
	Remark         *string
}

func (u *EconomyUsecase) changePoints(ctx context.Context, req *economyChangePointsReq) (*EconomyRecordResp, error) {
	if existing, err := u.recordRepo.Get(ctx, &repo.RecordGetReq{IdempotencyKey: new(req.IdempotencyKey)}); err == nil {
		if !existing.SameBusiness(req.UserID, req.RecordType, req.Amount, req.SourceID) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_IDEMPOTENCY_CONFLICT)
		}
		account, accountErr := u.GetAccount(ctx, req.UserID)
		return &EconomyRecordResp{Account: account, Record: existing}, accountErr
	}
	var account *model.Account
	var record *model.Record
	err := u.tx(ctx, func(ctx context.Context) error {
		if existing, err := u.recordRepo.Get(ctx, &repo.RecordGetReq{IdempotencyKey: new(req.IdempotencyKey)}); err == nil {
			if !existing.SameBusiness(req.UserID, req.RecordType, req.Amount, req.SourceID) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_IDEMPOTENCY_CONFLICT)
			}
			record = existing
			accountResp, accountErr := u.GetAccount(ctx, req.UserID)
			account = accountResp
			return accountErr
		}
		accountResp, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.UserID)})
		if err != nil {
			accountResp, err = u.accountRepo.Save(ctx, &model.Account{UserID: req.UserID})
			if err != nil {
				return err
			}
		}
		direction, err := req.RecordType.Direction()
		if err != nil {
			return err
		}
		balanceDelta := req.Amount
		incomeDelta := req.Amount
		expenseDelta := int64(0)
		if direction == enum.EconomyRecordDirectionExpense {
			if !accountResp.CanDeduct(req.Amount) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_INSUFFICIENT_BALANCE)
			}
			balanceDelta = -req.Amount
			incomeDelta = 0
			expenseDelta = req.Amount
		}
		balanceBefore := accountResp.Balance
		accountResp, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{
			UserID:       req.UserID,
			BalanceDelta: balanceDelta,
			IncomeDelta:  incomeDelta,
			ExpenseDelta: expenseDelta,
		})
		if err != nil {
			return err
		}
		account = accountResp
		record, err = u.recordRepo.Save(ctx, &model.Record{
			TransactionNo:  req.IdempotencyKey,
			UserID:         req.UserID,
			RecordType:     req.RecordType,
			Direction:      direction,
			Amount:         req.Amount,
			BalanceBefore:  balanceBefore,
			BalanceAfter:   accountResp.Balance,
			SourceID:       req.SourceID,
			IdempotencyKey: req.IdempotencyKey,
			Remark:         req.Remark,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &EconomyRecordResp{Account: account, Record: record}, nil
}

type EconomyTransferPointsReq struct {
	FromUserID     int64
	ToUserID       int64
	Amount         int64
	OutRecordType  enum.EconomyRecordType
	InRecordType   enum.EconomyRecordType
	SourceID       *string
	IdempotencyKey string
	Remark         *string
}

type EconomyTransferPointsResp struct {
	FromAccount *model.Account
	ToAccount   *model.Account
	OutRecord   *model.Record
	InRecord    *model.Record
}

func (u *EconomyUsecase) TransferPoints(ctx context.Context, req *EconomyTransferPointsReq) (*EconomyTransferPointsResp, error) {
	if req.FromUserID <= 0 || req.ToUserID <= 0 || req.IdempotencyKey == "" || !req.OutRecordType.IsExpense() || !req.InRecordType.IsIncome() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.FromUserID == req.ToUserID {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_TRANSFER_SELF_NOT_ALLOWED)
	}
	if req.Amount <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_AMOUNT_INVALID)
	}
	outKey := req.IdempotencyKey + ":out"
	inKey := req.IdempotencyKey + ":in"
	var fromAccount *model.Account
	var toAccount *model.Account
	var outRecord *model.Record
	var inRecord *model.Record
	err := u.tx(ctx, func(ctx context.Context) error {
		if existing, err := u.recordRepo.Get(ctx, &repo.RecordGetReq{IdempotencyKey: new(outKey)}); err == nil {
			if !existing.SameBusiness(req.FromUserID, req.OutRecordType, req.Amount, req.SourceID) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_IDEMPOTENCY_CONFLICT)
			}
			outRecord = existing
			inRecord, err = u.recordRepo.Get(ctx, &repo.RecordGetReq{IdempotencyKey: new(inKey)})
			if err != nil {
				return err
			}
			fromAccount, err = u.GetAccount(ctx, req.FromUserID)
			if err != nil {
				return err
			}
			toAccount, err = u.GetAccount(ctx, req.ToUserID)
			return err
		}
		fromResp, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.FromUserID)})
		if err != nil {
			fromResp, err = u.accountRepo.Save(ctx, &model.Account{UserID: req.FromUserID})
			if err != nil {
				return err
			}
		}
		toResp, err := u.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: new(req.ToUserID)})
		if err != nil {
			toResp, err = u.accountRepo.Save(ctx, &model.Account{UserID: req.ToUserID})
			if err != nil {
				return err
			}
		}
		if !fromResp.CanDeduct(req.Amount) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_INSUFFICIENT_BALANCE)
		}
		transactionNo := req.IdempotencyKey
		fromBefore := fromResp.Balance
		toBefore := toResp.Balance
		fromAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.FromUserID, BalanceDelta: -req.Amount, ExpenseDelta: req.Amount})
		if err != nil {
			return err
		}
		toAccount, err = u.accountRepo.UpdateBalance(ctx, &repo.AccountUpdateBalanceReq{UserID: req.ToUserID, BalanceDelta: req.Amount, IncomeDelta: req.Amount})
		if err != nil {
			return err
		}
		outRecord, err = u.recordRepo.Save(ctx, &model.Record{
			TransactionNo: transactionNo, UserID: req.FromUserID, RecordType: req.OutRecordType, Direction: enum.EconomyRecordDirectionExpense,
			Amount: req.Amount, BalanceBefore: fromBefore, BalanceAfter: fromAccount.Balance, SourceID: req.SourceID, IdempotencyKey: outKey, Remark: req.Remark,
		})
		if err != nil {
			return err
		}
		inRecord, err = u.recordRepo.Save(ctx, &model.Record{
			TransactionNo: transactionNo, UserID: req.ToUserID, RecordType: req.InRecordType, Direction: enum.EconomyRecordDirectionIncome,
			Amount: req.Amount, BalanceBefore: toBefore, BalanceAfter: toAccount.Balance, SourceID: req.SourceID, IdempotencyKey: inKey, Remark: req.Remark,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &EconomyTransferPointsResp{FromAccount: fromAccount, ToAccount: toAccount, OutRecord: outRecord, InRecord: inRecord}, nil
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

func (u *EconomyUsecase) ListRecords(ctx context.Context, req *EconomyListRecordsReq) (*EconomyListRecordsResp, error) {
	if req.UserID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	page, err := u.recordRepo.Page(ctx, &repo.RecordGetReq{Page: req.Page, UserID: new(req.UserID), Direction: req.Direction, RecordType: req.RecordType})
	if err != nil {
		return nil, err
	}
	return &EconomyListRecordsResp{Rows: page.Rows, Page: page.Page}, nil
}
