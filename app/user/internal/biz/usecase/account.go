package usecase

import (
	"bytes"
	"context"
	"image/png"
	"user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/config"

	"log/slog"

	"github.com/MuhammadSaim/goavatar"
)

type AccountUsecase struct {
	conf            *config.Bootstrap
	tx              base.Tx
	accountRepo     repo.AccountRepo
	preferencesRepo repo.PreferencesRepo
}

func NewAccountUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	preferencesRepo repo.PreferencesRepo,
) (*AccountUsecase, error) {
	return &AccountUsecase{
		conf:            conf,
		tx:              tx,
		accountRepo:     accountRepo,
		preferencesRepo: preferencesRepo,
	}, nil
}

type GetAccountByUserIDReq struct {
	UserID int64
}

type GetAccountByUserIDResponse struct {
	Account *model.Account
}

func (s *AccountUsecase) GetByUserID(ctx context.Context, req *GetAccountByUserIDReq) (*GetAccountByUserIDResponse, error) {
	accountResp, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	return &GetAccountByUserIDResponse{Account: accountResp.Account}, nil
}

type CheckAccountAvailabilityReq struct {
	Availability *model.AccountAvailability
}

type CheckAccountAvailabilityResponse struct {
	Availability *model.AccountAvailability
}

func (s *AccountUsecase) CheckAvailability(ctx context.Context, req *CheckAccountAvailabilityReq) (*CheckAccountAvailabilityResponse, error) {
	if req == nil || req.Availability == nil {
		return &CheckAccountAvailabilityResponse{Availability: &model.AccountAvailability{}}, nil
	}
	availability := req.Availability
	result := &model.AccountAvailability{
		Name:  availability.Name,
		Email: availability.Email,
		Phone: availability.Phone,
	}
	var existResp *repo.AccountExistsByAccountResponse
	var err error
	if availability.Name != nil {
		existResp, err = s.accountRepo.ExistsByAccount(ctx, &repo.AccountExistsByAccountReq{Account: *availability.Name})
		if err != nil {
			return nil, err
		}
		result.NameAvailable = !existResp.Exists
	}
	if availability.Email != nil {
		existResp, err = s.accountRepo.ExistsByAccount(ctx, &repo.AccountExistsByAccountReq{Account: *availability.Email})
		if err != nil {
			return nil, err
		}
		result.EmailAvailable = !existResp.Exists
	}
	if availability.Phone != nil {
		existResp, err = s.accountRepo.ExistsByAccount(ctx, &repo.AccountExistsByAccountReq{Account: *availability.Phone})
		if err != nil {
			return nil, err
		}
		result.PhoneAvailable = !existResp.Exists
	}
	return &CheckAccountAvailabilityResponse{Availability: result}, nil
}

type ListAccountsByUserIDsReq struct {
	UserIDs []int64
}

type ListAccountsByUserIDsResponse struct {
	Accounts []*model.Account
}

func (s *AccountUsecase) ListByUserIDs(ctx context.Context, req *ListAccountsByUserIDsReq) (*ListAccountsByUserIDsResponse, error) {
	accountsResp, err := s.accountRepo.List(ctx, &repo.AccountGetReq{UserIds: req.UserIDs})
	if err != nil {
		return nil, err
	}
	return &ListAccountsByUserIDsResponse{Accounts: accountsResp.Rows}, nil
}

type MapAccountsByUserIDsReq struct {
	UserIDs []int64
}

type MapAccountsByUserIDsResponse struct {
	Accounts map[int64]*model.Account
}

func (s *AccountUsecase) MapByUserIDs(ctx context.Context, req *MapAccountsByUserIDsReq) (*MapAccountsByUserIDsResponse, error) {
	accountsResp, err := s.accountRepo.Map(ctx, &repo.AccountGetReq{UserIds: req.UserIDs})
	if err != nil {
		return nil, err
	}
	return &MapAccountsByUserIDsResponse{Accounts: accountsResp.Rows}, nil
}

type UpdateAccountProfileReq struct {
	Profile *model.AccountProfileUpdate
}

type UpdateAccountProfileResponse struct {
	Account *model.Account
}

func (s *AccountUsecase) UpdateProfile(ctx context.Context, req *UpdateAccountProfileReq) (*UpdateAccountProfileResponse, error) {
	accountResp, err := s.accountRepo.UpdateProfile(ctx, &repo.AccountUpdateProfileReq{Profile: req.Profile})
	if err != nil {
		return nil, err
	}
	return &UpdateAccountProfileResponse{Account: accountResp.Account}, nil
}

type UpdateAccountSettingReq struct {
	UserID      int64
	Account     *model.Account
	Preferences *model.Preferences
}

type UpdateAccountSettingResponse struct {
	Account *model.Account
}

// UpdateSetting 在同一事务中更新账号资料和偏好设置。
func (s *AccountUsecase) UpdateSetting(ctx context.Context, req *UpdateAccountSettingReq) (*UpdateAccountSettingResponse, error) {
	var fullAccount *model.Account
	err := s.tx(ctx, func(ctx context.Context) error {
		req.Account.ID = req.UserID
		if _, err := s.accountRepo.Update(ctx, &repo.AccountUpdateReq{Account: req.Account}); err != nil {
			return err
		}

		req.Preferences.UserID = req.UserID
		if _, err := s.preferencesRepo.UpsertByUserID(ctx, &repo.PreferencesUpsertByUserIDReq{Preferences: req.Preferences}); err != nil {
			return err
		}

		var err error
		accountResp, err := s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &req.UserID})
		if err != nil {
			return err
		}
		fullAccount = accountResp.Account
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &UpdateAccountSettingResponse{Account: fullAccount}, nil
}

type AvatarAccountReq struct {
	Name string
}

type AvatarAccountResponse struct {
	Data []byte
}

// Avatar 生成注册时使用的默认账号头像。
func (s *AccountUsecase) Avatar(ctx context.Context, req *AvatarAccountReq) (*AvatarAccountResponse, error) {
	buf := &bytes.Buffer{}
	avatar := goavatar.Make(req.Name, goavatar.WithSize(512))
	err := png.Encode(buf, avatar)
	if err != nil {
		return nil, err
	}
	return &AvatarAccountResponse{Data: buf.Bytes()}, nil
}
