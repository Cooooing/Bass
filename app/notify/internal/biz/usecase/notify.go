package usecase

import (
	"bytes"
	commonenum "common/pkg/enum"
	"context"
	"errors"
	"fmt"
	base "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"
	"text/template"

	"github.com/go-kratos/kratos/v2/log"
)

// UserContact 是 notify 从 user 服务读取的触达信息。
type UserContact struct {
	Phone string
	Email string
}

// UserClient 定义 notify 依赖的用户服务能力。
type UserClient interface {
	GetContacts(ctx context.Context, userIDs []int64) (map[int64]*UserContact, error)
	ListFollowerIDs(ctx context.Context, userID int64) ([]int64, error)
}

// NotificationIntent 是事件 handler 输出的通知意图。
// 事件 payload 只表达事件事实；各渠道目标由 handler 或 NotifyUsecase 按业务需要补齐。
type NotificationIntent struct {
	EventID   string
	EventType commonenum.EventType
	Vars      any

	Station []*StationInput
	Email   []*EmailInput
	SMS     []*SMSInput
	Webhook []*WebhookInput
}

type StationInput struct {
	UserID int64
}

type EmailInput struct {
	UserID int64
	Email  string
	Name   string
}

type SMSInput struct {
	UserID      int64
	Phone       string
	CountryCode string
}

type WebhookInput struct {
	UserID     int64
	EndpointID int64
	URL        string
}

type notificationPrepared struct {
	station    *stationPrepared
	deliveries []*model.NotificationDelivery
}

type stationPrepared struct {
	title       string
	content     string
	receiverIDs []int64
}

// NotifyUsecase 负责把通知意图写入站内信和外部投递任务。
type NotifyUsecase struct {
	log                      *log.Helper
	tx                       base.Tx
	userClient               UserClient
	notificationMetaRepo     repo.NotificationMetaRepo
	notificationRecordRepo   repo.NotificationRecordRepo
	notificationTemplateRepo repo.NotificationTemplateRepo
	notificationSettingRepo  repo.NotificationSettingRepo
	notificationDeliveryRepo repo.NotificationDeliveryRepo
}

func NewNotifyUsecase(
	logger log.Logger,
	tx base.Tx,
	userClient UserClient,
	notificationMetaRepo repo.NotificationMetaRepo,
	notificationRecordRepo repo.NotificationRecordRepo,
	notificationTemplateRepo repo.NotificationTemplateRepo,
	notificationSettingRepo repo.NotificationSettingRepo,
	notificationDeliveryRepo repo.NotificationDeliveryRepo,
) *NotifyUsecase {
	return &NotifyUsecase{
		log:                      log.NewHelper(logger),
		tx:                       tx,
		userClient:               userClient,
		notificationMetaRepo:     notificationMetaRepo,
		notificationRecordRepo:   notificationRecordRepo,
		notificationTemplateRepo: notificationTemplateRepo,
		notificationSettingRepo:  notificationSettingRepo,
		notificationDeliveryRepo: notificationDeliveryRepo,
	}
}

func (s *NotifyUsecase) Create(ctx context.Context, intent *NotificationIntent) error {
	prepared, err := s.prepare(ctx, intent)
	if err != nil {
		return err
	}
	if prepared == nil || (prepared.station == nil && len(prepared.deliveries) == 0) {
		return nil
	}
	return s.tx(ctx, func(ctx context.Context) error {
		return s.savePrepared(ctx, prepared)
	})
}

func (s *NotifyUsecase) prepare(ctx context.Context, intent *NotificationIntent) (*notificationPrepared, error) {
	prepared := &notificationPrepared{}
	if intent == nil || (len(intent.Station) == 0 && len(intent.Email) == 0 && len(intent.SMS) == 0 && len(intent.Webhook) == 0) {
		return prepared, nil
	}
	if intent.EventID == "" {
		return nil, errors.New("event id is required")
	}
	protoEventType, ok := commonenum.EventTypeMap.ToProto(intent.EventType)
	if !ok {
		return nil, errors.New("event type is invalid")
	}

	templates, err := s.notificationTemplateRepo.GetTemplates(ctx, protoEventType, string(notifyenum.LanguageZhCN))
	if err != nil {
		return nil, err
	}
	templatesByChannel := make(map[notifyenum.NotificationChannel]*model.NotificationTemplate, len(templates))
	for _, tpl := range templates {
		if tpl != nil {
			templatesByChannel[tpl.Channel] = tpl
		}
	}
	if len(templatesByChannel) == 0 {
		return prepared, nil
	}

	settingUserIDs := make([]int64, 0)
	contactUserIDs := make([]int64, 0)
	seenSettingUserIDs := make(map[int64]struct{})
	seenContactUserIDs := make(map[int64]struct{})
	for _, input := range intent.Station {
		if input == nil || input.UserID == 0 {
			continue
		}
		if _, exists := seenSettingUserIDs[input.UserID]; exists {
			continue
		}
		seenSettingUserIDs[input.UserID] = struct{}{}
		settingUserIDs = append(settingUserIDs, input.UserID)
	}
	for _, input := range intent.Email {
		if input == nil {
			continue
		}
		if input.UserID != 0 {
			if _, exists := seenSettingUserIDs[input.UserID]; !exists {
				seenSettingUserIDs[input.UserID] = struct{}{}
				settingUserIDs = append(settingUserIDs, input.UserID)
			}
		}
		if input.Email == "" && input.UserID != 0 {
			if _, exists := seenContactUserIDs[input.UserID]; exists {
				continue
			}
			seenContactUserIDs[input.UserID] = struct{}{}
			contactUserIDs = append(contactUserIDs, input.UserID)
		}
	}
	for _, input := range intent.SMS {
		if input == nil {
			continue
		}
		if input.UserID != 0 {
			if _, exists := seenSettingUserIDs[input.UserID]; !exists {
				seenSettingUserIDs[input.UserID] = struct{}{}
				settingUserIDs = append(settingUserIDs, input.UserID)
			}
		}
		if input.Phone == "" && input.UserID != 0 {
			if _, exists := seenContactUserIDs[input.UserID]; exists {
				continue
			}
			seenContactUserIDs[input.UserID] = struct{}{}
			contactUserIDs = append(contactUserIDs, input.UserID)
		}
	}
	for _, input := range intent.Webhook {
		if input == nil || input.UserID == 0 {
			continue
		}
		if _, exists := seenSettingUserIDs[input.UserID]; exists {
			continue
		}
		seenSettingUserIDs[input.UserID] = struct{}{}
		settingUserIDs = append(settingUserIDs, input.UserID)
	}

	// 直接邮箱、手机号和 Webhook URL 可以没有系统用户；用户偏好只作用于带 UserID 的目标。
	settingsByUserID := make(map[int64]map[notifyenum.NotificationChannel]bool)
	if len(settingUserIDs) > 0 {
		settings, err := s.notificationSettingRepo.List(ctx, &repo.NotificationSettingGetReq{
			UserIDs:   settingUserIDs,
			EventType: &protoEventType,
		})
		if err != nil {
			return nil, err
		}
		for _, setting := range settings {
			if setting == nil {
				continue
			}
			if _, hasTemplate := templatesByChannel[setting.Channel]; !hasTemplate {
				continue
			}
			if settingsByUserID[setting.UserID] == nil {
				settingsByUserID[setting.UserID] = make(map[notifyenum.NotificationChannel]bool)
			}
			settingsByUserID[setting.UserID][setting.Channel] = setting.Enable
		}
	}

	contactsByUserID := map[int64]*UserContact{}
	if len(contactUserIDs) > 0 && s.userClient != nil {
		contactsByUserID, err = s.userClient.GetContacts(ctx, contactUserIDs)
		if err != nil {
			return nil, err
		}
	}

	// 模板 channel 决定本次实际生成哪些目标；没有模板的 channel 不产生通知数据。
	if tpl := templatesByChannel[notifyenum.NotificationChannelStation]; tpl != nil {
		receiverIDs := make([]int64, 0, len(intent.Station))
		seenReceiverIDs := make(map[int64]struct{}, len(intent.Station))
		for _, input := range intent.Station {
			if input == nil || input.UserID == 0 {
				continue
			}
			if settingsByChannel := settingsByUserID[input.UserID]; settingsByChannel != nil {
				if enable, exists := settingsByChannel[notifyenum.NotificationChannelStation]; exists && !enable {
					continue
				}
			}
			if _, exists := seenReceiverIDs[input.UserID]; exists {
				continue
			}
			seenReceiverIDs[input.UserID] = struct{}{}
			receiverIDs = append(receiverIDs, input.UserID)
		}
		if len(receiverIDs) > 0 {
			title, err := renderTemplate(tpl.Title, intent.Vars)
			if err != nil {
				return nil, err
			}
			content, err := renderTemplate(tpl.Content, intent.Vars)
			if err != nil {
				return nil, err
			}
			prepared.station = &stationPrepared{
				title:       title,
				content:     content,
				receiverIDs: receiverIDs,
			}
		}
	}

	if tpl := templatesByChannel[notifyenum.NotificationChannelEmail]; tpl != nil {
		title, err := renderTemplate(tpl.Title, intent.Vars)
		if err != nil {
			return nil, err
		}
		content, err := renderTemplate(tpl.Content, intent.Vars)
		if err != nil {
			return nil, err
		}
		seenTargets := make(map[string]struct{}, len(intent.Email))
		for _, input := range intent.Email {
			if input == nil {
				continue
			}
			if input.UserID != 0 {
				if settingsByChannel := settingsByUserID[input.UserID]; settingsByChannel != nil {
					if enable, exists := settingsByChannel[notifyenum.NotificationChannelEmail]; exists && !enable {
						continue
					}
				}
			}
			target := input.Email
			if target == "" {
				if contact := contactsByUserID[input.UserID]; contact != nil {
					target = contact.Email
				}
			}
			if target == "" {
				continue
			}
			if _, exists := seenTargets[target]; exists {
				continue
			}
			seenTargets[target] = struct{}{}
			delivery := &model.NotificationDelivery{
				EventID:   intent.EventID,
				EventType: intent.EventType,
				Channel:   notifyenum.NotificationChannelEmail,
				Target:    target,
				Title:     title,
				Content:   content,
				Status:    notifyenum.NotificationDeliveryStatusPending,
			}
			if input.UserID != 0 {
				delivery.ReceiverID = new(input.UserID)
			}
			prepared.deliveries = append(prepared.deliveries, delivery)
		}
	}

	if tpl := templatesByChannel[notifyenum.NotificationChannelSMS]; tpl != nil {
		title, err := renderTemplate(tpl.Title, intent.Vars)
		if err != nil {
			return nil, err
		}
		content, err := renderTemplate(tpl.Content, intent.Vars)
		if err != nil {
			return nil, err
		}
		seenTargets := make(map[string]struct{}, len(intent.SMS))
		for _, input := range intent.SMS {
			if input == nil {
				continue
			}
			if input.UserID != 0 {
				if settingsByChannel := settingsByUserID[input.UserID]; settingsByChannel != nil {
					if enable, exists := settingsByChannel[notifyenum.NotificationChannelSMS]; exists && !enable {
						continue
					}
				}
			}
			target := input.Phone
			if target == "" {
				if contact := contactsByUserID[input.UserID]; contact != nil {
					target = contact.Phone
				}
			}
			if target == "" {
				continue
			}
			if input.CountryCode != "" {
				target = input.CountryCode + target
			}
			if _, exists := seenTargets[target]; exists {
				continue
			}
			seenTargets[target] = struct{}{}
			delivery := &model.NotificationDelivery{
				EventID:   intent.EventID,
				EventType: intent.EventType,
				Channel:   notifyenum.NotificationChannelSMS,
				Target:    target,
				Title:     title,
				Content:   content,
				Status:    notifyenum.NotificationDeliveryStatusPending,
			}
			if input.UserID != 0 {
				delivery.ReceiverID = new(input.UserID)
			}
			prepared.deliveries = append(prepared.deliveries, delivery)
		}
	}

	if tpl := templatesByChannel[notifyenum.NotificationChannelWebhook]; tpl != nil {
		title, err := renderTemplate(tpl.Title, intent.Vars)
		if err != nil {
			return nil, err
		}
		content, err := renderTemplate(tpl.Content, intent.Vars)
		if err != nil {
			return nil, err
		}
		seenTargets := make(map[string]struct{}, len(intent.Webhook))
		for _, input := range intent.Webhook {
			if input == nil || input.URL == "" {
				continue
			}
			if input.UserID != 0 {
				if settingsByChannel := settingsByUserID[input.UserID]; settingsByChannel != nil {
					if enable, exists := settingsByChannel[notifyenum.NotificationChannelWebhook]; exists && !enable {
						continue
					}
				}
			}
			if _, exists := seenTargets[input.URL]; exists {
				continue
			}
			seenTargets[input.URL] = struct{}{}
			delivery := &model.NotificationDelivery{
				EventID:   intent.EventID,
				EventType: intent.EventType,
				Channel:   notifyenum.NotificationChannelWebhook,
				Target:    input.URL,
				Title:     title,
				Content:   content,
				Status:    notifyenum.NotificationDeliveryStatusPending,
			}
			if input.UserID != 0 {
				delivery.ReceiverID = new(input.UserID)
			}
			prepared.deliveries = append(prepared.deliveries, delivery)
		}
	}
	return prepared, nil
}

func (s *NotifyUsecase) savePrepared(ctx context.Context, prepared *notificationPrepared) error {
	if prepared == nil {
		return nil
	}
	if prepared.station != nil && len(prepared.station.receiverIDs) > 0 {
		meta, err := s.notificationMetaRepo.Save(ctx, &model.NotificationMeta{
			Title:   prepared.station.title,
			Content: prepared.station.content,
			Status:  notifyenum.NotificationStatusNormal,
		})
		if err != nil {
			return err
		}
		records := make([]*model.NotificationRecord, 0, len(prepared.station.receiverIDs))
		for _, receiverID := range prepared.station.receiverIDs {
			records = append(records, &model.NotificationRecord{
				NotificationID: meta.ID,
				ReceiverID:     receiverID,
			})
		}
		if _, err := s.notificationRecordRepo.Saves(ctx, records); err != nil {
			return err
		}
	}
	if len(prepared.deliveries) == 0 {
		return nil
	}
	_, err := s.notificationDeliveryRepo.Saves(ctx, prepared.deliveries)
	return err
}

func renderTemplate(tplStr string, variables any) (string, error) {
	tpl, err := template.New("").Option("missingkey=error").Parse(tplStr)
	if err != nil {
		return "", fmt.Errorf("parse notification template: %w", err)
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, variables); err != nil {
		return "", fmt.Errorf("execute notification template: %w", err)
	}
	return buf.String(), nil
}
