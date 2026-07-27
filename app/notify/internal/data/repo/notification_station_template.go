package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationstationtemplate"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.NotificationStationTemplateRepo = (*NotificationStationTemplateRepo)(nil)

type NotificationStationTemplateRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationStationTemplateRepo(
	db *gen.Client,
) bizrepo.NotificationStationTemplateRepo {
	return &NotificationStationTemplateRepo{
		db: db,
	}
}

func (r *NotificationStationTemplateRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationStationTemplateRepo) Upsert(ctx context.Context, template *model.NotificationStationTemplate) (*model.NotificationStationTemplate, error) {
	err := r.getClient(ctx).NotificationStationTemplate.Create().
		SetRuleID(template.RuleID).
		SetTitleTemplate(template.TitleTemplate).
		SetContentTemplate(template.ContentTemplate).
		OnConflict(entsql.ConflictColumns(notificationstationtemplate.FieldRuleID)).
		UpdateTitleTemplate().
		UpdateContentTemplate().
		UpdateUpdatedAt().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, &bizrepo.NotificationStationTemplateQuery{
		RuleID: &template.RuleID,
	})
}

func (r *NotificationStationTemplateRepo) BulkUpsert(ctx context.Context, templates []*model.NotificationStationTemplate) error {
	if len(templates) == 0 {
		return nil
	}
	creates := make([]*gen.NotificationStationTemplateCreate, 0, len(templates))
	for _, template := range templates {
		if template == nil {
			continue
		}
		creates = append(creates, r.getClient(ctx).NotificationStationTemplate.Create().
			SetRuleID(template.RuleID).
			SetTitleTemplate(template.TitleTemplate).
			SetContentTemplate(template.ContentTemplate))
	}
	if len(creates) == 0 {
		return nil
	}
	return r.getClient(ctx).NotificationStationTemplate.CreateBulk(creates...).
		OnConflict(entsql.ConflictColumns(notificationstationtemplate.FieldRuleID)).
		UpdateTitleTemplate().
		UpdateContentTemplate().
		UpdateUpdatedAt().
		Exec(ctx)
}

func (r *NotificationStationTemplateRepo) Get(ctx context.Context, req *bizrepo.NotificationStationTemplateQuery) (*model.NotificationStationTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationStationTemplateRepo) List(ctx context.Context, req *bizrepo.NotificationStationTemplateQuery) ([]*model.NotificationStationTemplate, error) {
	query := r.getClient(ctx).NotificationStationTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationStationTemplate, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationStationTemplate{
			ID:              item.ID,
			RuleID:          item.RuleID,
			TitleTemplate:   item.TitleTemplate,
			ContentTemplate: item.ContentTemplate,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationStationTemplateRepo) Map(ctx context.Context, req *bizrepo.NotificationStationTemplateQuery) (map[int64]*model.NotificationStationTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationStationTemplate, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationStationTemplateRepo) Count(ctx context.Context, req *bizrepo.NotificationStationTemplateQuery) (int, error) {
	query := r.getClient(ctx).NotificationStationTemplate.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *NotificationStationTemplateRepo) Page(ctx context.Context, req *bizrepo.NotificationStationTemplateQuery) (*bizrepo.NotificationStationTemplatePageResp, error) {
	var pageReq *base.PageRequest
	if req != nil {
		pageReq = req.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationStationTemplate.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationStationTemplate, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationStationTemplate{
			ID:              item.ID,
			RuleID:          item.RuleID,
			TitleTemplate:   item.TitleTemplate,
			ContentTemplate: item.ContentTemplate,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}
	return &bizrepo.NotificationStationTemplatePageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationStationTemplateRepo) getQuery(query *gen.NotificationStationTemplateQuery, req *bizrepo.NotificationStationTemplateQuery) *gen.NotificationStationTemplateQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationstationtemplate.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationstationtemplate.IDIn(req.IDs...))
	}
	if req.RuleID != nil {
		query = query.Where(notificationstationtemplate.RuleIDEQ(*req.RuleID))
	}
	if len(req.RuleIDs) > 0 {
		query = query.Where(notificationstationtemplate.RuleIDIn(req.RuleIDs...))
	}
	return query
}
