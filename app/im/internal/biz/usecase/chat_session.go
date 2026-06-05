package usecase

import (
	"context"
	"im/internal/biz/model"
	"im/internal/biz/repo"
)

type ChatSessionUsecase struct {
	chatSessionRepo repo.ChatSessionRepo
}

func NewChatSessionUsecase(
	chatSessionRepo repo.ChatSessionRepo,
) (*ChatSessionUsecase, error) {
	return &ChatSessionUsecase{
		chatSessionRepo: chatSessionRepo,
	}, nil
}

func (s *ChatSessionUsecase) MarkMuted(ctx context.Context, id int64, disturb bool) (*model.ChatSession, error) {
	return s.chatSessionRepo.UpdateMuted(ctx, id, disturb)
}

func (s *ChatSessionUsecase) MarkRead(ctx context.Context) error {
	return nil
}

func (s *ChatSessionUsecase) MarkPinned(ctx context.Context) error {
	return nil
}

func (s *ChatSessionUsecase) Page(ctx context.Context) error {
	return nil
}
