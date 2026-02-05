package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	basedata "content/internal/data/base"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/tag"
	"context"
)

type TagRepo struct {
	*basedata.BaseData
}

func NewTagRepo(BaseData *basedata.BaseData) repo.TagRepo {
	return &TagRepo{
		BaseData: BaseData,
	}
}

func (r *TagRepo) Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error) {
	save, err := tx.Tag.Create().
		SetName(tag.Name).
		SetNillableDomainID(tag.DomainID).
		SetStatus(int32(v1.TagStatus_TAG_STATUS_NORMAL)).
		Save(ctx)
	return &model.Tag{Tag: save}, err
}

func (r *TagRepo) Saves(ctx context.Context, tx *gen.Client, tags []*model.Tag) ([]*model.Tag, error) {
	creates := make([]*gen.TagCreate, 0, len(tags))
	for i := range tags {
		creates = append(creates,
			tx.Tag.Create().
				SetName(tags[i].Name).
				SetNillableDomainID(tags[i].DomainID).
				SetStatus(int32(v1.TagStatus_TAG_STATUS_NORMAL)),
		)
	}
	save, err := tx.Tag.CreateBulk(creates...).Save(ctx)
	res := make([]*model.Tag, 0, len(save))
	for _, item := range save {
		res = append(res, &model.Tag{Tag: item})
	}
	return res, err
}

func (r *TagRepo) Update(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error) {
	update := tx.Tag.UpdateOneID(tag.ID).
		SetName(tag.Name).
		SetNillableDescription(tag.Description).
		SetNillableDomainID(tag.DomainID).
		SetStatus(int32(v1.TagStatus_TAG_STATUS_NORMAL))
	save, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Tag{Tag: save}, nil
}

func (r *TagRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.TagGetReq) (*model.Tag, error) {
	query := tx.Tag.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("tag is not found")
	}
	return &model.Tag{Tag: t}, err
}

func (r *TagRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.TagGetReq) ([]*model.Tag, error) {
	var (
		tags []*model.Tag
		err  error
	)
	query := tx.Tag.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		tags = append(tags, &model.Tag{Tag: item})
	}
	return tags, nil
}

func (r *TagRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *cv1.PageReply, error) {
	var (
		tags []*model.Tag
		err  error
	)
	page = constant.PageValid(page)
	query := tx.Tag.Query()
	query = r.getQuery(query, req)
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
		tags = append(tags, &model.Tag{Tag: item})
	}
	return tags, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *TagRepo) getQuery(query *gen.TagQuery, req *repo.TagGetReq) *gen.TagQuery {
	if req.TagId != nil {
		query = query.Where(tag.IDEQ(*req.TagId))
	}
	if req.TagIds != nil {
		query = query.Where(tag.IDIn(req.TagIds...))
	}
	if req.UserId != nil {
		query = query.Where(tag.CreatedBy(*req.UserId))
	}
	if req.Name != nil {
		query = query.Where(tag.NameContains(*req.Name))
	}
	if len(req.Names) > 0 {
		query = query.Where(tag.NameIn(req.Names...))
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
