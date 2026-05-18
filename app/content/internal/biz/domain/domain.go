package domain

import (
	"common/api/gen/common"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/client"
	"content/internal/data/gen"
	"context"
)

type DomainDomain struct {
	db *gen.Client

	domainRepo repo.DomainRepo
}

func NewDomainDomain(
	db *gen.Client,
	domainRepo repo.DomainRepo,
) *DomainDomain {
	return &DomainDomain{
		db:         db,
		domainRepo: domainRepo,
	}
}

func (d *DomainDomain) Adds(ctx context.Context, domains []*model.Domain) ([]*model.Domain, error) {
	var (
		reply []*model.Domain
		err   error
	)
	err = client.WithTx(ctx, d.db, func(tx *gen.Client) error {
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
	err = client.WithTx(ctx, d.db, func(tx *gen.Client) error {
		reply, err = d.domainRepo.Update(ctx, tx, domain)
		return err
	})
	return reply, err
}

func (d *DomainDomain) Page(ctx context.Context, page *common.PageRequest, req *repo.DomainGetReq) ([]*model.Domain, *common.PageReply, error) {
	return d.domainRepo.GetPage(ctx, d.db, page, req)
}
