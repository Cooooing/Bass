package service

import (
	"common/api/gen/common"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"user/internal/biz/doamin"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	v1.UnimplementedUserUserServiceServer
	*BaseService
	authenticationDomain *doamin.AuthenticationDomain
	userDomain           *doamin.UserDomain
	userRepo             repo.UserRepo
}

func NewUserService(baseService *BaseService, authenticationDomain *doamin.AuthenticationDomain, userDomain *doamin.UserDomain, userRepo repo.UserRepo) *UserService {
	return &UserService{
		BaseService:          baseService,
		authenticationDomain: authenticationDomain,
		userDomain:           userDomain,
		userRepo:             userRepo,
	}
}

func (s *UserService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserUserServiceServer(gs, s)
}

func (s *UserService) RegisterHttp(hs *http.Server) {
	v1.RegisterUserUserServiceHTTPServer(hs, s)
}

func (s *UserService) UpdateSetting(ctx context.Context, req *v1.UpdateSettingUser_Request) (rsp *v1.UpdateSettingUser_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	update, err := s.userRepo.Update(ctx, s.Db, &model.User{User: &gen.User{
		ID:                   user.ID,
		AvatarURL:            req.AvatarUrl,
		Language:             req.Language,
		Nickname:             req.Nickname,
		Timezone:             req.Timezone,
		Theme:                req.Theme,
		MobileTheme:          req.MobileTheme,
		EnableWebNotify:      req.EnableWebNotify,
		EnableEmailSubscribe: req.EnableEmailSubscribe,
	}})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateSettingUser_Reply{
		User: update.ConvertToRpc(),
	}, err
}

func (s *UserService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserUser_Request) (rsp *v1.GetCurrentUserUser_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, common.ErrorUnauthorized("user not login")
	}
	u, err := s.userRepo.GetOne(ctx, s.Db, &repo.UserGetReq{UserId: new(user.ID)})
	if err != nil {
		return nil, err
	}
	return &v1.GetCurrentUserUser_Reply{
		User: u.ConvertToRpc(),
	}, err
}

func (s *UserService) GetOne(ctx context.Context, req *v1.GetOneUser_Request) (rsp *v1.GetOneUser_Reply, err error) {
	res := &v1.GetOneUser_Reply{
		User: &v1.User{},
	}
	user, err := s.userRepo.GetOne(ctx, s.Db, &repo.UserGetReq{
		UserId: req.UserId,
		Name:   req.Name,
	})
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

func (s *UserService) GetList(ctx context.Context, req *v1.GetListUser_Request) (rsp *v1.GetListUser_Reply, err error) {
	res := &v1.GetListUser_Reply{
		Users: []*v1.User{},
	}
	list, err := s.userRepo.GetList(ctx, s.Db, &repo.UserGetReq{
		UserIds:  req.Query.UserIds,
		Name:     req.Query.Name,
		Nickname: req.Query.Nickname,
		Email:    req.Query.Email,
		Phone:    req.Query.Phone,
	})
	if err != nil {
		return nil, err
	}
	for i := range list {
		item := &v1.User{}
		err = copier.Copy(item, list[i])
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

func (s *UserService) GetMap(ctx context.Context, req *v1.GetMapUser_Request) (rsp *v1.GetMapUser_Reply, err error) {
	res := &v1.GetMapUser_Reply{
		Users: make(map[int64]*v1.User),
	}
	list, err := s.userRepo.GetList(ctx, s.Db, &repo.UserGetReq{
		UserIds:  req.Query.UserIds,
		Name:     req.Query.Name,
		Nickname: req.Query.Nickname,
		Email:    req.Query.Email,
		Phone:    req.Query.Phone,
	})
	if err != nil {
		return nil, err
	}
	for i := range list {
		item := &v1.User{}
		err = copier.Copy(item, list[i])
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

func (s *UserService) Page(ctx context.Context, req *v1.PageUser_Request) (rsp *v1.PageUser_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.UserQueryParams{})
	reply := make([]*v1.User, 0)
	users, page, err := s.userRepo.GetPage(ctx, s.Db, req.Page, &repo.UserGetReq{
		UserIds:   req.Query.UserIds,
		Name:      req.Query.Name,
		Names:     req.Query.Names,
		Nickname:  req.Query.Nickname,
		Nicknames: req.Query.Nicknames,
		Email:     req.Query.Email,
		Emails:    req.Query.Emails,
		Phone:     req.Query.Phone,
		Phones:    req.Query.Phones,
	})
	for _, user := range users {
		reply = append(reply, user.ConvertToRpc())
	}
	return &v1.PageUser_Reply{
		Page: page,
		Rows: reply,
	}, nil
}

func (s *UserService) Avatar(ctx context.Context, req *v1.AvatarUser_Request) (rsp *common.ImageReply, err error) {
	buf, err := s.userDomain.Avatar(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &common.ImageReply{
		Data:        buf,
		ContentType: "image/png",
	}, nil
}
