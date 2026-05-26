package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type RelationService struct {
	bbsuserv1.UnimplementedRelationServiceServer
	userUsecase *usecase.UserUsecase
}

func NewRelationService(userUsecase *usecase.UserUsecase) *RelationService {
	return &RelationService{userUsecase: userUsecase}
}

func (s *RelationService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterRelationServiceServer(gs, s)
}

func (s *RelationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterRelationServiceHTTPServer(hs, s)
}

func (s *RelationService) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	return s.userUsecase.Follow(ctx, req)
}

func (s *RelationService) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	return s.userUsecase.Unfollow(ctx, req)
}

func (s *RelationService) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	return s.userUsecase.Block(ctx, req)
}

func (s *RelationService) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	return s.userUsecase.Unblock(ctx, req)
}

func (s *RelationService) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Reply, error) {
	return s.userUsecase.ListFollowing(ctx, req)
}

func (s *RelationService) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Reply, error) {
	return s.userUsecase.ListFollowers(ctx, req)
}

func (s *RelationService) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Reply, error) {
	return s.userUsecase.ListBlocked(ctx, req)
}

func (s *RelationService) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Reply, error) {
	return s.userUsecase.GetStatus(ctx, req)
}
