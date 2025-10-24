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

func NewDomainRepo(baseRepo *BaseRepo) repo.DomainRepo {
	return &DomainRepo{
		BaseRepo: baseRepo,
	}
}

func (r *DomainRepo) Save(ctx context.Context, db *gen.Client, domain *model.Domain) (*model.Domain, error) {
	create := db.Domain.Create().
		SetName(domain.Name).
		SetDescription(domain.Description).
		SetStatus(domain.Status).
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

func (r *DomainRepo) AddTagCount(ctx context.Context, db *gen.Client, id int, num int) (*model.Domain, error) {
	domain, err := db.Domain.UpdateOneID(id).
		AddTagCount(num).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Domain)(domain), nil
}

func (r *DomainRepo) Get(ctx context.Context, db *gen.Client) ([]*model.Domain, error) {
	query := db.Domain.Query()

	domains, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.Domain, len(domains))
	for i, domain := range domains {
		res[i] = (*model.Domain)(domain)
	}
	return res, nil
}
