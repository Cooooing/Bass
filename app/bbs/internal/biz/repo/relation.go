package repo

import "context"

type RelationClient interface {
	Follow(ctx context.Context, req *FollowRelationReq) (*FollowRelationResponse, error)
	Unfollow(ctx context.Context, req *UnfollowRelationReq) (*UnfollowRelationResponse, error)
	Block(ctx context.Context, req *BlockRelationReq) (*BlockRelationResponse, error)
	Unblock(ctx context.Context, req *UnblockRelationReq) (*UnblockRelationResponse, error)
	ListFollowing(ctx context.Context, req *ListFollowingRelationsReq) (*ListFollowingRelationsResponse, error)
	ListFollowers(ctx context.Context, req *ListFollowersRelationsReq) (*ListFollowersRelationsResponse, error)
	ListBlocked(ctx context.Context, req *ListBlockedRelationsReq) (*ListBlockedRelationsResponse, error)
	GetStatus(ctx context.Context, req *GetStatusRelationReq) (*GetStatusRelationResponse, error)
}

type FollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

type FollowRelationResponse struct{}

type UnfollowRelationReq struct {
	ActorID  int64
	TargetID int64
}

type UnfollowRelationResponse struct{}

type BlockRelationReq struct {
	ActorID  int64
	TargetID int64
}

type BlockRelationResponse struct{}

type UnblockRelationReq struct {
	ActorID  int64
	TargetID int64
}

type UnblockRelationResponse struct{}

type ListFollowingRelationsReq struct {
	ActorID int64
	Page    *PageReq
}

type ListFollowingRelationsResponse struct {
	Page *PageResponse
	Rows []*Relation
}

type ListFollowersRelationsReq struct {
	ActorID int64
	Page    *PageReq
}

type ListFollowersRelationsResponse struct {
	Page *PageResponse
	Rows []*Relation
}

type ListBlockedRelationsReq struct {
	ActorID int64
	Page    *PageReq
}

type ListBlockedRelationsResponse struct {
	Page *PageResponse
	Rows []*Relation
}

type GetStatusRelationReq struct {
	ActorID  int64
	TargetID int64
}

type GetStatusRelationResponse struct {
	Status *RelationStatus
}

type Relation struct {
	ID        int64
	Type      int32
	ActorID   int64
	TargetID  int64
	CreatedAt string
	UpdatedAt string
}

type RelationStatus struct {
	TargetID   int64
	Following  bool
	FollowedBy bool
	Blocking   bool
	BlockedBy  bool
}
