package usecase

import (
	"bytes"
	"common/pkg/util"
	"context"
	"image/png"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"

	"github.com/MuhammadSaim/goavatar"
	"log/slog"
)

type AccountUsecase struct {
	conf            *conf.Bootstrap
	log             *util.LogHelper
	tx              base.Tx
	accountRepo     repo.AccountRepo
	preferencesRepo repo.PreferencesRepo
}

func NewAccountUsecase(
	conf *conf.Bootstrap,
	logger *slog.Logger,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	preferencesRepo repo.PreferencesRepo,
) (*AccountUsecase, error) {
	return &AccountUsecase{
		conf:            conf,
		log:             util.NewLogHelper(logger),
		tx:              tx,
		accountRepo:     accountRepo,
		preferencesRepo: preferencesRepo,
	}, nil
}

func (s *AccountUsecase) GetByUserID(ctx context.Context, userID int64) (*model.Account, error) {
	return s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &userID})
}

func (s *AccountUsecase) CheckAvailability(ctx context.Context, req *model.AccountAvailability) (*model.AccountAvailability, error) {
	if req == nil {
		return &model.AccountAvailability{}, nil
	}
	result := &model.AccountAvailability{
		Name:  req.Name,
		Email: req.Email,
		Phone: req.Phone,
	}
	if req.Name != nil {
		exist, err := s.accountRepo.ExistsByAccount(ctx, *req.Name)
		if err != nil {
			return nil, err
		}
		result.NameAvailable = !exist
	}
	if req.Email != nil {
		exist, err := s.accountRepo.ExistsByAccount(ctx, *req.Email)
		if err != nil {
			return nil, err
		}
		result.EmailAvailable = !exist
	}
	if req.Phone != nil {
		exist, err := s.accountRepo.ExistsByAccount(ctx, *req.Phone)
		if err != nil {
			return nil, err
		}
		result.PhoneAvailable = !exist
	}
	return result, nil
}

func (s *AccountUsecase) ListByUserIDs(ctx context.Context, userIDs []int64) ([]*model.Account, error) {
	return s.accountRepo.List(ctx, &repo.AccountGetReq{UserIds: userIDs})
}

func (s *AccountUsecase) MapByUserIDs(ctx context.Context, userIDs []int64) (map[int64]*model.Account, error) {
	return s.accountRepo.Map(ctx, &repo.AccountGetReq{UserIds: userIDs})
}

func (s *AccountUsecase) UpdateProfile(ctx context.Context, req *model.AccountProfileUpdate) (*model.Account, error) {
	return s.accountRepo.UpdateProfile(ctx, req)
}

// UpdateSetting 在同一事务中更新账号资料和偏好设置。
func (s *AccountUsecase) UpdateSetting(ctx context.Context, userID int64, account *model.Account, prefs *model.Preferences) (*model.Account, error) {
	var fullAccount *model.Account
	err := s.tx(ctx, func(ctx context.Context) error {
		account.ID = userID
		if _, err := s.accountRepo.Update(ctx, account); err != nil {
			return err
		}

		prefs.UserID = userID
		if _, err := s.preferencesRepo.UpsertByUserID(ctx, prefs); err != nil {
			return err
		}

		var err error
		fullAccount, err = s.accountRepo.Get(ctx, &repo.AccountGetReq{UserID: &userID})
		return err
	})
	if err != nil {
		return nil, err
	}
	return fullAccount, nil
}

// Avatar 生成注册时使用的默认账号头像。
func (s *AccountUsecase) Avatar(ctx context.Context, name string) ([]byte, error) {
	buf := &bytes.Buffer{}
	avatar := goavatar.Make(name, goavatar.WithSize(512))
	err := png.Encode(buf, avatar)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
