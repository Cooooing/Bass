package repo

import (
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationtencentsmstemplate"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.NotificationTencentSMSTemplateRepo = (*NotificationTencentSMSTemplateRepo)(nil)

type NotificationTencentSMSTemplateRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationTencentSMSTemplateRepo(
	db *gen.Client,
) bizrepo.NotificationTencentSMSTemplateRepo {
	return &NotificationTencentSMSTemplateRepo{
		db: db,
	}
}

func (r *NotificationTencentSMSTemplateRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationTencentSMSTemplateRepo) Upsert(ctx context.Context, template *model.NotificationTencentSMSTemplate) (*model.NotificationTencentSMSTemplate, error) {
	err := r.getClient(ctx).NotificationTencentSMSTemplate.Create().
		SetRuleID(template.RuleID).
		SetSmsSdkAppID(template.SMSSDKAppID).
		SetSignName(template.SignName).
		SetProviderTemplateID(template.ProviderTemplateID).
		SetParamTemplates(template.ParamTemplates).
		OnConflict(entsql.ConflictColumns(notificationtencentsmstemplate.FieldRuleID)).
		UpdateSmsSdkAppID().
		UpdateSignName().
		UpdateProviderTemplateID().
		UpdateParamTemplates().
		UpdateUpdatedAt().
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return r.Get(ctx, &bizrepo.NotificationTencentSMSTemplateQuery{
		RuleID: &template.RuleID,
	})
}

func (r *NotificationTencentSMSTemplateRepo) BulkUpsert(ctx context.Context, templates []*model.NotificationTencentSMSTemplate) error {
	if len(templates) == 0 {
		return nil
	}
	creates := make([]*gen.NotificationTencentSMSTemplateCreate, 0, len(templates))
	for _, template := range templates {
		if template == nil {
			continue
		}
		creates = append(creates, r.getClient(ctx).NotificationTencentSMSTemplate.Create().
			SetRuleID(template.RuleID).
			SetSmsSdkAppID(template.SMSSDKAppID).
			SetSignName(template.SignName).
			SetProviderTemplateID(template.ProviderTemplateID).
			SetParamTemplates(template.ParamTemplates))
	}
	if len(creates) == 0 {
		return nil
	}
	return r.getClient(ctx).NotificationTencentSMSTemplate.CreateBulk(creates...).
		OnConflict(entsql.ConflictColumns(notificationtencentsmstemplate.FieldRuleID)).
		UpdateSmsSdkAppID().
		UpdateSignName().
		UpdateProviderTemplateID().
		UpdateParamTemplates().
		UpdateUpdatedAt().
		Exec(ctx)
}

func (r *NotificationTencentSMSTemplateRepo) Get(ctx context.Context, req *bizrepo.NotificationTencentSMSTemplateQuery) (*model.NotificationTencentSMSTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationTencentSMSTemplateRepo) List(ctx context.Context, req *bizrepo.NotificationTencentSMSTemplateQuery) ([]*model.NotificationTencentSMSTemplate, error) {
	query := r.getClient(ctx).NotificationTencentSMSTemplate.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.NotificationTencentSMSTemplate, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationTencentSMSTemplate{
			ID:                 item.ID,
			RuleID:             item.RuleID,
			SMSSDKAppID:        item.SmsSdkAppID,
			SignName:           item.SignName,
			ProviderTemplateID: item.ProviderTemplateID,
			ParamTemplates:     item.ParamTemplates,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		})
	}
	return result, nil
}

func (r *NotificationTencentSMSTemplateRepo) Map(ctx context.Context, req *bizrepo.NotificationTencentSMSTemplateQuery) (map[int64]*model.NotificationTencentSMSTemplate, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationTencentSMSTemplate, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationTencentSMSTemplateRepo) Count(ctx context.Context, req *bizrepo.NotificationTencentSMSTemplateQuery) (int, error) {
	query := r.getClient(ctx).NotificationTencentSMSTemplate.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *NotificationTencentSMSTemplateRepo) Page(ctx context.Context, req *bizrepo.NotificationTencentSMSTemplateQuery) (*bizrepo.NotificationTencentSMSTemplatePageResp, error) {
	var pageReq *base.PageRequest
	if req != nil {
		pageReq = req.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationTencentSMSTemplate.Query()
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
	result := make([]*model.NotificationTencentSMSTemplate, 0, len(list))
	for _, item := range list {
		result = append(result, &model.NotificationTencentSMSTemplate{
			ID:                 item.ID,
			RuleID:             item.RuleID,
			SMSSDKAppID:        item.SmsSdkAppID,
			SignName:           item.SignName,
			ProviderTemplateID: item.ProviderTemplateID,
			ParamTemplates:     item.ParamTemplates,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		})
	}
	return &bizrepo.NotificationTencentSMSTemplatePageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationTencentSMSTemplateRepo) getQuery(query *gen.NotificationTencentSMSTemplateQuery, req *bizrepo.NotificationTencentSMSTemplateQuery) *gen.NotificationTencentSMSTemplateQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationtencentsmstemplate.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationtencentsmstemplate.IDIn(req.IDs...))
	}
	if req.RuleID != nil {
		query = query.Where(notificationtencentsmstemplate.RuleIDEQ(*req.RuleID))
	}
	if len(req.RuleIDs) > 0 {
		query = query.Where(notificationtencentsmstemplate.RuleIDIn(req.RuleIDs...))
	}
	return query
}
