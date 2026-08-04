package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	v1enum "common/proto/gen/user/v1/enum"
	"context"
	"net/mail"
	"regexp"
	"strings"
	"unicode/utf8"
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
	phoneRe        *regexp.Regexp
	codeRe         *regexp.Regexp
}

func NewAccountService(
	accountUsecase *usecase.AccountUsecase,
) *AccountService {
	return &AccountService{
		accountUsecase: accountUsecase,
		phoneRe:        regexp.MustCompile("^1[3-9]\\d{9}$"),
		codeRe:         regexp.MustCompile("^[A-Za-z0-9]{6}$"),
	}
}
func (s *AccountService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterAccountServiceServer(gs, s)
}

func (s *AccountService) RegisterHttp(hs *http.Server) {
}

func (s *AccountService) Get(ctx context.Context, req *v1.GetAccount_Req) (*v1.GetAccount_Resp, error) {
	req = util.OrDefault(req, &v1.GetAccount_Req{})
	res, err := s.accountUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	account := res
	basic := &v1.AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarAssetId: account.AvatarAssetID,
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
	replyAccount := &v1.AccountInfo{
		Basic: basic,
		Contact: &v1.AccountContact{
			UserId: account.ID,
			Email:  account.Email,
			Phone:  account.Phone,
		},
	}
	return &v1.GetAccount_Resp{
		Account: replyAccount,
	}, nil
}

func (s *AccountService) List(ctx context.Context, req *v1.ListAccounts_Req) (*v1.ListAccounts_Resp, error) {
	req = util.OrDefault(req, &v1.ListAccounts_Req{})
	query := util.OrDefault(req.Query, &v1.ListAccounts_Req_AccountQuery{})
	if len(query.UserIds) == 0 {
		return &v1.ListAccounts_Resp{
			Rows: []*v1.AccountInfo{},
		}, nil
	}
	res, err := s.accountUsecase.ListByUserIDs(ctx, query.UserIds)
	if err != nil {
		return nil, err
	}
	accounts := res
	rows := make([]*v1.AccountInfo, 0, len(accounts))
	for _, account := range accounts {
		basic := &v1.AccountBasic{
			Id:            account.ID,
			Name:          account.Name,
			Nickname:      account.Nickname,
			Url:           account.URL,
			AvatarAssetId: account.AvatarAssetID,
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
		replyAccount := &v1.AccountInfo{
			Basic: basic,
			Contact: &v1.AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		}
		rows = append(rows, replyAccount)
	}
	return &v1.ListAccounts_Resp{
		Rows: rows,
	}, nil
}

func (s *AccountService) Map(ctx context.Context, req *v1.MapAccounts_Req) (*v1.MapAccounts_Resp, error) {
	req = util.OrDefault(req, &v1.MapAccounts_Req{})
	query := util.OrDefault(req.Query, &v1.MapAccounts_Req_AccountQuery{})
	if len(query.UserIds) == 0 {
		return &v1.MapAccounts_Resp{
			Accounts: map[int64]*v1.AccountInfo{},
		}, nil
	}
	res, err := s.accountUsecase.MapByUserIDs(ctx, query.UserIds)
	if err != nil {
		return nil, err
	}
	accounts := res
	rows := make(map[int64]*v1.AccountInfo, len(accounts))
	for userID, account := range accounts {
		basic := &v1.AccountBasic{
			Id:            account.ID,
			Name:          account.Name,
			Nickname:      account.Nickname,
			Url:           account.URL,
			AvatarAssetId: account.AvatarAssetID,
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
		replyAccount := &v1.AccountInfo{
			Basic: basic,
			Contact: &v1.AccountContact{
				UserId: account.ID,
				Email:  account.Email,
				Phone:  account.Phone,
			},
		}
		rows[userID] = replyAccount
	}
	return &v1.MapAccounts_Resp{
		Accounts: rows,
	}, nil
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
		UserID:        req.GetUserId(),
		AvatarAssetID: req.AvatarAssetId,
		Nickname:      req.Nickname,
		URL:           req.Url,
		Introduction:  req.Introduction,
		Mbti:          mbti,
		ClearMBTI:     clearMBTI,
	})
	if err != nil {
		return nil, err
	}
	account := res
	basic := &v1.AccountBasic{
		Id:            account.ID,
		Name:          account.Name,
		Nickname:      account.Nickname,
		Url:           account.URL,
		AvatarAssetId: account.AvatarAssetID,
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
	return &v1.UpdateProfileAccount_Resp{
		Account: basic,
	}, nil
}

func (s *AccountService) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordAccount_Req) (*v1.UpdatePasswordAccount_Resp, error) {
	if req == nil || req.GetUserId() == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	oldPassword := req.GetOldPassword()
	newPassword := req.GetNewPassword()
	if len(oldPassword) < 6 || len(oldPassword) > 64 || len(newPassword) < 6 || len(newPassword) > 64 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	for _, r := range newPassword {
		if r < '!' || r > '~' {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		switch {
		case r >= 'A' && r <= 'Z':
			hasUpper = true
		case r >= 'a' && r <= 'z':
			hasLower = true
		case r >= '0' && r <= '9':
			hasDigit = true
		default:
			hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit || !hasSpecial {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.accountUsecase.UpdatePassword(ctx, &usecase.UpdateAccountPasswordReq{
		UserID:      req.GetUserId(),
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}); err != nil {
		return nil, err
	}
	return &v1.UpdatePasswordAccount_Resp{}, nil
}

func (s *AccountService) UpdateEmail(ctx context.Context, req *v1.UpdateEmailAccount_Req) (*v1.UpdateEmailAccount_Resp, error) {
	if req == nil || req.GetUserId() == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	email := strings.ToLower(strings.TrimSpace(req.GetEmail()))
	parsed, err := mail.ParseAddress(email)
	if err != nil || parsed.Address != email || !strings.Contains(email, "@") || utf8.RuneCountInString(email) > 254 || !s.codeRe.MatchString(strings.TrimSpace(req.GetCode())) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.accountUsecase.UpdateEmail(ctx, &usecase.UpdateAccountEmailReq{
		UserID: req.GetUserId(),
		Email:  email,
		Code:   strings.TrimSpace(req.GetCode()),
	}); err != nil {
		return nil, err
	}
	return &v1.UpdateEmailAccount_Resp{}, nil
}

func (s *AccountService) UpdatePhone(ctx context.Context, req *v1.UpdatePhoneAccount_Req) (*v1.UpdatePhoneAccount_Resp, error) {
	if req == nil || req.GetUserId() == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	phone := strings.TrimSpace(req.GetPhone())
	if !s.phoneRe.MatchString(phone) || !s.codeRe.MatchString(strings.TrimSpace(req.GetCode())) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.accountUsecase.UpdatePhone(ctx, &usecase.UpdateAccountPhoneReq{
		UserID: req.GetUserId(),
		Phone:  phone,
		Code:   strings.TrimSpace(req.GetCode()),
	}); err != nil {
		return nil, err
	}
	return &v1.UpdatePhoneAccount_Resp{}, nil
}
func (s *AccountService) Avatar(ctx context.Context, req *v1.AvatarAccount_Req) (*common.ImageResp, error) {
	res, err := s.accountUsecase.Avatar(ctx, req.GetName())
	if err != nil {
		return nil, err
	}
	return &common.ImageResp{
		Data:        res,
		ContentType: "image/png",
	}, nil
}
