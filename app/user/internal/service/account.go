package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountService struct {
	v1.UnimplementedAccountServiceServer
	accountUsecase *usecase.AccountUsecase
}

func NewAccountService(accountUsecase *usecase.AccountUsecase) *AccountService {
	return &AccountService{accountUsecase: accountUsecase}
}

func (s *AccountService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAccountServiceServer(gs, s)
}

func (s *AccountService) RegisterHttp(hs *http.Server) {}

func (s *AccountService) Get(ctx context.Context, req *v1.GetAccount_Request) (*v1.GetAccount_Reply, error) {
	req = util.OrDefault(req, &v1.GetAccount_Request{})
	account, err := s.accountUsecase.GetByUserID(ctx, req.GetUserId())
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
	replyAccount := &v1.Account{
		Basic: basic,
		Contact: &v1.AccountContact{
			UserId: account.ID,
			Email:  account.Email,
			Phone:  account.Phone,
		},
	}
	return &v1.GetAccount_Reply{Account: replyAccount}, nil
}

func (s *AccountService) List(ctx context.Context, req *v1.ListAccounts_Request) (*v1.ListAccounts_Reply, error) {
	req = util.OrDefault(req, &v1.ListAccounts_Request{})
	query := util.OrDefault(req.Query, &v1.AccountQuery{})
	if len(query.UserIds) == 0 {
		return &v1.ListAccounts_Reply{Rows: []*v1.Account{}}, nil
	}
	accounts, err := s.accountUsecase.ListByUserIDs(ctx, query.UserIds)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Account, 0, len(accounts))
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
		replyAccount := &v1.Account{
			Basic: basic,
			Contact: &v1.AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		}
		rows = append(rows, replyAccount)
	}
	return &v1.ListAccounts_Reply{Rows: rows}, nil
}

func (s *AccountService) Map(ctx context.Context, req *v1.MapAccounts_Request) (*v1.MapAccounts_Reply, error) {
	req = util.OrDefault(req, &v1.MapAccounts_Request{})
	query := util.OrDefault(req.Query, &v1.AccountQuery{})
	if len(query.UserIds) == 0 {
		return &v1.MapAccounts_Reply{Accounts: map[int64]*v1.Account{}}, nil
	}
	accounts, err := s.accountUsecase.MapByUserIDs(ctx, query.UserIds)
	if err != nil {
		return nil, err
	}
	rows := make(map[int64]*v1.Account, len(accounts))
	for userID, account := range accounts {
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
		replyAccount := &v1.Account{
			Basic: basic,
			Contact: &v1.AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		}
		rows[userID] = replyAccount
	}
	return &v1.MapAccounts_Reply{Accounts: rows}, nil
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileAccount_Request) (*v1.UpdateProfileAccount_Reply, error) {
	req = util.OrDefault(req, &v1.UpdateProfileAccount_Request{})
	var mbti *enum.MBTI
	clearMBTI := false
	if req.Mbti != nil {
		if *req.Mbti == v1.MBTI_MBTI_UNSPECIFIED {
			clearMBTI = true
		} else {
			value, ok := enum.MBTIMap.ToEnum(*req.Mbti)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			mbti = new(value)
		}
	}
	account, err := s.accountUsecase.UpdateProfile(ctx, &model.AccountProfileUpdate{
		UserID:       req.GetUserId(),
		AvatarURL:    req.AvatarUrl,
		Nickname:     req.Nickname,
		URL:          req.Url,
		Introduction: req.Introduction,
		Mbti:         mbti,
		ClearMBTI:    clearMBTI,
	})
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

func (s *AccountService) Avatar(ctx context.Context, req *v1.AvatarAccount_Request) (*common.ImageReply, error) {
	data, err := s.accountUsecase.Avatar(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &common.ImageReply{Data: data, ContentType: "image/png"}, nil
}
