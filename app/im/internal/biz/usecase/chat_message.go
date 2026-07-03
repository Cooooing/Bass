package usecase

import (
	"context"

	"common/proto/gen/common"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"

	"log/slog"
)

type ChatMessageUsecase struct {
	chatMessageRepo     repo.ChatMessageRepo
	chatSessionRepo     repo.ChatSessionRepo
	chatGroupRepo       repo.ChatGroupRepo
	chatGroupMemberRepo repo.ChatGroupMemberRepo
	log                 *slog.Logger
}

func NewChatMessageUsecase(chatMessageRepo repo.ChatMessageRepo, chatSessionRepo repo.ChatSessionRepo, chatGroupRepo repo.ChatGroupRepo, chatGroupMemberRepo repo.ChatGroupMemberRepo, logger *slog.Logger) (*ChatMessageUsecase, error) {
	return &ChatMessageUsecase{chatMessageRepo: chatMessageRepo, chatSessionRepo: chatSessionRepo, chatGroupRepo: chatGroupRepo, chatGroupMemberRepo: chatGroupMemberRepo, log: logger}, nil
}

func (u *ChatMessageUsecase) Send(ctx context.Context, senderID int64, receiverID int64, receiverType enum.ReceiverType, msgType enum.MessageType, content string) error {
	switch receiverType {
	case enum.ReceiverTypeUser:
		return u.sendToUser(ctx, senderID, receiverID, msgType, content)
	case enum.ReceiverTypeGroup:
		return u.sendToGroup(ctx, senderID, receiverID, msgType, content)
	default:
		return nil
	}
}

func (u *ChatMessageUsecase) sendToUser(ctx context.Context, senderID int64, receiverID int64, msgType enum.MessageType, content string) error {
	session, err := u.findOrCreateSession(ctx, senderID, receiverID, 0)
	if err != nil {
		return err
	}
	msg, err := u.chatMessageRepo.Save(ctx, &model.ChatMessage{SenderID: senderID, ReceiverID: &receiverID, SessionID: &session.ID, Type: msgType, Content: content, Status: enum.MessageStatusNormal, CreatedBy: &senderID, UpdatedBy: &senderID})
	if err != nil {
		return err
	}
	_, err = u.chatSessionRepo.UpdateLastReadMessage(ctx, session.ID, msg.ID, 0, senderID)
	return err
}

func (u *ChatMessageUsecase) sendToGroup(ctx context.Context, senderID int64, groupID int64, msgType enum.MessageType, content string) error {
	if _, err := u.chatGroupMemberRepo.Get(ctx, &repo.ChatGroupMemberGetReq{GroupID: &groupID, UserID: &senderID}); err != nil {
		return err
	}
	msg, err := u.chatMessageRepo.Save(ctx, &model.ChatMessage{SenderID: senderID, GroupID: &groupID, Type: msgType, Content: content, Status: enum.MessageStatusNormal, CreatedBy: &senderID, UpdatedBy: &senderID})
	if err != nil {
		return err
	}
	_, err = u.chatGroupRepo.UpdateLastMessage(ctx, msg, senderID)
	return err
}

func (u *ChatMessageUsecase) findOrCreateSession(ctx context.Context, userID int64, receiverID int64, _ int64) (*model.ChatSession, error) {
	list, err := u.chatSessionRepo.List(ctx, &repo.ChatSessionGetReq{CreatedBy: &userID, ReceiverID: &receiverID})
	if err == nil && len(list) > 0 {
		return list[0], nil
	}
	list, err = u.chatSessionRepo.List(ctx, &repo.ChatSessionGetReq{CreatedBy: &receiverID, ReceiverID: &userID})
	if err == nil && len(list) > 0 {
		return list[0], nil
	}
	return u.chatSessionRepo.Save(ctx, &model.ChatSession{ReceiverID: &receiverID, CreatedBy: &userID, UpdatedBy: &userID})
}

func (u *ChatMessageUsecase) Revoke(ctx context.Context, messageID int64, senderID int64) error {
	msg, err := u.chatMessageRepo.Get(ctx, &repo.ChatMessageGetReq{IDs: []int64{messageID}})
	if err != nil {
		return err
	}
	if msg.SenderID != senderID {
		return nil
	}
	_, err = u.chatMessageRepo.UpdateStatus(ctx, msg.ID, enum.MessageStatusRevoke, senderID)
	return err
}

func (u *ChatMessageUsecase) List(ctx context.Context, page *common.PageRequest, ids []int64, sessionID *int64, senderID *int64) ([]*model.ChatMessage, *common.PageReply, error) {
	return u.chatMessageRepo.Page(ctx, page, &repo.ChatMessageGetReq{IDs: ids, SessionID: sessionID, SenderID: senderID})
}
