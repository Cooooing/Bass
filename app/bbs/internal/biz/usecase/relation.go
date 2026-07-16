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

func NewRelationUsecase(relationClient repo.RelationClient) *RelationUsecase {
	return &RelationUsecase{relationClient: relationClient}
}

type FollowReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Follow(ctx context.Context, req *FollowReq) error {
	if err := validateRelationTarget(req); err != nil {
		return err
	}
	_, err := u.relationClient.Follow(ctx, &repo.FollowRelationReq{ActorID: req.ActorID, TargetID: req.TargetID})
	return err
}

type UnfollowReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Unfollow(ctx context.Context, req *UnfollowReq) error {
	if err := validateRelationTarget(req); err != nil {
		return err
	}
	_, err := u.relationClient.Unfollow(ctx, &repo.UnfollowRelationReq{ActorID: req.ActorID, TargetID: req.TargetID})
	return err
}

type BlockReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Block(ctx context.Context, req *BlockReq) error {
	if err := validateRelationTarget(req); err != nil {
		return err
	}
	_, err := u.relationClient.Block(ctx, &repo.BlockRelationReq{ActorID: req.ActorID, TargetID: req.TargetID})
	return err
}

type UnblockReq struct {
	ActorID  int64
	TargetID int64
}

func (u *RelationUsecase) Unblock(ctx context.Context, req *UnblockReq) error {
	if err := validateRelationTarget(req); err != nil {
		return err
	}
	_, err := u.relationClient.Unblock(ctx, &repo.UnblockRelationReq{ActorID: req.ActorID, TargetID: req.TargetID})
	return err
}

type ListFollowingReq struct {
	ActorID int64
	Page    *common.PageRequest
}

type ListFollowingResponse struct {
	Page *repo.PageResponse
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListFollowing(ctx context.Context, req *ListFollowingReq) (*ListFollowingResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	response, err := u.relationClient.ListFollowing(ctx, &repo.ListFollowingRelationsReq{ActorID: req.ActorID, Page: page})
	if err != nil {
		return nil, err
	}
	return &ListFollowingResponse{Page: response.Page, Rows: response.Rows}, nil
}

type ListFollowersReq struct {
	ActorID int64
	Page    *common.PageRequest
}

type ListFollowersResponse struct {
	Page *repo.PageResponse
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListFollowers(ctx context.Context, req *ListFollowersReq) (*ListFollowersResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	response, err := u.relationClient.ListFollowers(ctx, &repo.ListFollowersRelationsReq{ActorID: req.ActorID, Page: page})
	if err != nil {
		return nil, err
	}
	return &ListFollowersResponse{Page: response.Page, Rows: response.Rows}, nil
}

type ListBlockedReq struct {
	ActorID int64
	Page    *common.PageRequest
}

type ListBlockedResponse struct {
	Page *repo.PageResponse
	Rows []*repo.Relation
}

func (u *RelationUsecase) ListBlocked(ctx context.Context, req *ListBlockedReq) (*ListBlockedResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	response, err := u.relationClient.ListBlocked(ctx, &repo.ListBlockedRelationsReq{ActorID: req.ActorID, Page: page})
	if err != nil {
		return nil, err
	}
	return &ListBlockedResponse{Page: response.Page, Rows: response.Rows}, nil
}

type GetStatusReq struct {
	ActorID  int64
	TargetID int64
}

type GetStatusResponse struct {
	Status *repo.RelationStatus
}

func (u *RelationUsecase) GetStatus(ctx context.Context, req *GetStatusReq) (*GetStatusResponse, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	response, err := u.relationClient.GetStatus(ctx, &repo.GetStatusRelationReq{ActorID: req.ActorID, TargetID: req.TargetID})
	if err != nil {
		return nil, err
	}
	return &GetStatusResponse{Status: response.Status}, nil
}

type relationTarget interface {
	getActorID() int64
	getTargetID() int64
}

func (req *FollowReq) getActorID() int64    { return req.ActorID }
func (req *FollowReq) getTargetID() int64   { return req.TargetID }
func (req *UnfollowReq) getActorID() int64  { return req.ActorID }
func (req *UnfollowReq) getTargetID() int64 { return req.TargetID }
func (req *BlockReq) getActorID() int64     { return req.ActorID }
func (req *BlockReq) getTargetID() int64    { return req.TargetID }
func (req *UnblockReq) getActorID() int64   { return req.ActorID }
func (req *UnblockReq) getTargetID() int64  { return req.TargetID }

func validateRelationTarget(req relationTarget) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.getActorID() == req.getTargetID() {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return nil
}
