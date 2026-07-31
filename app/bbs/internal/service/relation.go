package service

import (
	"bbs/internal/biz/usecase"
	"bbs/internal/enum"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"

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

func (s *RelationService) RegisterHttp(hs *http.Server) {
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
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.relationUsecase.ListFollowing(ctx, &usecase.ListFollowingReq{
		ActorID: user.ID,
		Page:    req.GetPage(),
	})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{
			Page:  resp.Page.Page,
			Size:  resp.Page.Size,
			Total: resp.Page.Total,
		}
	}
	rows := make([]*bbsuserv1.ListFollowingRelations_Resp_Relation, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.ListFollowingRelations_Resp_Relation{
			Id:        row.ID,
			Type:      enum.RelationTypeMap.MustToProto(row.Type),
			ActorId:   row.ActorID,
			TargetId:  row.TargetID,
			CreatedAt: timestamppb.New(*row.CreatedAt),
			UpdatedAt: timestamppb.New(*row.UpdatedAt),
		})
	}
	return &bbsuserv1.ListFollowingRelations_Resp{
		Page: page,
		Rows: rows,
	}, nil
}

func (s *RelationService) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Req) (*bbsuserv1.ListFollowersRelations_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.relationUsecase.ListFollowers(ctx, &usecase.ListFollowersReq{
		ActorID: user.ID,
		Page:    req.GetPage(),
	})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{
			Page:  resp.Page.Page,
			Size:  resp.Page.Size,
			Total: resp.Page.Total,
		}
	}
	rows := make([]*bbsuserv1.ListFollowersRelations_Resp_Relation, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.ListFollowersRelations_Resp_Relation{
			Id:        row.ID,
			Type:      enum.RelationTypeMap.MustToProto(row.Type),
			ActorId:   row.ActorID,
			TargetId:  row.TargetID,
			CreatedAt: timestamppb.New(*row.CreatedAt),
			UpdatedAt: timestamppb.New(*row.UpdatedAt),
		})
	}
	return &bbsuserv1.ListFollowersRelations_Resp{
		Page: page,
		Rows: rows,
	}, nil
}

func (s *RelationService) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Req) (*bbsuserv1.ListBlockedRelations_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.relationUsecase.ListBlocked(ctx, &usecase.ListBlockedReq{
		ActorID: user.ID,
		Page:    req.GetPage(),
	})
	if err != nil {
		return nil, err
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{
			Page:  resp.Page.Page,
			Size:  resp.Page.Size,
			Total: resp.Page.Total,
		}
	}
	rows := make([]*bbsuserv1.ListBlockedRelations_Resp_Relation, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		if row == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.ListBlockedRelations_Resp_Relation{
			Id:        row.ID,
			Type:      enum.RelationTypeMap.MustToProto(row.Type),
			ActorId:   row.ActorID,
			TargetId:  row.TargetID,
			CreatedAt: timestamppb.New(*row.CreatedAt),
			UpdatedAt: timestamppb.New(*row.UpdatedAt),
		})
	}
	return &bbsuserv1.ListBlockedRelations_Resp{
		Page: page,
		Rows: rows,
	}, nil
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
