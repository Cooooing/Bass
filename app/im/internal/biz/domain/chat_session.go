package domain

import (
	"context"
	"im/internal/biz/model"
	"im/internal/biz/repo"
)

type ChatSessionDomain struct {
	db              *gen.Client
	chatSessionRepo repo.ChatSessionRepo
}

func NewChatSessionDomain(
	db *gen.Client,
	chatSessionRepo repo.ChatSessionRepo,
) (*ChatSessionDomain, error) {
	return &ChatSessionDomain{
		db:              db,
		chatSessionRepo: chatSessionRepo,
	}, nil
}

func (s *ChatSessionDomain) MarkMuted(ctx context.Context, id int64, disturb bool) (*model.ChatSession, error) {
	return s.chatSessionRepo.UpdateMuted(ctx, s.db, id, disturb)
}

func (s *ChatSessionDomain) MarkRead(ctx context.Context) error {
	return nil
}

func (s *ChatSessionDomain) MarkPinned(ctx context.Context) error {
	return nil
}

func (s *ChatSessionDomain) Page(ctx context.Context) error {
	return nil
}
