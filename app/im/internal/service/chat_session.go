package service

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/base"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/im/v1"
	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ChatSessionService struct {
	v1.UnimplementedIMChatSessionServiceServer
	chatSessionUsecase *usecase.ChatSessionUsecase
}

func NewChatSessionService(
	chatSessionUsecase *usecase.ChatSessionUsecase,
) *ChatSessionService {
	return &ChatSessionService{
		chatSessionUsecase: chatSessionUsecase,
	}
}

func (s *ChatSessionService) RegisterGrpc(
	gs *grpc.Server,
) {
	v1.RegisterIMChatSessionServiceServer(gs, s)
}

func (s *ChatSessionService) RegisterHttp(
	hs *http.Server,
) {
}

// MarkMuted 设置免打扰状态。
func (s *ChatSessionService) MarkMuted(
	ctx context.Context,
	req *v1.MarkMutedChatSession_Req,
) (*v1.MarkMutedChatSession_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatSessionUsecase.MarkMuted(ctx, &usecase.MarkMutedReq{
		IDs:     req.GetIds(),
		Disturb: req.GetDisturb(),
		UserID:  req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.MarkMutedChatSession_Resp{}, nil
}

// MarkPinned 设置置顶状态。
func (s *ChatSessionService) MarkPinned(
	ctx context.Context,
	req *v1.MarkPinnedChatSession_Req,
) (*v1.MarkPinnedChatSession_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatSessionUsecase.MarkPinned(ctx, &usecase.MarkPinnedReq{
		IDs:    req.GetIds(),
		Top:    req.GetTop(),
		UserID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.MarkPinnedChatSession_Resp{}, nil
}

// MarkRead 标记已读。
func (s *ChatSessionService) MarkRead(
	ctx context.Context,
	req *v1.MarkReadChatSession_Req,
) (*v1.MarkReadChatSession_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatSessionUsecase.MarkRead(ctx, &usecase.MarkReadReq{
		IDs:    req.GetIds(),
		UserID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.MarkReadChatSession_Resp{}, nil
}

// List 查询会话列表。
func (s *ChatSessionService) List(
	ctx context.Context,
	req *v1.ListChatSessions_Req,
) (*v1.ListChatSessions_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var queryIDs []int64
	if req.GetQuery() != nil {
		queryIDs = req.GetQuery().GetIds()
	}
	resp, err := s.chatSessionUsecase.Page(ctx, &usecase.ChatSessionPageReq{
		Page: &base.PageRequest{
			Page: int64(req.GetPage().GetPage()),
			Size: int64(req.GetPage().GetSize()),
		},
		QueryIDs: queryIDs,
		UserID:   req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ListChatSessions_Resp_ChatSession, 0, len(resp.List))
	for _, item := range resp.List {
		row := &v1.ListChatSessions_Resp_ChatSession{
			Id:          item.ID,
			IsMuted:     item.IsMuted,
			IsPinned:    item.IsPinned,
			UnreadCount: item.UnreadCount,
		}
		if item.ReceiverID != nil {
			row.RelationId = *item.ReceiverID
		}
		if item.GroupID != nil {
			row.GroupId = *item.GroupID
		}
		if item.LastReadMessageID != nil {
			row.LastReadMessageId = *item.LastReadMessageID
		}
		rows = append(rows, row)
	}
	return &v1.ListChatSessions_Resp{
		Page: &common.PageResp{
			Page:  uint32(resp.Page.Page),
			Size:  uint32(resp.Page.Size),
			Total: uint32(resp.Page.Total),
		},
		Rows: rows,
	}, nil
}
