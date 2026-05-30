package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/totp"

	utilent "common/pkg/util/ent"
)

var _ repo.TotpRepo = (*TotpRepo)(nil)

type TotpRepo struct {
	db *gen.Client
}

func NewTotpRepo(db *gen.Client) repo.TotpRepo {
	return &TotpRepo{
		db: db,
	}
}

func (r *TotpRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *TotpRepo) FindByUserID(ctx context.Context, userID int64) (*model.Totp, error) {
	tx := r.getClient(ctx)
	row, err := tx.Totp.Query().Where(totp.UserID(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Totp{
		ID:         row.ID,
		UserID:     row.UserID,
		Enable:     row.Enable,
		EnableTime: row.EnableTime,
		Secret:     row.Secret,
	}, nil
}

func (r *TotpRepo) UpsertEnabledByUserID(ctx context.Context, userID int64, secret string) (*model.Totp, error) {
	tx := r.getClient(ctx)
	exist, err := tx.Totp.Query().Where(totp.UserID(userID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		existing, err := tx.Totp.Query().Where(totp.UserID(userID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		saved, err := tx.Totp.UpdateOneID(existing.ID).
			SetEnable(true).
			SetEnableTime(time.Now()).
			SetSecret(secret).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Totp{
			ID:         saved.ID,
			UserID:     saved.UserID,
			Enable:     saved.Enable,
			EnableTime: saved.EnableTime,
			Secret:     saved.Secret,
		}, nil
	}
	saved, err := tx.Totp.Create().
		SetUserID(userID).
		SetEnable(true).
		SetEnableTime(time.Now()).
		SetSecret(secret).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Totp{
		ID:         saved.ID,
		UserID:     saved.UserID,
		Enable:     saved.Enable,
		EnableTime: saved.EnableTime,
		Secret:     saved.Secret,
	}, nil
}

func (r *TotpRepo) DisableByUserID(ctx context.Context, userID int64) (*model.Totp, error) {
	tx := r.getClient(ctx)
	existing, err := tx.Totp.Query().Where(totp.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	saved, err := tx.Totp.UpdateOneID(existing.ID).
		SetEnable(false).
		SetSecret("").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Totp{
		ID:         saved.ID,
		UserID:     saved.UserID,
		Enable:     saved.Enable,
		EnableTime: saved.EnableTime,
		Secret:     saved.Secret,
	}, nil
}
