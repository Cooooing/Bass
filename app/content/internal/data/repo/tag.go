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
	tagent "content/internal/data/gen/tag"
	"content/internal/enum"

	"github.com/samber/lo"
)

var _ repo.TagRepo = (*TagRepo)(nil)

type TagRepo struct {
	db *gen.Client
}

func NewTagRepo(db *gen.Client) repo.TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *TagRepo) Save(ctx context.Context, req *repo.TagSaveReq) (*repo.TagSaveResponse, error) {
	tag := req.Tag
	save, err := r.getClient(ctx).Tag.Create().
		SetName(tag.Name).
		SetNillableDescription(tag.Description).
		SetNillableDomainID(tag.DomainID).
		SetStatus(tagent.Status(tag.Status)).
		SetNillableCreatedBy(tag.CreatedBy).
		SetNillableUpdatedBy(tag.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.TagSaveResponse{Tag: &model.Tag{
		ID:          save.ID,
		Name:        save.Name,
		Description: save.Description,
		DomainID:    save.DomainID,
		Status:      enum.TagStatus(save.Status),
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
		CreatedBy:   save.CreatedBy,
		UpdatedBy:   save.UpdatedBy,
	}}, nil
}

func (r *TagRepo) Saves(ctx context.Context, req *repo.TagSavesReq) (*repo.TagSavesResponse, error) {
	tags := req.Tags
	client := r.getClient(ctx)
	creates := make([]*gen.TagCreate, 0, len(tags))
	for i := range tags {
		creates = append(creates,
			client.Tag.Create().
				SetName(tags[i].Name).
				SetNillableDescription(tags[i].Description).
				SetNillableDomainID(tags[i].DomainID).
				SetStatus(tagent.Status(tags[i].Status)).
				SetNillableCreatedBy(tags[i].CreatedBy).
				SetNillableUpdatedBy(tags[i].UpdatedBy),
		)
	}
	save, err := client.Tag.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*model.Tag, 0, len(save))
	for _, item := range save {
		res = append(res, &model.Tag{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			DomainID:    item.DomainID,
			Status:      enum.TagStatus(item.Status),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
		})
	}
	return &repo.TagSavesResponse{Rows: res}, nil
}

func (r *TagRepo) Update(ctx context.Context, req *repo.TagUpdateReq) (*repo.TagUpdateResponse, error) {
	tag := req.Tag
	save, err := r.getClient(ctx).Tag.UpdateOneID(tag.ID).
		SetName(tag.Name).
		SetNillableDescription(tag.Description).
		SetNillableDomainID(tag.DomainID).
		SetStatus(tagent.Status(tag.Status)).
		SetNillableUpdatedBy(tag.UpdatedBy).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.TagUpdateResponse{Tag: &model.Tag{
		ID:          save.ID,
		Name:        save.Name,
		Description: save.Description,
		DomainID:    save.DomainID,
		Status:      enum.TagStatus(save.Status),
		CreatedAt:   save.CreatedAt,
		UpdatedAt:   save.UpdatedAt,
		CreatedBy:   save.CreatedBy,
		UpdatedBy:   save.UpdatedBy,
	}}, nil
}

func (r *TagRepo) Get(ctx context.Context, req *repo.TagGetReq) (*repo.TagGetResponse, error) {
	query := r.getClient(ctx).Tag.Query()
	query = r.getQuery(query, req)
	t, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &repo.TagGetResponse{Tag: &model.Tag{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		DomainID:    t.DomainID,
		Status:      enum.TagStatus(t.Status),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		CreatedBy:   t.CreatedBy,
		UpdatedBy:   t.UpdatedBy,
	}}, nil
}

func (r *TagRepo) List(ctx context.Context, req *repo.TagGetReq) (*repo.TagListResponse, error) {
	query := r.getClient(ctx).Tag.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	tags := make([]*model.Tag, 0, len(list))
	for _, item := range list {
		tags = append(tags, &model.Tag{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			DomainID:    item.DomainID,
			Status:      enum.TagStatus(item.Status),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
		})
	}
	return &repo.TagListResponse{Rows: tags}, nil
}

func (r *TagRepo) Map(ctx context.Context, req *repo.TagGetReq) (*repo.TagMapResponse, error) {
	listResponse, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return &repo.TagMapResponse{Rows: lo.SliceToMap(listResponse.Rows, func(item *model.Tag) (int64, *model.Tag) {
		return item.ID, item
	})}, nil
}

func (r *TagRepo) Count(ctx context.Context, req *repo.TagGetReq) (*repo.TagCountResponse, error) {
	query := r.getClient(ctx).Tag.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &repo.TagCountResponse{Count: count}, nil
}

func (r *TagRepo) Page(ctx context.Context, req *repo.TagGetReq) (*repo.TagPageResponse, error) {
	page := normalizePage(req.Page)
	query := r.getClient(ctx).Tag.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}

	tags := make([]*model.Tag, 0, len(list))
	for _, item := range list {
		tags = append(tags, &model.Tag{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			DomainID:    item.DomainID,
			Status:      enum.TagStatus(item.Status),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
		})
	}
	return &repo.TagPageResponse{
		Rows: tags,
		Page: &base.PageResponse{
			Total: int64(count),
			Size:  page.Size,
			Page:  page.Page,
		},
	}, nil
}

func (r *TagRepo) getQuery(query *gen.TagQuery, req *repo.TagGetReq) *gen.TagQuery {
	if req.TagId != nil {
		query = query.Where(tagent.IDEQ(*req.TagId))
	}
	if len(req.TagIds) > 0 {
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
		query = query.Where(tagent.StatusEQ(tagent.Status(*req.Status)))
	}
	if req.DomainId != nil {
		query = query.Where(tagent.DomainIDEQ(*req.DomainId))
	}
	return query
}
