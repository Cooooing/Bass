package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	commonv1 "common/api/gen/common"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type RelationService struct {
	bbsuserv1.UnimplementedRelationServiceServer
	userClient *rpc.UserClient
}

func NewRelationService(userClient *rpc.UserClient) *RelationService {
	return &RelationService{userClient: userClient}
}

func (s *RelationService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterRelationServiceServer(gs, s)
}

func (s *RelationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterRelationServiceHTTPServer(hs, s)
}

func (s *RelationService) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	_, err := s.userClient.Relation.Follow(forwardAuth(ctx), &userv1.FollowRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.FollowRelation_Reply{}, nil
}

func (s *RelationService) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	_, err := s.userClient.Relation.Unfollow(forwardAuth(ctx), &userv1.UnfollowRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UnfollowRelation_Reply{}, nil
}

func (s *RelationService) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	_, err := s.userClient.Relation.Block(forwardAuth(ctx), &userv1.BlockRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BlockRelation_Reply{}, nil
}

func (s *RelationService) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	_, err := s.userClient.Relation.Unblock(forwardAuth(ctx), &userv1.UnblockRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UnblockRelation_Reply{}, nil
}

func (s *RelationService) PageFollowing(ctx context.Context, req *bbsuserv1.PageFollowingRelation_Request) (*bbsuserv1.PageFollowingRelation_Reply, error) {
	reply, err := s.userClient.Relation.PageFollowing(forwardAuth(ctx), &userv1.PageFollowingRelation_Request{Page: toUserPageRequest(req.GetPage())})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.PageFollowingRelation_Reply{
		Page: toBFFPageReply(reply.GetPage()),
		Rows: toBFFRelations(reply.GetRows()),
	}, nil
}

func (s *RelationService) PageFollowers(ctx context.Context, req *bbsuserv1.PageFollowersRelation_Request) (*bbsuserv1.PageFollowersRelation_Reply, error) {
	reply, err := s.userClient.Relation.PageFollowers(forwardAuth(ctx), &userv1.PageFollowersRelation_Request{Page: toUserPageRequest(req.GetPage())})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.PageFollowersRelation_Reply{
		Page: toBFFPageReply(reply.GetPage()),
		Rows: toBFFRelations(reply.GetRows()),
	}, nil
}

func (s *RelationService) PageBlocked(ctx context.Context, req *bbsuserv1.PageBlockedRelation_Request) (*bbsuserv1.PageBlockedRelation_Reply, error) {
	reply, err := s.userClient.Relation.PageBlocked(forwardAuth(ctx), &userv1.PageBlockedRelation_Request{Page: toUserPageRequest(req.GetPage())})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.PageBlockedRelation_Reply{
		Page: toBFFPageReply(reply.GetPage()),
		Rows: toBFFRelations(reply.GetRows()),
	}, nil
}

func (s *RelationService) BatchGetStatus(ctx context.Context, req *bbsuserv1.BatchGetStatusRelation_Request) (*bbsuserv1.BatchGetStatusRelation_Reply, error) {
	reply, err := s.userClient.Relation.BatchGetStatus(forwardAuth(ctx), &userv1.BatchGetStatusRelation_Request{TargetIds: req.GetTargetIds()})
	if err != nil {
		return nil, err
	}
	statuses := make(map[int64]*bbsuserv1.RelationStatus, len(reply.GetStatuses()))
	for id, status := range reply.GetStatuses() {
		statuses[id] = toBFFRelationStatus(status)
	}
	return &bbsuserv1.BatchGetStatusRelation_Reply{Statuses: statuses}, nil
}

func toUserPageRequest(in *bbsuserv1.PageRequest) *commonv1.PageRequest {
	if in == nil {
		return nil
	}
	return &commonv1.PageRequest{
		Page: uint32(in.GetCurrent()),
		Size: uint32(in.GetPageSize()),
	}
}

func toBFFPageReply(in *commonv1.PageReply) *bbsuserv1.PageReply {
	if in == nil {
		return nil
	}
	return &bbsuserv1.PageReply{
		Current:  int64(in.GetPage()),
		PageSize: int64(in.GetSize()),
		Total:    int64(in.GetTotal()),
	}
}

func toBFFRelations(in []*userv1.Relation) []*bbsuserv1.Relation {
	rows := make([]*bbsuserv1.Relation, 0, len(in))
	for _, item := range in {
		rows = append(rows, toBFFRelation(item))
	}
	return rows
}

func toBFFRelation(in *userv1.Relation) *bbsuserv1.Relation {
	if in == nil {
		return nil
	}
	return &bbsuserv1.Relation{
		Id:        in.GetId(),
		Type:      int32(in.GetType()),
		ActorId:   in.GetActorId(),
		TargetId:  in.GetTargetId(),
		CreatedAt: formatProtoTime(in.GetCreatedAt()),
		UpdatedAt: formatProtoTime(in.GetUpdatedAt()),
	}
}

func toBFFRelationStatus(in *userv1.RelationStatus) *bbsuserv1.RelationStatus {
	if in == nil {
		return nil
	}
	return &bbsuserv1.RelationStatus{
		TargetId:   in.GetTargetId(),
		Following:  in.GetFollowing(),
		FollowedBy: in.GetFollowedBy(),
		Blocking:   in.GetBlocking(),
		BlockedBy:  in.GetBlockedBy(),
	}
}
