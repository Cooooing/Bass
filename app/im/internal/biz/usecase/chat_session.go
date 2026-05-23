package usecase

import (
	"context"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/data/gen"
)

type ChatSessionUsecase struct {
	db              *gen.Client
	chatSessionRepo repo.ChatSessionRepo
}

func NewChatSessionUsecase(
	db *gen.Client,
	chatSessionRepo repo.ChatSessionRepo,
) (*ChatSessionUsecase, error) {
	return &ChatSessionUsecase{
		db:              db,
		chatSessionRepo: chatSessionRepo,
	}, nil
}

func (s *ChatSessionUsecase) MarkMuted(ctx context.Context, id int64, disturb bool) (*model.ChatSession, error) {
	return s.chatSessionRepo.UpdateMuted(ctx, s.db, id, disturb)
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
