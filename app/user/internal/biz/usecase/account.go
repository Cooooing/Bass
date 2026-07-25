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

func (s *AccountUsecase) GetByUserID(
	ctx context.Context,
	userID int64,
) (*model.Account, error) {
	return s.accountRepo.Get(ctx, &repo.AccountGetReq{
		UserID: &userID,
	})
}

func (s *AccountUsecase) CheckAvailability(
	ctx context.Context,
	availability *model.AccountAvailability,
) (*model.AccountAvailability, error) {
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

func (s *AccountUsecase) ListByUserIDs(
	ctx context.Context,
	userIDs []int64,
) ([]*model.Account, error) {
	return s.accountRepo.List(ctx, &repo.AccountGetReq{
		UserIds: userIDs,
	})
}

func (s *AccountUsecase) MapByUserIDs(
	ctx context.Context,
	userIDs []int64,
) (map[int64]*model.Account, error) {
	return s.accountRepo.Map(ctx, &repo.AccountGetReq{
		UserIds: userIDs,
	})
}

func (s *AccountUsecase) UpdateProfile(
	ctx context.Context,
	profile *model.AccountProfileUpdate,
) (*model.Account, error) {
	return s.accountRepo.UpdateProfile(ctx, profile)
}

type UpdateAccountSettingReq struct {
	UserID      int64
	Account     *model.Account
	Preferences *model.Preferences
}

func (s *AccountUsecase) UpdateSetting(
	ctx context.Context,
	req *UpdateAccountSettingReq,
) (*model.Account, error) {
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
			UserID: &req.UserID,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return fullAccount, nil
}

func (s *AccountUsecase) Avatar(
	ctx context.Context,
	name string,
) ([]byte, error) {
	buf := &bytes.Buffer{}
	avatar := goavatar.Make(name, goavatar.WithSize(512))
	err := png.Encode(buf, avatar)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
