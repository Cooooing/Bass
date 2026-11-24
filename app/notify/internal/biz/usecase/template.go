package usecase

import (
	"bytes"
	utilent "common/pkg/util/ent"
	event "common/proto/gen/common/enums/event/v1"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sort"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"
)

type TemplateUsecase struct {
	log                                 *slog.Logger
	tx                                  utilent.TxFunc
	notificationRuleRepo                repo.NotificationRuleRepo
	notificationStationTemplateRepo     repo.NotificationStationTemplateRepo
	notificationEmailTemplateRepo       repo.NotificationEmailTemplateRepo
	notificationTencentSMSTemplateRepo  repo.NotificationTencentSMSTemplateRepo
	notificationLarkWebhookTemplateRepo repo.NotificationLarkWebhookTemplateRepo
}

func NewTemplateUsecase(
	logger *slog.Logger,
	tx utilent.TxFunc,
	notificationRuleRepo repo.NotificationRuleRepo,
	notificationStationTemplateRepo repo.NotificationStationTemplateRepo,
	notificationEmailTemplateRepo repo.NotificationEmailTemplateRepo,
	notificationTencentSMSTemplateRepo repo.NotificationTencentSMSTemplateRepo,
	notificationLarkWebhookTemplateRepo repo.NotificationLarkWebhookTemplateRepo,
) *TemplateUsecase {
	return &TemplateUsecase{
		log:                                 logger,
		tx:                                  tx,
		notificationRuleRepo:                notificationRuleRepo,
		notificationStationTemplateRepo:     notificationStationTemplateRepo,
		notificationEmailTemplateRepo:       notificationEmailTemplateRepo,
		notificationTencentSMSTemplateRepo:  notificationTencentSMSTemplateRepo,
		notificationLarkWebhookTemplateRepo: notificationLarkWebhookTemplateRepo,
	}
}

type TemplatePreviewReq struct {
	Template         string
	TemplateDataJSON string
}

type TemplatePreviewResp struct {
	Rendered string
}

type ListEnabledRulesWithTemplatesReq struct {
	NotificationType notifyenum.NotificationType
	Language         notifyenum.Language
	Channels         []notifyenum.NotificationChannel
}

func (u *TemplateUsecase) Preview(ctx context.Context, req *TemplatePreviewReq) (*TemplatePreviewResp, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if req == nil {
		return nil, errors.New("template preview request is required")
	}
	decoder := json.NewDecoder(bytes.NewBufferString(req.TemplateDataJSON))
	decoder.UseNumber()
	var data map[string]any
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("template_data_json must be a valid json object: %w", err)
	}
	if data == nil {
		return nil, errors.New("template_data_json must be a json object")
	}
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			err = errors.New("unexpected trailing json value")
		}
		return nil, fmt.Errorf("template_data_json must contain one json object: %w", err)
	}
	rendered, err := (&model.NotificationTemplatePreview{
		Template: req.Template,
	}).Render(ctx, data)
	if err != nil {
		return nil, err
	}
	return &TemplatePreviewResp{
		Rendered: rendered,
	}, nil
}

func (u *TemplateUsecase) ListEnabledRulesWithTemplates(ctx context.Context, req *ListEnabledRulesWithTemplatesReq) ([]*model.NotificationRule, error) {
	if req == nil {
		return nil, nil
	}
	rules, err := u.notificationRuleRepo.List(ctx, &repo.NotificationRuleQuery{
		NotificationType: &req.NotificationType,
		Language:         &req.Language,
		Channels:         req.Channels,
		Enabled:          new(true),
	})
	if err != nil || len(rules) == 0 {
		return rules, err
	}
	rulesByID, ruleIDsByChannel := u.indexNotificationRules(rules)
	if err := u.loadStationTemplates(ctx, ruleIDsByChannel[notifyenum.NotificationChannelStation], rulesByID); err != nil {
		return nil, err
	}
	if err := u.loadEmailTemplates(ctx, ruleIDsByChannel[notifyenum.NotificationChannelEmail], rulesByID); err != nil {
		return nil, err
	}
	if err := u.loadTencentSMSTemplates(ctx, ruleIDsByChannel[notifyenum.NotificationChannelTencentSMS], rulesByID); err != nil {
		return nil, err
	}
	if err := u.loadLarkWebhookTemplates(ctx, ruleIDsByChannel[notifyenum.NotificationChannelLarkWebhook], rulesByID); err != nil {
		return nil, err
	}
	return rules, nil
}

func (*TemplateUsecase) indexNotificationRules(rules []*model.NotificationRule) (map[int64]*model.NotificationRule, map[notifyenum.NotificationChannel][]int64) {
	rulesByID := make(map[int64]*model.NotificationRule, len(rules))
	ruleIDsByChannel := make(map[notifyenum.NotificationChannel][]int64)
	for _, rule := range rules {
		if rule == nil || rule.ID == 0 {
			continue
		}
		rulesByID[rule.ID] = rule
		ruleIDsByChannel[rule.Channel] = append(ruleIDsByChannel[rule.Channel], rule.ID)
	}
	return rulesByID, ruleIDsByChannel
}

func (u *TemplateUsecase) loadStationTemplates(ctx context.Context, ruleIDs []int64, rulesByID map[int64]*model.NotificationRule) error {
	if len(ruleIDs) == 0 {
		return nil
	}
	templates, err := u.notificationStationTemplateRepo.List(ctx, &repo.NotificationStationTemplateQuery{RuleIDs: ruleIDs})
	if err != nil {
		return err
	}
	for _, item := range templates {
		if item != nil && rulesByID[item.RuleID] != nil {
			rulesByID[item.RuleID].StationTemplate = item
		}
	}
	return nil
}

func (u *TemplateUsecase) loadEmailTemplates(ctx context.Context, ruleIDs []int64, rulesByID map[int64]*model.NotificationRule) error {
	if len(ruleIDs) == 0 {
		return nil
	}
	templates, err := u.notificationEmailTemplateRepo.List(ctx, &repo.NotificationEmailTemplateQuery{RuleIDs: ruleIDs})
	if err != nil {
		return err
	}
	for _, item := range templates {
		if item != nil && rulesByID[item.RuleID] != nil {
			rulesByID[item.RuleID].EmailTemplate = item
		}
	}
	return nil
}

func (u *TemplateUsecase) loadTencentSMSTemplates(ctx context.Context, ruleIDs []int64, rulesByID map[int64]*model.NotificationRule) error {
	if len(ruleIDs) == 0 {
		return nil
	}
	templates, err := u.notificationTencentSMSTemplateRepo.List(ctx, &repo.NotificationTencentSMSTemplateQuery{RuleIDs: ruleIDs})
	if err != nil {
		return err
	}
	for _, item := range templates {
		if item != nil && rulesByID[item.RuleID] != nil {
			rulesByID[item.RuleID].TencentSMSTemplate = item
		}
	}
	return nil
}

func (u *TemplateUsecase) loadLarkWebhookTemplates(ctx context.Context, ruleIDs []int64, rulesByID map[int64]*model.NotificationRule) error {
	if len(ruleIDs) == 0 {
		return nil
	}
	templates, err := u.notificationLarkWebhookTemplateRepo.List(ctx, &repo.NotificationLarkWebhookTemplateQuery{RuleIDs: ruleIDs})
	if err != nil {
		return err
	}
	for _, item := range templates {
		if item != nil && rulesByID[item.RuleID] != nil {
			rulesByID[item.RuleID].LarkWebhookTemplate = item
		}
	}
	return nil
}

func (u *TemplateUsecase) InitDefaultTemplates(ctx context.Context, eventHandlers map[event.EventSubject][]EventHandlerBinding) error {
	if u == nil || u.notificationRuleRepo == nil {
		return nil
	}
	definitions := u.collectDefaultTemplateDefinitions(eventHandlers)
	if len(definitions.templates) == 0 {
		return nil
	}
	sort.Strings(definitions.keys)
	if err := u.saveDefaultTemplateDefinitions(ctx, definitions); err != nil {
		return err
	}
	if u.log != nil {
		u.log.Info("notification templates initialized", slog.Int("count", len(definitions.templates)))
	}
	return nil
}

type defaultTemplateDefinitions struct {
	templates         map[string]*model.NotificationTemplateDefinition
	keys              []string
	rules             []*model.NotificationRule
	notificationTypes []notifyenum.NotificationType
	channels          []notifyenum.NotificationChannel
	languages         []notifyenum.Language
}

func (*TemplateUsecase) collectDefaultTemplateDefinitions(eventHandlers map[event.EventSubject][]EventHandlerBinding) *defaultTemplateDefinitions {
	definitions := &defaultTemplateDefinitions{templates: make(map[string]*model.NotificationTemplateDefinition)}
	for _, bindings := range eventHandlers {
		for _, binding := range bindings {
			if binding.Handler == nil {
				continue
			}
			for _, template := range binding.Handler.Templates() {
				definitions.add(template)
			}
		}
	}
	return definitions
}

func (d *defaultTemplateDefinitions) add(template *model.NotificationTemplateDefinition) {
	if template == nil {
		return
	}
	key := fmt.Sprintf("%s:%s:%s", template.NotificationType, template.Channel, template.Language)
	if _, exists := d.templates[key]; exists {
		return
	}
	d.templates[key] = template
	d.keys = append(d.keys, key)
	d.rules = append(d.rules, &model.NotificationRule{
		NotificationType: template.NotificationType,
		Channel:          template.Channel,
		Language:         template.Language,
		Enabled:          template.Enabled,
	})
	d.notificationTypes = append(d.notificationTypes, template.NotificationType)
	d.channels = append(d.channels, template.Channel)
	d.languages = append(d.languages, template.Language)
}

func (u *TemplateUsecase) saveDefaultTemplateDefinitions(ctx context.Context, definitions *defaultTemplateDefinitions) error {
	return u.tx(ctx, func(txCtx context.Context) error {
		if err := u.notificationRuleRepo.BulkUpsert(txCtx, definitions.rules); err != nil {
			return err
		}
		rules, err := u.notificationRuleRepo.List(txCtx, &repo.NotificationRuleQuery{
			NotificationTypes: definitions.notificationTypes,
			Channels:          definitions.channels,
			Languages:         definitions.languages,
		})
		if err != nil {
			return err
		}
		batches, err := definitions.buildTemplateBatches(rules)
		if err != nil {
			return err
		}
		return u.bulkUpsertTemplateBatches(txCtx, batches)
	})
}

type notificationTemplateBatches struct {
	station     []*model.NotificationStationTemplate
	email       []*model.NotificationEmailTemplate
	tencentSMS  []*model.NotificationTencentSMSTemplate
	larkWebhook []*model.NotificationLarkWebhookTemplate
}

func (d *defaultTemplateDefinitions) buildTemplateBatches(rules []*model.NotificationRule) (*notificationTemplateBatches, error) {
	rulesByKey := make(map[string]*model.NotificationRule, len(rules))
	for _, rule := range rules {
		if rule != nil {
			rulesByKey[fmt.Sprintf("%s:%s:%s", rule.NotificationType, rule.Channel, rule.Language)] = rule
		}
	}
	batches := &notificationTemplateBatches{}
	for _, key := range d.keys {
		if err := batches.add(d.templates[key], rulesByKey[key]); err != nil {
			return nil, err
		}
	}
	return batches, nil
}

func (b *notificationTemplateBatches) add(template *model.NotificationTemplateDefinition, rule *model.NotificationRule) error {
	if rule == nil || rule.ID == 0 {
		return fmt.Errorf("notification rule not saved: notification_type=%s channel=%s language=%s", template.NotificationType, template.Channel, template.Language)
	}
	switch template.Channel {
	case notifyenum.NotificationChannelStation:
		if template.StationTemplate == nil {
			return fmt.Errorf("station template is required: rule_id=%d", rule.ID)
		}
		b.station = append(b.station, &model.NotificationStationTemplate{RuleID: rule.ID, TitleTemplate: template.StationTemplate.TitleTemplate, ContentTemplate: template.StationTemplate.ContentTemplate})
	case notifyenum.NotificationChannelEmail:
		if template.EmailTemplate == nil {
			return fmt.Errorf("email template is required: rule_id=%d", rule.ID)
		}
		b.email = append(b.email, &model.NotificationEmailTemplate{RuleID: rule.ID, SubjectTemplate: template.EmailTemplate.SubjectTemplate, BodyTemplate: template.EmailTemplate.BodyTemplate, ContentType: template.EmailTemplate.ContentType})
	case notifyenum.NotificationChannelTencentSMS:
		if template.TencentSMSTemplate == nil {
			return fmt.Errorf("tencent sms template is required: rule_id=%d", rule.ID)
		}
		b.tencentSMS = append(b.tencentSMS, &model.NotificationTencentSMSTemplate{RuleID: rule.ID, SMSSDKAppID: template.TencentSMSTemplate.SMSSDKAppID, SignName: template.TencentSMSTemplate.SignName, ProviderTemplateID: template.TencentSMSTemplate.ProviderTemplateID, ParamTemplates: template.TencentSMSTemplate.ParamTemplates})
	case notifyenum.NotificationChannelLarkWebhook:
		if template.LarkWebhookTemplate == nil {
			return fmt.Errorf("lark webhook template is required: rule_id=%d", rule.ID)
		}
		b.larkWebhook = append(b.larkWebhook, &model.NotificationLarkWebhookTemplate{RuleID: rule.ID, WebhookID: template.LarkWebhookTemplate.WebhookID, Token: template.LarkWebhookTemplate.Token, Secret: template.LarkWebhookTemplate.Secret, MsgType: template.LarkWebhookTemplate.MsgType, ContentTemplate: template.LarkWebhookTemplate.ContentTemplate})
	default:
		return fmt.Errorf("unsupported notification channel: %s", template.Channel)
	}
	return nil
}

func (u *TemplateUsecase) bulkUpsertTemplateBatches(ctx context.Context, batches *notificationTemplateBatches) error {
	if err := u.notificationStationTemplateRepo.BulkUpsert(ctx, batches.station); err != nil {
		return err
	}
	if err := u.notificationEmailTemplateRepo.BulkUpsert(ctx, batches.email); err != nil {
		return err
	}
	if err := u.notificationTencentSMSTemplateRepo.BulkUpsert(ctx, batches.tencentSMS); err != nil {
		return err
	}
	return u.notificationLarkWebhookTemplateRepo.BulkUpsert(ctx, batches.larkWebhook)
}
