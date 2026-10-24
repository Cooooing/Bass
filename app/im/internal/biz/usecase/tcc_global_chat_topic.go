package usecase

import (
	"context"
	"time"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"

	"im/internal/biz/model"
	"im/internal/biz/repo"
)

// TCCGlobalChatTopicUsecase 执行全局聊天室话题设置的 TCC 分支操作。
type TCCGlobalChatTopicUsecase struct {
	topics repo.GlobalChatTopicRepo
	mutes  repo.GlobalChatMuteRepo
}

func NewTCCGlobalChatTopicUsecase(topics repo.GlobalChatTopicRepo, mutes repo.GlobalChatMuteRepo) *TCCGlobalChatTopicUsecase {
	return &TCCGlobalChatTopicUsecase{topics: topics, mutes: mutes}
}

func (u *TCCGlobalChatTopicUsecase) Try(ctx context.Context, userID int64, clientRequestID, content string, pointsCost int64) error {
	if userID <= 0 || clientRequestID == "" || content == "" || pointsCost <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	mute, err := u.mutes.Get(ctx, userID)
	if err != nil {
		return err
	}
	if mute != nil && mute.MutedUntil.After(time.Now()) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RBAC_PERMISSION_DENIED)
	}
	existing, err := u.topics.GetByClientRequestID(ctx, clientRequestID)
	if err != nil {
		return err
	}
	if existing != nil {
		if existing.UserID != userID || existing.Content != content || existing.PointsCost != pointsCost {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		return nil
	}
	_, err = u.topics.CreatePending(ctx, &model.GlobalChatTopic{UserID: userID, ClientRequestID: clientRequestID, Content: content, PointsCost: pointsCost})
	return err
}

func (u *TCCGlobalChatTopicUsecase) Confirm(ctx context.Context, clientRequestID string) error {
	if clientRequestID == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := u.topics.Publish(ctx, clientRequestID)
	return err
}

func (u *TCCGlobalChatTopicUsecase) Cancel(ctx context.Context, clientRequestID string) error {
	if clientRequestID == "" {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.topics.Cancel(ctx, clientRequestID)
}
