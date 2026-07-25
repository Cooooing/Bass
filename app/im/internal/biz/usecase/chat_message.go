package usecase

import (
	"context"
	"log/slog"

	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"
)

type ChatMessageUsecase struct {
	chatMessageRepo     repo.ChatMessageRepo
	chatSessionRepo     repo.ChatSessionRepo
	chatGroupRepo       repo.ChatGroupRepo
	chatGroupMemberRepo repo.ChatGroupMemberRepo
	log                 *slog.Logger
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
		log:                 logger,
	}, nil
}

type SendReq struct {
	SenderID     int64
	ReceiverID   int64
	ReceiverType enum.ReceiverType
	MessageType  enum.MessageType
	Content      string
}

func (u *ChatMessageUsecase) Send(ctx context.Context, req *SendReq) error {
	switch req.ReceiverType {
	case enum.ReceiverTypeUser:
		return u.sendToUser(ctx, &sendToUserReq{
			SenderID:    req.SenderID,
			ReceiverID:  req.ReceiverID,
			MessageType: req.MessageType,
			Content:     req.Content,
		})
	case enum.ReceiverTypeGroup:
		return u.sendToGroup(ctx, &sendToGroupReq{
			SenderID:    req.SenderID,
			GroupID:     req.ReceiverID,
			MessageType: req.MessageType,
			Content:     req.Content,
		})
	default:
		return nil
	}
}

type sendToUserReq struct {
	SenderID    int64
	ReceiverID  int64
	MessageType enum.MessageType
	Content     string
}

func (u *ChatMessageUsecase) sendToUser(ctx context.Context, req *sendToUserReq) error {
	senderID := req.SenderID
	receiverID := req.ReceiverID
	msgType := req.MessageType
	content := req.Content
	session, err := u.findOrCreateSession(ctx, &findOrCreateSessionReq{
		UserID:     senderID,
		ReceiverID: receiverID,
	})
	if err != nil {
		return err
	}
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
	_, err = u.chatSessionRepo.UpdateLastReadMessage(ctx, &repo.ChatSessionUpdateLastReadMessageReq{
		ChatSessionID:      session.ID,
		MessageID:          msg.ID,
		OperationReadCount: 0,
		UpdatedBy:          senderID,
	})
	return err
}

type sendToGroupReq struct {
	SenderID    int64
	GroupID     int64
	MessageType enum.MessageType
	Content     string
}

func (u *ChatMessageUsecase) sendToGroup(ctx context.Context, req *sendToGroupReq) error {
	senderID := req.SenderID
	groupID := req.GroupID
	msgType := req.MessageType
	content := req.Content
	if _, err := u.chatGroupMemberRepo.Get(ctx, &repo.ChatGroupMemberQuery{
		GroupID: &groupID,
		UserID:  &senderID,
	}); err != nil {
		return err
	}
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
	_, err = u.chatGroupRepo.UpdateLastMessage(ctx, &repo.ChatGroupUpdateLastMessageReq{
		Message:   msg,
		UpdatedBy: senderID,
	})
	return err
}

type findOrCreateSessionReq struct {
	UserID     int64
	ReceiverID int64
}

func (u *ChatMessageUsecase) findOrCreateSession(ctx context.Context, req *findOrCreateSessionReq) (*model.ChatSession, error) {
	userID := req.UserID
	receiverID := req.ReceiverID
	listResp, err := u.chatSessionRepo.List(ctx, &repo.ChatSessionQuery{
		CreatedBy:  &userID,
		ReceiverID: &receiverID,
	})
	if err == nil && len(listResp) > 0 {
		return listResp[0], nil
	}
	listResp, err = u.chatSessionRepo.List(ctx, &repo.ChatSessionQuery{
		CreatedBy:  &receiverID,
		ReceiverID: &userID,
	})
	if err == nil && len(listResp) > 0 {
		return listResp[0], nil
	}
	session, err := u.chatSessionRepo.Save(ctx, &model.ChatSession{
		ReceiverID: &receiverID,
		CreatedBy:  &userID,
		UpdatedBy:  &userID,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}

type RevokeReq struct {
	MessageID int64
	SenderID  int64
}

func (u *ChatMessageUsecase) Revoke(ctx context.Context, req *RevokeReq) error {
	msg, err := u.chatMessageRepo.Get(ctx, &repo.ChatMessageQuery{
		IDs: []int64{req.MessageID},
	})
	if err != nil {
		return err
	}
	if msg.SenderID != req.SenderID {
		return nil
	}
	_, err = u.chatMessageRepo.UpdateStatus(ctx, &repo.ChatMessageUpdateStatusReq{
		ChatMessageID: msg.ID,
		Status:        enum.MessageStatusRevoke,
		UpdatedBy:     req.SenderID,
	})
	return err
}

type ChatMessageListReq struct {
	Page      *base.PageRequest
	IDs       []int64
	SessionID *int64
	SenderID  *int64
}

type ChatMessageListResp struct {
	List []*model.ChatMessage
	Page *base.PageResp
}

func (u *ChatMessageUsecase) List(ctx context.Context, req *ChatMessageListReq) (*ChatMessageListResp, error) {
	pageResp, err := u.chatMessageRepo.Page(ctx, &repo.ChatMessageQuery{
		Page:      req.Page,
		IDs:       req.IDs,
		SessionID: req.SessionID,
		SenderID:  req.SenderID,
	})
	if err != nil {
		return nil, err
	}
	return &ChatMessageListResp{
		List: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}
