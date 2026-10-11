package repo

import (
	commonmodel "common/pkg/model"
	"context"

	"bff_bbs/internal/biz/model"
)

type RelationClient interface {
	Follow(ctx context.Context, req *FollowRelationReq) error
	Unfollow(ctx context.Context, req *UnfollowRelationReq) error
	Block(ctx context.Context, req *BlockRelationReq) error
	Unblock(ctx context.Context, req *UnblockRelationReq) error
	ListFollowing(ctx context.Context, req *ListFollowingRelationsReq) (*ListFollowingRelationsResp, error)
	ListFollowers(ctx context.Context, req *ListFollowersRelationsReq) (*ListFollowersRelationsResp, error)
	ListBlocked(ctx context.Context, req *ListBlockedRelationsReq) (*ListBlockedRelationsResp, error)
	GetStatus(ctx context.Context, req *GetStatusRelationReq) (*RelationStatus, error)
	MapStatus(ctx context.Context, actorID int64, targetIDs []int64) (map[int64]*RelationStatus, error)
	ListBlockedCounterparts(ctx context.Context, userID int64) ([]int64, error)
}

type FollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

type UnfollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

type BlockRelationReq struct {
	ActorID  int64
	TargetID int64
}

type UnblockRelationReq struct {
	ActorID  int64
	TargetID int64
}

type ListFollowingRelationsReq struct {
	ActorID         int64
	Page            *commonmodel.PageReq
	ExcludedUserIDs []int64
}

type ListFollowingRelationsResp struct {
	Page *commonmodel.PageResp
	Rows []*model.Relation
}

type ListFollowersRelationsReq struct {
	ActorID         int64
	Page            *commonmodel.PageReq
	ExcludedUserIDs []int64
}

type ListFollowersRelationsResp struct {
	Page *commonmodel.PageResp
	Rows []*model.Relation
}

type ListBlockedRelationsReq struct {
	ActorID int64
	Page    *commonmodel.PageReq
}

type ListBlockedRelationsResp struct {
	Page *commonmodel.PageResp
	Rows []*model.Relation
}

type GetStatusRelationReq struct {
	ActorID  int64
	TargetID int64
}

type RelationStatus struct {
	TargetID   int64
	Following  bool
	FollowedBy bool
	Blocking   bool
	BlockedBy  bool
}
