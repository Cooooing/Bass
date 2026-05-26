package usecase

import (
	"context"
	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
)

type NotificationSettingUsecase struct {
	tx                      base.Tx
	notificationSettingRepo repo.NotificationSettingRepo
}

func NewNotificationSettingUsecase(
	tx base.Tx,
	notificationSettingRepo repo.NotificationSettingRepo,
) *NotificationSettingUsecase {
	return &NotificationSettingUsecase{
		tx:                      tx,
		notificationSettingRepo: notificationSettingRepo,
	}
}

func (u *NotificationSettingUsecase) List(ctx context.Context, req *repo.NotificationSettingGetReq) ([]*model.NotificationSetting, error) {
	return u.notificationSettingRepo.List(ctx, req)
}

func (u *NotificationSettingUsecase) Upsert(ctx context.Context, setting *model.NotificationSetting) (*model.NotificationSetting, error) {
	var save *model.NotificationSetting
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		save, err = u.notificationSettingRepo.Upsert(ctx, setting)
		return err
	})
	return save, err
}
