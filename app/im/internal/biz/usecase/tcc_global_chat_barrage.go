package usecase

import (
	"context"
	"time"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"

	"im/internal/biz/model"
	"im/internal/biz/repo"
)

// TCCGlobalChatBarrageUsecase 执行弹幕创建的 TCC 本地分支。
type TCCGlobalChatBarrageUsecase struct {
	barrages  repo.GlobalChatBarrageRepo
	mutes     repo.GlobalChatMuteRepo
	tccBranch commonclient.TCCBranchExecutor
}

func NewTCCGlobalChatBarrageUsecase(barrages repo.GlobalChatBarrageRepo, mutes repo.GlobalChatMuteRepo, tccBranch commonclient.TCCBranchExecutor) *TCCGlobalChatBarrageUsecase {
	return &TCCGlobalChatBarrageUsecase{barrages: barrages, mutes: mutes, tccBranch: tccBranch}
}

func (u *TCCGlobalChatBarrageUsecase) Try(ctx context.Context, userID int64, clientRequestID, content string, pointsCost int64) error {
	if userID <= 0 || clientRequestID == "" || content == "" || pointsCost <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		mute, err := u.mutes.Get(ctx, userID)
		if err != nil {
			return err
		}
		if mute != nil && mute.MutedUntil.After(time.Now()) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RBAC_PERMISSION_DENIED)
		}
		existing, err := u.barrages.GetByClientRequestID(ctx, clientRequestID)
		if err != nil {
			return err
		}
		if existing != nil {
			if existing.UserID != userID || existing.Content != content || existing.PointsCost != pointsCost {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
			}
			return nil
		}
		_, err = u.barrages.CreatePending(ctx, &model.GlobalChatBarrage{
			UserID:          userID,
			ClientRequestID: clientRequestID,
			Content:         content,
			PointsCost:      pointsCost,
		})
		return err
	})
}

func (u *TCCGlobalChatBarrageUsecase) Confirm(ctx context.Context, clientRequestID string) error {
	if clientRequestID == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		_, err := u.barrages.Publish(ctx, clientRequestID)
		return err
	})
}

func (u *TCCGlobalChatBarrageUsecase) Cancel(ctx context.Context, clientRequestID string) error {
	if clientRequestID == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		return u.barrages.Cancel(ctx, clientRequestID)
	})
}
