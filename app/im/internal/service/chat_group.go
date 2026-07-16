package service

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/base"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/im/v1"
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
func (s *ChatGroupService) Create(ctx context.Context, req *v1.CreateChatGroup_Request) (*v1.CreateChatGroup_Response, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := s.chatGroupUsecase.Create(ctx, &usecase.CreateReq{
		Name:         req.GetName(),
		Avatar:       req.Avatar,
		Introduction: req.Introduction,
		OwnerID:      req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateChatGroup_Response{}, nil
}

// Dismiss 解散群组。
func (s *ChatGroupService) Dismiss(ctx context.Context, req *v1.DismissChatGroup_Request) (*v1.DismissChatGroup_Response, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatGroupUsecase.Dismiss(ctx, &usecase.DismissReq{
		GroupID:    req.GetId(),
		OperatorID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.DismissChatGroup_Response{}, nil
}

// List 查询群组列表。
func (s *ChatGroupService) List(ctx context.Context, req *v1.ListChatGroups_Request) (*v1.ListChatGroups_Response, error) {
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
	resp, err := s.chatGroupUsecase.List(ctx, &usecase.ChatGroupListReq{
		Page:   &base.PageRequest{Page: int64(req.GetPage().GetPage()), Size: int64(req.GetPage().GetSize())},
		IDs:    ids,
		Status: status,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ListChatGroups_Response_ChatGroup, 0, len(resp.List))
	for _, item := range resp.List {
		status := enum.ChatGroupStatusMap.MustToProto(item.Status)
		rows = append(rows, &v1.ListChatGroups_Response_ChatGroup{
			Id:           item.ID,
			Name:         item.Name,
			Avatar:       item.Avatar,
			Introduction: item.Introduction,
			Status:       status,
			MemberCount:  item.MemberCount,
		})
	}
	return &v1.ListChatGroups_Response{
		Page: &common.PageResponse{Page: uint32(resp.Page.Page), Size: uint32(resp.Page.Size), Total: uint32(resp.Page.Total)},
		Rows: rows,
	}, nil
}
