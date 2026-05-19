package usecase

import (
	"common/api/gen/common"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/client"
	"content/internal/data/gen"
	"context"
)

type TagUsecase struct {
	db *gen.Client

	tagRepo repo.TagRepo
}

func NewTagUsecase(
	db *gen.Client,
	tagRepo repo.TagRepo,
) *TagUsecase {
	return &TagUsecase{
		db:      db,
		tagRepo: tagRepo,
	}
}

func (t *TagUsecase) Saves(ctx context.Context, tags []*model.Tag) ([]*model.Tag, error) {
	var (
		reply []*model.Tag
		err   error
	)
	err = client.WithTx(ctx, t.db, func(tx *gen.Client) error {
		reply, err = t.tagRepo.Saves(ctx, tx, tags)
		return err
	})
	return reply, err
}

func (t *TagUsecase) Update(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	var (
		reply *model.Tag
		err   error
	)
	err = client.WithTx(ctx, t.db, func(tx *gen.Client) error {
		reply, err = t.tagRepo.Update(ctx, tx, tag)
		return err
	})
	return reply, err
}

func (t *TagUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *common.PageReply, error) {
	return t.tagRepo.GetPage(ctx, t.db, page, req)
}
