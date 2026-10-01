package usecase

import (
	commonmodel "common/pkg/model"
	"context"
	"log/slog"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"
)

type ChatGroupUsecase struct {
	chatGroupRepo       repo.ChatGroupRepo
	chatGroupMemberRepo repo.ChatGroupMemberRepo
	assetClient         repo.AssetClient
	log                 *slog.Logger
}

func NewChatGroupUsecase(
	chatGroupRepo repo.ChatGroupRepo,
	chatGroupMemberRepo repo.ChatGroupMemberRepo,
	assetClient repo.AssetClient,
	logger *slog.Logger,
) (*ChatGroupUsecase, error) {
	return &ChatGroupUsecase{
		chatGroupRepo:       chatGroupRepo,
		chatGroupMemberRepo: chatGroupMemberRepo,
		assetClient:         assetClient,
		log:                 logger,
	}, nil
}

type CreateReq struct {
	Name          string
	AvatarAssetID *int64
	Introduction  *string
	OwnerID       int64
}

func (u *ChatGroupUsecase) Create(ctx context.Context, req *CreateReq) (int64, error) {
	if req == nil || req.OwnerID <= 0 {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.AvatarAssetID != nil && *req.AvatarAssetID > 0 {
		if err := u.assetClient.ValidateAvailable(ctx, *req.AvatarAssetID); err != nil {
			return 0, err
		}
	}
	group, err := u.chatGroupRepo.Save(ctx, &model.ChatGroup{
		Name:          req.Name,
		AvatarAssetID: req.AvatarAssetID,
		Introduction:  req.Introduction,
		OwnerID:       req.OwnerID,
		MemberCount:   1,
		CreatedBy:     &req.OwnerID,
		UpdatedBy:     &req.OwnerID,
	})
	if err != nil {
		return 0, err
	}
	_, err = u.chatGroupMemberRepo.Save(ctx, &model.ChatGroupMember{
		GroupID:   group.ID,
		UserID:    req.OwnerID,
		Role:      enum.ChatGroupMemberRoleOwner,
		CreatedBy: &req.OwnerID,
		UpdatedBy: &req.OwnerID,
	})
	if err != nil {
		return 0, err
	}
	return group.ID, nil
}

type DismissReq struct {
	GroupID    int64
	OperatorID int64
}

func (u *ChatGroupUsecase) Dismiss(ctx context.Context, req *DismissReq) error {
	group, err := u.chatGroupRepo.Get(ctx, &repo.ChatGroupQuery{
		IDs: []int64{req.GroupID},
	})
	if err != nil {
		return err
	}
	if group.OwnerID != req.OperatorID {
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
	Page   *commonmodel.PageReq
	IDs    []int64
	Status *enum.ChatGroupStatus
}

type ChatGroupListResp struct {
	List []*model.ChatGroup
	Page *commonmodel.PageResp
}

func (u *ChatGroupUsecase) List(ctx context.Context, req *ChatGroupListReq) (*ChatGroupListResp, error) {
	pageResp, err := u.chatGroupRepo.Page(ctx, &repo.ChatGroupQuery{
		Page:   req.Page,
		IDs:    req.IDs,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &ChatGroupListResp{
		List: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}
