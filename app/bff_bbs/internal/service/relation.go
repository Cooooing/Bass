package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
	relationmodel "common/proto/gen/bff_bbs/v1/user/model"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_bbs/internal/biz/usecase"
	"bff_bbs/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RelationService struct {
	bbsuserv1.UnimplementedRelationServiceServer
	relationUsecase *usecase.RelationUsecase
}

func NewRelationService(
	relationUsecase *usecase.RelationUsecase,
) *RelationService {
	return &RelationService{
		relationUsecase: relationUsecase,
	}
}

func (s *RelationService) RegisterGrpc(gs *grpc.Server) {
}

func (s *RelationService) RegisterHTTP(hs *http.Server) {
	bbsuserv1.RegisterRelationServiceHTTPServer(hs, s)
}

func (s *RelationService) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Req) (*bbsuserv1.FollowRelation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.relationUsecase.Follow(ctx, &usecase.FollowReq{
		ActorID:  user.ID,
		TargetID: req.GetTargetId(),
	})
	return &bbsuserv1.FollowRelation_Resp{}, err
}

func (s *RelationService) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Req) (*bbsuserv1.UnfollowRelation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.relationUsecase.Unfollow(ctx, &usecase.UnfollowReq{
		ActorID:  user.ID,
		TargetID: req.GetTargetId(),
	})
	return &bbsuserv1.UnfollowRelation_Resp{}, err
}

func (s *RelationService) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Req) (*bbsuserv1.BlockRelation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.relationUsecase.Block(ctx, &usecase.BlockReq{
		ActorID:  user.ID,
		TargetID: req.GetTargetId(),
	})
	return &bbsuserv1.BlockRelation_Resp{}, err
}

func (s *RelationService) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Req) (*bbsuserv1.UnblockRelation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.relationUsecase.Unblock(ctx, &usecase.UnblockReq{
		ActorID:  user.ID,
		TargetID: req.GetTargetId(),
	})
	return &bbsuserv1.UnblockRelation_Resp{}, err
}

func (s *RelationService) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Req) (*bbsuserv1.ListFollowingRelations_Resp, error) {
	result, err := s.listRelations(ctx, req.GetPage(), func(ctx context.Context, userID int64, page *commonmodel.PageReq) (*usecase.RelationListResp, error) {
		return s.relationUsecase.ListFollowing(ctx, &usecase.ListFollowingReq{ActorID: userID, Page: page})
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ListFollowingRelations_Resp{
		Page: result.Page.ToProto(),
		Rows: result.Rows,
	}, nil
}

func (s *RelationService) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Req) (*bbsuserv1.ListFollowersRelations_Resp, error) {
	result, err := s.listRelations(ctx, req.GetPage(), func(ctx context.Context, userID int64, page *commonmodel.PageReq) (*usecase.RelationListResp, error) {
		return s.relationUsecase.ListFollowers(ctx, &usecase.ListFollowersReq{ActorID: userID, Page: page})
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ListFollowersRelations_Resp{
		Page: result.Page.ToProto(),
		Rows: result.Rows,
	}, nil
}

func (s *RelationService) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Req) (*bbsuserv1.ListBlockedRelations_Resp, error) {
	result, err := s.listRelations(ctx, req.GetPage(), func(ctx context.Context, userID int64, page *commonmodel.PageReq) (*usecase.RelationListResp, error) {
		return s.relationUsecase.ListBlocked(ctx, &usecase.ListBlockedReq{ActorID: userID, Page: page})
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ListBlockedRelations_Resp{
		Page: result.Page.ToProto(),
		Rows: result.Rows,
	}, nil
}

type relationListResult struct {
	Page *commonmodel.PageResp
	Rows []*relationmodel.Relation
}

func (s *RelationService) listRelations(ctx context.Context, page *common.PageReq, list func(context.Context, int64, *commonmodel.PageReq) (*usecase.RelationListResp, error)) (*relationListResult, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	result, err := list(ctx, user.ID, commonmodel.PageReqFromProto(page))
	if err != nil {
		return nil, err
	}
	rows := make([]*relationmodel.Relation, 0, len(result.Rows))
	for _, row := range result.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &relationmodel.Relation{Id: row.ID, Type: relationmodel.RelationType(enum.RelationTypeMap.MustToProto(row.Type)), ActorId: row.ActorID, TargetId: row.TargetID, CreatedAt: timestamppb.New(*row.CreatedAt), UpdatedAt: timestamppb.New(*row.UpdatedAt)})
	}
	return &relationListResult{Page: result.Page, Rows: rows}, nil
}

func (s *RelationService) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Req) (*bbsuserv1.GetStatusRelation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return &bbsuserv1.GetStatusRelation_Resp{
			Status: &bbsuserv1.GetStatusRelation_Resp_RelationStatus{
				TargetId: req.GetTargetId(),
			},
		}, nil
	}
	resp, err := s.relationUsecase.GetStatus(ctx, &usecase.GetStatusReq{
		ActorID:  user.ID,
		TargetID: req.GetTargetId(),
	})
	if err != nil {
		return nil, err
	}
	var status *bbsuserv1.GetStatusRelation_Resp_RelationStatus
	if resp != nil {
		status = &bbsuserv1.GetStatusRelation_Resp_RelationStatus{
			TargetId:   resp.TargetID,
			Following:  resp.Following,
			FollowedBy: resp.FollowedBy,
			Blocking:   resp.Blocking,
			BlockedBy:  resp.BlockedBy,
		}
	}
	return &bbsuserv1.GetStatusRelation_Resp{
		Status: status,
	}, nil
}
