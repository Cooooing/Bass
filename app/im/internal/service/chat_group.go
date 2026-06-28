package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/im/v1"
	"context"
	"im/internal/biz/usecase"
	"im/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ChatGroupService struct {
	v1.UnimplementedIMChatGroupServiceServer
	chatGroupUsecase *usecase.ChatGroupUsecase
}

func NewChatGroupService(chatGroupUsecase *usecase.ChatGroupUsecase) *ChatGroupService {
	return &ChatGroupService{
		chatGroupUsecase: chatGroupUsecase,
	}
}

func (s *ChatGroupService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatGroupServiceServer(gs, s)
}

func (s *ChatGroupService) RegisterHttp(hs *http.Server) {
}

// Create 创建群组。
func (s *ChatGroupService) Create(ctx context.Context, req *v1.CreateChatGroup_Request) (*v1.CreateChatGroup_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	groupID, err := s.chatGroupUsecase.Create(ctx, req.GetName(), req.Avatar, req.Introduction, req.GetUserId())
	if err != nil {
		return nil, err
	}
	_ = groupID
	return &v1.CreateChatGroup_Reply{}, nil
}

// Dismiss 解散群组。
func (s *ChatGroupService) Dismiss(ctx context.Context, req *v1.DismissChatGroup_Request) (*v1.DismissChatGroup_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatGroupUsecase.Dismiss(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.DismissChatGroup_Reply{}, nil
}

// List 查询群组列表。
func (s *ChatGroupService) List(ctx context.Context, req *v1.ListChatGroups_Request) (*v1.ListChatGroups_Reply, error) {
	var ids []int64
	var status *enum.ChatGroupStatus
	if req.GetQuery() != nil {
		ids = req.GetQuery().GetIds()
		if req.GetQuery().Status != nil {
			queryStatus, ok := enum.ChatGroupStatusMap.ToEnum(*req.GetQuery().Status)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_STATUS_INVALID)
			}
			status = &queryStatus
		}
	}
	list, page, err := s.chatGroupUsecase.List(ctx, req.GetPage(), ids, status)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ChatGroup, 0, len(list))
	for _, item := range list {
		rows = append(rows, toProtoChatGroup(item))
	}
	return &v1.ListChatGroups_Reply{
		Page: page,
		Rows: rows,
	}, nil
}
