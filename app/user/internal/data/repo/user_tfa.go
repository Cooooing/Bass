package repo

import (
	"context"
	"time"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/usertfa"

	utilent "common/pkg/util/ent"
)

type UserTfaRepo struct {
	*base.BaseData
}

func NewUserTfaRepo(repo *base.BaseData) repo.UserTfaRepo {
	return &UserTfaRepo{
		BaseData: repo,
	}
}

func (r *UserTfaRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.Db
}

func tfaToDomain(t *gen.UserTFA) *model.UserTFA {
	return &model.UserTFA{
		ID:         t.ID,
		UserID:     t.UserID,
		Enable:     t.Enable,
		EnableTime: t.EnableTime,
		Secret:     t.Secret,
	}
}

func (r *UserTfaRepo) GetByUserID(ctx context.Context, userID int64) (*model.UserTFA, error) {
	tx := r.getClient(ctx)
	t, err := tx.UserTFA.Query().Where(usertfa.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return tfaToDomain(t), nil
}

func (r *UserTfaRepo) Enable(ctx context.Context, userID int64, secret string) (*model.UserTFA, error) {
	tx := r.getClient(ctx)
	exist, err := tx.UserTFA.Query().Where(usertfa.UserID(userID)).Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		existing, err := tx.UserTFA.Query().Where(usertfa.UserID(userID)).Only(ctx)
		if err != nil {
			return nil, err
		}
		saved, err := tx.UserTFA.UpdateOneID(existing.ID).
			SetEnable(true).
			SetEnableTime(time.Now()).
			SetSecret(secret).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return tfaToDomain(saved), nil
	}
	saved, err := tx.UserTFA.Create().
		SetUserID(userID).
		SetEnable(true).
		SetEnableTime(time.Now()).
		SetSecret(secret).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return tfaToDomain(saved), nil
}

func (r *UserTfaRepo) Disable(ctx context.Context, userID int64) (*model.UserTFA, error) {
	tx := r.getClient(ctx)
	existing, err := tx.UserTFA.Query().Where(usertfa.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	saved, err := tx.UserTFA.UpdateOneID(existing.ID).
		SetEnable(false).
		SetSecret("").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return tfaToDomain(saved), nil
}
