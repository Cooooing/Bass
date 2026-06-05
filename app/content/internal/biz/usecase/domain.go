package usecase

import (
	"context"

	"common/api/gen/common"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
)

type ContentUsecase struct {
	tx         base.Tx
	domainRepo repo.DomainRepo
}

func NewContentUsecase(
	tx base.Tx,
	domainRepo repo.DomainRepo,
) *ContentUsecase {
	return &ContentUsecase{
		tx:         tx,
		domainRepo: domainRepo,
	}
}

func (d *ContentUsecase) Adds(ctx context.Context, domains []*model.Domain) ([]*model.Domain, error) {
	var (
		reply []*model.Domain
		err   error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		reply, err = d.domainRepo.Saves(ctx, domains)
		return err
	})
	return reply, err
}

func (d *ContentUsecase) Update(ctx context.Context, domain *model.Domain) (*model.Domain, error) {
	var (
		reply *model.Domain
		err   error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		reply, err = d.domainRepo.Update(ctx, domain)
		return err
	})
	return reply, err
}

func (d *ContentUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.DomainGetReq) ([]*model.Domain, *common.PageReply, error) {
	return d.domainRepo.Page(ctx, page, req)
}
