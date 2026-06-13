package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type RelationService struct {
	bbsuserv1.UnimplementedRelationServiceServer
	relationUsecase *usecase.RelationUsecase
}

func NewRelationService(relationUsecase *usecase.RelationUsecase) *RelationService {
	return &RelationService{relationUsecase: relationUsecase}
}

func (s *RelationService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterRelationServiceServer(gs, s)
}

func (s *RelationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterRelationServiceHTTPServer(hs, s)
}

func (s *RelationService) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil && user.ID == req.GetTargetId() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return s.relationUsecase.Follow(ctx, req)
}

func (s *RelationService) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil && user.ID == req.GetTargetId() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return s.relationUsecase.Unfollow(ctx, req)
}

func (s *RelationService) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil && user.ID == req.GetTargetId() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return s.relationUsecase.Block(ctx, req)
}

func (s *RelationService) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil && user.ID == req.GetTargetId() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_SELF_OPERATION_NOT_ALLOWED)
	}
	return s.relationUsecase.Unblock(ctx, req)
}

func (s *RelationService) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Reply, error) {
	return s.relationUsecase.ListFollowing(ctx, req)
}

func (s *RelationService) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Reply, error) {
	return s.relationUsecase.ListFollowers(ctx, req)
}

func (s *RelationService) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Reply, error) {
	return s.relationUsecase.ListBlocked(ctx, req)
}

func (s *RelationService) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Reply, error) {
	return s.relationUsecase.GetStatus(ctx, req)
}
