package usecase

import (
	"context"
	"log/slog"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"
)

type ChatGroupUsecase struct {
	chatGroupRepo       repo.ChatGroupRepo
	chatGroupMemberRepo repo.ChatGroupMemberRepo
	log                 *slog.Logger
}

func NewChatGroupUsecase(chatGroupRepo repo.ChatGroupRepo, chatGroupMemberRepo repo.ChatGroupMemberRepo, logger *slog.Logger) (*ChatGroupUsecase, error) {
	return &ChatGroupUsecase{chatGroupRepo: chatGroupRepo, chatGroupMemberRepo: chatGroupMemberRepo, log: logger}, nil
}

type CreateReq struct {
	Name         string
	Avatar       *string
	Introduction *string
	OwnerID      int64
}

type CreateResponse struct {
	GroupID int64
}

func (u *ChatGroupUsecase) Create(ctx context.Context, req *CreateReq) (*CreateResponse, error) {
	groupResp, err := u.chatGroupRepo.Save(ctx, &repo.ChatGroupSaveReq{ChatGroup: &model.ChatGroup{
		Name:         req.Name,
		Avatar:       req.Avatar,
		Introduction: req.Introduction,
		OwnerID:      req.OwnerID,
		MemberCount:  1,
		CreatedBy:    &req.OwnerID,
		UpdatedBy:    &req.OwnerID,
	}})
	if err != nil {
		return nil, err
	}
	group := groupResp.ChatGroup
	_, err = u.chatGroupMemberRepo.Save(ctx, &repo.ChatGroupMemberSaveReq{ChatGroupMember: &model.ChatGroupMember{
		GroupID:   group.ID,
		UserID:    req.OwnerID,
		Role:      enum.ChatGroupMemberRoleOwner,
		CreatedBy: &req.OwnerID,
		UpdatedBy: &req.OwnerID,
	}})
	if err != nil {
		return nil, err
	}
	return &CreateResponse{GroupID: group.ID}, nil
}

type DismissReq struct {
	GroupID    int64
	OperatorID int64
}

func (u *ChatGroupUsecase) Dismiss(ctx context.Context, req *DismissReq) error {
	groupResp, err := u.chatGroupRepo.Get(ctx, &repo.ChatGroupGetReq{ChatGroupQuery: repo.ChatGroupQuery{IDs: []int64{req.GroupID}}})
	if err != nil {
		return err
	}
	if groupResp.ChatGroup.OwnerID != req.OperatorID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_STATUS_INVALID)
	}
	_, err = u.chatGroupRepo.UpdateStatus(ctx, &repo.ChatGroupUpdateStatusReq{
		ChatGroupID: req.GroupID,
		Status:      enum.ChatGroupStatusDissolve,
		UpdatedBy:   req.OperatorID,
	})
	return err
}

type ChatGroupListReq struct {
	Page   *base.PageRequest
	IDs    []int64
	Status *enum.ChatGroupStatus
}

type ChatGroupListResponse struct {
	List []*model.ChatGroup
	Page *base.PageResponse
}

func (u *ChatGroupUsecase) List(ctx context.Context, req *ChatGroupListReq) (*ChatGroupListResponse, error) {
	pageResponse, err := u.chatGroupRepo.Page(ctx, &repo.ChatGroupPageReq{ChatGroupQuery: repo.ChatGroupQuery{Page: req.Page, IDs: req.IDs, Status: req.Status}})
	if err != nil {
		return nil, err
	}
	return &ChatGroupListResponse{List: pageResponse.Rows, Page: pageResponse.Page}, nil
}
