package usecase

import (
	"context"
	"log/slog"
	"time"

	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"

	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"
)

type GlobalChatUsecase struct {
	log      *slog.Logger
	messages repo.GlobalChatMessageRepo
	mutes    repo.GlobalChatMuteRepo
	events   repo.RealtimeEventPublisher
}

func NewGlobalChatUsecase(log *slog.Logger, messages repo.GlobalChatMessageRepo, mutes repo.GlobalChatMuteRepo, events repo.RealtimeEventPublisher) *GlobalChatUsecase {
	return &GlobalChatUsecase{log: log, messages: messages, mutes: mutes, events: events}
}

func (u *GlobalChatUsecase) Send(ctx context.Context, senderID int64, clientMessageID, content string) (*model.GlobalChatMessage, error) {
	if senderID <= 0 || clientMessageID == "" || content == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	mute, err := u.mutes.Get(ctx, senderID)
	if err != nil {
		return nil, err
	}
	if mute != nil && mute.MutedUntil.After(time.Now()) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RBAC_PERMISSION_DENIED)
	}
	existing, err := u.messages.GetByClientMessageID(ctx, senderID, clientMessageID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}
	message, err := u.messages.Save(ctx, &model.GlobalChatMessage{SenderID: senderID, ClientMessageID: clientMessageID, Content: content, Type: enum.MessageTypeText})
	if err != nil {
		return nil, err
	}
	if err = u.events.PublishGlobalChatMessageCreated(ctx, message); err != nil {
		u.log.WarnContext(ctx, "publish global chat message event failed", "message_id", message.ID, "error", err)
	}
	return message, nil
}

func (u *GlobalChatUsecase) Revoke(ctx context.Context, userID, messageID int64, administrator bool) (*model.GlobalChatMessage, error) {
	message, err := u.messages.Get(ctx, messageID)
	if err != nil {
		return nil, err
	}
	if message == nil || message.Revoked {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if !administrator && (message.SenderID != userID || message.CreatedAt == nil || time.Since(*message.CreatedAt) > 2*time.Minute) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RBAC_PERMISSION_DENIED)
	}
	return u.messages.Revoke(ctx, messageID, userID)
}

func (u *GlobalChatUsecase) List(ctx context.Context, beforeID int64, page *commonmodel.PageReq) ([]*model.GlobalChatMessage, *commonmodel.PageResp, error) {
	p := commonmodel.NormalizePage(page)
	rows, total, err := u.messages.List(ctx, beforeID, p.Limit(), p.Offset())
	if err != nil {
		return nil, nil, err
	}
	return rows, &commonmodel.PageResp{Page: p.Page, Size: p.Size, Total: total}, nil
}

func (u *GlobalChatUsecase) Mute(ctx context.Context, operatorUserID, userID int64, duration time.Duration, reason *string) error {
	if operatorUserID <= 0 || userID <= 0 || duration <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.mutes.Save(ctx, &model.GlobalChatMute{UserID: userID, MutedUntil: time.Now().Add(duration), Reason: reason, OperatorUserID: operatorUserID})
}

func (u *GlobalChatUsecase) Unmute(ctx context.Context, userID int64) error {
	return u.mutes.Delete(ctx, userID)
}
