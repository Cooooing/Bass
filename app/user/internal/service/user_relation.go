package service

import (
	"common/api/gen/common"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"user/internal/biz/doamin"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type UserRelationService struct {
	v1.UnimplementedUserUserRelationServiceServer
	*BaseService
	userRelationDomain *doamin.UserRelationDomain
}

func NewUserRelationService(baseService *BaseService, userRelationDomain *doamin.UserRelationDomain) *UserRelationService {
	return &UserRelationService{
		BaseService:        baseService,
		userRelationDomain: userRelationDomain,
	}
}

func (s *UserRelationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserUserRelationServiceServer(gs, s)
}

func (s *UserRelationService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserUserRelationServiceHTTPServer(hs, s)
}

func (s *UserRelationService) Block(ctx context.Context, req *v1.BlockUserRelation_Request) (rsp *v1.BlockUserRelation_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	if user.ID == req.BlockUserId {
		return nil, common.ErrorBadRequest("can not block yourself")
	}
	err = s.userRelationDomain.UpdateUserRelation(ctx, v1.UserRelationType_USER_RELATION_TYPE_BLOCK, req.Block, user.ID, req.BlockUserId)
	return &v1.BlockUserRelation_Reply{}, err
}

func (s *UserRelationService) Follow(ctx context.Context, req *v1.FollowUserRelation_Request) (rsp *v1.FollowUserRelation_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	if user.ID == req.FollowUserId {
		return nil, common.ErrorBadRequest("can not follow yourself")
	}
	err = s.userRelationDomain.UpdateUserRelation(ctx, v1.UserRelationType_USER_RELATION_TYPE_FOLLOW, req.Follow, user.ID, req.FollowUserId)
	return &v1.FollowUserRelation_Reply{}, err
}

func (s *UserRelationService) Page(ctx context.Context, req *v1.PageUserRelation_Request) (rsp *v1.PageUserRelation_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	req.Page = util.OrDefault(req.Page, &common.PageRequest{})
	userRelations, page, err := s.userRelationDomain.Page(ctx, req.Page, &repo.UserRelationGetReq{
		ActorId:    new(user.ID),
		Type:       req.Query.Type,
		WithTarget: true,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PageUserRelation_Reply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(userRelations),
	}, nil
}
