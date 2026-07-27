package usecase

import (
	"bytes"
	"common/pkg/enum"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"
	"sort"
)

type TemplateUsecase struct {
	log                                 *slog.Logger
	tx                                  base.Tx
	notificationRuleRepo                repo.NotificationRuleRepo
	notificationStationTemplateRepo     repo.NotificationStationTemplateRepo
	notificationEmailTemplateRepo       repo.NotificationEmailTemplateRepo
	notificationTencentSMSTemplateRepo  repo.NotificationTencentSMSTemplateRepo
	notificationLarkWebhookTemplateRepo repo.NotificationLarkWebhookTemplateRepo
}

func NewTemplateUsecase(
	logger *slog.Logger,
	tx base.Tx,
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
	EventType enum.EventType
	Language  notifyenum.Language
	Channels  []notifyenum.NotificationChannel
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
	if err := decoder.Decode(&extra); err != io.EOF {
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
		EventType: &req.EventType,
		Language:  &req.Language,
		Channels:  req.Channels,
		Enabled:   new(true),
	})
	if err != nil || len(rules) == 0 {
		return rules, err
	}
	rulesByID := make(map[int64]*model.NotificationRule, len(rules))
	ruleIDsByChannel := make(map[notifyenum.NotificationChannel][]int64)
	for _, rule := range rules {
		if rule == nil || rule.ID == 0 {
			continue
		}
		rulesByID[rule.ID] = rule
		ruleIDsByChannel[rule.Channel] = append(ruleIDsByChannel[rule.Channel], rule.ID)
	}
	if ruleIDs := ruleIDsByChannel[notifyenum.NotificationChannelStation]; len(ruleIDs) > 0 {
		templates, err := u.notificationStationTemplateRepo.List(ctx, &repo.NotificationStationTemplateQuery{
			RuleIDs: ruleIDs,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range templates {
			if item != nil && rulesByID[item.RuleID] != nil {
				rulesByID[item.RuleID].StationTemplate = item
			}
		}
	}
	if ruleIDs := ruleIDsByChannel[notifyenum.NotificationChannelEmail]; len(ruleIDs) > 0 {
		templates, err := u.notificationEmailTemplateRepo.List(ctx, &repo.NotificationEmailTemplateQuery{
			RuleIDs: ruleIDs,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range templates {
			if item != nil && rulesByID[item.RuleID] != nil {
				rulesByID[item.RuleID].EmailTemplate = item
			}
		}
	}
	if ruleIDs := ruleIDsByChannel[notifyenum.NotificationChannelTencentSMS]; len(ruleIDs) > 0 {
		templates, err := u.notificationTencentSMSTemplateRepo.List(ctx, &repo.NotificationTencentSMSTemplateQuery{
			RuleIDs: ruleIDs,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range templates {
			if item != nil && rulesByID[item.RuleID] != nil {
				rulesByID[item.RuleID].TencentSMSTemplate = item
			}
		}
	}
	if ruleIDs := ruleIDsByChannel[notifyenum.NotificationChannelLarkWebhook]; len(ruleIDs) > 0 {
		templates, err := u.notificationLarkWebhookTemplateRepo.List(ctx, &repo.NotificationLarkWebhookTemplateQuery{
			RuleIDs: ruleIDs,
		})
		if err != nil {
			return nil, err
		}
		for _, item := range templates {
			if item != nil && rulesByID[item.RuleID] != nil {
				rulesByID[item.RuleID].LarkWebhookTemplate = item
			}
		}
	}
	return rules, nil
}

func (u *TemplateUsecase) InitDefaultTemplates(ctx context.Context, eventHandlers map[enum.EventType]EventHandler) error {
	if u == nil || u.notificationRuleRepo == nil {
		return nil
	}
	defaultTemplates := make(map[string]*model.NotificationTemplateDefinition)
	defaultTemplateKeys := make([]string, 0)
	defaultRules := make([]*model.NotificationRule, 0)
	eventTypes := make([]enum.EventType, 0)
	channels := make([]notifyenum.NotificationChannel, 0)
	languages := make([]notifyenum.Language, 0)
	for eventType, handler := range eventHandlers {
		if handler == nil {
			continue
		}
		for _, tpl := range handler.Templates() {
			if tpl == nil || tpl.EventType != eventType {
				continue
			}
			key := fmt.Sprintf("%s:%s:%s", tpl.EventType, tpl.Channel, tpl.Language)
			if _, ok := defaultTemplates[key]; ok {
				return fmt.Errorf("duplicate notification template definition: %s", key)
			}
			defaultTemplates[key] = tpl
			defaultTemplateKeys = append(defaultTemplateKeys, key)
			defaultRules = append(defaultRules, &model.NotificationRule{
				EventType: tpl.EventType,
				Channel:   tpl.Channel,
				Language:  tpl.Language,
				Enabled:   tpl.Enabled,
			})
			eventTypes = append(eventTypes, tpl.EventType)
			channels = append(channels, tpl.Channel)
			languages = append(languages, tpl.Language)
		}
	}
	if len(defaultTemplates) == 0 {
		return nil
	}
	sort.Strings(defaultTemplateKeys)
	if err := u.tx(ctx, func(txCtx context.Context) error {
		if err := u.notificationRuleRepo.BulkUpsert(txCtx, defaultRules); err != nil {
			return err
		}
		rules, err := u.notificationRuleRepo.List(txCtx, &repo.NotificationRuleQuery{
			EventTypes: eventTypes,
			Channels:   channels,
			Languages:  languages,
		})
		if err != nil {
			return err
		}
		rulesByKey := make(map[string]*model.NotificationRule, len(rules))
		for _, rule := range rules {
			rulesByKey[fmt.Sprintf("%s:%s:%s", rule.EventType, rule.Channel, rule.Language)] = rule
		}
		stationTemplates := make([]*model.NotificationStationTemplate, 0)
		emailTemplates := make([]*model.NotificationEmailTemplate, 0)
		tencentSMSTemplates := make([]*model.NotificationTencentSMSTemplate, 0)
		larkWebhookTemplates := make([]*model.NotificationLarkWebhookTemplate, 0)
		for _, key := range defaultTemplateKeys {
			tpl := defaultTemplates[key]
			rule := rulesByKey[key]
			if rule == nil || rule.ID == 0 {
				return fmt.Errorf("notification rule not saved: event_type=%s channel=%s language=%s", tpl.EventType, tpl.Channel, tpl.Language)
			}
			switch tpl.Channel {
			case notifyenum.NotificationChannelStation:
				if tpl.StationTemplate == nil {
					return fmt.Errorf("station template is required: rule_id=%d", rule.ID)
				}
				stationTemplates = append(stationTemplates, &model.NotificationStationTemplate{
					RuleID:          rule.ID,
					TitleTemplate:   tpl.StationTemplate.TitleTemplate,
					ContentTemplate: tpl.StationTemplate.ContentTemplate,
				})
			case notifyenum.NotificationChannelEmail:
				if tpl.EmailTemplate == nil {
					return fmt.Errorf("email template is required: rule_id=%d", rule.ID)
				}
				emailTemplates = append(emailTemplates, &model.NotificationEmailTemplate{
					RuleID:          rule.ID,
					SubjectTemplate: tpl.EmailTemplate.SubjectTemplate,
					BodyTemplate:    tpl.EmailTemplate.BodyTemplate,
					ContentType:     tpl.EmailTemplate.ContentType,
				})
			case notifyenum.NotificationChannelTencentSMS:
				if tpl.TencentSMSTemplate == nil {
					return fmt.Errorf("tencent sms template is required: rule_id=%d", rule.ID)
				}
				tencentSMSTemplates = append(tencentSMSTemplates, &model.NotificationTencentSMSTemplate{
					RuleID:             rule.ID,
					SMSSDKAppID:        tpl.TencentSMSTemplate.SMSSDKAppID,
					SignName:           tpl.TencentSMSTemplate.SignName,
					ProviderTemplateID: tpl.TencentSMSTemplate.ProviderTemplateID,
					ParamTemplates:     tpl.TencentSMSTemplate.ParamTemplates,
				})
			case notifyenum.NotificationChannelLarkWebhook:
				if tpl.LarkWebhookTemplate == nil {
					return fmt.Errorf("lark webhook template is required: rule_id=%d", rule.ID)
				}
				larkWebhookTemplates = append(larkWebhookTemplates, &model.NotificationLarkWebhookTemplate{
					RuleID:          rule.ID,
					WebhookID:       tpl.LarkWebhookTemplate.WebhookID,
					Token:           tpl.LarkWebhookTemplate.Token,
					Secret:          tpl.LarkWebhookTemplate.Secret,
					MsgType:         tpl.LarkWebhookTemplate.MsgType,
					ContentTemplate: tpl.LarkWebhookTemplate.ContentTemplate,
				})
			default:
				return fmt.Errorf("unsupported notification channel: %s", tpl.Channel)
			}
		}
		if err := u.notificationStationTemplateRepo.BulkUpsert(txCtx, stationTemplates); err != nil {
			return err
		}
		if err := u.notificationEmailTemplateRepo.BulkUpsert(txCtx, emailTemplates); err != nil {
			return err
		}
		if err := u.notificationTencentSMSTemplateRepo.BulkUpsert(txCtx, tencentSMSTemplates); err != nil {
			return err
		}
		if err := u.notificationLarkWebhookTemplateRepo.BulkUpsert(txCtx, larkWebhookTemplates); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	if u.log != nil {
		u.log.Info("notification templates initialized", slog.Int("count", len(defaultTemplates)))
	}
	return nil
}
