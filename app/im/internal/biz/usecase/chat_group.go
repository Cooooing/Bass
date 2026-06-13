package usecase

import (
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"common/pkg/apperror"
	"context"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
)

type ChatGroupUsecase struct {
	chatGroupRepo       repo.ChatGroupRepo
	chatGroupMemberRepo repo.ChatGroupMemberRepo
	log                 *log.Helper
}

func NewChatGroupUsecase(
	chatGroupRepo repo.ChatGroupRepo,
	chatGroupMemberRepo repo.ChatGroupMemberRepo,
	logger log.Logger,
) (*ChatGroupUsecase, error) {
	return &ChatGroupUsecase{
		chatGroupRepo:       chatGroupRepo,
		chatGroupMemberRepo: chatGroupMemberRepo,
		log:                 log.NewHelper(logger),
	}, nil
}

// Create 创建群组，并将创建者设为群主。
func (u *ChatGroupUsecase) Create(ctx context.Context, name string, avatar *string, introduction *string, ownerID int64) (int64, error) {
	group, err := u.chatGroupRepo.Save(ctx, &model.ChatGroup{
		Name:         name,
		Avatar:       avatar,
		Introduction: introduction,
		OwnerID:      ownerID,
		MemberCount:  1,
		CreatedBy:    &ownerID,
		UpdatedBy:    &ownerID,
	})
	if err != nil {
		return 0, err
	}
	// 将创建者加入群组，角色为群主
	_, err = u.chatGroupMemberRepo.Save(ctx, &model.ChatGroupMember{
		GroupID:   group.ID,
		UserID:    ownerID,
		Role:      enum.ChatGroupMemberRoleOwner,
		CreatedBy: &ownerID,
		UpdatedBy: &ownerID,
	})
	if err != nil {
		return 0, err
	}
	return group.ID, nil
}

// Dismiss 解散群组，仅群主可操作。
func (u *ChatGroupUsecase) Dismiss(ctx context.Context, groupID int64, operatorID int64) error {
	group, err := u.chatGroupRepo.Get(ctx, &repo.ChatGroupGetReq{
		IDs: []int64{groupID},
	})
	if err != nil {
		return err
	}
	if group.OwnerID != operatorID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_STATUS_INVALID)
	}
	_, err = u.chatGroupRepo.UpdateStatus(ctx, groupID, enum.ChatGroupStatusDissolve, operatorID)
	return err
}

// List 查询群组列表。
func (u *ChatGroupUsecase) List(ctx context.Context, page *common.PageRequest, ids []int64, status *enum.ChatGroupStatus) ([]*model.ChatGroup, *common.PageReply, error) {
	getReq := &repo.ChatGroupGetReq{
		IDs:    ids,
		Status: status,
	}
	return u.chatGroupRepo.Page(ctx, page, getReq)
}
