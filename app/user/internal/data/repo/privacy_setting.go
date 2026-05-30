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
		PublicPoints:       new(p.PublicPoints),
		PublicFollowers:    new(p.PublicFollowers),
		PublicArticles:     new(p.PublicArticles),
		PublicComments:     new(p.PublicComments),
		PublicOnlineStatus: new(p.PublicOnlineStatus),
		PublicLocation:     new(p.PublicLocation),
	}, nil
}

func (r *PrivacySettingRepo) UpsertByUserID(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error) {
	existing, err := r.FindByUserID(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		create := tx.PrivacySetting.Create().
			SetUserID(p.UserID)
		if p.PublicPoints != nil {
			create.SetPublicPoints(*p.PublicPoints)
		}
		if p.PublicFollowers != nil {
			create.SetPublicFollowers(*p.PublicFollowers)
		}
		if p.PublicArticles != nil {
			create.SetPublicArticles(*p.PublicArticles)
		}
		if p.PublicComments != nil {
			create.SetPublicComments(*p.PublicComments)
		}
		if p.PublicOnlineStatus != nil {
			create.SetPublicOnlineStatus(*p.PublicOnlineStatus)
		}
		if p.PublicLocation != nil {
			create.SetPublicLocation(*p.PublicLocation)
		}
		saved, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.PrivacySetting{
			ID:                 saved.ID,
			UserID:             saved.UserID,
			PublicPoints:       new(saved.PublicPoints),
			PublicFollowers:    new(saved.PublicFollowers),
			PublicArticles:     new(saved.PublicArticles),
			PublicComments:     new(saved.PublicComments),
			PublicOnlineStatus: new(saved.PublicOnlineStatus),
			PublicLocation:     new(saved.PublicLocation),
		}, nil
	}
	p.ID = existing.ID
	return r.Update(ctx, p)
}

func (r *PrivacySettingRepo) Update(ctx context.Context, p *model.PrivacySetting) (*model.PrivacySetting, error) {
	tx := r.getClient(ctx)
	if p.PublicPoints == nil &&
		p.PublicFollowers == nil &&
		p.PublicArticles == nil &&
		p.PublicComments == nil &&
		p.PublicOnlineStatus == nil &&
		p.PublicLocation == nil {
		saved, err := tx.PrivacySetting.Get(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		return &model.PrivacySetting{
			ID:                 saved.ID,
			UserID:             saved.UserID,
			PublicPoints:       new(saved.PublicPoints),
			PublicFollowers:    new(saved.PublicFollowers),
			PublicArticles:     new(saved.PublicArticles),
			PublicComments:     new(saved.PublicComments),
			PublicOnlineStatus: new(saved.PublicOnlineStatus),
			PublicLocation:     new(saved.PublicLocation),
		}, nil
	}
	update := tx.PrivacySetting.UpdateOneID(p.ID)
	if p.PublicPoints != nil {
		update.SetPublicPoints(*p.PublicPoints)
	}
	if p.PublicFollowers != nil {
		update.SetPublicFollowers(*p.PublicFollowers)
	}
	if p.PublicArticles != nil {
		update.SetPublicArticles(*p.PublicArticles)
	}
	if p.PublicComments != nil {
		update.SetPublicComments(*p.PublicComments)
	}
	if p.PublicOnlineStatus != nil {
		update.SetPublicOnlineStatus(*p.PublicOnlineStatus)
	}
	if p.PublicLocation != nil {
		update.SetPublicLocation(*p.PublicLocation)
	}
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PrivacySetting{
		ID:                 saved.ID,
		UserID:             saved.UserID,
		PublicPoints:       new(saved.PublicPoints),
		PublicFollowers:    new(saved.PublicFollowers),
		PublicArticles:     new(saved.PublicArticles),
		PublicComments:     new(saved.PublicComments),
		PublicOnlineStatus: new(saved.PublicOnlineStatus),
		PublicLocation:     new(saved.PublicLocation),
	}, nil
}
