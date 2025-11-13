package biz

import (
	cv1 "common/api/common/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
)

type TagDomain struct {
	*BaseDomain
	tagRepo repo.TagRepo
}

func NewTagDomain(baseDomain *BaseDomain, tagRepo repo.TagRepo) *TagDomain {
	return &TagDomain{
		BaseDomain: baseDomain,
		tagRepo:    tagRepo,
	}
}

func (t *TagDomain) Saves(ctx context.Context, tags []*model.Tag) ([]*model.Tag, error) {
	var (
		reply []*model.Tag
		err   error
	)
	err = ent.WithTx(ctx, t.db, func(tx *gen.Client) error {
		reply, err = t.tagRepo.Saves(ctx, tx, tags)
		return err
	})
	return reply, err
}

func (t *TagDomain) Update(ctx context.Context, tag *model.Tag) (*model.Tag, error) {
	var (
		reply *model.Tag
		err   error
	)
	err = ent.WithTx(ctx, t.db, func(tx *gen.Client) error {
		reply, err = t.tagRepo.Update(ctx, tx, tag)
		return err
	})
	return reply, err
}

func (t *TagDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *cv1.PageReply, error) {
	return t.tagRepo.GetPage(ctx, t.db, page, req)
}
