package usecase

import (
	"context"

	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
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
		rows []*model.Domain
		err  error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		savesResp, saveErr := d.domainRepo.Saves(ctx, domains)
		if saveErr != nil {
			return saveErr
		}
		rows = savesResp
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *ContentUsecase) Update(ctx context.Context, domain *model.Domain) (*model.Domain, error) {
	var (
		updated *model.Domain
		err     error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		updateResp, updateErr := d.domainRepo.Update(ctx, domain)
		if updateErr != nil {
			return updateErr
		}
		updated = updateResp
		return err
	})
	if err != nil {
		return nil, err
	}
	return updated, nil
}

type DomainPageReq struct {
	Page        *base.PageRequest
	DomainIDs   []int64
	Name        *string
	Description *string
	Status      *enum.DomainStatus
	URL         *string
	Icon        *string
	IsNav       *bool
}

type DomainPageResp struct {
	Rows []*model.Domain
	Page *base.PageResp
}

func (d *ContentUsecase) Page(ctx context.Context, req *DomainPageReq) (*DomainPageResp, error) {
	if req == nil {
		req = &DomainPageReq{}
	}
	pageResp, err := d.domainRepo.Page(ctx, &repo.DomainGetReq{
		Page:        req.Page,
		DomainIds:   req.DomainIDs,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		Url:         req.URL,
		Icon:        req.Icon,
		IsNav:       req.IsNav,
	})
	if err != nil {
		return nil, err
	}
	return &DomainPageResp{
		Rows: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}
