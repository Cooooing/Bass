package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v3/transport/http"
)

const (
	maxProfileURLLength          = 2048
	maxProfileIntroductionLength = 512
)

type AccountService struct {
	bbsuserv1.UnimplementedAccountServiceServer
	accountUsecase *usecase.AccountUsecase
}

func NewAccountService(accountUsecase *usecase.AccountUsecase) *AccountService {
	return &AccountService{accountUsecase: accountUsecase}
}

func (s *AccountService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAccountServiceHTTPServer(hs, s)
}

func (s *AccountService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.accountUsecase.GetCurrentAccount(ctx, &usecase.GetCurrentAccountReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentAccount_Response{Account: response.Account}, nil
}

func (s *AccountService) GetProfile(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Response, error) {
	response, err := s.accountUsecase.GetProfileAccount(ctx, &usecase.GetProfileAccountReq{UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetProfileAccount_Response{Profile: response.Profile}, nil
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
	}
	if req.AvatarUrl != nil {
		value := strings.TrimSpace(*req.AvatarUrl)
		if utf8.RuneCountInString(value) > maxProfileURLLength {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
		}
		req.AvatarUrl = new(value)
	}
	if req.Nickname != nil {
		value := strings.TrimSpace(*req.Nickname)
		if value != "" {
			length := utf8.RuneCountInString(value)
			if length < 2 || length > 32 {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
			}
			hasNonDigit := false
			for _, r := range value {
				if !unicode.IsDigit(r) {
					hasNonDigit = true
					break
				}
			}
			if !hasNonDigit {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
			}
		}
		req.Nickname = new(value)
	}
	if req.Url != nil {
		value := strings.TrimSpace(*req.Url)
		if utf8.RuneCountInString(value) > maxProfileURLLength {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
		}
		req.Url = new(value)
	}
	if req.Introduction != nil && utf8.RuneCountInString(*req.Introduction) > maxProfileIntroductionLength {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
	}
	if req.Mbti != nil {
		if _, ok := bbsuserv1.MBTI_name[int32(*req.Mbti)]; !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_INVALID)
		}
	}
	response, err := s.accountUsecase.UpdateProfileAccount(ctx, &usecase.UpdateProfileAccountReq{UserID: userID, AvatarURL: req.AvatarUrl, Nickname: req.Nickname, URL: req.Url, Introduction: req.Introduction, Mbti: req.Mbti})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateProfileAccount_Response{Profile: response.Profile}, nil
}

func (s *AccountService) Avatar(ctx context.Context, req *bbsuserv1.AvatarAccount_Request) (*common.ImageResponse, error) {
	response, err := s.accountUsecase.AvatarAccount(ctx, &usecase.AvatarAccountReq{Name: req.GetName()})
	if err != nil {
		return nil, err
	}
	return &common.ImageResponse{Data: response.Data, ContentType: response.ContentType}, nil
}
