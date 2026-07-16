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

type TagSavesReq struct {
	Tags []*model.Tag
}

type TagSavesResponse struct {
	Rows []*model.Tag
}

func (t *TagUsecase) Saves(ctx context.Context, req *TagSavesReq) (*TagSavesResponse, error) {
	var (
		rows []*model.Tag
		err  error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		savesResponse, saveErr := t.tagRepo.Saves(ctx, &repo.TagSavesReq{Tags: req.Tags})
		if saveErr != nil {
			return saveErr
		}
		rows = savesResponse.Rows
		return err
	})
	if err != nil {
		return nil, err
	}
	return &TagSavesResponse{Rows: rows}, nil
}

type TagUpdateReq struct {
	Tag *model.Tag
}

type TagUpdateResponse struct {
	Tag *model.Tag
}

func (t *TagUsecase) Update(ctx context.Context, req *TagUpdateReq) (*TagUpdateResponse, error) {
	var (
		tag *model.Tag
		err error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		updateResponse, updateErr := t.tagRepo.Update(ctx, &repo.TagUpdateReq{Tag: req.Tag})
		if updateErr != nil {
			return updateErr
		}
		tag = updateResponse.Tag
		return err
	})
	if err != nil {
		return nil, err
	}
	return &TagUpdateResponse{Tag: tag}, nil
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

type TagPageResponse struct {
	Rows []*model.Tag
	Page *base.PageResponse
}

func (t *TagUsecase) Page(ctx context.Context, req *TagPageReq) (*TagPageResponse, error) {
	if req == nil {
		req = &TagPageReq{}
	}
	pageResponse, err := t.tagRepo.Page(ctx, &repo.TagGetReq{
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
	return &TagPageResponse{Rows: pageResponse.Rows, Page: pageResponse.Page}, nil
}
