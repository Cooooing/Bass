package usecase

import (
	"context"

	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"im/internal/biz/model"
	"im/internal/biz/repo"
	"im/internal/enum"

	"log/slog"
)

type ChatGroupUsecase struct {
	chatGroupRepo       repo.ChatGroupRepo
	chatGroupMemberRepo repo.ChatGroupMemberRepo
	log                 *slog.Logger
}

func NewChatGroupUsecase(chatGroupRepo repo.ChatGroupRepo, chatGroupMemberRepo repo.ChatGroupMemberRepo, logger *slog.Logger) (*ChatGroupUsecase, error) {
	return &ChatGroupUsecase{chatGroupRepo: chatGroupRepo, chatGroupMemberRepo: chatGroupMemberRepo, log: logger}, nil
}

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

func (u *ChatGroupUsecase) Dismiss(ctx context.Context, groupID int64, operatorID int64) error {
	group, err := u.chatGroupRepo.Get(ctx, &repo.ChatGroupGetReq{IDs: []int64{groupID}})
	if err != nil {
		return err
	}
	if group.OwnerID != operatorID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_IM_CHAT_GROUP_STATUS_INVALID)
	}
	_, err = u.chatGroupRepo.UpdateStatus(ctx, groupID, enum.ChatGroupStatusDissolve, operatorID)
	return err
}

func (u *ChatGroupUsecase) List(ctx context.Context, page *common.PageRequest, ids []int64, status *enum.ChatGroupStatus) ([]*model.ChatGroup, *common.PageReply, error) {
	return u.chatGroupRepo.Page(ctx, page, &repo.ChatGroupGetReq{IDs: ids, Status: status})
}
