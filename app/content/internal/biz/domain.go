package biz

import (
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"context"
)

type DomainDomain struct {
	*BaseDomain
	domainRepo repo.Domain
}

func NewDomainDomain(baseDomain *BaseDomain, domainRepo repo.Domain) *DomainDomain {
	return &DomainDomain{
		BaseDomain: baseDomain,
		domainRepo: domainRepo,
	}
}

func (d *DomainDomain) AddDomain(ctx context.Context, domain *model.Domain) (*model.Domain, error) {
	return d.domainRepo.Save(ctx, domain)
}

func (d *DomainDomain) UpdateDomain(ctx context.Context, domain *model.Domain) (*model.Domain, error) {
	return d.domainRepo.Update(ctx, domain)
}
