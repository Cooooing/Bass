package biz

import (
	cv1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
)

type DomainDomain struct {
	*BaseDomain
	domainRepo repo.DomainRepo
}

func NewDomainDomain(baseDomain *BaseDomain, domainRepo repo.DomainRepo) *DomainDomain {
	return &DomainDomain{
		BaseDomain: baseDomain,
		domainRepo: domainRepo,
	}
}

func (d *DomainDomain) Adds(ctx context.Context, domains []*model.Domain) ([]*model.Domain, error) {
	var (
		reply []*model.Domain
		err   error
	)
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		reply, err = d.domainRepo.Saves(ctx, tx, domains)
		return err
	})
	return reply, err
}

func (d *DomainDomain) Update(ctx context.Context, domain *model.Domain) (*model.Domain, error) {
	var (
		reply *model.Domain
		err   error
	)
	err = ent.WithTx(ctx, d.db, func(tx *gen.Client) error {
		reply, err = d.domainRepo.Update(ctx, tx, domain)
		return err
	})
	return reply, err
}

func (d *DomainDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.DomainGetReq) ([]*model.Domain, *cv1.PageReply, error) {
	return d.domainRepo.GetPage(ctx, d.db, page, req)
}
