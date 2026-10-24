package usecase

import (
	"common/pkg/apperror"
	commonClient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"
	"context"

	"economy/internal/biz/model"
	"economy/internal/enum"
)

type TCCPointsTransferUsecase struct {
	points    *PointsUsecase
	tccBranch commonClient.TCCBranchExecutor
}

func NewTCCPointsTransferUsecase(points *PointsUsecase, tccBranch commonClient.TCCBranchExecutor) *TCCPointsTransferUsecase {
	return &TCCPointsTransferUsecase{points: points, tccBranch: tccBranch}
}

type TCCPointsTransferReq struct {
	FromUserID, ToUserID, Amount int64
	OutRecordType, InRecordType  enum.EconomyRecordType
	Remark                       *string
}

type TCCPointsTransferResp struct {
	FromAccount, ToAccount *model.Account
	OutRecord, InRecord    *model.Record
}

func (u *TCCPointsTransferUsecase) Try(ctx context.Context, req *TCCPointsTransferReq) (*TCCPointsTransferResp, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	response := &TCCPointsTransferResp{}
	err := u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		toAccount, err := u.points.EnsureAccount(ctx, req.ToUserID)
		if err != nil {
			return err
		}
		fromAccount, err := u.points.FreezePoints(ctx, &EconomyFreezePointsReq{UserID: req.FromUserID, Amount: req.Amount})
		if err != nil {
			return err
		}
		response.FromAccount = fromAccount
		response.ToAccount = toAccount
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *TCCPointsTransferUsecase) Confirm(ctx context.Context, req *TCCPointsTransferReq) (*TCCPointsTransferResp, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	response := &TCCPointsTransferResp{}
	err := u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		result, err := u.points.CommitFrozenTransfer(ctx, &EconomyCommitFrozenTransferReq{
			FromUserID:    req.FromUserID,
			ToUserID:      req.ToUserID,
			Amount:        req.Amount,
			OutRecordType: req.OutRecordType,
			InRecordType:  req.InRecordType,
			Remark:        req.Remark,
		})
		if err != nil {
			return err
		}
		response.FromAccount = result.FromAccount
		response.ToAccount = result.ToAccount
		response.OutRecord = result.OutRecord
		response.InRecord = result.InRecord
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *TCCPointsTransferUsecase) Cancel(ctx context.Context, req *TCCPointsTransferReq) (*TCCPointsTransferResp, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	response := &TCCPointsTransferResp{}
	err := u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		fromAccount, err := u.points.UnfreezePoints(ctx, &EconomyFreezePointsReq{UserID: req.FromUserID, Amount: req.Amount})
		if err != nil {
			return err
		}
		response.FromAccount = fromAccount
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *TCCPointsTransferUsecase) validate(req *TCCPointsTransferReq) error {
	if req == nil || req.FromUserID <= 0 || req.ToUserID <= 0 || req.Amount <= 0 || !req.OutRecordType.IsExpense() || !req.InRecordType.IsIncome() {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	if req.FromUserID == req.ToUserID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_TRANSFER_SELF_NOT_ALLOWED)
	}
	return nil
}
