package data

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type DomainRepo struct {
	*BaseRepo
}

func NewDomainRepo(baseRepo *BaseRepo) repo.Domain {
	return &DomainRepo{
		BaseRepo: baseRepo,
	}
}

func (r *DomainRepo) Save(ctx context.Context, db *gen.Client, domain *model.Domain) (*model.Domain, error) {
	create := db.Domain.Create().
		SetName(domain.Name).
		SetDescription(domain.Description).
		SetNillableURL(domain.URL).
		SetNillableIcon(domain.Icon).
		SetIsNav(domain.IsNav)
	save, err := create.Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Domain)(save), nil
}

func (r *DomainRepo) Update(ctx context.Context, db *gen.Client, domain *model.Domain) (*model.Domain, error) {
	update := db.Domain.UpdateOneID(domain.ID).
		SetName(domain.Name).
		SetDescription(domain.Description).
		SetStatus(domain.Status).
		SetNillableURL(domain.URL).
		SetNillableIcon(domain.Icon).
		SetIsNav(domain.IsNav)
	save, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Domain)(save), nil
}

func (r *DomainRepo) AddTagCount(ctx context.Context, db *gen.Client, id int) (*model.Domain, error) {
	domain, err := db.Domain.UpdateOneID(id).
		AddTagCount(1).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Domain)(domain), nil
}
