package domain

import (
	"context"
	domainbase "im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/biz/repo"
)

type ChatSessionDomain struct {
	*domainbase.BaseDomain
	chatSessionRepo repo.ChatSessionRepo
}

func NewChatSessionDomain(base *domainbase.BaseDomain, chatSessionRepo repo.ChatSessionRepo) (*ChatSessionDomain, error) {
	return &ChatSessionDomain{
		BaseDomain:      base,
		chatSessionRepo: chatSessionRepo,
	}, nil
}

func (s *ChatSessionDomain) MarkMuted(ctx context.Context, id int64, disturb bool) (*model.ChatSession, error) {
	return s.chatSessionRepo.UpdateMuted(ctx, s.Db, id, disturb)
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
