package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
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

func (r *RelationClient) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Relation.Follow(ctx, &userv1.FollowRelation_Request{ActorId: userID, TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.FollowRelation_Reply{}, nil
}

func (r *RelationClient) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Relation.Unfollow(ctx, &userv1.UnfollowRelation_Request{ActorId: userID, TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UnfollowRelation_Reply{}, nil
}

func (r *RelationClient) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Relation.Block(ctx, &userv1.BlockRelation_Request{ActorId: userID, TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BlockRelation_Reply{}, nil
}

func (r *RelationClient) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.userClient.Relation.Unblock(ctx, &userv1.UnblockRelation_Request{ActorId: userID, TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UnblockRelation_Reply{}, nil
}

func (r *RelationClient) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Relation.ListFollowing(ctx, &userv1.ListFollowingRelations_Request{Page: req.GetPage(), UserId: userID})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbsuserv1.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.Relation{
			Id:        item.GetId(),
			Type:      int32(item.GetType()),
			ActorId:   item.GetActorId(),
			TargetId:  item.GetTargetId(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsuserv1.ListFollowingRelations_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *RelationClient) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Relation.ListFollowers(ctx, &userv1.ListFollowersRelations_Request{Page: req.GetPage(), UserId: userID})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbsuserv1.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.Relation{
			Id:        item.GetId(),
			Type:      int32(item.GetType()),
			ActorId:   item.GetActorId(),
			TargetId:  item.GetTargetId(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsuserv1.ListFollowersRelations_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *RelationClient) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Relation.ListBlocked(ctx, &userv1.ListBlockedRelations_Request{Page: req.GetPage(), UserId: userID})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbsuserv1.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.Relation{
			Id:        item.GetId(),
			Type:      int32(item.GetType()),
			ActorId:   item.GetActorId(),
			TargetId:  item.GetTargetId(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsuserv1.ListBlockedRelations_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *RelationClient) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Relation.MapStatus(ctx, &userv1.MapRelationStatuses_Request{ActorId: userID, TargetIds: []int64{req.GetTargetId()}})
	if err != nil {
		return nil, err
	}
	var out *bbsuserv1.RelationStatus
	if status := reply.GetStatuses()[req.GetTargetId()]; status != nil {
		out = &bbsuserv1.RelationStatus{
			TargetId:   status.GetTargetId(),
			Following:  status.GetFollowing(),
			FollowedBy: status.GetFollowedBy(),
			Blocking:   status.GetBlocking(),
			BlockedBy:  status.GetBlockedBy(),
		}
	}
	return &bbsuserv1.GetStatusRelation_Reply{Status: out}, nil
}
