package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"strings"
)

const chatMessageMaxRuneLen = 500

type ChatUsecase struct {
	characterRepo     repo.CharacterRepo
	chatMessageRepo   repo.ChatMessageRepo
	gameIdleEventRepo repo.GameIdleEventRepo
}

func NewChatUsecase(
	characterRepo repo.CharacterRepo,
	chatMessageRepo repo.ChatMessageRepo,
	gameIdleEventRepo repo.GameIdleEventRepo,
) *ChatUsecase {
	return &ChatUsecase{
		characterRepo:     characterRepo,
		chatMessageRepo:   chatMessageRepo,
		gameIdleEventRepo: gameIdleEventRepo,
	}
}

type SendMessageReq struct {
	CharacterID         int64
	ChannelType         enum.ChatChannelType
	ChannelID           string
	ReceiverCharacterID *int64
	Content             string
}

func (u *ChatUsecase) Send(ctx context.Context, req *SendMessageReq) (*model.ChatMessage, error) {
	content := strings.TrimSpace(req.Content)
	if content == "" || len([]rune(content)) > chatMessageMaxRuneLen {
		return nil, model.ErrChatMessageInvalid
	}
	if _, err := u.characterRepo.Get(ctx, req.CharacterID); err != nil {
		return nil, err
	}
	channelType := req.ChannelType
	channelID := strings.TrimSpace(req.ChannelID)
	if channelType == "" || channelID == "" {
		return nil, model.ErrChatMessageInvalid
	}
	message, err := u.chatMessageRepo.Create(ctx, &model.ChatMessage{
		ChannelType:         channelType,
		ChannelID:           channelID,
		SenderCharacterID:   req.CharacterID,
		ReceiverCharacterID: req.ReceiverCharacterID,
		Content:             content,
		Status:              enum.ChatMessageStatusNormal,
	})
	if err != nil {
		return nil, err
	}
	return message, u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
		ChatMessage: message,
	})
}

type ListMessagesReq struct {
	ChannelType enum.ChatChannelType
	ChannelID   string
	BeforeID    int64
	Size        int
}

func (u *ChatUsecase) List(ctx context.Context, req *ListMessagesReq) ([]*model.ChatMessage, error) {
	size := req.Size
	if size <= 0 || size > 100 {
		size = 50
	}
	channelType := req.ChannelType
	channelID := strings.TrimSpace(req.ChannelID)
	if channelType == "" || channelID == "" {
		return nil, model.ErrChatMessageInvalid
	}
	return u.chatMessageRepo.List(ctx, &repo.ChatMessageListReq{
		ChannelType: channelType,
		ChannelID:   channelID,
		BeforeID:    req.BeforeID,
		Size:        size,
	})
}
