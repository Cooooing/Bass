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

type TagUsecase struct {
	tx base.Tx

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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		reply, err = t.tagRepo.Saves(ctx, c, tags)
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
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return errors.New("no transaction in context")
		}
		reply, err = t.tagRepo.Update(ctx, c, tag)
		return err
	})
	return reply, err
}

func (t *TagUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *common.PageReply, error) {
	c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
	if !ok {
		return nil, nil, errors.New("no client in context")
	}
	return t.tagRepo.GetPage(ctx, c, page, req)
}
