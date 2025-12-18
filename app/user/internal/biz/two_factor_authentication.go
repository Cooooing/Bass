package biz

import (
	"bytes"
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"context"
	"image/png"
	"time"
	"user/internal/biz/repo"
	"user/internal/data/ent"
	"user/internal/data/ent/gen"

	"github.com/pquerna/otp/totp"
)

type TwoFactorAuthenticationDomain struct {
	*BaseDomain
	userRepo repo.UserRepo
}

func NewTwoFactorAuthenticationDomain(base *BaseDomain, userRepo repo.UserRepo) (*TwoFactorAuthenticationDomain, error) {
	return &TwoFactorAuthenticationDomain{
		BaseDomain: base,
		userRepo:   userRepo,
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
		return cv1.ErrorBadRequest("2FA code invalid")
	}
	err := ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		_, err := d.userRepo.DisableTwoFactorAuthentication(ctx, tx, name)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (d *TwoFactorAuthenticationDomain) Confirm(ctx context.Context, name string, code string) error {
	secret, err := d.redis.Client.Get(ctx, constant.GetKeyTwoFactorAuthentication(name)).Result()
	if err != nil {
		return err
	}
	if !totp.Validate(code, secret) {
		return cv1.ErrorBadRequest("2FA code invalid")
	}
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		_, err = d.userRepo.EnableTwoFactorAuthentication(ctx, tx, name, secret)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
