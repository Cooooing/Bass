package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/im/v1"
	"context"
	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ChatSessionService struct {
	v1.UnimplementedIMChatSessionServiceServer
	chatSessionUsecase *usecase.ChatSessionUsecase
}

func NewChatSessionService(chatSessionUsecase *usecase.ChatSessionUsecase) *ChatSessionService {
	return &ChatSessionService{
		chatSessionUsecase: chatSessionUsecase,
	}
}

func (s *ChatSessionService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatSessionServiceServer(gs, s)
}

func (s *ChatSessionService) RegisterHttp(hs *http.Server) {
}

// MarkMuted 设置免打扰状态。
func (s *ChatSessionService) MarkMuted(ctx context.Context, req *v1.MarkMutedChatSession_Request) (*v1.MarkMutedChatSession_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatSessionUsecase.MarkMuted(ctx, req.GetIds(), req.GetDisturb(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.MarkMutedChatSession_Reply{}, nil
}

// MarkPinned 设置置顶状态。
func (s *ChatSessionService) MarkPinned(ctx context.Context, req *v1.MarkPinnedChatSession_Request) (*v1.MarkPinnedChatSession_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatSessionUsecase.MarkPinned(ctx, req.GetIds(), req.GetTop(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.MarkPinnedChatSession_Reply{}, nil
}

// MarkRead 标记已读。
func (s *ChatSessionService) MarkRead(ctx context.Context, req *v1.MarkReadChatSession_Request) (*v1.MarkReadChatSession_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatSessionUsecase.MarkRead(ctx, req.GetIds(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.MarkReadChatSession_Reply{}, nil
}

// List 查询会话列表。
func (s *ChatSessionService) List(ctx context.Context, req *v1.ListChatSessions_Request) (*v1.ListChatSessions_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var queryIds []int64
	if req.GetQuery() != nil {
		queryIds = req.GetQuery().GetIds()
	}
	list, page, err := s.chatSessionUsecase.Page(ctx, req.GetPage(), queryIds, req.GetUserId())
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ChatSession, 0, len(list))
	for _, item := range list {
		rows = append(rows, toProtoChatSession(item))
	}
	return &v1.ListChatSessions_Reply{
		Page: page,
		Rows: rows,
	}, nil
}
