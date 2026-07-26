package repo

import (
	"bbs/internal/enum"
	"context"
	"time"
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
	ActorID int64
	Page    *PageReq
}

type ListFollowingRelationsResp struct {
	Page *PageResp
	Rows []*Relation
}

type ListFollowersRelationsReq struct {
	ActorID int64
	Page    *PageReq
}

type ListFollowersRelationsResp struct {
	Page *PageResp
	Rows []*Relation
}

type ListBlockedRelationsReq struct {
	ActorID int64
	Page    *PageReq
}

type ListBlockedRelationsResp struct {
	Page *PageResp
	Rows []*Relation
}

type GetStatusRelationReq struct {
	ActorID  int64
	TargetID int64
}

type Relation struct {
	ID        int64
	Type      enum.RelationType
	ActorID   int64
	TargetID  int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type RelationStatus struct {
	TargetID   int64
	Following  bool
	FollowedBy bool
	Blocking   bool
	BlockedBy  bool
}
