package usecase

import (
	"context"

	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
)

type TagUsecase struct {
	tx      base.Tx
	tagRepo repo.TagRepo
}

func NewTagUsecase(
	tx base.Tx,
	tagRepo repo.TagRepo,
) *TagUsecase {
	return &TagUsecase{
		tx:      tx,
		tagRepo: tagRepo,
	}
}

func (t *TagUsecase) Saves(
	ctx context.Context,
	tags []*model.Tag,
) ([]*model.Tag, error) {
	var (
		rows []*model.Tag
		err  error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		savesResp, saveErr := t.tagRepo.Saves(ctx, tags)
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

func (t *TagUsecase) Update(
	ctx context.Context,
	tag *model.Tag,
) (*model.Tag, error) {
	var (
		updated *model.Tag
		err     error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		updateResp, updateErr := t.tagRepo.Update(ctx, tag)
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

type TagPageReq struct {
	Page        *base.PageRequest
	TagIDs      []int64
	Name        *string
	Names       []string
	Description *string
	Status      *enum.TagStatus
	DomainID    *int64
}

type TagPageResp struct {
	Rows []*model.Tag
	Page *base.PageResp
}

func (t *TagUsecase) Page(
	ctx context.Context,
	req *TagPageReq,
) (*TagPageResp, error) {
	if req == nil {
		req = &TagPageReq{}
	}
	pageResp, err := t.tagRepo.Page(ctx, &repo.TagGetReq{
		Page:        req.Page,
		TagIds:      req.TagIDs,
		Name:        req.Name,
		Names:       req.Names,
		Description: req.Description,
		Status:      req.Status,
		DomainId:    req.DomainID,
	})
	if err != nil {
		return nil, err
	}
	return &TagPageResp{
		Rows: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}
