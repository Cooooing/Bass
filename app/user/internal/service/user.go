package service

import (
	v1 "common/api/user/v1"
	"common/pkg/constant"
	"context"
	"user/internal/biz"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	v1.UnimplementedUserUserServiceServer
	*BaseService
	authenticationDomain *biz.AuthenticationDomain
	userRepo             repo.UserRepo
}

func NewUserService(baseService *BaseService, authenticationDomain *biz.AuthenticationDomain, userRepo repo.UserRepo) *UserService {
	return &UserService{
		BaseService:          baseService,
		authenticationDomain: authenticationDomain,
		userRepo:             userRepo,
	}
}

func (s *UserService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserUserServiceServer(gs, s)
}

func (s *UserService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserUserServiceHTTPServer(hs, s)
}

func (s *UserService) GetList(ctx context.Context, req *v1.GetListRequest) (rsp *v1.GetListReply, err error) {
	res := &v1.GetListReply{
		Users: []*v1.User{},
	}
	list, err := s.userRepo.GetUserList(ctx, s.db, constant.GetPageMax(), req.Ids)
	if err != nil {
		return nil, err
	}
	for i := range list {
		item := &v1.User{}
		err = copier.Copy(&item, list[i])
		if err != nil {
			return nil, err
		}
		if list[i].LastLoginTime != nil {
			item.LastLoginTime = timestamppb.New(*list[i].LastLoginTime)
		}
		if list[i].LastCheckinTime != nil {
			item.LastCheckinTime = timestamppb.New(*list[i].LastCheckinTime)
		}
		item.CreatedAt = timestamppb.New(*list[i].CreatedAt)
		item.UpdatedAt = timestamppb.New(*list[i].UpdatedAt)
		res.Users = append(res.Users, item)
	}
	return res, nil
}

func (s *UserService) GetMap(ctx context.Context, req *v1.GetMapRequest) (rsp *v1.GetMapReply, err error) {
	res := &v1.GetMapReply{
		Users: make(map[int64]*v1.User),
	}
	list, err := s.userRepo.GetUserList(ctx, s.db, constant.GetPageMax(), req.Ids)
	if err != nil {
		return nil, err
	}
	for i := range list {
		item := &v1.User{}
		err = copier.Copy(&item, list[i])
		if err != nil {
			return nil, err
		}
		if list[i].LastLoginTime != nil {
			item.LastLoginTime = timestamppb.New(*list[i].LastLoginTime)
		}
		if list[i].LastCheckinTime != nil {
			item.LastCheckinTime = timestamppb.New(*list[i].LastCheckinTime)
		}
		item.CreatedAt = timestamppb.New(*list[i].CreatedAt)
		item.UpdatedAt = timestamppb.New(*list[i].UpdatedAt)
		res.Users[list[i].ID] = item
	}
	return res, nil
}

func (s *UserService) GetOne(ctx context.Context, req *v1.GetOneRequest) (rsp *v1.GetOneReply, err error) {
	res := &v1.GetOneReply{
		User: &v1.User{},
	}
	user, err := s.userRepo.GetUserById(ctx, s.db, req.Id)
	if err != nil {
		return nil, err
	}
	err = copier.Copy(res.User, user)
	if err != nil {
		return nil, err
	}
	if user.LastLoginTime != nil {
		res.User.LastLoginTime = timestamppb.New(*user.LastLoginTime)
	}
	if user.LastCheckinTime != nil {
		res.User.LastCheckinTime = timestamppb.New(*user.LastCheckinTime)
	}
	res.User.CreatedAt = timestamppb.New(*user.CreatedAt)
	res.User.UpdatedAt = timestamppb.New(*user.UpdatedAt)
	return res, nil
}
