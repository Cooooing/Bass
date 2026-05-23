package usecase

import (
	"bytes"
	cerrors "common/api/gen/common/errors"
	"common/pkg/client"
	"common/pkg/constant"
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
	redis       *client.RedisClient
	tx          base.Tx
	accountRepo repo.AccountRepo
	tfaRepo     repo.TfaRepo
}

func NewTfaUsecase(
	conf *conf.Bootstrap,
	redis *client.RedisClient,
	tx base.Tx,
	accountRepo repo.AccountRepo,
	tfaRepo repo.TfaRepo,
) (*TfaUsecase, error) {
	return &TfaUsecase{
		conf:        conf,
		redis:       redis,
		tx:          tx,
		accountRepo: accountRepo,
		tfaRepo:     tfaRepo,
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

	err = d.redis.Client.SetEx(ctx, constant.GetKeyTwoFactorAuth(name), generate.Secret(), 5*time.Minute).Err()
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
		return err
	})
	return err
}

func (d *TfaUsecase) Confirm(ctx context.Context, name string, code string) error {
	secret, err := d.redis.Client.Get(ctx, constant.GetKeyTwoFactorAuth(name)).Result()
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
		return err
	})
	return err
}
