package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationrule"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationRuleRepo = (*NotificationRuleRepo)(nil)

type NotificationRuleRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewNotificationRuleRepo(
	db *gen.Client,
) bizrepo.NotificationRuleRepo {
	return &NotificationRuleRepo{
		db: db,
	}
}

func (r *NotificationRuleRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationRuleRepo) Get(ctx context.Context, req *bizrepo.NotificationRuleQuery) (*model.NotificationRule, error) {
	list, err := r.List(ctx, req)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return list[0], nil
}

func (r *NotificationRuleRepo) List(ctx context.Context, req *bizrepo.NotificationRuleQuery) ([]*model.NotificationRule, error) {
	query := r.getClient(ctx).NotificationRule.Query()
	query = r.getQuery(query, req)
	list, err := query.
		WithStationTemplate().
		WithEmailTemplate().
		WithTencentSmsTemplate().
		WithLarkWebhookTemplate().
		All(ctx)
	if err != nil {
		return nil, err
	}

	rules := make([]*model.NotificationRule, 0, len(list))
	for _, item := range list {
		rules = append(rules, r.notificationRule(item))
	}
	return rules, nil
}

func (r *NotificationRuleRepo) Map(ctx context.Context, req *bizrepo.NotificationRuleQuery) (map[int64]*model.
	NotificationRule, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.NotificationRule, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *NotificationRuleRepo) Count(ctx context.Context, req *bizrepo.NotificationRuleQuery) (int, error) {
	query := r.getClient(ctx).NotificationRule.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *NotificationRuleRepo) Page(ctx context.Context, req *bizrepo.NotificationRuleQuery) (*bizrepo.NotificationRulePageResp, error) {
	queryReq := req
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).NotificationRule.Query()
	query = r.getQuery(query, queryReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.
		WithStationTemplate().
		WithEmailTemplate().
		WithTencentSmsTemplate().
		WithLarkWebhookTemplate().
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	rules := make([]*model.NotificationRule, 0, len(list))
	for _, item := range list {
		rules = append(rules, r.notificationRule(item))
	}
	return &bizrepo.NotificationRulePageResp{
		Rows: rules,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *NotificationRuleRepo) getQuery(query *gen.NotificationRuleQuery, req *bizrepo.NotificationRuleQuery) *gen.NotificationRuleQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(notificationrule.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(notificationrule.IDIn(req.IDs...))
	}
	if req.EventType != nil {
		query = query.Where(notificationrule.EventTypeEQ(notificationrule.EventType(*req.EventType)))
	}
	if req.Channel != nil {
		query = query.Where(notificationrule.ChannelEQ(notificationrule.Channel(*req.Channel)))
	}
	if req.Language != nil {
		query = query.Where(notificationrule.LanguageEQ(notificationrule.Language(*req.Language)))
	}
	if req.Enabled != nil {
		query = query.Where(notificationrule.EnabledEQ(*req.Enabled))
	}
	return query
}

func (r *NotificationRuleRepo) notificationRule(item *gen.NotificationRule) *model.NotificationRule {
	rule := &model.NotificationRule{
		ID:        item.ID,
		EventType: commonenum.EventType(item.EventType),
		Channel:   notifyenum.NotificationChannel(item.Channel),
		Language:  notifyenum.Language(item.Language),
		Enabled:   item.Enabled,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
	if item.Edges.StationTemplate != nil {
		rule.StationTemplate = &model.NotificationStationTemplate{
			ID:              item.Edges.StationTemplate.ID,
			RuleID:          item.Edges.StationTemplate.RuleID,
			TitleTemplate:   item.Edges.StationTemplate.TitleTemplate,
			ContentTemplate: item.Edges.StationTemplate.ContentTemplate,
			CreatedAt:       item.Edges.StationTemplate.CreatedAt,
			UpdatedAt:       item.Edges.StationTemplate.UpdatedAt,
		}
	}
	if item.Edges.EmailTemplate != nil {
		rule.EmailTemplate = &model.NotificationEmailTemplate{
			ID:              item.Edges.EmailTemplate.ID,
			RuleID:          item.Edges.EmailTemplate.RuleID,
			SubjectTemplate: item.Edges.EmailTemplate.SubjectTemplate,
			BodyTemplate:    item.Edges.EmailTemplate.BodyTemplate,
			ContentType:     item.Edges.EmailTemplate.ContentType,
			CreatedAt:       item.Edges.EmailTemplate.CreatedAt,
			UpdatedAt:       item.Edges.EmailTemplate.UpdatedAt,
		}
	}
	if item.Edges.TencentSmsTemplate != nil {
		rule.TencentSMSTemplate = &model.NotificationTencentSMSTemplate{
			ID:                 item.Edges.TencentSmsTemplate.ID,
			RuleID:             item.Edges.TencentSmsTemplate.RuleID,
			SMSSDKAppID:        item.Edges.TencentSmsTemplate.SmsSdkAppID,
			SignName:           item.Edges.TencentSmsTemplate.SignName,
			ProviderTemplateID: item.Edges.TencentSmsTemplate.ProviderTemplateID,
			ParamTemplates:     item.Edges.TencentSmsTemplate.ParamTemplates,
			CreatedAt:          item.Edges.TencentSmsTemplate.CreatedAt,
			UpdatedAt:          item.Edges.TencentSmsTemplate.UpdatedAt,
		}
	}
	if item.Edges.LarkWebhookTemplate != nil {
		secret := ""
		if item.Edges.LarkWebhookTemplate.Secret != nil {
			secret = *item.Edges.LarkWebhookTemplate.Secret
		}
		rule.LarkWebhookTemplate = &model.NotificationLarkWebhookTemplate{
			ID:              item.Edges.LarkWebhookTemplate.ID,
			RuleID:          item.Edges.LarkWebhookTemplate.RuleID,
			WebhookID:       item.Edges.LarkWebhookTemplate.WebhookID,
			Token:           item.Edges.LarkWebhookTemplate.Token,
			Secret:          secret,
			MsgType:         item.Edges.LarkWebhookTemplate.MsgType,
			ContentTemplate: item.Edges.LarkWebhookTemplate.ContentTemplate,
			CreatedAt:       item.Edges.LarkWebhookTemplate.CreatedAt,
			UpdatedAt:       item.Edges.LarkWebhookTemplate.UpdatedAt,
		}
	}
	return rule
}
