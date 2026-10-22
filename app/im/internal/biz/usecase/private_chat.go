package usecase

import (
	"context"
	"log/slog"
	"time"

	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"

	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"
)

type PrivateChatUsecase struct {
	log       *slog.Logger
	tx        base.Tx
	sessions  repo.PrivateChatSessionRepo
	messages  repo.PrivateChatMessageRepo
	relations repo.RelationClient
	events    repo.RealtimeEventPublisher
}

func NewPrivateChatUsecase(log *slog.Logger, tx base.Tx, sessions repo.PrivateChatSessionRepo, messages repo.PrivateChatMessageRepo, relations repo.RelationClient, events repo.RealtimeEventPublisher) *PrivateChatUsecase {
	return &PrivateChatUsecase{log: log, tx: tx, sessions: sessions, messages: messages, relations: relations, events: events}
}

func (u *PrivateChatUsecase) Send(ctx context.Context, senderID, receiverID int64, clientMessageID, content string) (*model.PrivateChatMessage, error) {
	if senderID <= 0 || receiverID <= 0 || senderID == receiverID || clientMessageID == "" || content == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := u.ensureVisible(ctx, senderID, receiverID); err != nil {
		return nil, err
	}
	var result *model.PrivateChatMessage
	created := false
	err := u.tx(ctx, func(txCtx context.Context) error {
		existing, err := u.messages.GetByClientMessageID(txCtx, senderID, clientMessageID)
		if err != nil {
			return err
		}
		if existing != nil {
			result = existing
			return nil
		}
		for _, item := range [][2]int64{{senderID, receiverID}, {receiverID, senderID}} {
			session, err := u.sessions.Get(txCtx, item[0], item[1])
			if err != nil {
				return err
			}
			if session == nil {
				if _, err = u.sessions.Save(txCtx, &model.PrivateChatSession{UserID: item[0], PeerUserID: item[1]}); err != nil {
					return err
				}
			}
		}
		result, err = u.messages.Save(txCtx, &model.PrivateChatMessage{SenderID: senderID, ReceiverID: receiverID, ClientMessageID: clientMessageID, Content: content, Type: enum.MessageTypeText})
		if err != nil {
			return err
		}
		if err = u.sessions.UpdateAfterSend(txCtx, senderID, receiverID, result.ID, false); err != nil {
			return err
		}
		if err = u.sessions.UpdateAfterSend(txCtx, receiverID, senderID, result.ID, true); err != nil {
			return err
		}
		created = true
		return nil
	})
	if err != nil || result == nil || !created {
		return result, err
	}
	if err = u.events.PublishPrivateChatMessageCreated(ctx, result); err != nil {
		u.log.WarnContext(ctx, "publish private chat message event failed", "message_id", result.ID, "error", err)
	}
	return result, nil
}

func (u *PrivateChatUsecase) Revoke(ctx context.Context, userID, messageID int64) (*model.PrivateChatMessage, error) {
	message, err := u.messages.Get(ctx, messageID)
	if err != nil {
		return nil, err
	}
	if message == nil || message.SenderID != userID || message.Revoked || message.CreatedAt == nil || time.Since(*message.CreatedAt) > 2*time.Minute {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err = u.ensureVisible(ctx, userID, message.ReceiverID); err != nil {
		return nil, err
	}
	return u.messages.Revoke(ctx, messageID, userID)
}

func (u *PrivateChatUsecase) ListSessions(ctx context.Context, userID int64, page *commonmodel.PageReq) ([]*model.PrivateChatSession, *commonmodel.PageResp, error) {
	p := commonmodel.NormalizePage(page)
	rows, total, err := u.sessions.List(ctx, userID, p.Limit(), p.Offset())
	if err != nil {
		return nil, nil, err
	}
	blocked, err := u.blocked(ctx, userID)
	if err != nil {
		return nil, nil, err
	}
	visible := make([]*model.PrivateChatSession, 0, len(rows))
	messageIDs := make([]int64, 0, len(rows))
	for _, row := range rows {
		if _, found := blocked[row.PeerUserID]; !found {
			visible = append(visible, row)
			if row.LastMessageID != nil {
				messageIDs = append(messageIDs, *row.LastMessageID)
			}
		}
	}
	messages, err := u.messages.Map(ctx, messageIDs)
	if err != nil {
		return nil, nil, err
	}
	for _, row := range visible {
		if row.LastMessageID != nil {
			row.LastMessage = messages[*row.LastMessageID]
		}
	}
	return visible, &commonmodel.PageResp{Page: p.Page, Size: p.Size, Total: total}, nil
}

func (u *PrivateChatUsecase) ListMessages(ctx context.Context, userID, peerUserID, beforeID int64, page *commonmodel.PageReq) ([]*model.PrivateChatMessage, *commonmodel.PageResp, error) {
	if err := u.ensureVisible(ctx, userID, peerUserID); err != nil {
		return nil, nil, err
	}
	p := commonmodel.NormalizePage(page)
	rows, total, err := u.messages.List(ctx, userID, peerUserID, beforeID, p.Limit(), p.Offset())
	if err != nil {
		return nil, nil, err
	}
	return rows, &commonmodel.PageResp{Page: p.Page, Size: p.Size, Total: total}, nil
}

func (u *PrivateChatUsecase) CountUnread(ctx context.Context, userID int64) (uint32, error) {
	if userID <= 0 {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.sessions.CountUnread(ctx, userID)
}

func (u *PrivateChatUsecase) MarkRead(ctx context.Context, userID, peerUserID, messageID int64) error {
	if err := u.ensureVisible(ctx, userID, peerUserID); err != nil {
		return err
	}
	message, err := u.messages.Get(ctx, messageID)
	if err != nil {
		return err
	}
	if message == nil || ((message.SenderID != userID || message.ReceiverID != peerUserID) && (message.SenderID != peerUserID || message.ReceiverID != userID)) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_SESSION_INVALID)
	}
	return u.sessions.MarkRead(ctx, userID, peerUserID, messageID)
}

func (u *PrivateChatUsecase) SetMuted(ctx context.Context, userID, peerUserID int64, muted bool) error {
	if err := u.ensureVisible(ctx, userID, peerUserID); err != nil {
		return err
	}
	return u.sessions.SetMuted(ctx, userID, peerUserID, muted)
}
func (u *PrivateChatUsecase) SetPinned(ctx context.Context, userID, peerUserID int64, pinned bool) error {
	if err := u.ensureVisible(ctx, userID, peerUserID); err != nil {
		return err
	}
	return u.sessions.SetPinned(ctx, userID, peerUserID, pinned)
}

func (u *PrivateChatUsecase) ensureVisible(ctx context.Context, userID, peerUserID int64) error {
	blocked, err := u.blocked(ctx, userID)
	if err != nil {
		return err
	}
	if _, found := blocked[peerUserID]; found {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RBAC_PERMISSION_DENIED)
	}
	return nil
}
func (u *PrivateChatUsecase) blocked(ctx context.Context, userID int64) (map[int64]struct{}, error) {
	rows, err := u.relations.BlockedCounterparts(ctx, userID)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]struct{}, len(rows))
	for _, id := range rows {
		result[id] = struct{}{}
	}
	return result, nil
}
