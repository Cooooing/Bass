package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type RelationService struct {
	bbsuserv1.UnimplementedRelationServiceServer
	relationUsecase *usecase.RelationUsecase
}

func NewRelationService(relationUsecase *usecase.RelationUsecase) *RelationService {
	return &RelationService{relationUsecase: relationUsecase}
}
func (s *RelationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterRelationServiceHTTPServer(hs, s)
}
func (s *RelationService) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.relationUsecase.Follow(ctx, &usecase.FollowReq{ActorID: actorID, TargetID: req.GetTargetId()})
	return &bbsuserv1.FollowRelation_Response{}, err
}
func (s *RelationService) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.relationUsecase.Unfollow(ctx, &usecase.UnfollowReq{ActorID: actorID, TargetID: req.GetTargetId()})
	return &bbsuserv1.UnfollowRelation_Response{}, err
}
func (s *RelationService) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.relationUsecase.Block(ctx, &usecase.BlockReq{ActorID: actorID, TargetID: req.GetTargetId()})
	return &bbsuserv1.BlockRelation_Response{}, err
}
func (s *RelationService) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.relationUsecase.Unblock(ctx, &usecase.UnblockReq{ActorID: actorID, TargetID: req.GetTargetId()})
	return &bbsuserv1.UnblockRelation_Response{}, err
}
func (s *RelationService) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.relationUsecase.ListFollowing(ctx, &usecase.ListFollowingReq{ActorID: actorID, Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbsuserv1.ListFollowingRelations_Response_Relation, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.ListFollowingRelations_Response_Relation{Id: row.ID, Type: userv1.RelationType(row.Type), ActorId: row.ActorID, TargetId: row.TargetID, CreatedAt: protoTime(row.CreatedAt), UpdatedAt: protoTime(row.UpdatedAt)})
	}
	return &bbsuserv1.ListFollowingRelations_Response{Page: page, Rows: rows}, nil
}
func (s *RelationService) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.relationUsecase.ListFollowers(ctx, &usecase.ListFollowersReq{ActorID: actorID, Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbsuserv1.ListFollowersRelations_Response_Relation, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.ListFollowersRelations_Response_Relation{Id: row.ID, Type: userv1.RelationType(row.Type), ActorId: row.ActorID, TargetId: row.TargetID, CreatedAt: protoTime(row.CreatedAt), UpdatedAt: protoTime(row.UpdatedAt)})
	}
	return &bbsuserv1.ListFollowersRelations_Response{Page: page, Rows: rows}, nil
}
func (s *RelationService) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.relationUsecase.ListBlocked(ctx, &usecase.ListBlockedReq{ActorID: actorID, Page: req.GetPage()})
	if err != nil {
		return nil, err
	}
	var page *common.PageResponse
	if response.Page != nil {
		page = &common.PageResponse{Page: response.Page.Page, Size: response.Page.Size, Total: response.Page.Total}
	}
	rows := make([]*bbsuserv1.ListBlockedRelations_Response_Relation, 0, len(response.Rows))
	for _, row := range response.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.ListBlockedRelations_Response_Relation{Id: row.ID, Type: userv1.RelationType(row.Type), ActorId: row.ActorID, TargetId: row.TargetID, CreatedAt: protoTime(row.CreatedAt), UpdatedAt: protoTime(row.UpdatedAt)})
	}
	return &bbsuserv1.ListBlockedRelations_Response{Page: page, Rows: rows}, nil
}
func (s *RelationService) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Response, error) {
	actorID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.relationUsecase.GetStatus(ctx, &usecase.GetStatusReq{ActorID: actorID, TargetID: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	var status *bbsuserv1.GetStatusRelation_Response_RelationStatus
	if response.Status != nil {
		status = &bbsuserv1.GetStatusRelation_Response_RelationStatus{TargetId: response.Status.TargetID, Following: response.Status.Following, FollowedBy: response.Status.FollowedBy, Blocking: response.Status.Blocking, BlockedBy: response.Status.BlockedBy}
	}
	return &bbsuserv1.GetStatusRelation_Response{Status: status}, nil
}
