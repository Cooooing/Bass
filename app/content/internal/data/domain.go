package data

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"
)

type DomainRepo struct {
	*BaseRepo
	client *gen.Client
}

func NewDomainRepo(baseRepo *BaseRepo, client *gen.Client) repo.Domain {
	return &DomainRepo{
		BaseRepo: baseRepo,
		client:   client,
	}
}

func (r *DomainRepo) Save(ctx context.Context, d *model.Domain) (*model.Domain, error) {
	return nil, nil
}
