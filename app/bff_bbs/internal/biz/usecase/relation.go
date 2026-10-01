package usecase

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
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
	Page    *commonmodel.PageReq
}

type ListFollowingResp struct {
	Page *commonmodel.PageResp
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListFollowing(ctx context.Context, req *ListFollowingReq) (*ListFollowingResp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
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
	Page    *commonmodel.PageReq
}

type ListFollowersResp struct {
	Page *commonmodel.PageResp
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListFollowers(ctx context.Context, req *ListFollowersReq) (*ListFollowersResp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
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
	Page    *commonmodel.PageReq
}

type ListBlockedResp struct {
	Page *commonmodel.PageResp
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListBlocked(ctx context.Context, req *ListBlockedReq) (*ListBlockedResp, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
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
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_RELATION_INVALID)
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
