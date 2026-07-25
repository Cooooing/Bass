package data

import (
	"bbs/internal/biz/repo"
	"bbs/internal/enum"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"
	userv1enum "common/proto/gen/user/v1/enum"
	"context"
)

var _ repo.RelationClient = (*RelationClient)(nil)

type RelationClient struct {
	protoTimeFormatter
	userClient *rpc.UserClient
}

func NewRelationClient(
	userClient *rpc.UserClient,
) repo.RelationClient {
	return &RelationClient{
		userClient: userClient,
	}
}

func (r *RelationClient) Follow(ctx context.Context, req *repo.FollowRelationReq) error {
	_, err := r.userClient.Relation.Follow(ctx, &userv1.FollowRelation_Req{
		ActorId:  req.ActorID,
		TargetId: req.TargetID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RelationClient) Unfollow(ctx context.Context, req *repo.UnfollowRelationReq) error {
	_, err := r.userClient.Relation.Unfollow(ctx, &userv1.UnfollowRelation_Req{
		ActorId:  req.ActorID,
		TargetId: req.TargetID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RelationClient) Block(ctx context.Context, req *repo.BlockRelationReq) error {
	_, err := r.userClient.Relation.Block(ctx, &userv1.BlockRelation_Req{
		ActorId:  req.ActorID,
		TargetId: req.TargetID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RelationClient) Unblock(ctx context.Context, req *repo.UnblockRelationReq) error {
	_, err := r.userClient.Relation.Unblock(ctx, &userv1.UnblockRelation_Req{
		ActorId:  req.ActorID,
		TargetId: req.TargetID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RelationClient) ListFollowing(ctx context.Context, req *repo.ListFollowingRelationsReq) (*repo.ListFollowingRelationsResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.userClient.Relation.ListFollowing(ctx, &userv1.ListFollowingRelations_Req{
		Page:   pageReq,
		UserId: req.ActorID,
	})
	if err != nil {
		return nil, err
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	rows := make([]*repo.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &repo.Relation{
			ID:        item.GetId(),
			Type:      r.relationTypeFromUser(item.GetType()),
			ActorID:   item.GetActorId(),
			TargetID:  item.GetTargetId(),
			CreatedAt: r.formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: r.formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &repo.ListFollowingRelationsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *RelationClient) ListFollowers(ctx context.Context, req *repo.ListFollowersRelationsReq) (*repo.ListFollowersRelationsResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.userClient.Relation.ListFollowers(ctx, &userv1.ListFollowersRelations_Req{
		Page:   pageReq,
		UserId: req.ActorID,
	})
	if err != nil {
		return nil, err
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	rows := make([]*repo.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &repo.Relation{
			ID:        item.GetId(),
			Type:      r.relationTypeFromUser(item.GetType()),
			ActorID:   item.GetActorId(),
			TargetID:  item.GetTargetId(),
			CreatedAt: r.formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: r.formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &repo.ListFollowersRelationsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *RelationClient) ListBlocked(ctx context.Context, req *repo.ListBlockedRelationsReq) (*repo.ListBlockedRelationsResp, error) {
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.userClient.Relation.ListBlocked(ctx, &userv1.ListBlockedRelations_Req{
		Page:   pageReq,
		UserId: req.ActorID,
	})
	if err != nil {
		return nil, err
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	rows := make([]*repo.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &repo.Relation{
			ID:        item.GetId(),
			Type:      r.relationTypeFromUser(item.GetType()),
			ActorID:   item.GetActorId(),
			TargetID:  item.GetTargetId(),
			CreatedAt: r.formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: r.formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &repo.ListBlockedRelationsResp{
		Page: page,
		Rows: rows,
	}, nil
}

func (r *RelationClient) GetStatus(ctx context.Context, req *repo.GetStatusRelationReq) (*repo.RelationStatus, error) {
	reply, err := r.userClient.Relation.MapStatus(ctx, &userv1.MapRelationStatuses_Req{
		ActorId:   req.ActorID,
		TargetIds: []int64{req.TargetID},
	})
	if err != nil {
		return nil, err
	}
	var out *repo.RelationStatus
	if status := reply.GetStatuses()[req.TargetID]; status != nil {
		out = &repo.RelationStatus{
			TargetID:   status.GetTargetId(),
			Following:  status.GetFollowing(),
			FollowedBy: status.GetFollowedBy(),
			Blocking:   status.GetBlocking(),
			BlockedBy:  status.GetBlockedBy(),
		}
	}
	return out, nil
}

func (r *RelationClient) relationTypeFromUser(value userv1enum.RelationType) enum.RelationType {
	switch value {
	case userv1enum.RelationType_RELATION_TYPE_BLOCK:
		return enum.RelationTypeBlock
	default:
		return enum.RelationTypeFollow
	}
}
