package usecase

import (
	"context"
	"time"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"

	"user/internal/biz/repo"
	"user/internal/enum"
)

// TCCLivenessRewardUsecase 执行活跃奖励领取的本地 TCC 分支。
type TCCLivenessRewardUsecase struct {
	records repo.LivenessRecordRepo
	branch  commonclient.TCCBranchExecutor
}

// NewTCCLivenessRewardUsecase 创建活跃奖励 TCC 分支用例。
func NewTCCLivenessRewardUsecase(records repo.LivenessRecordRepo, branch commonclient.TCCBranchExecutor) *TCCLivenessRewardUsecase {
	return &TCCLivenessRewardUsecase{records: records, branch: branch}
}

// Try 将未领取奖励预留为领取中。
func (u *TCCLivenessRewardUsecase) Try(ctx context.Context, req *userv1.TCCLivenessReward_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if req.GetUserId() <= 0 || req.GetDate() == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		row, err := u.records.Get(ctx, req.GetUserId(), req.GetDate().AsTime())
		if err != nil {
			return err
		}
		if row == nil || row.Score <= 0 || (row.RewardStatus != enum.LivenessRewardStatusUnclaimed && row.RewardStatus != enum.LivenessRewardStatusCanceled) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		_, err = u.records.UpdateRewardStatus(ctx, row.ID, enum.LivenessRewardStatusPending, nil)
		return err
	})
}

// Confirm 将预留奖励标记为已领取。
func (u *TCCLivenessRewardUsecase) Confirm(ctx context.Context, req *userv1.TCCLivenessReward_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		row, err := u.records.Get(ctx, req.GetUserId(), req.GetDate().AsTime())
		if err != nil {
			return err
		}
		if row == nil || row.RewardStatus == enum.LivenessRewardStatusCanceled {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		if row.RewardStatus == enum.LivenessRewardStatusCompleted {
			return nil
		}
		now := time.Now()
		_, err = u.records.UpdateRewardStatus(ctx, row.ID, enum.LivenessRewardStatusCompleted, &now)
		return err
	})
}

// Cancel 释放领取中状态，以便后续请求重新领取。
func (u *TCCLivenessRewardUsecase) Cancel(ctx context.Context, req *userv1.TCCLivenessReward_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		row, err := u.records.Get(ctx, req.GetUserId(), req.GetDate().AsTime())
		if err != nil {
			return err
		}
		if row == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		if row.RewardStatus == enum.LivenessRewardStatusCanceled || row.RewardStatus == enum.LivenessRewardStatusCompleted {
			return nil
		}
		_, err = u.records.UpdateRewardStatus(ctx, row.ID, enum.LivenessRewardStatusCanceled, nil)
		return err
	})
}
