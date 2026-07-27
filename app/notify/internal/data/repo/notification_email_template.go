package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationemailtemplate"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.NotificationEmailTemplateRepo = (*NotificationEmailTemplateRepo)(nil)

type NotificationEmailTemplateRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationEmailTemplateRepo(
	db *gen.Client,
) bizrepo.NotificationEmailTemplateRepo {
	return &NotificationEmailTemplateRepo{
		db: db,
	}
}

func (r *NotificationEmailTemplateRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationEmailTemplateRepo) Upsert(ctx context.Context, template *model.NotificationEmailTemplate) (*model.NotificationEmailTemplate, error) {
	err := r.getClient(ctx).NotificationEmailTemplate.Create().
		SetRuleID(template.RuleID).
		SetSubjectTemplate(template.SubjectTemplate).
		SetBodyTemplate(template.BodyTemplate).
		SetContentType(template.ContentType).
		OnConflict(entsql.ConflictColumns(notificationemailtemplate.FieldRuleID)).
		UpdateSubjectTemplate().
		UpdateBodyTemplate().
		UpdateContentType().
		UpdateUpdatedAt().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, &bizrepo.NotificationEmailTemplateQuery{
		RuleID: &template.RuleID,
	})
}

func (r *NotificationEmailTemplateRepo) BulkUpsert(ctx context.Context, templates []*model.NotificationEmailTemplate) error {
	if len(templates) == 0 {
		return nil
	}
	creates := make([]*gen.NotificationEmailTemplateCreate, 0, len(templates))
	for _, template := range templates {
		if template == nil {
			continue
		}
		creates = append(creates, r.getClient(ctx).NotificationEmailTemplate.Create().
			SetRuleID(template.RuleID).
			SetSubjectTemplate(template.SubjectTemplate).
			SetBodyTemplate(template.BodyTemplate).
			SetContentType(template.ContentType))
	}
	if len(creates) == 0 {
		return nil
	}
	return r.getClient(ctx).NotificationEmailTemplate.CreateBulk(creates...).
		OnConflict(entsql.ConflictColumns(notificationemailtemplate.FieldRuleID)).
		UpdateSubjectTemplate().
		UpdateBodyTemplate().
		UpdateContentType().
		UpdateUpdatedAt().
		Exec(ctx)
}

func (r *NotificationEmailTemplateRepo) Get(ctx context.Context, req *bizrepo.NotificationEmailTemplateQuery) (*model.NotificationEmailTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationEmailTemplateRepo) List(ctx context.Context, req *bizrepo.NotificationEmailTemplateQuery) ([]*model.NotificationEmailTemplate, error) {
	query := r.getClient(ctx).NotificationEmailTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationEmailTemplate, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationEmailTemplate{
			ID:              item.ID,
			RuleID:          item.RuleID,
			SubjectTemplate: item.SubjectTemplate,
			BodyTemplate:    item.BodyTemplate,
			ContentType:     item.ContentType,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationEmailTemplateRepo) Map(ctx context.Context, req *bizrepo.NotificationEmailTemplateQuery) (map[int64]*model.NotificationEmailTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationEmailTemplate, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationEmailTemplateRepo) Count(ctx context.Context, req *bizrepo.NotificationEmailTemplateQuery) (int, error) {
	query := r.getClient(ctx).NotificationEmailTemplate.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *NotificationEmailTemplateRepo) Page(ctx context.Context, req *bizrepo.NotificationEmailTemplateQuery) (*bizrepo.NotificationEmailTemplatePageResp, error) {
	var pageReq *base.PageRequest
	if req != nil {
		pageReq = req.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationEmailTemplate.Query()
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
	result := make([]*model.NotificationEmailTemplate, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationEmailTemplate{
			ID:              item.ID,
			RuleID:          item.RuleID,
			SubjectTemplate: item.SubjectTemplate,
			BodyTemplate:    item.BodyTemplate,
			ContentType:     item.ContentType,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		})
	}
	return &bizrepo.NotificationEmailTemplatePageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationEmailTemplateRepo) getQuery(query *gen.NotificationEmailTemplateQuery, req *bizrepo.NotificationEmailTemplateQuery) *gen.NotificationEmailTemplateQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationemailtemplate.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationemailtemplate.IDIn(req.IDs...))
	}
	if req.RuleID != nil {
		query = query.Where(notificationemailtemplate.RuleIDEQ(*req.RuleID))
	}
	if len(req.RuleIDs) > 0 {
		query = query.Where(notificationemailtemplate.RuleIDIn(req.RuleIDs...))
	}
	return query
}
