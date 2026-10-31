package usecase

import (
	"common/pkg/apperror"
	commonClient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"
	"context"

	"economy/internal/biz/model"
	"economy/internal/enum"
)

type TCCPointsUsecase struct {
	points    *PointsUsecase
	tccBranch commonClient.TCCBranchExecutor
}

func NewTCCPointsUsecase(points *PointsUsecase, tccBranch commonClient.TCCBranchExecutor) *TCCPointsUsecase {
	return &TCCPointsUsecase{points: points, tccBranch: tccBranch}
}

type TCCPointsReq struct {
	UserID, Amount int64
	RecordType     enum.EconomyRecordType
	Remark         *string
	IdempotencyKey *string
}

type TCCPointsResp struct {
	Account *model.Account
	Record  *model.Record
}

func (u *TCCPointsUsecase) Try(ctx context.Context, req *TCCPointsReq) (*TCCPointsResp, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	response := &TCCPointsResp{}
	err := u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		if req.RecordType.IsIncome() {
			return nil
		}
		account, err := u.points.FreezePoints(ctx, &EconomyFreezePointsReq{UserID: req.UserID, Amount: req.Amount})
		if err != nil {
			return err
		}
		response.Account = account
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *TCCPointsUsecase) Confirm(ctx context.Context, req *TCCPointsReq) (*TCCPointsResp, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	response := &TCCPointsResp{}
	err := u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		if req.RecordType.IsIncome() {
			result, err := u.points.AddPoints(ctx, &EconomyAddPointsReq{
				UserID:         req.UserID,
				Amount:         req.Amount,
				RecordType:     req.RecordType,
				Remark:         req.Remark,
				IdempotencyKey: req.IdempotencyKey,
			})
			if err != nil {
				return err
			}
			response.Account = result.Account
			response.Record = result.Record
			return nil
		}
		result, err := u.points.CommitFrozenPoints(ctx, &EconomyCommitFrozenPointsReq{
			UserID:         req.UserID,
			Amount:         req.Amount,
			RecordType:     req.RecordType,
			Remark:         req.Remark,
			IdempotencyKey: req.IdempotencyKey,
		})
		if err != nil {
			return err
		}
		response.Account = result.Account
		response.Record = result.Record
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *TCCPointsUsecase) Cancel(ctx context.Context, req *TCCPointsReq) (*TCCPointsResp, error) {
	if err := u.validate(req); err != nil {
		return nil, err
	}
	response := &TCCPointsResp{}
	err := u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		if req.RecordType.IsIncome() {
			return nil
		}
		account, err := u.points.UnfreezePoints(ctx, &EconomyFreezePointsReq{UserID: req.UserID, Amount: req.Amount})
		if err != nil {
			return err
		}
		response.Account = account
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u *TCCPointsUsecase) validate(req *TCCPointsReq) error {
	if req == nil || req.UserID <= 0 || req.Amount <= 0 || (!req.RecordType.IsIncome() && !req.RecordType.IsExpense()) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_ECONOMY_RECORD_TYPE_INVALID)
	}
	return nil
}
