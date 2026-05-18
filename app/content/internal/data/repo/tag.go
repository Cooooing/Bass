package repo

import (
	common "common/api/gen/common"
	cerrors "common/api/gen/common/errors"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/conf"
	"content/internal/data/gen"
	tagent "content/internal/data/gen/tag"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TagRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewTagRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.TagRepo {
	return &TagRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *TagRepo) Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error) {
	save, err := tx.Tag.Create().
		SetName(tag.Name).
		SetNillableDomainID(tag.DomainID).
		SetStatus(tagent.StatusNormal).
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
				SetStatus(tagent.StatusNormal),
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
		SetStatus(tagent.StatusNormal)
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
		return nil, cerrors.ErrorBadRequest("tag is not found")
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

func (r *TagRepo) GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *repo.TagGetReq) ([]*model.Tag, *common.PageReply, error) {
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
	return tags, &common.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *TagRepo) getQuery(query *gen.TagQuery, req *repo.TagGetReq) *gen.TagQuery {
	if req.TagId != nil {
		query = query.Where(tagent.IDEQ(*req.TagId))
	}
	if req.TagIds != nil {
		query = query.Where(tagent.IDIn(req.TagIds...))
	}
	if req.UserId != nil {
		query = query.Where(tagent.CreatedBy(*req.UserId))
	}
	if req.Name != nil {
		query = query.Where(tagent.NameContains(*req.Name))
	}
	if len(req.Names) > 0 {
		query = query.Where(tagent.NameIn(req.Names...))
	}
	if req.Description != nil {
		query = query.Where(tagent.DescriptionContains(*req.Description))
	}
	if req.Status != nil {
		dbStatus, _ := enum.TagStatusMap.ToEnum(*req.Status)
		query = query.Where(tagent.StatusEQ(tagent.Status(dbStatus)))
	}
	if req.DomainId != nil {
		query = query.Where(tagent.DomainIDEQ(*req.DomainId))
	}
	if req.ArticleCount != nil {
		if req.ArticleCount.Start != nil {
			query = query.Where(tagent.ArticleCountGTE(*req.ArticleCount.Start))
		}
		if req.ArticleCount.End != nil {
			query = query.Where(tagent.ArticleCountLTE(*req.ArticleCount.End))
		}
	}
	return query
}
