package service

import (
	v1 "common/api/gen/im/v1"
	"context"
	"im/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ChatSessionService struct {
	v1.UnimplementedIMChatSessionServiceServer
	chatSessionDomain *domain.ChatSessionDomain
}

func NewChatSessionService(chatSessionDomain *domain.ChatSessionDomain) *ChatSessionService {
	return &ChatSessionService{
		chatSessionDomain: chatSessionDomain,
	}
}

func (s *ChatSessionService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatSessionServiceServer(gs, s)
}

func (s *ChatSessionService) RegisterHttp(hs *http.Server) {
	v1.RegisterIMChatSessionServiceHTTPServer(hs, s)
}

func (s *ChatSessionService) MarkMuted(ctx context.Context, req *v1.MarkMutedChatSession_Request) (rsp *v1.MarkMutedChatSession_Reply, err error) {
	// 最小实现：直接返回空的应答
	return &v1.MarkMutedChatSession_Reply{}, nil
}

func (s *ChatSessionService) MarkRead(ctx context.Context, req *v1.MarkReadChatSession_Request) (rsp *v1.MarkReadChatSession_Reply, err error) {
	// 最小实现：直接返回空的应答
	return &v1.MarkReadChatSession_Reply{}, nil
}

func (s *ChatSessionService) MarkPinned(ctx context.Context, req *v1.MarkPinnedChatSession_Request) (rsp *v1.MarkPinnedChatSession_Reply, err error) {
	// 最小实现：直接返回空的应答
	return &v1.MarkPinnedChatSession_Reply{}, nil
}

func (s *ChatSessionService) Page(ctx context.Context, req *v1.PageChatSession_Request) (rsp *v1.PageChatSession_Reply, err error) {
	// 占位实现：返回空的分页结果，Page 字段暂不暴露具体值
	return &v1.PageChatSession_Reply{Page: nil, Rows: []*v1.ChatSession{}}, nil
}
