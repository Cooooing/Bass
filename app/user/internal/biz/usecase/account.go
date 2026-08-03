package usecase

import (
	"bytes"
	"common/pkg/apperror"
	"common/pkg/util/str"
	commonerrors "common/proto/gen/common/errors"
	"context"
	"image/png"
	"log/slog"
	"strings"
	"user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"
	"user/internal/enum"

	"github.com/MuhammadSaim/goavatar"
)

type AccountUsecase struct {
	conf            *config.Bootstrap
	tx              base.Tx
	accountRepo     repo.AccountRepo
	preferencesRepo repo.PreferencesRepo
	authCacheRepo   repo.AuthCacheRepo
	emailOtpUsecase *EmailOtpUsecase
	smsOtpUsecase   *SmsOtpUsecase
}

func NewAccountUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	preferencesRepo repo.PreferencesRepo,
	authCacheRepo repo.AuthCacheRepo,
	emailOtpUsecase *EmailOtpUsecase,
	smsOtpUsecase *SmsOtpUsecase,
) (*AccountUsecase, error) {
	return &AccountUsecase{
		conf:            conf,
		tx:              tx,
		accountRepo:     accountRepo,
		preferencesRepo: preferencesRepo,
		authCacheRepo:   authCacheRepo,
		emailOtpUsecase: emailOtpUsecase,
		smsOtpUsecase:   smsOtpUsecase,
	}, nil
}

func (s *AccountUsecase) GetByUserID(ctx context.Context, userID int64) (*model.Account, error) {
	return s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: new(userID),
	})
}

func (s *AccountUsecase) CheckAvailability(ctx context.Context, availability *model.AccountAvailability) (*model.AccountAvailability, error) {
	if availability == nil {
		return &model.AccountAvailability{}, nil
	}
	result := &model.AccountAvailability{
		Name:  availability.Name,
		Email: availability.Email,
		Phone: availability.Phone,
	}
	var exists bool
	var err error
	if availability.Name != nil {
		exists, err = s.accountRepo.ExistsByAccount(ctx, *availability.Name)
		if err != nil {
			return nil, err
		}
		result.NameAvailable = !exists
	}
	if availability.Email != nil {
		exists, err = s.accountRepo.ExistsByAccount(ctx, *availability.Email)
		if err != nil {
			return nil, err
		}
		result.EmailAvailable = !exists
	}
	if availability.Phone != nil {
		exists, err = s.accountRepo.ExistsByAccount(ctx, *availability.Phone)
		if err != nil {
			return nil, err
		}
		result.PhoneAvailable = !exists
	}
	return result, nil
}

func (s *AccountUsecase) ListByUserIDs(ctx context.Context, userIDs []int64) ([]*model.Account, error) {
	return s.accountRepo.List(ctx, &repo.AccountGetReq{
		UserIds: userIDs,
	})
}

func (s *AccountUsecase) MapByUserIDs(ctx context.Context, userIDs []int64) (map[int64]*model.Account, error) {
	return s.accountRepo.Map(ctx, &repo.AccountGetReq{
		UserIds: userIDs,
	})
}

func (s *AccountUsecase) UpdateProfile(ctx context.Context, profile *model.AccountProfileUpdate) (*model.Account, error) {
	return s.accountRepo.UpdateProfile(ctx, profile)
}

type UpdateAccountPasswordReq struct {
	UserID      int64
	OldPassword string
	NewPassword string
}

func (s *AccountUsecase) UpdatePassword(ctx context.Context, req *UpdateAccountPasswordReq) error {
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: new(req.UserID),
	})
	if err != nil {
		return err
	}
	if account.Status == nil || *account.Status != enum.AccountStatusNormal {
		switch {
		case account.Status != nil && *account.Status == enum.AccountStatusBanned:
			return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_BANNED)
		case account.Status != nil && *account.Status == enum.AccountStatusCancelled:
			return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_CANCELLED)
		default:
			return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
	}
	if !str.VerifyPassword(account.Password, req.OldPassword) {
		return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	passwordHash, err := str.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}
	if err := s.accountRepo.UpdatePassword(ctx, req.UserID, passwordHash); err != nil {
		return err
	}
	return s.authCacheRepo.DeleteUserSessions(ctx, req.UserID)
}

type UpdateAccountEmailReq struct {
	UserID int64
	Email  string
	Code   string
}

func (s *AccountUsecase) UpdateEmail(ctx context.Context, req *UpdateAccountEmailReq) error {
	email := strings.ToLower(strings.TrimSpace(req.Email))
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: new(req.UserID),
	})
	if err != nil {
		return err
	}
	if account.Status == nil || *account.Status != enum.AccountStatusNormal {
		return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	if account.Email == nil || *account.Email != email {
		exists, err := s.accountRepo.ExistsByAccount(ctx, email)
		if err != nil {
			return err
		}
		if exists {
			return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_ALREADY_EXISTS)
		}
	}
	if err := s.emailOtpUsecase.VerifyEmailOtp(ctx, &VerifyEmailOtpReq{
		UserID: new(req.UserID),
		Email:  email,
		Code:   req.Code,
	}); err != nil {
		return err
	}
	return s.accountRepo.UpdateEmail(ctx, req.UserID, email)
}

type UpdateAccountPhoneReq struct {
	UserID int64
	Phone  string
	Code   string
}

func (s *AccountUsecase) UpdatePhone(ctx context.Context, req *UpdateAccountPhoneReq) error {
	phone := strings.TrimSpace(req.Phone)
	account, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: new(req.UserID),
	})
	if err != nil {
		return err
	}
	if account.Status == nil || *account.Status != enum.AccountStatusNormal {
		return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_INVALID_CREDENTIALS)
	}
	if account.Phone == nil || *account.Phone != phone {
		exists, err := s.accountRepo.ExistsByAccount(ctx, phone)
		if err != nil {
			return err
		}
		if exists {
			return apperror.New(commonerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_ACCOUNT_ALREADY_EXISTS)
		}
	}
	if err := s.smsOtpUsecase.VerifyPhoneOtp(ctx, &VerifyPhoneOtpReq{
		UserID: new(req.UserID),
		Phone:  phone,
		Code:   req.Code,
	}); err != nil {
		return err
	}
	return s.accountRepo.UpdatePhone(ctx, req.UserID, phone)
}

type UpdateAccountSettingReq struct {
	UserID      int64
	Account     *model.Account
	Preferences *model.Preferences
}

func (s *AccountUsecase) UpdateSetting(ctx context.Context, req *UpdateAccountSettingReq) (*model.Account, error) {
	var fullAccount *model.Account
	err := s.tx(ctx, func(ctx context.Context) error {
		req.Account.ID = req.UserID
		if _, err := s.accountRepo.Update(ctx, req.Account); err != nil {
			return err
		}

		req.Preferences.UserID = req.UserID
		if _, err := s.preferencesRepo.UpsertByUserID(ctx, req.Preferences); err != nil {
			return err
		}

		var err error
		fullAccount, err = s.accountRepo.Get(ctx, &repo.AccountGetReq{
			UserID: new(req.UserID),
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return fullAccount, nil
}

func (s *AccountUsecase) Avatar(ctx context.Context, name string) ([]byte, error) {
	buf := &bytes.Buffer{}
	avatar := goavatar.Make(name, goavatar.WithSize(512))
	err := png.Encode(buf, avatar)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
