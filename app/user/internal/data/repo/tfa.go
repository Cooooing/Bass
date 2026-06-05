package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/tfa"

	utilent "common/pkg/util/ent"
)

var _ repo.TfaRepo = (*TfaRepo)(nil)

type TfaRepo struct {
	db *gen.Client
}

func NewTfaRepo(db *gen.Client) repo.TfaRepo {
	return &TfaRepo{
		db: db,
	}
}

func (r *TfaRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *TfaRepo) FindByUserID(ctx context.Context, userID int64) (*model.TFA, error) {
	tx := r.getClient(ctx)
	t, err := tx.TFA.Query().Where(tfa.UserID(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.TFA{
		ID:         t.ID,
		UserID:     t.UserID,
		Enable:     t.Enable,
		EnableTime: t.EnableTime,
		Secret:     t.Secret,
	}, nil
}

func (r *TfaRepo) UpsertEnabledByUserID(ctx context.Context, userID int64, secret string) (*model.TFA, error) {
	tx := r.getClient(ctx)
	exist, err := tx.TFA.Query().Where(tfa.UserID(userID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		existing, err := tx.TFA.Query().Where(tfa.UserID(userID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		saved, err := tx.TFA.UpdateOneID(existing.ID).
			SetEnable(true).
			SetEnableTime(time.Now()).
			SetSecret(secret).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.TFA{
			ID:         saved.ID,
			UserID:     saved.UserID,
			Enable:     saved.Enable,
			EnableTime: saved.EnableTime,
			Secret:     saved.Secret,
		}, nil
	}
	saved, err := tx.TFA.Create().
		SetUserID(userID).
		SetEnable(true).
		SetEnableTime(time.Now()).
		SetSecret(secret).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.TFA{
		ID:         saved.ID,
		UserID:     saved.UserID,
		Enable:     saved.Enable,
		EnableTime: saved.EnableTime,
		Secret:     saved.Secret,
	}, nil
}

func (r *TfaRepo) DisableByUserID(ctx context.Context, userID int64) (*model.TFA, error) {
	tx := r.getClient(ctx)
	existing, err := tx.TFA.Query().Where(tfa.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	saved, err := tx.TFA.UpdateOneID(existing.ID).
		SetEnable(false).
		SetSecret("").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.TFA{
		ID:         saved.ID,
		UserID:     saved.UserID,
		Enable:     saved.Enable,
		EnableTime: saved.EnableTime,
		Secret:     saved.Secret,
	}, nil
}
