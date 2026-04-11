package service

import (
	v1 "common/api/im/v1"
	"context"
	"im/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ChatSessionService struct {
	v1.UnimplementedIMChatSessionServiceServer
	*BaseService
	chatSessionDomain *domain.ChatSessionDomain
}

func NewChatSessionService(baseService *BaseService, chatSessionDomain *domain.ChatSessionDomain) *ChatSessionService {
	return &ChatSessionService{
		BaseService:       baseService,
		chatSessionDomain: chatSessionDomain,
	}
}

func (s *ChatSessionService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatSessionServiceServer(gs, s)
}

func (s *ChatSessionService) RegisterHttp(hs *http.Server) {
	v1.RegisterIMChatSessionServiceHTTPServer(hs, s)
}

func (s *ChatSessionService) MarkMuted(ctx context.Context, req *v1.ChatSessionMarkMutedRequest) (rsp *v1.ChatSessionMarkMutedReply, err error) {
	// 最小实现：直接返回空的应答
	return &v1.ChatSessionMarkMutedReply{}, nil
}

func (s *ChatSessionService) MarkRead(ctx context.Context, req *v1.ChatSessionMarkReadRequest) (rsp *v1.ChatSessionMarkReadReply, err error) {
	// 最小实现：直接返回空的应答
	return &v1.ChatSessionMarkReadReply{}, nil
}

func (s *ChatSessionService) MarkPinned(ctx context.Context, req *v1.ChatSessionMarkPinnedRequest) (rsp *v1.ChatSessionMarkPinnedReply, err error) {
	// 最小实现：直接返回空的应答
	return &v1.ChatSessionMarkPinnedReply{}, nil
}

func (s *ChatSessionService) Page(ctx context.Context, req *v1.ChatSessionPageRequest) (rsp *v1.ChatSessionPageReply, err error) {
	// 占位实现：返回空的分页结果，Page 字段暂不暴露具体值
	return &v1.ChatSessionPageReply{Page: nil, Rows: []*v1.ChatSession{}}, nil
}
