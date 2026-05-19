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

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type UserRelationService struct {
	v1.UnimplementedUserUserRelationServiceServer
	userRelationUsecase *usecase.UserRelationUsecase
}

func NewUserRelationService(userRelationUsecase *usecase.UserRelationUsecase) *UserRelationService {
	return &UserRelationService{
		userRelationUsecase: userRelationUsecase,
	}
}

func (s *UserRelationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserUserRelationServiceServer(gs, s)
}

func (s *UserRelationService) RegisterHttp(hs *http.Server) {}

func (s *UserRelationService) Block(ctx context.Context, req *v1.BlockUserRelation_Request) (rsp *v1.BlockUserRelation_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if user.ID == req.BlockUserId {
		return nil, cerrors.ErrorBadRequest("can not block yourself")
	}
	err = s.userRelationUsecase.UpdateUserRelation(ctx, v1.UserRelationType_USER_RELATION_TYPE_BLOCK, req.Block, user.ID, req.BlockUserId)
	return &v1.BlockUserRelation_Reply{}, err
}

func (s *UserRelationService) Follow(ctx context.Context, req *v1.FollowUserRelation_Request) (rsp *v1.FollowUserRelation_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	if user.ID == req.FollowUserId {
		return nil, cerrors.ErrorBadRequest("can not follow yourself")
	}
	err = s.userRelationUsecase.UpdateUserRelation(ctx, v1.UserRelationType_USER_RELATION_TYPE_FOLLOW, req.Follow, user.ID, req.FollowUserId)
	return &v1.FollowUserRelation_Reply{}, err
}

func (s *UserRelationService) Page(ctx context.Context, req *v1.PageUserRelation_Request) (rsp *v1.PageUserRelation_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	userRelations, page, err := s.userRelationUsecase.Page(ctx, req.Page, &repo.UserRelationGetReq{
		ActorId:    new(user.ID),
		Type:       req.Query.Type,
		WithTarget: true,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PageUserRelation_Reply{
		Page: page,
		Rows: userRelationsToProto(userRelations),
	}, nil
}
