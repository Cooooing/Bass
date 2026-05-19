package usecase

import (
	"common/api/gen/common"
	utilent "common/pkg/util/ent"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"context"
	"errors"
)

type ContentUsecase struct {
	tx base.Tx

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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		reply, err = d.domainRepo.Saves(ctx, c, domains)
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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		reply, err = d.domainRepo.Update(ctx, c, domain)
		return err
	})
	return reply, err
}

func (d *ContentUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.DomainGetReq) ([]*model.Domain, *common.PageReply, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	return d.domainRepo.GetPage(ctx, c, page, req)
}
