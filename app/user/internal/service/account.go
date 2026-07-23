package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	v1enum "common/proto/gen/user/v1/enum"
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

func (s *AccountService) Get(ctx context.Context, req *v1.GetAccount_Req) (*v1.GetAccount_Resp, error) {
	req = util.OrDefault(req, &v1.GetAccount_Req{})
	res, err := s.accountUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	account := res
	basic := &v1.GetAccount_Resp_AccountBasic{
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
	replyAccount := &v1.GetAccount_Resp_Account{
		Basic: basic,
		Contact: &v1.GetAccount_Resp_AccountContact{
			UserId: account.ID,
			Email:  account.Email,
			Phone:  account.Phone,
		},
	}
	return &v1.GetAccount_Resp{Account: replyAccount}, nil
}

func (s *AccountService) List(ctx context.Context, req *v1.ListAccounts_Req) (*v1.ListAccounts_Resp, error) {
	req = util.OrDefault(req, &v1.ListAccounts_Req{})
	query := util.OrDefault(req.Query, &v1.ListAccounts_Req_AccountQuery{})
	if len(query.UserIds) == 0 {
		return &v1.ListAccounts_Resp{Rows: []*v1.ListAccounts_Resp_Account{}}, nil
	}
	res, err := s.accountUsecase.ListByUserIDs(ctx, query.UserIds)
	if err != nil {
		return nil, err
	}
	accounts := res
	rows := make([]*v1.ListAccounts_Resp_Account, 0, len(accounts))
	for _, account := range accounts {
		basic := &v1.ListAccounts_Resp_AccountBasic{
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
		replyAccount := &v1.ListAccounts_Resp_Account{
			Basic: basic,
			Contact: &v1.ListAccounts_Resp_AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		}
		rows = append(rows, replyAccount)
	}
	return &v1.ListAccounts_Resp{Rows: rows}, nil
}

func (s *AccountService) Map(ctx context.Context, req *v1.MapAccounts_Req) (*v1.MapAccounts_Resp, error) {
	req = util.OrDefault(req, &v1.MapAccounts_Req{})
	query := util.OrDefault(req.Query, &v1.MapAccounts_Req_AccountQuery{})
	if len(query.UserIds) == 0 {
		return &v1.MapAccounts_Resp{Accounts: map[int64]*v1.MapAccounts_Resp_Account{}}, nil
	}
	res, err := s.accountUsecase.MapByUserIDs(ctx, query.UserIds)
	if err != nil {
		return nil, err
	}
	accounts := res
	rows := make(map[int64]*v1.MapAccounts_Resp_Account, len(accounts))
	for userID, account := range accounts {
		basic := &v1.MapAccounts_Resp_AccountBasic{
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
		replyAccount := &v1.MapAccounts_Resp_Account{
			Basic: basic,
			Contact: &v1.MapAccounts_Resp_AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		}
		rows[userID] = replyAccount
	}
	return &v1.MapAccounts_Resp{Accounts: rows}, nil
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileAccount_Req) (*v1.UpdateProfileAccount_Resp, error) {
	req = util.OrDefault(req, &v1.UpdateProfileAccount_Req{})
	var mbti *enum.MBTI
	clearMBTI := false
	if req.Mbti != nil {
		if *req.Mbti == v1enum.MBTI_MBTI_UNSPECIFIED {
			clearMBTI = true
		} else {
			value, ok := enum.MBTIMap.ToEnum(*req.Mbti)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			mbti = new(value)
		}
	}
	res, err := s.accountUsecase.UpdateProfile(ctx, &model.AccountProfileUpdate{
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
	account := res
	basic := &v1.UpdateProfileAccount_Resp_AccountBasic{
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
	return &v1.UpdateProfileAccount_Resp{Account: basic}, nil
}

func (s *AccountService) Avatar(ctx context.Context, req *v1.AvatarAccount_Req) (*common.ImageResp, error) {
	res, err := s.accountUsecase.Avatar(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &common.ImageResp{Data: res, ContentType: "image/png"}, nil
}
