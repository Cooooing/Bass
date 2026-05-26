package service

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz/repo"
	"user/internal/biz/usecase"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RelationService struct {
	v1.UnimplementedRelationServiceServer
	relationUsecase *usecase.RelationUsecase
	relationRepo    repo.RelationRepo
}

func NewRelationService(relationUsecase *usecase.RelationUsecase, relationRepo repo.RelationRepo) *RelationService {
	return &RelationService{
		relationUsecase: relationUsecase,
		relationRepo:    relationRepo,
	}
}

func (s *RelationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterRelationServiceServer(gs, s)
}

func (s *RelationService) RegisterHttp(hs *http.Server) {}

func (s *RelationService) Follow(ctx context.Context, req *v1.FollowRelation_Request) (*v1.FollowRelation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if current.ID == req.TargetId {
		return nil, cerrors.ErrorBadRequest("can not follow yourself")
	}
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_FOLLOW, true, current.ID, req.TargetId)
	return &v1.FollowRelation_Reply{}, err
}

func (s *RelationService) Unfollow(ctx context.Context, req *v1.UnfollowRelation_Request) (*v1.UnfollowRelation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_FOLLOW, false, current.ID, req.TargetId)
	return &v1.UnfollowRelation_Reply{}, err
}

func (s *RelationService) Block(ctx context.Context, req *v1.BlockRelation_Request) (*v1.BlockRelation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if current.ID == req.TargetId {
		return nil, cerrors.ErrorBadRequest("can not block yourself")
	}
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_BLOCK, true, current.ID, req.TargetId)
	return &v1.BlockRelation_Reply{}, err
}

func (s *RelationService) Unblock(ctx context.Context, req *v1.UnblockRelation_Request) (*v1.UnblockRelation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	err := s.relationUsecase.UpdateRelation(ctx, v1.RelationType_RELATION_TYPE_BLOCK, false, current.ID, req.TargetId)
	return &v1.UnblockRelation_Reply{}, err
}

func (s *RelationService) ListFollowing(ctx context.Context, req *v1.ListFollowingRelations_Request) (*v1.ListFollowingRelations_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	req = util.OrDefault(req, &v1.ListFollowingRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	relationType := v1.RelationType_RELATION_TYPE_FOLLOW
	rows, page, err := s.relationUsecase.Page(ctx, req.Page, &repo.RelationGetReq{ActorId: &current.ID, Type: &relationType})
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
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	req = util.OrDefault(req, &v1.ListFollowersRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	relationType := v1.RelationType_RELATION_TYPE_FOLLOW
	rows, page, err := s.relationUsecase.Page(ctx, req.Page, &repo.RelationGetReq{TargetId: &current.ID, Type: &relationType})
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

func (s *RelationService) ListFollowerIds(ctx context.Context, req *v1.ListFollowerIdsRelation_Request) (*v1.ListFollowerIdsRelation_Reply, error) {
	req = util.OrDefault(req, &v1.ListFollowerIdsRelation_Request{})
	if req.UserId == 0 {
		return nil, cerrors.ErrorBadRequest("user id is required")
	}
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	relationType := v1.RelationType_RELATION_TYPE_FOLLOW
	rows, page, err := s.relationUsecase.Page(ctx, req.Page, &repo.RelationGetReq{TargetId: &req.UserId, Type: &relationType})
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(rows))
	for _, row := range rows {
		userIDs = append(userIDs, row.ActorID)
	}
	return &v1.ListFollowerIdsRelation_Reply{Page: page, UserIds: userIDs}, nil
}

func (s *RelationService) ListBlocked(ctx context.Context, req *v1.ListBlockedRelations_Request) (*v1.ListBlockedRelations_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	req = util.OrDefault(req, &v1.ListBlockedRelations_Request{})
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	relationType := v1.RelationType_RELATION_TYPE_BLOCK
	rows, page, err := s.relationUsecase.Page(ctx, req.Page, &repo.RelationGetReq{ActorId: &current.ID, Type: &relationType})
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

func (s *RelationService) BatchGetStatus(ctx context.Context, req *v1.BatchGetStatusRelation_Request) (*v1.BatchGetStatusRelation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	statuses := make(map[int64]*v1.RelationStatus, len(req.TargetIds))
	for _, targetID := range req.TargetIds {
		statuses[targetID] = &v1.RelationStatus{TargetId: targetID}
	}
	rows, err := s.relationRepo.List(ctx, &repo.RelationGetReq{ActorId: &current.ID})
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		if status, ok := statuses[row.TargetID]; ok {
			switch row.Type {
			case enum.RelationTypeFollow:
				status.Following = true
			case enum.RelationTypeBlock:
				status.Blocking = true
			}
		}
	}
	rows, err = s.relationRepo.List(ctx, &repo.RelationGetReq{TargetId: &current.ID})
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		if status, ok := statuses[row.ActorID]; ok {
			switch row.Type {
			case enum.RelationTypeFollow:
				status.FollowedBy = true
			case enum.RelationTypeBlock:
				status.BlockedBy = true
			}
		}
	}
	return &v1.BatchGetStatusRelation_Reply{Statuses: statuses}, nil
}
