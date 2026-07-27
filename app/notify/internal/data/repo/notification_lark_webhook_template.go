package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationlarkwebhooktemplate"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.NotificationLarkWebhookTemplateRepo = (*NotificationLarkWebhookTemplateRepo)(nil)

type NotificationLarkWebhookTemplateRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationLarkWebhookTemplateRepo(
	db *gen.Client,
) bizrepo.NotificationLarkWebhookTemplateRepo {
	return &NotificationLarkWebhookTemplateRepo{
		db: db,
	}
}

func (r *NotificationLarkWebhookTemplateRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationLarkWebhookTemplateRepo) Upsert(ctx context.Context, template *model.NotificationLarkWebhookTemplate) (*model.NotificationLarkWebhookTemplate, error) {
	secret := &template.Secret
	if template.Secret == "" {
		secret = nil
	}
	upsert := r.getClient(ctx).NotificationLarkWebhookTemplate.Create().
		SetRuleID(template.RuleID).
		SetWebhookID(template.WebhookID).
		SetToken(template.Token).
		SetNillableSecret(secret).
		SetMsgType(template.MsgType).
		SetContentTemplate(template.ContentTemplate).
		OnConflict(entsql.ConflictColumns(notificationlarkwebhooktemplate.FieldRuleID)).
		UpdateWebhookID().
		UpdateToken().
		UpdateMsgType().
		UpdateContentTemplate().
		UpdateUpdatedAt()
	if template.Secret == "" {
		upsert.ClearSecret()
	} else {
		upsert.UpdateSecret()
	}
	err := upsert.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, &bizrepo.NotificationLarkWebhookTemplateQuery{
		RuleID: &template.RuleID,
	})
}

func (r *NotificationLarkWebhookTemplateRepo) BulkUpsert(ctx context.Context, templates []*model.NotificationLarkWebhookTemplate) error {
	if len(templates) == 0 {
		return nil
	}
	creates := make([]*gen.NotificationLarkWebhookTemplateCreate, 0, len(templates))
	for _, template := range templates {
		if template == nil {
			continue
		}
		secret := &template.Secret
		if template.Secret == "" {
			secret = nil
		}
		creates = append(creates, r.getClient(ctx).NotificationLarkWebhookTemplate.Create().
			SetRuleID(template.RuleID).
			SetWebhookID(template.WebhookID).
			SetToken(template.Token).
			SetNillableSecret(secret).
			SetMsgType(template.MsgType).
			SetContentTemplate(template.ContentTemplate))
	}
	if len(creates) == 0 {
		return nil
	}
	return r.getClient(ctx).NotificationLarkWebhookTemplate.CreateBulk(creates...).
		OnConflict(entsql.ConflictColumns(notificationlarkwebhooktemplate.FieldRuleID)).
		UpdateWebhookID().
		UpdateToken().
		UpdateSecret().
		UpdateMsgType().
		UpdateContentTemplate().
		UpdateUpdatedAt().
		Exec(ctx)
}

func (r *NotificationLarkWebhookTemplateRepo) Get(ctx context.Context, req *bizrepo.NotificationLarkWebhookTemplateQuery) (*model.NotificationLarkWebhookTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationLarkWebhookTemplateRepo) List(ctx context.Context, req *bizrepo.NotificationLarkWebhookTemplateQuery) ([]*model.NotificationLarkWebhookTemplate, error) {
	query := r.getClient(ctx).NotificationLarkWebhookTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationLarkWebhookTemplate, 0, len(list))
	for _, item := range list {
		template := &model.NotificationLarkWebhookTemplate{
			ID:              item.ID,
			RuleID:          item.RuleID,
			WebhookID:       item.WebhookID,
			Token:           item.Token,
			MsgType:         item.MsgType,
			ContentTemplate: item.ContentTemplate,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		}
		if item.Secret != nil {
			template.Secret = *item.Secret
		}
		result = append(result, template)
	}
	return result, nil
}

func (r *NotificationLarkWebhookTemplateRepo) Map(ctx context.Context, req *bizrepo.NotificationLarkWebhookTemplateQuery) (map[int64]*model.NotificationLarkWebhookTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationLarkWebhookTemplate, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationLarkWebhookTemplateRepo) Count(ctx context.Context, req *bizrepo.NotificationLarkWebhookTemplateQuery) (int, error) {
	query := r.getClient(ctx).NotificationLarkWebhookTemplate.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *NotificationLarkWebhookTemplateRepo) Page(ctx context.Context, req *bizrepo.NotificationLarkWebhookTemplateQuery) (*bizrepo.NotificationLarkWebhookTemplatePageResp, error) {
	var pageReq *base.PageRequest
	if req != nil {
		pageReq = req.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationLarkWebhookTemplate.Query()
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
	result := make([]*model.NotificationLarkWebhookTemplate, 0, len(list))
	for _, item := range list {
		template := &model.NotificationLarkWebhookTemplate{
			ID:              item.ID,
			RuleID:          item.RuleID,
			WebhookID:       item.WebhookID,
			Token:           item.Token,
			MsgType:         item.MsgType,
			ContentTemplate: item.ContentTemplate,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
		}
		if item.Secret != nil {
			template.Secret = *item.Secret
		}
		result = append(result, template)
	}
	return &bizrepo.NotificationLarkWebhookTemplatePageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationLarkWebhookTemplateRepo) getQuery(query *gen.NotificationLarkWebhookTemplateQuery, req *bizrepo.NotificationLarkWebhookTemplateQuery) *gen.NotificationLarkWebhookTemplateQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationlarkwebhooktemplate.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationlarkwebhooktemplate.IDIn(req.IDs...))
	}
	if req.RuleID != nil {
		query = query.Where(notificationlarkwebhooktemplate.RuleIDEQ(*req.RuleID))
	}
	if len(req.RuleIDs) > 0 {
		query = query.Where(notificationlarkwebhooktemplate.RuleIDIn(req.RuleIDs...))
	}
	return query
}
