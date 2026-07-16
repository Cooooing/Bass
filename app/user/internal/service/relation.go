package service

import (
	"common/pkg/util"
	"common/proto/gen/common"
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/usecase"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
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

func (s *RelationService) Follow(ctx context.Context, req *v1.FollowRelation_Request) (*v1.FollowRelation_Response, error) {
	err := s.relationUsecase.Follow(ctx, &usecase.FollowRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.FollowRelation_Response{}, err
}

func (s *RelationService) Unfollow(ctx context.Context, req *v1.UnfollowRelation_Request) (*v1.UnfollowRelation_Response, error) {
	err := s.relationUsecase.Unfollow(ctx, &usecase.UnfollowRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.UnfollowRelation_Response{}, err
}

func (s *RelationService) Block(ctx context.Context, req *v1.BlockRelation_Request) (*v1.BlockRelation_Response, error) {
	err := s.relationUsecase.Block(ctx, &usecase.BlockRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.BlockRelation_Response{}, err
}

func (s *RelationService) Unblock(ctx context.Context, req *v1.UnblockRelation_Request) (*v1.UnblockRelation_Response, error) {
	err := s.relationUsecase.Unblock(ctx, &usecase.UnblockRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.UnblockRelation_Response{}, err
}

func (s *RelationService) ListFollowing(ctx context.Context, req *v1.ListFollowingRelations_Request) (*v1.ListFollowingRelations_Response, error) {
	req = util.OrDefault(req, &v1.ListFollowingRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	res, err := s.relationUsecase.ListFollowing(ctx, &usecase.ListFollowingRelationsReq{
		Page:    usecase.RelationPageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()},
		ActorID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.ListFollowingRelations_Response_Relation, 0, len(res.Rows))
	for _, row := range res.Rows {
		reply := &v1.ListFollowingRelations_Response_Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListFollowingRelations_Response{Page: &common.PageResponse{Total: res.Page.Total, Page: res.Page.Page, Size: res.Page.Size}, Rows: replyRows}, nil
}

func (s *RelationService) ListFollowers(ctx context.Context, req *v1.ListFollowersRelations_Request) (*v1.ListFollowersRelations_Response, error) {
	req = util.OrDefault(req, &v1.ListFollowersRelations_Request{})
	targetID := req.GetUserId()
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	res, err := s.relationUsecase.ListFollowers(ctx, &usecase.ListFollowersRelationsReq{
		Page:     usecase.RelationPageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()},
		TargetID: targetID,
	})
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.ListFollowersRelations_Response_Relation, 0, len(res.Rows))
	for _, row := range res.Rows {
		reply := &v1.ListFollowersRelations_Response_Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListFollowersRelations_Response{Page: &common.PageResponse{Total: res.Page.Total, Page: res.Page.Page, Size: res.Page.Size}, Rows: replyRows}, nil
}

func (s *RelationService) ListBlocked(ctx context.Context, req *v1.ListBlockedRelations_Request) (*v1.ListBlockedRelations_Response, error) {
	req = util.OrDefault(req, &v1.ListBlockedRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	res, err := s.relationUsecase.ListBlocked(ctx, &usecase.ListBlockedRelationsReq{
		Page:    usecase.RelationPageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()},
		ActorID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.ListBlockedRelations_Response_Relation, 0, len(res.Rows))
	for _, row := range res.Rows {
		reply := &v1.ListBlockedRelations_Response_Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListBlockedRelations_Response{Page: &common.PageResponse{Total: res.Page.Total, Page: res.Page.Page, Size: res.Page.Size}, Rows: replyRows}, nil
}

func (s *RelationService) MapStatus(ctx context.Context, req *v1.MapRelationStatuses_Request) (*v1.MapRelationStatuses_Response, error) {
	req = util.OrDefault(req, &v1.MapRelationStatuses_Request{})
	res, err := s.relationUsecase.MapStatus(ctx, &usecase.MapRelationStatusReq{ActorID: req.GetActorId(), TargetIDs: req.TargetIds})
	if err != nil {
		return nil, err
	}
	statuses := make(map[int64]*v1.MapRelationStatuses_Response_RelationStatus, len(res.Statuses))
	for targetID, row := range res.Statuses {
		if row == nil {
			continue
		}
		statuses[targetID] = &v1.MapRelationStatuses_Response_RelationStatus{
			TargetId:   row.TargetID,
			Following:  row.Following,
			FollowedBy: row.FollowedBy,
			Blocking:   row.Blocking,
			BlockedBy:  row.BlockedBy,
		}
	}
	return &v1.MapRelationStatuses_Response{Statuses: statuses}, nil
}
