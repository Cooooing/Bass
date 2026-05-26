package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/privacysetting"

	utilent "common/pkg/util/ent"
)

var _ repo.PrivacySettingRepo = (*PrivacySettingRepo)(nil)

type PrivacySettingRepo struct {
	db *gen.Client
}

func NewPrivacySettingRepo(db *gen.Client) repo.PrivacySettingRepo {
	return &PrivacySettingRepo{
		db: db,
	}
}

func (r *PrivacySettingRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *PrivacySettingRepo) FindByUserID(ctx context.Context, userID int64) (*model.PrivacySetting, error) {
	tx := r.getClient(ctx)
	p, err := tx.PrivacySetting.Query().Where(privacysetting.UserID(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.PrivacySetting{
		ID:                 p.ID,
		UserID:             p.UserID,
		PublicPoints:       p.PublicPoints,
		PublicFollowers:    p.PublicFollowers,
		PublicArticles:     p.PublicArticles,
		PublicComments:     p.PublicComments,
		PublicOnlineStatus: p.PublicOnlineStatus,
		PublicLocation:     p.PublicLocation,
	}, nil
}

func (r *PrivacySettingRepo) UpsertByUserID(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error) {
	existing, err := r.FindByUserID(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		saved, err := tx.PrivacySetting.Create().
			SetUserID(p.UserID).
			SetNillablePublicPoints(p.PublicPoints).
			SetNillablePublicFollowers(p.PublicFollowers).
			SetNillablePublicArticles(p.PublicArticles).
			SetNillablePublicComments(p.PublicComments).
			SetNillablePublicOnlineStatus(p.PublicOnlineStatus).
			SetNillablePublicLocation(p.PublicLocation).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.PrivacySetting{
			ID:                 saved.ID,
			UserID:             saved.UserID,
			PublicPoints:       saved.PublicPoints,
			PublicFollowers:    saved.PublicFollowers,
			PublicArticles:     saved.PublicArticles,
			PublicComments:     saved.PublicComments,
			PublicOnlineStatus: saved.PublicOnlineStatus,
			PublicLocation:     saved.PublicLocation,
		}, nil
	}
	p.ID = existing.ID
	return r.Update(ctx, p)
}

func (r *PrivacySettingRepo) Update(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error) {
	tx := r.getClient(ctx)
	saved, err := tx.PrivacySetting.UpdateOneID(p.ID).
		SetNillablePublicPoints(p.PublicPoints).
		SetNillablePublicFollowers(p.PublicFollowers).
		SetNillablePublicArticles(p.PublicArticles).
		SetNillablePublicComments(p.PublicComments).
		SetNillablePublicOnlineStatus(p.PublicOnlineStatus).
		SetNillablePublicLocation(p.PublicLocation).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PrivacySetting{
		ID:                 saved.ID,
		UserID:             saved.UserID,
		PublicPoints:       saved.PublicPoints,
		PublicFollowers:    saved.PublicFollowers,
		PublicArticles:     saved.PublicArticles,
		PublicComments:     saved.PublicComments,
		PublicOnlineStatus: saved.PublicOnlineStatus,
		PublicLocation:     saved.PublicLocation,
	}, nil
}
