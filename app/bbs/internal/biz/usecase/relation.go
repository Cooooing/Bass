package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type RelationUsecase struct {
	relationClient repo.RelationClient
}

func NewRelationUsecase(
	relationClient repo.RelationClient,
) *RelationUsecase {
	return &RelationUsecase{
		relationClient: relationClient,
	}
}

type FollowReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Follow(ctx context.Context, req *FollowReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return u.relationClient.Follow(ctx, &repo.FollowRelationReq{
		ActorID:  req.ActorID,
		TargetID: req.TargetID,
	})
}

type UnfollowReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Unfollow(ctx context.Context, req *UnfollowReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return u.relationClient.Unfollow(ctx, &repo.UnfollowRelationReq{
		ActorID:  req.ActorID,
		TargetID: req.TargetID,
	})
}

type BlockReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Block(ctx context.Context, req *BlockReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return u.relationClient.Block(ctx, &repo.BlockRelationReq{
		ActorID:  req.ActorID,
		TargetID: req.TargetID,
	})
}

type UnblockReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Unblock(ctx context.Context, req *UnblockReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.ActorID == req.TargetID {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return u.relationClient.Unblock(ctx, &repo.UnblockRelationReq{
		ActorID:  req.ActorID,
		TargetID: req.TargetID,
	})
}

type ListFollowingReq struct {
	ActorID int64
	Page    *common.PageReq
}

type ListFollowingResp struct {
	Page *repo.PageResp
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListFollowing(ctx context.Context, req *ListFollowingReq) (*ListFollowingResp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	resp, err := u.relationClient.ListFollowing(ctx, &repo.ListFollowingRelationsReq{
		ActorID: req.ActorID,
		Page:    page,
	})
	if err != nil {
		return nil, err
	}
	return &ListFollowingResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type ListFollowersReq struct {
	ActorID int64
	Page    *common.PageReq
}

type ListFollowersResp struct {
	Page *repo.PageResp
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListFollowers(ctx context.Context, req *ListFollowersReq) (*ListFollowersResp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	resp, err := u.relationClient.ListFollowers(ctx, &repo.ListFollowersRelationsReq{
		ActorID: req.ActorID,
		Page:    page,
	})
	if err != nil {
		return nil, err
	}
	return &ListFollowersResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type ListBlockedReq struct {
	ActorID int64
	Page    *common.PageReq
}

type ListBlockedResp struct {
	Page *repo.PageResp
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListBlocked(ctx context.Context, req *ListBlockedReq) (*ListBlockedResp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	resp, err := u.relationClient.ListBlocked(ctx, &repo.ListBlockedRelationsReq{
		ActorID: req.ActorID,
		Page:    page,
	})
	if err != nil {
		return nil, err
	}
	return &ListBlockedResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type GetStatusReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) GetStatus(ctx context.Context, req *GetStatusReq) (*repo.RelationStatus, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.relationClient.GetStatus(ctx, &repo.GetStatusRelationReq{
		ActorID:  req.ActorID,
		TargetID: req.TargetID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
