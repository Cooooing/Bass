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

func (s *RelationService) Follow(ctx context.Context, req *v1.FollowRelation_Req) (*v1.FollowRelation_Resp, error) {
	err := s.relationUsecase.Follow(ctx, &usecase.FollowRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.FollowRelation_Resp{}, err
}

func (s *RelationService) Unfollow(ctx context.Context, req *v1.UnfollowRelation_Req) (*v1.UnfollowRelation_Resp, error) {
	err := s.relationUsecase.Unfollow(ctx, &usecase.UnfollowRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.UnfollowRelation_Resp{}, err
}

func (s *RelationService) Block(ctx context.Context, req *v1.BlockRelation_Req) (*v1.BlockRelation_Resp, error) {
	err := s.relationUsecase.Block(ctx, &usecase.BlockRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.BlockRelation_Resp{}, err
}

func (s *RelationService) Unblock(ctx context.Context, req *v1.UnblockRelation_Req) (*v1.UnblockRelation_Resp, error) {
	err := s.relationUsecase.Unblock(ctx, &usecase.UnblockRelationReq{ActorID: req.GetActorId(), TargetID: req.GetTargetId()})
	return &v1.UnblockRelation_Resp{}, err
}

func (s *RelationService) ListFollowing(ctx context.Context, req *v1.ListFollowingRelations_Req) (*v1.ListFollowingRelations_Resp, error) {
	req = util.OrDefault(req, &v1.ListFollowingRelations_Req{})
	req.Page = util.OrDefault(req.Page, &common.PageReq{})
	res, err := s.relationUsecase.ListFollowing(ctx, &usecase.ListFollowingRelationsReq{
		Page:    usecase.RelationPageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()},
		ActorID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.ListFollowingRelations_Resp_Relation, 0, len(res.Rows))
	for _, row := range res.Rows {
		reply := &v1.ListFollowingRelations_Resp_Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListFollowingRelations_Resp{Page: &common.PageResp{Total: res.Page.Total, Page: res.Page.Page, Size: res.Page.Size}, Rows: replyRows}, nil
}

func (s *RelationService) ListFollowers(ctx context.Context, req *v1.ListFollowersRelations_Req) (*v1.ListFollowersRelations_Resp, error) {
	req = util.OrDefault(req, &v1.ListFollowersRelations_Req{})
	targetID := req.GetUserId()
	req.Page = util.OrDefault(req.Page, &common.PageReq{})
	res, err := s.relationUsecase.ListFollowers(ctx, &usecase.ListFollowersRelationsReq{
		Page:     usecase.RelationPageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()},
		TargetID: targetID,
	})
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.ListFollowersRelations_Resp_Relation, 0, len(res.Rows))
	for _, row := range res.Rows {
		reply := &v1.ListFollowersRelations_Resp_Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListFollowersRelations_Resp{Page: &common.PageResp{Total: res.Page.Total, Page: res.Page.Page, Size: res.Page.Size}, Rows: replyRows}, nil
}

func (s *RelationService) ListBlocked(ctx context.Context, req *v1.ListBlockedRelations_Req) (*v1.ListBlockedRelations_Resp, error) {
	req = util.OrDefault(req, &v1.ListBlockedRelations_Req{})
	req.Page = util.OrDefault(req.Page, &common.PageReq{})
	res, err := s.relationUsecase.ListBlocked(ctx, &usecase.ListBlockedRelationsReq{
		Page:    usecase.RelationPageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()},
		ActorID: req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	replyRows := make([]*v1.ListBlockedRelations_Resp_Relation, 0, len(res.Rows))
	for _, row := range res.Rows {
		reply := &v1.ListBlockedRelations_Resp_Relation{Id: row.ID, Type: enum.RelationTypeMap.MustToProto(row.Type), ActorId: row.ActorID, TargetId: row.TargetID}
		if row.CreatedAt != nil {
			reply.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, reply)
	}
	return &v1.ListBlockedRelations_Resp{Page: &common.PageResp{Total: res.Page.Total, Page: res.Page.Page, Size: res.Page.Size}, Rows: replyRows}, nil
}

func (s *RelationService) MapStatus(ctx context.Context, req *v1.MapRelationStatuses_Req) (*v1.MapRelationStatuses_Resp, error) {
	req = util.OrDefault(req, &v1.MapRelationStatuses_Req{})
	res, err := s.relationUsecase.MapStatus(ctx, &usecase.MapRelationStatusReq{ActorID: req.GetActorId(), TargetIDs: req.TargetIds})
	if err != nil {
		return nil, err
	}
	statuses := make(map[int64]*v1.MapRelationStatuses_Resp_RelationStatus, len(res))
	for targetID, row := range res {
		if row == nil {
			continue
		}
		statuses[targetID] = &v1.MapRelationStatuses_Resp_RelationStatus{
			TargetId:   row.TargetID,
			Following:  row.Following,
			FollowedBy: row.FollowedBy,
			Blocking:   row.Blocking,
			BlockedBy:  row.BlockedBy,
		}
	}
	return &v1.MapRelationStatuses_Resp{Statuses: statuses}, nil
}
