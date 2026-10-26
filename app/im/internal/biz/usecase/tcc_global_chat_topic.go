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

// TCCGlobalChatTopicUsecase 执行全局聊天室话题设置的 TCC 分支操作。
type TCCGlobalChatTopicUsecase struct {
	topics    repo.GlobalChatTopicRepo
	mutes     repo.GlobalChatMuteRepo
	tccBranch commonclient.TCCBranchExecutor
}

func NewTCCGlobalChatTopicUsecase(topics repo.GlobalChatTopicRepo, mutes repo.GlobalChatMuteRepo, tccBranch commonclient.TCCBranchExecutor) *TCCGlobalChatTopicUsecase {
	return &TCCGlobalChatTopicUsecase{topics: topics, mutes: mutes, tccBranch: tccBranch}
}

func (u *TCCGlobalChatTopicUsecase) Try(ctx context.Context, topicID, userID int64, content string, pointsCost int64) error {
	if topicID <= 0 || userID <= 0 || content == "" || pointsCost <= 0 {
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
		existing, err := u.topics.Get(ctx, topicID)
		if err != nil {
			return err
		}
		if existing != nil {
			if existing.UserID != userID || existing.Content != content || existing.PointsCost != pointsCost {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
			}
			return nil
		}
		_, err = u.topics.CreatePending(ctx, &model.GlobalChatTopic{ID: topicID, UserID: userID, Content: content, PointsCost: pointsCost})
		return err
	})
}

func (u *TCCGlobalChatTopicUsecase) Confirm(ctx context.Context, topicID int64) error {
	if topicID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		_, err := u.topics.Publish(ctx, topicID)
		return err
	})
}

func (u *TCCGlobalChatTopicUsecase) Cancel(ctx context.Context, topicID int64) error {
	if topicID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.tccBranch.Execute(ctx, func(ctx context.Context) error {
		return u.topics.Cancel(ctx, topicID)
	})
}
