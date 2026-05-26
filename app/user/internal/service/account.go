package service

import (
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

type AccountService struct {
	v1.UnimplementedAccountServiceServer
	accountValidationUsecase *usecase.AccountValidationUsecase
	accountRepo              repo.AccountRepo
}

func NewAccountService(
	accountValidationUsecase *usecase.AccountValidationUsecase,
	accountRepo repo.AccountRepo,
) *AccountService {
	return &AccountService{
		accountValidationUsecase: accountValidationUsecase,
		accountRepo:              accountRepo,
	}
}

func (s *AccountService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAccountServiceServer(gs, s)
}

func (s *AccountService) RegisterHttp(hs *http.Server) {}

func (s *AccountService) GetCurrent(ctx context.Context, req *v1.GetCurrentAccount_Request) (*v1.GetCurrentAccount_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &current.ID})
	if err != nil {
		return nil, err
	}
	basic := &v1.AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarUrl:     account.AvatarURL,
		Introduction:  account.Introduction,
		FollowCount:   account.FollowCount,
		FollowerCount: account.FollowerCount,
	}
	if account.Mbti != nil {
		basic.Mbti = enum.MBTIMap.MustToProto(*account.Mbti)
	}
	if account.Status != nil {
		basic.Status = enum.AccountStatusMap.MustToProto(*account.Status)
	}
	if account.CreatedAt != nil {
		basic.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		basic.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return &v1.GetCurrentAccount_Reply{
		Account: &v1.Account{
			Basic: basic,
			Contact: &v1.AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		},
	}, nil
}

func (s *AccountService) GetBasic(ctx context.Context, req *v1.GetBasicAccount_Request) (*v1.GetBasicAccount_Reply, error) {
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.UserId})
	if err != nil {
		return nil, err
	}
	basic := &v1.AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarUrl:     account.AvatarURL,
		Introduction:  account.Introduction,
		FollowCount:   account.FollowCount,
		FollowerCount: account.FollowerCount,
	}
	if account.Mbti != nil {
		basic.Mbti = enum.MBTIMap.MustToProto(*account.Mbti)
	}
	if account.Status != nil {
		basic.Status = enum.AccountStatusMap.MustToProto(*account.Status)
	}
	if account.CreatedAt != nil {
		basic.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		basic.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return &v1.GetBasicAccount_Reply{Account: basic}, nil
}

func (s *AccountService) BatchGetBasic(ctx context.Context, req *v1.BatchGetBasicAccount_Request) (*v1.BatchGetBasicAccount_Reply, error) {
	req = util.OrDefault(req, &v1.BatchGetBasicAccount_Request{})
	accounts, err := s.accountRepo.List(ctx, &repo.AccountGetReq{UserIds: req.UserIds})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*v1.AccountBasic, len(accounts))
	for _, account := range accounts {
		basic := &v1.AccountBasic{
			Id:            account.ID,
			Name:          account.Name,
			Nickname:      account.Nickname,
			Url:           account.URL,
			AvatarUrl:     account.AvatarURL,
			Introduction:  account.Introduction,
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
		}
		if account.Mbti != nil {
			basic.Mbti = enum.MBTIMap.MustToProto(*account.Mbti)
		}
		if account.Status != nil {
			basic.Status = enum.AccountStatusMap.MustToProto(*account.Status)
		}
		if account.CreatedAt != nil {
			basic.CreatedAt = timestamppb.New(*account.CreatedAt)
		}
		if account.UpdatedAt != nil {
			basic.UpdatedAt = timestamppb.New(*account.UpdatedAt)
		}
		result[account.ID] = basic
	}
	return &v1.BatchGetBasicAccount_Reply{Accounts: result}, nil
}

func (s *AccountService) BatchGetContact(ctx context.Context, req *v1.BatchGetContactAccount_Request) (*v1.BatchGetContactAccount_Reply, error) {
	req = util.OrDefault(req, &v1.BatchGetContactAccount_Request{})
	accounts, err := s.accountRepo.List(ctx, &repo.AccountGetReq{UserIds: req.UserIds})
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*v1.AccountContact, len(accounts))
	for _, account := range accounts {
		result[account.ID] = &v1.AccountContact{
			UserId: account.ID,
			Email:  account.Email,
			Phone:  account.Phone,
		}
	}
	return &v1.BatchGetContactAccount_Reply{Contacts: result}, nil
}

func (s *AccountService) ExistsEmail(ctx context.Context, req *v1.ExistsEmail_Request) (*v1.ExistsEmail_Reply, error) {
	exists, err := s.accountRepo.ExistsByAccount(ctx, req.Email)
	return &v1.ExistsEmail_Reply{Exists: exists}, err
}

func (s *AccountService) ExistsPhone(ctx context.Context, req *v1.ExistsPhone_Request) (*v1.ExistsPhone_Reply, error) {
	exists, err := s.accountRepo.ExistsByAccount(ctx, req.Phone)
	return &v1.ExistsPhone_Reply{Exists: exists}, err
}

func (s *AccountService) ExistsName(ctx context.Context, req *v1.ExistsName_Request) (*v1.ExistsName_Reply, error) {
	exists, err := s.accountRepo.ExistsByAccount(ctx, req.Name)
	return &v1.ExistsName_Reply{Exists: exists}, err
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileAccount_Request) (*v1.UpdateProfileAccount_Reply, error) {
	req = util.OrDefault(req, &v1.UpdateProfileAccount_Request{})
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	patch := &repo.AccountProfilePatch{
		UserID:       current.ID,
		AvatarURL:    repo.NewStringPatch(req.AvatarUrl),
		Nickname:     repo.NewStringPatch(req.Nickname),
		URL:          repo.NewStringPatch(req.Url),
		Introduction: repo.NewStringPatch(req.Introduction),
	}
	if req.Mbti != nil {
		if *req.Mbti == v1.MBTI_MBTI_UNSPECIFIED {
			patch.Mbti = repo.MBTIPatch{Set: true, Clear: true}
		} else {
			mbti, ok := enum.MBTIMap.ToEnum(*req.Mbti)
			if !ok {
				return nil, cerrors.ErrorBadRequest("mbti is invalid")
			}
			patch.Mbti = repo.MBTIPatch{Set: true, Value: mbti}
		}
	}
	if err := s.accountValidationUsecase.ValidateProfileUpdate(patch); err != nil {
		return nil, err
	}
	account, err := s.accountRepo.UpdateProfile(ctx, patch)
	if err != nil {
		return nil, err
	}
	basic := &v1.AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarUrl:     account.AvatarURL,
		Introduction:  account.Introduction,
		FollowCount:   account.FollowCount,
		FollowerCount: account.FollowerCount,
	}
	if account.Mbti != nil {
		basic.Mbti = enum.MBTIMap.MustToProto(*account.Mbti)
	}
	if account.Status != nil {
		basic.Status = enum.AccountStatusMap.MustToProto(*account.Status)
	}
	if account.CreatedAt != nil {
		basic.CreatedAt = timestamppb.New(*account.CreatedAt)
	}
	if account.UpdatedAt != nil {
		basic.UpdatedAt = timestamppb.New(*account.UpdatedAt)
	}
	return &v1.UpdateProfileAccount_Reply{Account: basic}, nil
}
