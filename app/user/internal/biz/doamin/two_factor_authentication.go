package doamin

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

type TwoFactorAuthenticationDomain struct {
	conf        *conf.Bootstrap
	redis       *client.RedisClient
	txRunner    base.TxRunner
	userRepo    repo.UserRepo
	userTfaRepo repo.UserTfaRepo
}

func NewTwoFactorAuthenticationDomain(
	conf *conf.Bootstrap,
	redis *client.RedisClient,
	txRunner base.TxRunner,
	userRepo repo.UserRepo,
	userTfaRepo repo.UserTfaRepo,
) (*TwoFactorAuthenticationDomain, error) {
	return &TwoFactorAuthenticationDomain{
		conf:        conf,
		redis:       redis,
		txRunner:    txRunner,
		userRepo:    userRepo,
		userTfaRepo: userTfaRepo,
	}, nil
}

func (d *TwoFactorAuthenticationDomain) Validate(ctx context.Context, secret string, code string) bool {
	return totp.Validate(code, secret)
}

func (d *TwoFactorAuthenticationDomain) Enable(ctx context.Context, name string) ([]byte, error) {
	buf := &bytes.Buffer{}
	generate, err := totp.Generate(totp.GenerateOpts{
		Issuer:      d.conf.Server.App,
		AccountName: name,
	})
	if err != nil {
		return nil, err
	}

	err = d.redis.Client.SetEx(ctx, constant.GetKeyTwoFactorAuthentication(name), generate.Secret(), 5*time.Minute).Err()
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

func (d *TwoFactorAuthenticationDomain) Disable(ctx context.Context, name string, secret string, code string) error {
	if !totp.Validate(code, secret) {
		return cerrors.ErrorBadRequest("2FA code invalid")
	}
	err := d.txRunner(ctx, func(ctx context.Context) error {
		u, err := d.userRepo.GetByAccount(ctx, name)
		if err != nil {
			return err
		}
		_, err = d.userTfaRepo.Disable(ctx, u.ID)
		return err
	})
	return err
}

func (d *TwoFactorAuthenticationDomain) Confirm(ctx context.Context, name string, code string) error {
	secret, err := d.redis.Client.Get(ctx, constant.GetKeyTwoFactorAuthentication(name)).Result()
	if err != nil {
		return err
	}
	if !totp.Validate(code, secret) {
		return cerrors.ErrorBadRequest("2FA code invalid")
	}
	err = d.txRunner(ctx, func(ctx context.Context) error {
		u, err := d.userRepo.GetByAccount(ctx, name)
		if err != nil {
			return err
		}
		_, err = d.userTfaRepo.Enable(ctx, u.ID, secret)
		return err
	})
	return err
}
