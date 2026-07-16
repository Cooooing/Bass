package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.RelationClient = (*RelationClient)(nil)

type RelationClient struct {
	userClient *rpc.UserClient
}

func NewRelationClient(userClient *rpc.UserClient) repo.RelationClient {
	return &RelationClient{userClient: userClient}
}

func (r *RelationClient) Follow(ctx context.Context, req *repo.FollowRelationReq) (*repo.FollowRelationResponse, error) {
	_, err := r.userClient.Relation.Follow(ctx, &userv1.FollowRelation_Request{ActorId: req.ActorID, TargetId: req.TargetID})
	if err != nil {
		return nil, err
	}
	return &repo.FollowRelationResponse{}, nil
}

func (r *RelationClient) Unfollow(ctx context.Context, req *repo.UnfollowRelationReq) (*repo.UnfollowRelationResponse, error) {
	_, err := r.userClient.Relation.Unfollow(ctx, &userv1.UnfollowRelation_Request{ActorId: req.ActorID, TargetId: req.TargetID})
	if err != nil {
		return nil, err
	}
	return &repo.UnfollowRelationResponse{}, nil
}

func (r *RelationClient) Block(ctx context.Context, req *repo.BlockRelationReq) (*repo.BlockRelationResponse, error) {
	_, err := r.userClient.Relation.Block(ctx, &userv1.BlockRelation_Request{ActorId: req.ActorID, TargetId: req.TargetID})
	if err != nil {
		return nil, err
	}
	return &repo.BlockRelationResponse{}, nil
}

func (r *RelationClient) Unblock(ctx context.Context, req *repo.UnblockRelationReq) (*repo.UnblockRelationResponse, error) {
	_, err := r.userClient.Relation.Unblock(ctx, &userv1.UnblockRelation_Request{ActorId: req.ActorID, TargetId: req.TargetID})
	if err != nil {
		return nil, err
	}
	return &repo.UnblockRelationResponse{}, nil
}

func (r *RelationClient) ListFollowing(ctx context.Context, req *repo.ListFollowingRelationsReq) (*repo.ListFollowingRelationsResponse, error) {
	var pageReq *common.PageRequest
	if req.Page != nil {
		pageReq = &common.PageRequest{Page: req.Page.Page, Size: req.Page.Size}
	}
	reply, err := r.userClient.Relation.ListFollowing(ctx, &userv1.ListFollowingRelations_Request{Page: pageReq, UserId: req.ActorID})
	if err != nil {
		return nil, err
	}
	var page *repo.PageResponse
	if reply.GetPage() != nil {
		page = &repo.PageResponse{Page: reply.GetPage().GetPage(), Size: reply.GetPage().GetSize(), Total: reply.GetPage().GetTotal()}
	}
	rows := make([]*repo.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &repo.Relation{ID: item.GetId(), Type: int32(item.GetType()), ActorID: item.GetActorId(), TargetID: item.GetTargetId(), CreatedAt: formatProtoTime(item.GetCreatedAt()), UpdatedAt: formatProtoTime(item.GetUpdatedAt())})
	}
	return &repo.ListFollowingRelationsResponse{Page: page, Rows: rows}, nil
}

func (r *RelationClient) ListFollowers(ctx context.Context, req *repo.ListFollowersRelationsReq) (*repo.ListFollowersRelationsResponse, error) {
	var pageReq *common.PageRequest
	if req.Page != nil {
		pageReq = &common.PageRequest{Page: req.Page.Page, Size: req.Page.Size}
	}
	reply, err := r.userClient.Relation.ListFollowers(ctx, &userv1.ListFollowersRelations_Request{Page: pageReq, UserId: req.ActorID})
	if err != nil {
		return nil, err
	}
	var page *repo.PageResponse
	if reply.GetPage() != nil {
		page = &repo.PageResponse{Page: reply.GetPage().GetPage(), Size: reply.GetPage().GetSize(), Total: reply.GetPage().GetTotal()}
	}
	rows := make([]*repo.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &repo.Relation{ID: item.GetId(), Type: int32(item.GetType()), ActorID: item.GetActorId(), TargetID: item.GetTargetId(), CreatedAt: formatProtoTime(item.GetCreatedAt()), UpdatedAt: formatProtoTime(item.GetUpdatedAt())})
	}
	return &repo.ListFollowersRelationsResponse{Page: page, Rows: rows}, nil
}

func (r *RelationClient) ListBlocked(ctx context.Context, req *repo.ListBlockedRelationsReq) (*repo.ListBlockedRelationsResponse, error) {
	var pageReq *common.PageRequest
	if req.Page != nil {
		pageReq = &common.PageRequest{Page: req.Page.Page, Size: req.Page.Size}
	}
	reply, err := r.userClient.Relation.ListBlocked(ctx, &userv1.ListBlockedRelations_Request{Page: pageReq, UserId: req.ActorID})
	if err != nil {
		return nil, err
	}
	var page *repo.PageResponse
	if reply.GetPage() != nil {
		page = &repo.PageResponse{Page: reply.GetPage().GetPage(), Size: reply.GetPage().GetSize(), Total: reply.GetPage().GetTotal()}
	}
	rows := make([]*repo.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &repo.Relation{ID: item.GetId(), Type: int32(item.GetType()), ActorID: item.GetActorId(), TargetID: item.GetTargetId(), CreatedAt: formatProtoTime(item.GetCreatedAt()), UpdatedAt: formatProtoTime(item.GetUpdatedAt())})
	}
	return &repo.ListBlockedRelationsResponse{Page: page, Rows: rows}, nil
}

func (r *RelationClient) GetStatus(ctx context.Context, req *repo.GetStatusRelationReq) (*repo.GetStatusRelationResponse, error) {
	reply, err := r.userClient.Relation.MapStatus(ctx, &userv1.MapRelationStatuses_Request{ActorId: req.ActorID, TargetIds: []int64{req.TargetID}})
	if err != nil {
		return nil, err
	}
	var out *repo.RelationStatus
	if status := reply.GetStatuses()[req.TargetID]; status != nil {
		out = &repo.RelationStatus{TargetID: status.GetTargetId(), Following: status.GetFollowing(), FollowedBy: status.GetFollowedBy(), Blocking: status.GetBlocking(), BlockedBy: status.GetBlockedBy()}
	}
	return &repo.GetStatusRelationResponse{Status: out}, nil
}
