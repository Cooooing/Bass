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

type DomainAddsReq struct {
	Domains []*model.Domain
}

type DomainAddsResponse struct {
	Rows []*model.Domain
}

func (d *ContentUsecase) Adds(ctx context.Context, req *DomainAddsReq) (*DomainAddsResponse, error) {
	var (
		rows []*model.Domain
		err  error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		savesResponse, saveErr := d.domainRepo.Saves(ctx, &repo.DomainSavesReq{Domains: req.Domains})
		if saveErr != nil {
			return saveErr
		}
		rows = savesResponse.Rows
		return err
	})
	if err != nil {
		return nil, err
	}
	return &DomainAddsResponse{Rows: rows}, nil
}

type DomainUpdateReq struct {
	Domain *model.Domain
}

type DomainUpdateResponse struct {
	Domain *model.Domain
}

func (d *ContentUsecase) Update(ctx context.Context, req *DomainUpdateReq) (*DomainUpdateResponse, error) {
	var (
		domain *model.Domain
		err    error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		updateResponse, updateErr := d.domainRepo.Update(ctx, &repo.DomainUpdateReq{Domain: req.Domain})
		if updateErr != nil {
			return updateErr
		}
		domain = updateResponse.Domain
		return err
	})
	if err != nil {
		return nil, err
	}
	return &DomainUpdateResponse{Domain: domain}, nil
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

type DomainPageResponse struct {
	Rows []*model.Domain
	Page *base.PageResponse
}

func (d *ContentUsecase) Page(ctx context.Context, req *DomainPageReq) (*DomainPageResponse, error) {
	if req == nil {
		req = &DomainPageReq{}
	}
	pageResponse, err := d.domainRepo.Page(ctx, &repo.DomainGetReq{
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
	return &DomainPageResponse{Rows: pageResponse.Rows, Page: pageResponse.Page}, nil
}
