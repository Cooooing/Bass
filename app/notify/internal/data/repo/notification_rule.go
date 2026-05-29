package repo

import (
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationrule"
	notifyenum "notify/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.NotificationRuleRepo = (*NotificationRuleRepo)(nil)

type NotificationRuleRepo struct {
	db *gen.Client
}

func NewNotificationRuleRepo(db *gen.Client) bizrepo.NotificationRuleRepo {
	return &NotificationRuleRepo{db: db}
}

func (r *NotificationRuleRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *NotificationRuleRepo) ListEnabled(ctx context.Context, eventType commonenum.EventType, language notifyenum.Language) ([]*model.NotificationRule, error) {
	list, err := r.getClient(ctx).NotificationRule.Query().
		Where(
			notificationrule.EventTypeEQ(notificationrule.EventType(eventType)),
			notificationrule.LanguageEQ(notificationrule.Language(language)),
			notificationrule.EnabledEQ(true),
		).
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
			rule.LarkWebhookTemplate = &model.NotificationLarkWebhookTemplate{
				ID:           item.Edges.LarkWebhookTemplate.ID,
				RuleID:       item.Edges.LarkWebhookTemplate.RuleID,
				WebhookID:    item.Edges.LarkWebhookTemplate.WebhookID,
				Token:        item.Edges.LarkWebhookTemplate.Token,
				BodyTemplate: item.Edges.LarkWebhookTemplate.BodyTemplate,
				CreatedAt:    item.Edges.LarkWebhookTemplate.CreatedAt,
				UpdatedAt:    item.Edges.LarkWebhookTemplate.UpdatedAt,
			}
		}
		rules = append(rules, rule)
	}
	return rules, nil
}
