package repo

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	basedata "content/internal/data/base"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/domain"
	"context"
)

type DomainRepo struct {
	*basedata.BaseData
}

func NewDomainRepo(BaseData *basedata.BaseData) repo.DomainRepo {
	return &DomainRepo{
		BaseData: BaseData,
	}
}

func (r *DomainRepo) Save(ctx context.Context, tx *gen.Client, domain *model.Domain) (*model.Domain, error) {
	save, err := tx.Domain.Create().
		SetName(domain.Name).
		SetNillableDescription(domain.Description).
		SetStatus(domain.Status).
		SetNillableURL(domain.URL).
		SetNillableIcon(domain.Icon).
		SetIsNav(domain.IsNav).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Domain{Domain: save}, nil
}

func (r *DomainRepo) Saves(ctx context.Context, tx *gen.Client, domains []*model.Domain) ([]*model.Domain, error) {

	creates := make([]*gen.DomainCreate, 0, len(domains))
	for i := range domains {
		creates = append(creates,
			tx.Domain.Create().
				SetName(domains[i].Name).
				SetNillableDescription(domains[i].Description).
				SetStatus(domains[i].Status).
				SetNillableURL(domains[i].URL).
				SetNillableIcon(domains[i].Icon).
				SetIsNav(domains[i].IsNav),
		)
	}

	save, err := tx.Domain.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.Domain, len(save))
	for i := range save {
		res[i] = &model.Domain{Domain: save[i]}
	}
	return res, nil
}

func (r *DomainRepo) Update(ctx context.Context, tx *gen.Client, domain *model.Domain) (*model.Domain, error) {
	update := tx.Domain.UpdateOneID(domain.ID).
		SetName(domain.Name).
		SetNillableDescription(domain.Description).
		SetStatus(domain.Status).
		SetNillableURL(domain.URL).
		SetNillableIcon(domain.Icon).
		SetIsNav(domain.IsNav)
	save, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Domain{Domain: save}, nil
}

func (r *DomainRepo) AddTagCount(ctx context.Context, tx *gen.Client, id int64, num int32) (*model.Domain, error) {
	save, err := tx.Domain.UpdateOneID(id).
		AddTagCount(num).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Domain{Domain: save}, nil
}

func (r *DomainRepo) GetOne(ctx context.Context, tx *gen.Client, req *repo.DomainGetReq) (*model.Domain, error) {
	query := tx.Domain.Query()
	query = r.getQuery(query, req)
	d, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, cv1.ErrorBadRequest("domain is not found")
	}
	return &model.Domain{Domain: d}, err
}

func (r *DomainRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.DomainGetReq) ([]*model.Domain, error) {
	var (
		domains []*model.Domain
		err     error
	)
	query := tx.Domain.Query().WithTags()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		domains = append(domains, &model.Domain{Domain: list[i]})
	}
	return domains, nil
}

func (r *DomainRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.DomainGetReq) ([]*model.Domain, *cv1.PageReply, error) {
	var (
		domains []*model.Domain
		err     error
		total   int
	)
	page = constant.PageValid(page)
	query := tx.Domain.Query().WithTags()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err = countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	for i := range list {
		domains = append(domains, &model.Domain{Domain: list[i]})
	}
	return domains, &cv1.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *DomainRepo) getQuery(query *gen.DomainQuery, req *repo.DomainGetReq) *gen.DomainQuery {
	if req.DomainId != nil {
		query = query.Where(domain.IDEQ(*req.DomainId))
	}
	if len(req.DomainIds) > 0 {
		query = query.Where(domain.IDIn(req.DomainIds...))
	}
	if req.Name != nil {
		query = query.Where(domain.NameContains(*req.Name))
	}
	if req.Description != nil {
		query = query.Where(domain.DescriptionContains(*req.Description))
	}
	if req.Status != nil {
		query = query.Where(domain.StatusEQ(int32(*req.Status)))
	}
	if req.Url != nil {
		query = query.Where(domain.URLContains(*req.Url))
	}
	if req.Icon != nil {
		query = query.Where(domain.IconContains(*req.Icon))
	}
	if req.TagCount != nil {
		if req.TagCount.Start != nil {
			query = query.Where(domain.TagCountGTE(*req.TagCount.Start))
		}
		if req.TagCount.End != nil {
			query = query.Where(domain.TagCountLTE(*req.TagCount.End))
		}
	}
	if req.IsNav != nil {
		query = query.Where(domain.IsNavEQ(*req.IsNav))
	}
	return query
}
