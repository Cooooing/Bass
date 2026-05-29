package usecase

import (
	"bytes"
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	"context"
	"image/png"
	"time"
	base "user/internal/biz/base"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"

	"github.com/pquerna/otp/totp"
)

type TfaUsecase struct {
	conf           *conf.Bootstrap
	tfaSecretCache repo.TfaSecretCache
	tx             base.Tx
	tfaRepo        repo.TfaRepo
	outboxRepo     repo.OutboxEventRepo
}

func NewTfaUsecase(
	conf *conf.Bootstrap,
	tfaSecretCache repo.TfaSecretCache,
	tx base.Tx,
	tfaRepo repo.TfaRepo,
	outboxRepo repo.OutboxEventRepo,
) (*TfaUsecase, error) {
	return &TfaUsecase{
		conf:           conf,
		tfaSecretCache: tfaSecretCache,
		tx:             tx,
		tfaRepo:        tfaRepo,
		outboxRepo:     outboxRepo,
	}, nil
}

func (d *TfaUsecase) Validate(ctx context.Context, secret string, code string) bool {
	return totp.Validate(code, secret)
}

func (d *TfaUsecase) GetByUserID(ctx context.Context, userID int64) (*model.TFA, error) {
	return d.tfaRepo.FindByUserID(ctx, userID)
}

func (d *TfaUsecase) ValidateByUserID(ctx context.Context, userID int64, code string) (bool, error) {
	tfa, err := d.tfaRepo.FindByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	if tfa == nil || !tfa.Enable || tfa.Secret == "" {
		return false, nil
	}
	return d.Validate(ctx, tfa.Secret, code), nil
}

func (d *TfaUsecase) BeginEnable(ctx context.Context, userID int64, accountName string) ([]byte, error) {
	tfa, err := d.tfaRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if tfa != nil && tfa.Enable {
		return nil, cerrors.ErrorBadRequest("2FA already enabled")
	}
	return d.Enable(ctx, userID, accountName)
}

func (d *TfaUsecase) Enable(ctx context.Context, userID int64, accountName string) ([]byte, error) {
	buf := &bytes.Buffer{}
	generate, err := totp.Generate(totp.GenerateOpts{
		Issuer:      d.conf.Server.App,
		AccountName: accountName,
	})
	if err != nil {
		return nil, err
	}

	err = d.tfaSecretCache.Save(ctx, userID, generate.Secret(), 5*time.Minute)
	if err != nil {
		return nil, err
	}

	image, err := generate.Image(256, 256)
	if err != nil {
		return nil, err
	}
	err = png.Encode(buf, image)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (d *TfaUsecase) Disable(ctx context.Context, userID int64, secret string, code string) error {
	if !totp.Validate(code, secret) {
		return cerrors.ErrorBadRequest("2FA code invalid")
	}
	err := d.tx(ctx, func(ctx context.Context) error {
		_, err := d.tfaRepo.DisableByUserID(ctx, userID)
		if err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_TFA_DISABLE,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_TFA_DISABLE,
				Payload: &commonenums.Event_UserTfaDisable{
					UserTfaDisable: &commonenums.UserTfaDisablePayload{
						UserId: userID,
					},
				},
			},
		})
	})
	return err
}

func (d *TfaUsecase) DisableByUserID(ctx context.Context, userID int64, code string) error {
	tfa, err := d.tfaRepo.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if tfa == nil || !tfa.Enable {
		return cerrors.ErrorBadRequest("2FA already disabled")
	}
	return d.Disable(ctx, userID, tfa.Secret, code)
}

func (d *TfaUsecase) Confirm(ctx context.Context, userID int64, code string) error {
	secret, err := d.tfaSecretCache.Get(ctx, userID)
	if err != nil {
		return err
	}
	if !totp.Validate(code, secret) {
		return cerrors.ErrorBadRequest("2FA code invalid")
	}
	err = d.tx(ctx, func(ctx context.Context) error {
		_, err = d.tfaRepo.UpsertEnabledByUserID(ctx, userID, secret)
		if err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_USER_TFA_ENABLE,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_USER_TFA_ENABLE,
				Payload: &commonenums.Event_UserTfaEnable{
					UserTfaEnable: &commonenums.UserTfaEnablePayload{
						UserId: userID,
					},
				},
			},
		})
	})
	return err
}
