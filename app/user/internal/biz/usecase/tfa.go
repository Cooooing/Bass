package usecase

import (
	"bytes"
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	"common/pkg/client"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"context"
	"image/png"
	"time"
	base "user/internal/biz/base"
	"user/internal/biz/repo"
	"user/internal/conf"

	"github.com/pquerna/otp/totp"
)

type TfaUsecase struct {
	conf        *conf.Bootstrap
	redisClient *client.RedisClient
	tx          base.Tx
	accountRepo repo.AccountRepo
	tfaRepo     repo.TfaRepo
	outboxRepo  repo.OutboxEventRepo
}

func NewTfaUsecase(
	conf *conf.Bootstrap,
	redisClient *client.RedisClient,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	tfaRepo repo.TfaRepo,
	outboxRepo repo.OutboxEventRepo,
) (*TfaUsecase, error) {
	return &TfaUsecase{
		conf:        conf,
		redisClient: redisClient,
		tx:          tx,
		accountRepo: accountRepo,
		tfaRepo:     tfaRepo,
		outboxRepo:  outboxRepo,
	}, nil
}

func (d *TfaUsecase) Validate(ctx context.Context, secret string, code string) bool {
	return totp.Validate(code, secret)
}

func (d *TfaUsecase) Enable(ctx context.Context, name string) ([]byte, error) {
	buf := &bytes.Buffer{}
	generate, err := totp.Generate(totp.GenerateOpts{
		Issuer:      d.conf.Server.App,
		AccountName: name,
	})
	if err != nil {
		return nil, err
	}

	err = d.redisClient.Client.SetEx(ctx, constant.GetKeyTwoFactorAuth(name), generate.Secret(), 5*time.Minute).Err()
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

func (d *TfaUsecase) Disable(ctx context.Context, name string, secret string, code string) error {
	if !totp.Validate(code, secret) {
		return cerrors.ErrorBadRequest("2FA code invalid")
	}
	err := d.tx(ctx, func(ctx context.Context) error {
		u, err := d.accountRepo.GetByAccount(ctx, name)
		if err != nil {
			return err
		}
		_, err = d.tfaRepo.DisableByUserID(ctx, u.ID)
		if err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			EventType: commonenum.EventTypeUserTfaDisable,
			Subject:   commonenum.EventSubjectUserTfaDisable,
			Event: &commonenums.Event{
				Payload: &commonenums.Event_UserTfaDisable{
					UserTfaDisable: &commonenums.UserTfaDisablePayload{
						UserId: u.ID,
						Name:   u.Name,
					},
				},
			},
		})
	})
	return err
}

func (d *TfaUsecase) Confirm(ctx context.Context, name string, code string) error {
	secret, err := d.redisClient.Client.Get(ctx, constant.GetKeyTwoFactorAuth(name)).Result()
	if err != nil {
		return err
	}
	if !totp.Validate(code, secret) {
		return cerrors.ErrorBadRequest("2FA code invalid")
	}
	err = d.tx(ctx, func(ctx context.Context) error {
		u, err := d.accountRepo.GetByAccount(ctx, name)
		if err != nil {
			return err
		}
		_, err = d.tfaRepo.UpsertEnabledByUserID(ctx, u.ID, secret)
		if err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			EventType: commonenum.EventTypeUserTfaEnable,
			Subject:   commonenum.EventSubjectUserTfaEnable,
			Event: &commonenums.Event{
				Payload: &commonenums.Event_UserTfaEnable{
					UserTfaEnable: &commonenums.UserTfaEnablePayload{
						UserId: u.ID,
						Name:   u.Name,
					},
				},
			},
		})
	})
	return err
}
