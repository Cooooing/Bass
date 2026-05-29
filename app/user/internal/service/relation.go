package service

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/util"
	"context"
	"user/internal/biz/usecase"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RelationService struct {
	v1.UnimplementedRelationServiceServer
	relationUsecase *usecase.RelationUsecase
}

func NewRelationService(relationUsecase *usecase.RelationUsecase) *RelationService {
	return &RelationService{
		relationUsecase: relationUsecase,
	}
}

func (s *RelationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterRelationServiceServer(gs, s)
}

func (s *RelationService) RegisterHttp(hs *http.Server) {}

func (s *RelationService) Follow(ctx context.Context, req *v1.FollowRelation_Request) (*v1.FollowRelation_Reply, error) {
	if req.GetActorId() == req.GetTargetId() {
		return nil, cerrors.ErrorBadRequest("can not follow yourself")
	}
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_FOLLOW, true, req.GetActorId(), req.GetTargetId())
	return &v1.FollowRelation_Reply{}, err
}

func (s *RelationService) Unfollow(ctx context.Context, req *v1.UnfollowRelation_Request) (*v1.UnfollowRelation_Reply, error) {
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_FOLLOW, false, req.GetActorId(), req.GetTargetId())
	return &v1.UnfollowRelation_Reply{}, err
}

func (s *RelationService) Block(ctx context.Context, req *v1.BlockRelation_Request) (*v1.BlockRelation_Reply, error) {
	if req.GetActorId() == req.GetTargetId() {
		return nil, cerrors.ErrorBadRequest("can not block yourself")
	}
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_BLOCK, true, req.GetActorId(), req.GetTargetId())
	return &v1.BlockRelation_Reply{}, err
}

func (s *RelationService) Unblock(ctx context.Context, req *v1.UnblockRelation_Request) (*v1.UnblockRelation_Reply, error) {
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_BLOCK, false, req.GetActorId(), req.GetTargetId())
	return &v1.UnblockRelation_Reply{}, err
}

func (s *RelationService) ListFollowing(ctx context.Context, req *v1.ListFollowingRelations_Request) (*v1.ListFollowingRelations_Reply, error) {
	req = util.OrDefault(req, &v1.ListFollowingRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	rows, page, err := s.relationUsecase.ListFollowing(ctx, req.Page, req.GetUserId())
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.Relation, 0, len(rows))
	for _, row := range rows {
		reply := &v1.Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListFollowingRelations_Reply{Page: page, Rows: replyRows}, nil
}

func (s *RelationService) ListFollowers(ctx context.Context, req *v1.ListFollowersRelations_Request) (*v1.ListFollowersRelations_Reply, error) {
	req = util.OrDefault(req, &v1.ListFollowersRelations_Request{})
	targetID := req.GetUserId()
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	rows, page, err := s.relationUsecase.ListFollowers(ctx, req.Page, targetID)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.Relation, 0, len(rows))
	for _, row := range rows {
		reply := &v1.Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListFollowersRelations_Reply{Page: page, Rows: replyRows}, nil
}

func (s *RelationService) ListBlocked(ctx context.Context, req *v1.ListBlockedRelations_Request) (*v1.ListBlockedRelations_Reply, error) {
	req = util.OrDefault(req, &v1.ListBlockedRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	rows, page, err := s.relationUsecase.ListBlocked(ctx, req.Page, req.GetUserId())
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.Relation, 0, len(rows))
	for _, row := range rows {
		reply := &v1.Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListBlockedRelations_Reply{Page: page, Rows: replyRows}, nil
}

func (s *RelationService) MapStatus(ctx context.Context, req *v1.MapRelationStatuses_Request) (*v1.MapRelationStatuses_Reply, error) {
	req = util.OrDefault(req, &v1.MapRelationStatuses_Request{})
	rows, err := s.relationUsecase.MapStatus(ctx, req.GetActorId(), req.TargetIds)
	if err != nil {
		return nil, err
	}
	statuses := make(map[int64]*v1.RelationStatus, len(rows))
	for targetID, row := range rows {
		if row == nil {
			continue
		}
		statuses[targetID] = &v1.RelationStatus{
			TargetId:   row.TargetID,
			Following:  row.Following,
			FollowedBy: row.FollowedBy,
			Blocking:   row.Blocking,
			BlockedBy:  row.BlockedBy,
		}
	}
	return &v1.MapRelationStatuses_Reply{Statuses: statuses}, nil
}
