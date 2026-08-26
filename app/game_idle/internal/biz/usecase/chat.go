package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"strings"
)

const (
	ChatWorldChannelID    = "world"
	chatMessageMaxRuneLen = 500
)

type ChatUsecase struct {
	characterRepo        repo.CharacterRepo
	chatMessageRepo      repo.ChatMessageRepo
	chatMessageEventRepo repo.ChatMessageEventRepo
}

func NewChatUsecase(
	characterRepo repo.CharacterRepo,
	chatMessageRepo repo.ChatMessageRepo,
	chatMessageEventRepo repo.ChatMessageEventRepo,
) *ChatUsecase {
	return &ChatUsecase{
		characterRepo:        characterRepo,
		chatMessageRepo:      chatMessageRepo,
		chatMessageEventRepo: chatMessageEventRepo,
	}
}

type SendWorldMessageReq struct {
	CharacterID int64
	Content     string
}

func (u *ChatUsecase) SendWorldMessage(ctx context.Context, req *SendWorldMessageReq) (*model.ChatMessage, error) {
	content := strings.TrimSpace(req.Content)
	if content == "" || len([]rune(content)) > chatMessageMaxRuneLen {
		return nil, model.ErrChatMessageInvalid
	}
	character, err := u.characterRepo.Get(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	if character.Status != enum.CharacterStatusActive {
		return nil, model.ErrCharacterInvalid
	}
	message, err := u.chatMessageRepo.Create(ctx, &model.ChatMessage{
		ChannelType:       enum.ChatChannelTypeWorld,
		ChannelID:         ChatWorldChannelID,
		SenderCharacterID: req.CharacterID,
		Content:           content,
		Status:            enum.ChatMessageStatusNormal,
	})
	if err != nil {
		return nil, err
	}
	return message, u.chatMessageEventRepo.PublishWorldMessage(ctx, message)
}

type ListWorldMessagesReq struct {
	BeforeID int64
	Size     int
}

func (u *ChatUsecase) ListWorldMessages(ctx context.Context, req *ListWorldMessagesReq) ([]*model.ChatMessage, error) {
	size := req.Size
	if size <= 0 || size > 100 {
		size = 50
	}
	return u.chatMessageRepo.List(ctx, &repo.ChatMessageListReq{
		ChannelType: enum.ChatChannelTypeWorld,
		ChannelID:   ChatWorldChannelID,
		BeforeID:    req.BeforeID,
		Size:        size,
	})
}
