package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/userprivacy"

	utilent "common/pkg/util/ent"
)

type UserPrivacyRepo struct {
	*base.BaseData
}

func NewUserPrivacyRepo(repo *base.BaseData) repo.UserPrivacyRepo {
	return &UserPrivacyRepo{
		BaseData: repo,
	}
}

func (r *UserPrivacyRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.Db
}

func privacyToDomain(p *gen.UserPrivacy) *model.UserPrivacy {
	return &model.UserPrivacy{
		ID:                 p.ID,
		UserID:             p.UserID,
		PublicPoints:       p.PublicPoints,
		PublicFollowers:    p.PublicFollowers,
		PublicArticles:     p.PublicArticles,
		PublicComments:     p.PublicComments,
		PublicOnlineStatus: p.PublicOnlineStatus,
		PublicLocation:     p.PublicLocation,
	}
}

func (r *UserPrivacyRepo) GetByUserID(ctx context.Context, userID int64) (*model.UserPrivacy, error) {
	tx := r.getClient(ctx)
	p, err := tx.UserPrivacy.Query().Where(userprivacy.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return privacyToDomain(p), nil
}

func (r *UserPrivacyRepo) Update(ctx context.Context, p *model.UserPrivacy) (*model.UserPrivacy, error) {
	tx := r.getClient(ctx)
	saved, err := tx.UserPrivacy.UpdateOneID(p.ID).
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
	return privacyToDomain(saved), nil
}
