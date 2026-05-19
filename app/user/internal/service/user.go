package service

import (
	"common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/biz/usecase"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type UserService struct {
	v1.UnimplementedUserUserServiceServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	authUsecase         *usecase.AuthUsecase
	userUsecase         *usecase.UserUsecase
	userRepo            repo.UserRepo
	userPreferencesRepo repo.UserPreferencesRepo
	userPrivacyRepo     repo.UserPrivacyRepo
	userLocationRepo    repo.UserLocationRepo
	userTfaRepo         repo.UserTfaRepo
	userCheckinRepo     repo.UserCheckinRepo
}

func NewUserService(conf *conf.Bootstrap, logger log.Logger, authUsecase *usecase.AuthUsecase,
	userUsecase *usecase.UserUsecase,
	userRepo repo.UserRepo,
	userPreferencesRepo repo.UserPreferencesRepo,
	userPrivacyRepo repo.UserPrivacyRepo,
	userLocationRepo repo.UserLocationRepo,
	userTfaRepo repo.UserTfaRepo,
	userCheckinRepo repo.UserCheckinRepo) *UserService {
	return &UserService{
		conf:                conf,
		log:                 log.NewHelper(logger),
		authUsecase:         authUsecase,
		userUsecase:         userUsecase,
		userRepo:            userRepo,
		userPreferencesRepo: userPreferencesRepo,
		userPrivacyRepo:     userPrivacyRepo,
		userLocationRepo:    userLocationRepo,
		userTfaRepo:         userTfaRepo,
		userCheckinRepo:     userCheckinRepo,
	}
}

func (s *UserService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterUserUserServiceServer(gs, s)
}

func (s *UserService) RegisterHttp(hs *http.Server) {}

func (s *UserService) UpdateSetting(ctx context.Context, req *v1.UpdateSettingUser_Request) (rsp *v1.UpdateSettingUser_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	// Update base user fields
	_, err = s.userRepo.Update(ctx, &model.User{
		ID:        user.ID,
		AvatarURL: req.AvatarUrl,
		Nickname:  req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	// Update user preferences
	prefs, err := s.userPreferencesRepo.GetByUserID(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	prefs.Language = req.Language
	prefs.Timezone = req.Timezone
	prefs.Theme = req.Theme
	prefs.MobileTheme = req.MobileTheme
	prefs.EnableWebNotify = req.EnableWebNotify
	prefs.EnableEmailSubscribe = req.EnableEmailSubscribe
	_, err = s.userPreferencesRepo.Update(ctx, prefs)
	if err != nil {
		return nil, err
	}
	// Re-fetch user with edges for a complete reply
	fullUser, err := s.userRepo.GetOne(ctx, &repo.UserGetReq{UserId: &user.ID})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateSettingUser_Reply{
		User: assembleUserProto(ctx, fullUser, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo),
	}, nil
}

func (s *UserService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserUser_Request) (rsp *v1.GetCurrentUserUser_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	u, err := s.userRepo.GetOne(ctx, &repo.UserGetReq{UserId: new(user.ID)})
	if err != nil {
		return nil, err
	}
	return &v1.GetCurrentUserUser_Reply{
		User: assembleUserProto(ctx, u, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo),
	}, err
}

func (s *UserService) GetOne(ctx context.Context, req *v1.GetOneUser_Request) (rsp *v1.GetOneUser_Reply, err error) {
	user, err := s.userRepo.GetOne(ctx, &repo.UserGetReq{
		UserId: req.UserId,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetOneUser_Reply{
		User: assembleUserProto(ctx, user, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo),
	}, nil
}

func (s *UserService) GetList(ctx context.Context, req *v1.GetListUser_Request) (rsp *v1.GetListUser_Reply, err error) {
	list, err := s.userRepo.GetList(ctx, &repo.UserGetReq{
		UserIds:  req.Query.UserIds,
		Name:     req.Query.Name,
		Nickname: req.Query.Nickname,
		Email:    req.Query.Email,
		Phone:    req.Query.Phone,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*v1.User, 0, len(list))
	for _, u := range list {
		users = append(users, assembleUserProto(ctx, u, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo))
	}
	return &v1.GetListUser_Reply{Users: users}, nil
}

func (s *UserService) GetMap(ctx context.Context, req *v1.GetMapUser_Request) (rsp *v1.GetMapUser_Reply, err error) {
	list, err := s.userRepo.GetList(ctx, &repo.UserGetReq{
		UserIds:  req.Query.UserIds,
		Name:     req.Query.Name,
		Nickname: req.Query.Nickname,
		Email:    req.Query.Email,
		Phone:    req.Query.Phone,
	})
	if err != nil {
		return nil, err
	}
	users := make(map[int64]*v1.User)
	for _, u := range list {
		users[u.ID] = assembleUserProto(ctx, u, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo)
	}
	return &v1.GetMapUser_Reply{Users: users}, nil
}

func (s *UserService) Page(ctx context.Context, req *v1.PageUser_Request) (rsp *v1.PageUser_Reply, err error) {
	req.Query = util.OrDefault(req.Query, &v1.UserQueryParams{})
	reply := make([]*v1.User, 0)
	users, page, err := s.userRepo.GetPage(ctx, req.Page, &repo.UserGetReq{
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
		reply = append(reply, assembleUserProto(ctx, user, s.userPreferencesRepo, s.userPrivacyRepo, s.userLocationRepo, s.userTfaRepo, s.userCheckinRepo))
	}
	return &v1.PageUser_Reply{
		Page: page,
		Rows: reply,
	}, nil
}

func (s *UserService) Avatar(ctx context.Context, req *v1.AvatarUser_Request) (rsp *common.ImageReply, err error) {
	buf, err := s.userUsecase.Avatar(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &common.ImageReply{
		Data:        buf,
		ContentType: "image/png",
	}, nil
}
