package usecase

import (
	"common/pkg/util"
	"common/proto/gen/common"
	"context"
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
	log                 *util.LogHelper
}

func NewChatMessageUsecase(
	chatMessageRepo repo.ChatMessageRepo,
	chatSessionRepo repo.ChatSessionRepo,
	chatGroupRepo repo.ChatGroupRepo,
	chatGroupMemberRepo repo.ChatGroupMemberRepo,
	logger *slog.Logger,
) (*ChatMessageUsecase, error) {
	return &ChatMessageUsecase{
		chatMessageRepo:     chatMessageRepo,
		chatSessionRepo:     chatSessionRepo,
		chatGroupRepo:       chatGroupRepo,
		chatGroupMemberRepo: chatGroupMemberRepo,
		log:                 util.NewLogHelper(logger),
	}, nil
}

// Send 发送消息。
// USER 类型：查找或创建会话，保存消息，更新会话进度。
// GROUP 类型：验证群成员身份，保存消息，更新群组进度和所有成员会话。
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

// sendToUser 发送私聊消息。
func (u *ChatMessageUsecase) sendToUser(ctx context.Context, senderID int64, receiverID int64, msgType enum.MessageType, content string) error {
	// 查找或创建会话（两个用户之间只有一个会话）
	session, err := u.findOrCreateSession(ctx, senderID, receiverID, 0)
	if err != nil {
		return err
	}
	// 保存消息
	msg, err := u.chatMessageRepo.Save(ctx, &model.ChatMessage{
		SenderID:   senderID,
		ReceiverID: &receiverID,
		SessionID:  &session.ID,
		Type:       msgType,
		Content:    content,
		Status:     enum.MessageStatusNormal,
		CreatedBy:  &senderID,
		UpdatedBy:  &senderID,
	})
	if err != nil {
		return err
	}
	// 更新会话的最后消息和消息计数
	_, err = u.chatSessionRepo.UpdateLastReadMessage(ctx, session.ID, msg.ID, 0, senderID)
	if err != nil {
		return err
	}
	return nil
}

// sendToGroup 发送群聊消息。
func (u *ChatMessageUsecase) sendToGroup(ctx context.Context, senderID int64, groupID int64, msgType enum.MessageType, content string) error {
	// 验证发送者是群成员
	member, err := u.chatGroupMemberRepo.Get(ctx, &repo.ChatGroupMemberGetReq{
		GroupID: &groupID,
		UserID:  &senderID,
	})
	if err != nil {
		return err
	}
	_ = member // 验证通过
	// 保存消息
	msg, err := u.chatMessageRepo.Save(ctx, &model.ChatMessage{
		SenderID:  senderID,
		GroupID:   &groupID,
		Type:      msgType,
		Content:   content,
		Status:    enum.MessageStatusNormal,
		CreatedBy: &senderID,
		UpdatedBy: &senderID,
	})
	if err != nil {
		return err
	}
	// 更新群组的最后消息和消息计数
	_, err = u.chatGroupRepo.UpdateLastMessage(ctx, msg, senderID)
	if err != nil {
		return err
	}
	return nil
}

// findOrCreateSession 查找或创建私聊会话。
func (u *ChatMessageUsecase) findOrCreateSession(ctx context.Context, userID int64, receiverID int64, _ int64) (*model.ChatSession, error) {
	// 先尝试查找已有会话（由当前用户发起的、面向该接收者的私聊会话）
	list, err := u.chatSessionRepo.List(ctx, &repo.ChatSessionGetReq{
		CreatedBy:  &userID,
		ReceiverID: &receiverID,
	})
	if err == nil && len(list) > 0 {
		return list[0], nil
	}
	// 也尝试查找对方发起的会话（接收者视角的私聊会话）
	list, err = u.chatSessionRepo.List(ctx, &repo.ChatSessionGetReq{
		CreatedBy:  &receiverID,
		ReceiverID: &userID,
	})
	if err == nil && len(list) > 0 {
		return list[0], nil
	}
	// 不存在则创建新会话
	return u.chatSessionRepo.Save(ctx, &model.ChatSession{
		ReceiverID: &receiverID,
		CreatedBy:  &userID,
		UpdatedBy:  &userID,
	})
}

// Revoke 撤回消息，仅发送者可操作。
func (u *ChatMessageUsecase) Revoke(ctx context.Context, messageID int64, senderID int64) error {
	msg, err := u.chatMessageRepo.Get(ctx, &repo.ChatMessageGetReq{
		IDs: []int64{messageID},
	})
	if err != nil {
		return err
	}
	if msg.SenderID != senderID {
		return nil // 无权撤回，静默返回
	}
	_, err = u.chatMessageRepo.UpdateStatus(ctx, msg.ID, enum.MessageStatusRevoke, senderID)
	return err
}

// List 查询消息列表。
func (u *ChatMessageUsecase) List(ctx context.Context, page *common.PageRequest, ids []int64, sessionID *int64, senderID *int64) ([]*model.ChatMessage, *common.PageReply, error) {
	getReq := &repo.ChatMessageGetReq{
		IDs:       ids,
		SessionID: sessionID,
		SenderID:  senderID,
	}
	return u.chatMessageRepo.Page(ctx, page, getReq)
}
