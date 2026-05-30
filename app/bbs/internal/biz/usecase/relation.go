package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type RelationUsecase struct {
	relationRepo repo.RelationRepo
}

func NewRelationUsecase(relationRepo repo.RelationRepo) *RelationUsecase {
	return &RelationUsecase{relationRepo: relationRepo}
}

func (u *RelationUsecase) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	return u.relationRepo.Follow(ctx, req)
}

func (u *RelationUsecase) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	return u.relationRepo.Unfollow(ctx, req)
}

func (u *RelationUsecase) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	return u.relationRepo.Block(ctx, req)
}

func (u *RelationUsecase) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	return u.relationRepo.Unblock(ctx, req)
}

func (u *RelationUsecase) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Reply, error) {
	return u.relationRepo.ListFollowing(ctx, req)
}

func (u *RelationUsecase) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Reply, error) {
	return u.relationRepo.ListFollowers(ctx, req)
}

func (u *RelationUsecase) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Reply, error) {
	return u.relationRepo.ListBlocked(ctx, req)
}

func (u *RelationUsecase) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Reply, error) {
	return u.relationRepo.GetStatus(ctx, req)
}
