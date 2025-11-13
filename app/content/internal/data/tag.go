package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"common/pkg/util/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/tag"
	"context"
)

type TagRepo struct {
	*BaseRepo
}

func NewTagRepo(baseRepo *BaseRepo) repo.TagRepo {
	return &TagRepo{
		BaseRepo: baseRepo,
	}
}

func (t *TagRepo) Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error) {
	save, err := tx.Tag.Create().
		SetName(tag.Name).
		SetNillableDomainID(tag.DomainID).
		SetStatus(int32(cv1.TagStatus_TagNormal)).
		Save(ctx)
	return (*model.Tag)(save), err
}

func (t *TagRepo) Saves(ctx context.Context, tx *gen.Client, tags []*model.Tag) ([]*model.Tag, error) {
	creates := make([]*gen.TagCreate, 0, len(tags))
	for i := range tags {
		creates = append(creates,
			tx.Tag.Create().
				SetName(tags[i].Name).
				SetNillableDomainID(tags[i].DomainID).
				SetStatus(int32(cv1.TagStatus_TagNormal)),
		)
	}
	save, err := tx.Tag.CreateBulk(creates...).Save(ctx)
	res := make([]*model.Tag, 0, len(save))
	for _, item := range save {
		res = append(res, (*model.Tag)(item))
	}
	return res, err
}

func (t *TagRepo) Update(ctx context.Context, db *gen.Client, tag *model.Tag) (*model.Tag, error) {
	update := db.Tag.UpdateOneID(tag.ID).
		SetName(tag.Name).
		SetNillableDescription(tag.Description).
		SetNillableDomainID(tag.DomainID).
		SetStatus(int32(cv1.TagStatus_TagNormal))
	save, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return (*model.Tag)(save), nil
}

func (t *TagRepo) GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Tag, error) {
	query, err := tx.Tag.Query().Where(tag.IDEQ(id)).First(ctx)
	return (*model.Tag)(query), err
}

func (t *TagRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.TagGetReq) ([]*model.Tag, error) {
	var (
		tags []*model.Tag
		err  error
	)
	query := tx.Tag.Query()
	query = t.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		tags = append(tags, (*model.Tag)(item))
	}
	return tags, nil
}

func (t *TagRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *cv1.PageReply, error) {
	var (
		tags []*model.Tag
		err  error
	)
	page = base.OrDefault(page, constant.GetPageDefault())
	query := tx.Tag.Query()
	query = t.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		tags = append(tags, (*model.Tag)(item))
	}
	return tags,
		&cv1.PageReply{
			Total: uint32(count),
			Size:  page.Size,
			Page:  page.Page,
		}, nil
}

func (r *TagRepo) getQuery(query *gen.TagQuery, req *repo.TagGetReq) *gen.TagQuery {
	if req.Ids != nil {
		query = query.Where(tag.IDIn(req.Ids...))
	}
	if req.UserId != nil {
		query = query.Where(tag.CreatedBy(*req.UserId))
	}
	if req.Name != nil {
		query = query.Where(tag.NameContains(*req.Name))
	}
	if req.Description != nil {
		query = query.Where(tag.DescriptionContains(*req.Description))
	}
	if req.Status != nil {
		query = query.Where(tag.StatusEQ(int32(*req.Status)))
	}
	if req.DomainId != nil {
		query = query.Where(tag.DomainIDEQ(*req.DomainId))
	}
	if req.ArticleCount != nil {
		if req.ArticleCount.Start != nil {
			query = query.Where(tag.ArticleCountGTE(*req.ArticleCount.Start))
		}
		if req.ArticleCount.End != nil {
			query = query.Where(tag.ArticleCountLTE(*req.ArticleCount.End))
		}
	}
	return query
}
