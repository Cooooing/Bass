package usecase

import (
	"context"

	"common/proto/gen/common"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
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

func (t *TagUsecase) Saves(ctx context.Context, tags []*model.Tag) ([]*model.Tag, error) {
	var (
		reply []*model.Tag
		err   error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		reply, err = t.tagRepo.Saves(ctx, tags)
		return err
	})
	return reply, err
}

func (t *TagUsecase) Update(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	var (
		reply *model.Tag
		err   error
	)
	err = t.tx(ctx, func(ctx context.Context) error {
		reply, err = t.tagRepo.Update(ctx, tag)
		return err
	})
	return reply, err
}

func (t *TagUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *common.PageReply, error) {
	return t.tagRepo.Page(ctx, page, req)
}
