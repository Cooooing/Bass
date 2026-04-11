package domain

import (
	cv1 "common/api/gen/common/v1"
	domainbase "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent"
	"content/internal/data/ent/gen"
	"context"
)

type TagDomain struct {
	*domainbase.BaseDomain
	tagRepo repo.TagRepo
}

func NewTagDomain(baseDomain *domainbase.BaseDomain, tagRepo repo.TagRepo) *TagDomain {
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
	err = ent.WithTx(ctx, t.Db, func(tx *gen.Client) error {
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
	err = ent.WithTx(ctx, t.Db, func(tx *gen.Client) error {
		reply, err = t.tagRepo.Update(ctx, tx, tag)
		return err
	})
	return reply, err
}

func (t *TagDomain) Page(ctx context.Context, page *cv1.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *cv1.PageReply, error) {
	return t.tagRepo.GetPage(ctx, t.Db, page, req)
}
