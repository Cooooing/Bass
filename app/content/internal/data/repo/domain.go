package repo

import (
	cerrors "common/proto/gen/common/errors"
	"content/internal/biz/base"
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/gen"
	"content/internal/data/gen/domain"
	"content/internal/enum"

	"github.com/samber/lo"
)

var _ repo.DomainRepo = (*DomainRepo)(nil)

type DomainRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewDomainRepo(
	db *gen.Client,
) repo.DomainRepo {
	return &DomainRepo{
		db: db,
	}
}

func (r *DomainRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *DomainRepo) Save(ctx context.Context, domainModel *model.Domain) (*model.Domain, error) {
	save, err := r.getClient(ctx).Domain.Create().
		SetCode(domainModel.Code).
		SetName(domainModel.Name).
		SetNillableDescription(domainModel.Description).
		SetStatus(domain.Status(domainModel.Status)).
		SetNillableURL(domainModel.URL).
		SetNillableIcon(domainModel.Icon).
		SetIsNav(domainModel.IsNav).
		SetSort(domainModel.Sort).
		SetNillableCreatedBy(domainModel.CreatedBy).
		SetNillableUpdatedBy(domainModel.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Domain{
		ID:          save.ID,
		Code:        save.Code,
		Name:        save.Name,
		Description: save.Description,
		Status:      enum.DomainStatus(save.Status),
		URL:         save.URL,
		Icon:        save.Icon,
		IsNav:       save.IsNav,
		Sort:        save.Sort,
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
		CreatedBy:   save.CreatedBy,
		UpdatedBy:   save.UpdatedBy,
	}, nil
}

func (r *DomainRepo) Saves(ctx context.Context, domains []*model.Domain) ([]*model.Domain, error) {
	client := r.getClient(ctx)
	creates := make([]*gen.DomainCreate, 0, len(domains))
	for i := range domains {
		creates = append(creates,
			client.Domain.Create().
				SetCode(domains[i].Code).
				SetName(domains[i].Name).
				SetNillableDescription(domains[i].Description).
				SetStatus(domain.Status(domains[i].Status)).
				SetNillableURL(domains[i].URL).
				SetNillableIcon(domains[i].Icon).
				SetIsNav(domains[i].IsNav).
				SetSort(domains[i].Sort).
				SetNillableCreatedBy(domains[i].CreatedBy).
				SetNillableUpdatedBy(domains[i].UpdatedBy),
		)
	}

	save, err := client.Domain.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.Domain, len(save))
	for i := range save {
		res[i] = &model.Domain{
			ID:          save[i].ID,
			Code:        save[i].Code,
			Name:        save[i].Name,
			Description: save[i].Description,
			Status:      enum.DomainStatus(save[i].Status),
			URL:         save[i].URL,
			Icon:        save[i].Icon,
			IsNav:       save[i].IsNav,
			Sort:        save[i].Sort,
			CreatedAt:   save[i].CreatedAt,
			UpdatedAt:   save[i].UpdatedAt,
			CreatedBy:   save[i].CreatedBy,
			UpdatedBy:   save[i].UpdatedBy,
		}
	}
	return res, nil
}

func (r *DomainRepo) Update(ctx context.Context, domainModel *model.Domain) (*model.Domain, error) {
	save, err := r.getClient(ctx).Domain.UpdateOneID(domainModel.ID).
		SetCode(domainModel.Code).
		SetName(domainModel.Name).
		SetNillableDescription(domainModel.Description).
		SetStatus(domain.Status(domainModel.Status)).
		SetNillableURL(domainModel.URL).
		SetNillableIcon(domainModel.Icon).
		SetIsNav(domainModel.IsNav).
		SetSort(domainModel.Sort).
		SetNillableUpdatedBy(domainModel.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Domain{
		ID:          save.ID,
		Code:        save.Code,
		Name:        save.Name,
		Description: save.Description,
		Status:      enum.DomainStatus(save.Status),
		URL:         save.URL,
		Icon:        save.Icon,
		IsNav:       save.IsNav,
		Sort:        save.Sort,
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
		CreatedBy:   save.CreatedBy,
		UpdatedBy:   save.UpdatedBy,
	}, nil
}

func (r *DomainRepo) Get(ctx context.Context, req *repo.DomainGetReq) (*model.Domain, error) {
	query := r.getClient(ctx).Domain.Query()
	query = r.getQuery(query, req)
	d, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_DOMAIN_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Domain{
		ID:          d.ID,
		Code:        d.Code,
		Name:        d.Name,
		Description: d.Description,
		Status:      enum.DomainStatus(d.Status),
		URL:         d.URL,
		Icon:        d.Icon,
		IsNav:       d.IsNav,
		Sort:        d.Sort,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
		CreatedBy:   d.CreatedBy,
		UpdatedBy:   d.UpdatedBy,
	}, nil
}

func (r *DomainRepo) List(ctx context.Context, req *repo.DomainGetReq) ([]*model.Domain, error) {
	query := r.getClient(ctx).Domain.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	domains := make([]*model.Domain, 0, len(list))
	for i := range list {
		domainItem := &model.Domain{
			ID:          list[i].ID,
			Code:        list[i].Code,
			Name:        list[i].Name,
			Description: list[i].Description,
			Status:      enum.DomainStatus(list[i].Status),
			URL:         list[i].URL,
			Icon:        list[i].Icon,
			IsNav:       list[i].IsNav,
			Sort:        list[i].Sort,
			CreatedAt:   list[i].CreatedAt,
			UpdatedAt:   list[i].UpdatedAt,
			CreatedBy:   list[i].CreatedBy,
			UpdatedBy:   list[i].UpdatedBy,
		}
		domains = append(domains, domainItem)
	}
	return domains, nil
}

func (r *DomainRepo) Map(ctx context.Context, req *repo.DomainGetReq) (map[int64]*model.Domain, error) {
	listResp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return lo.SliceToMap(listResp, func(item *model.Domain) (int64, *model.Domain) {
		return item.ID, item
	}), nil
}

func (r *DomainRepo) Count(ctx context.Context, req *repo.DomainGetReq) (int, error) {
	query := r.getClient(ctx).Domain.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *DomainRepo) Page(ctx context.Context, req *repo.DomainGetReq) (*repo.DomainPageResp, error) {
	page := r.normalizePage(req.Page)
	query := r.getClient(ctx).Domain.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	total, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	domains := make([]*model.Domain, 0, len(list))
	for i := range list {
		domainItem := &model.Domain{
			ID:          list[i].ID,
			Code:        list[i].Code,
			Name:        list[i].Name,
			Description: list[i].Description,
			Status:      enum.DomainStatus(list[i].Status),
			URL:         list[i].URL,
			Icon:        list[i].Icon,
			IsNav:       list[i].IsNav,
			Sort:        list[i].Sort,
			CreatedAt:   list[i].CreatedAt,
			UpdatedAt:   list[i].UpdatedAt,
			CreatedBy:   list[i].CreatedBy,
			UpdatedBy:   list[i].UpdatedBy,
		}
		domains = append(domains, domainItem)
	}
	return &repo.DomainPageResp{
		Rows: domains,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *DomainRepo) getQuery(query *gen.DomainQuery, req *repo.DomainGetReq) *gen.DomainQuery {
	query = query.Where(domain.DeletedAtIsNil())
	if req == nil {
		return query
	}
	if req.DomainId != nil {
		query = query.Where(domain.IDEQ(*req.DomainId))
	}
	if len(req.DomainIds) > 0 {
		query = query.Where(domain.IDIn(req.DomainIds...))
	}
	if req.Code != nil {
		query = query.Where(domain.CodeContains(*req.Code))
	}
	if req.Name != nil {
		query = query.Where(domain.NameContains(*req.Name))
	}
	if req.Description != nil {
		query = query.Where(domain.DescriptionContains(*req.Description))
	}
	if req.Status != nil {
		query = query.Where(domain.StatusEQ(domain.Status(*req.Status)))
	}
	if req.Url != nil {
		query = query.Where(domain.URLContains(*req.Url))
	}
	if req.Icon != nil {
		query = query.Where(domain.IconContains(*req.Icon))
	}
	if req.IsNav != nil {
		query = query.Where(domain.IsNavEQ(*req.IsNav))
	}
	query = query.Order(gen.Asc(domain.FieldSort), gen.Asc(domain.FieldID))
	return query
}
