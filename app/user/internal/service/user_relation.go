package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/user/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type UserRelationService struct {
	v1.UnimplementedUserUserRelationServiceServer
	*BaseService
	userRelationDomain *biz.UserRelationDomain
}

func NewUserRelationService(baseService *BaseService, userRelationDomain *biz.UserRelationDomain) *UserRelationService {
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

func (s *UserRelationService) Block(ctx context.Context, req *v1.BlockRequest) (rsp *v1.BlockReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	if user.ID == req.BlockUserId {
		return nil, cv1.ErrorBadRequest("can not block yourself")
	}
	err = s.userRelationDomain.UpdateUserRelation(ctx, v1.UserRelationType_USER_RELATION_TYPE_BLOCK, req.Block, user.ID, req.BlockUserId)
	return &v1.BlockReply{}, err
}

func (s *UserRelationService) Follow(ctx context.Context, req *v1.FollowRequest) (rsp *v1.FollowReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	if user.ID == req.FollowUserId {
		return nil, cv1.ErrorBadRequest("can not follow yourself")
	}
	err = s.userRelationDomain.UpdateUserRelation(ctx, v1.UserRelationType_USER_RELATION_TYPE_FOLLOW, req.Follow, user.ID, req.FollowUserId)
	return &v1.FollowReply{}, err
}

func (s *UserRelationService) Page(ctx context.Context, req *v1.PageUserRelationRequest) (rsp *v1.PageUserRelationReply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	req.Page = base.OrDefault(req.Page, &cv1.PageRequest{})
	userRelations, page, err := s.userRelationDomain.Page(ctx, req.Page, &repo.UserRelationGetReq{
		ActorId:    base.Ptr(user.ID),
		Type:       (*v1.UserRelationType)(req.Query.Type),
		WithTarget: true,
	})
	return &v1.PageUserRelationReply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(userRelations),
	}, nil
}
