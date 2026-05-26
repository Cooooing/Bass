package usecase

import (
	commonv1 "common/api/gen/common"
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"
	"testing"
)

func TestPrepareEmailUsesDirectAddressWithoutUserID(t *testing.T) {
	s := &NotifyUsecase{
		notificationTemplateRepo: &fakeNotificationTemplateRepo{templates: []*model.NotificationTemplate{
			{
				Channel: notifyenum.NotificationChannelEmail,
				Title:   "注册验证码",
				Content: "验证码 {{.Code}}",
			},
		}},
		notificationSettingRepo: &fakeNotificationSettingRepo{},
	}
	intent := &NotificationIntent{
		EventID:   "event-register-email",
		EventType: commonenum.EventTypeUserRegister,
		Vars: map[string]string{
			"Code": "123456",
		},
		Email: []*EmailInput{
			{Email: "user@example.com"},
		},
	}

	prepared, err := s.prepare(context.Background(), intent)
	if err != nil {
		t.Fatalf("prepare: %v", err)
	}
	if len(prepared.deliveries) != 1 {
		t.Fatalf("deliveries length = %d, want 1", len(prepared.deliveries))
	}
	delivery := prepared.deliveries[0]
	if delivery.ReceiverID != nil {
		t.Fatalf("receiver id = %v, want nil", *delivery.ReceiverID)
	}
	if delivery.Channel != notifyenum.NotificationChannelEmail {
		t.Fatalf("channel = %s, want %s", delivery.Channel, notifyenum.NotificationChannelEmail)
	}
	if delivery.Target != "user@example.com" {
		t.Fatalf("target = %s, want user@example.com", delivery.Target)
	}
	if delivery.Content != "验证码 123456" {
		t.Fatalf("content = %s, want 验证码 123456", delivery.Content)
	}
}

func TestPrepareWebhookSkipsDisabledUserAndDuplicateURL(t *testing.T) {
	s := &NotifyUsecase{
		notificationTemplateRepo: &fakeNotificationTemplateRepo{templates: []*model.NotificationTemplate{
			{
				Channel: notifyenum.NotificationChannelWebhook,
				Title:   "事件 {{.ID}}",
				Content: "内容 {{.ID}}",
			},
		}},
		notificationSettingRepo: &fakeNotificationSettingRepo{settings: []*model.NotificationSetting{
			{
				UserID:  2,
				Channel: notifyenum.NotificationChannelWebhook,
				Enable:  false,
			},
		}},
	}
	intent := &NotificationIntent{
		EventID:   "event-webhook",
		EventType: commonenum.EventTypeContentArticlePublish,
		Vars: map[string]string{
			"ID": "1001",
		},
		Webhook: []*WebhookInput{
			{UserID: 1, URL: "https://example.com/hook"},
			{UserID: 1, URL: "https://example.com/hook"},
			{UserID: 2, URL: "https://example.com/disabled"},
		},
	}

	prepared, err := s.prepare(context.Background(), intent)
	if err != nil {
		t.Fatalf("prepare: %v", err)
	}
	if len(prepared.deliveries) != 1 {
		t.Fatalf("deliveries length = %d, want 1", len(prepared.deliveries))
	}
	delivery := prepared.deliveries[0]
	if delivery.ReceiverID == nil || *delivery.ReceiverID != 1 {
		t.Fatalf("receiver id = %v, want 1", delivery.ReceiverID)
	}
	if delivery.Channel != notifyenum.NotificationChannelWebhook {
		t.Fatalf("channel = %s, want %s", delivery.Channel, notifyenum.NotificationChannelWebhook)
	}
	if delivery.Target != "https://example.com/hook" {
		t.Fatalf("target = %s, want https://example.com/hook", delivery.Target)
	}
	if delivery.Title != "事件 1001" {
		t.Fatalf("title = %s, want 事件 1001", delivery.Title)
	}
}

type fakeNotificationTemplateRepo struct {
	templates []*model.NotificationTemplate
}

func (r *fakeNotificationTemplateRepo) Save(context.Context, *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	panic("not implemented")
}

func (r *fakeNotificationTemplateRepo) Update(context.Context, *model.NotificationTemplate) (*model.NotificationTemplate, error) {
	panic("not implemented")
}

func (r *fakeNotificationTemplateRepo) GetTemplates(context.Context, enums.EventType, string) ([]*model.NotificationTemplate, error) {
	return r.templates, nil
}

func (r *fakeNotificationTemplateRepo) Get(context.Context, *repo.NotificationTemplateGetReq) (*model.NotificationTemplate, error) {
	panic("not implemented")
}

func (r *fakeNotificationTemplateRepo) GetList(context.Context, *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, error) {
	panic("not implemented")
}

func (r *fakeNotificationTemplateRepo) GetMap(context.Context, *repo.NotificationTemplateGetReq) (map[string]*model.NotificationTemplate, error) {
	panic("not implemented")
}

func (r *fakeNotificationTemplateRepo) GetPage(context.Context, *commonv1.PageRequest, *repo.NotificationTemplateGetReq) ([]*model.NotificationTemplate, *commonv1.PageReply, error) {
	panic("not implemented")
}

type fakeNotificationSettingRepo struct {
	settings []*model.NotificationSetting
}

func (r *fakeNotificationSettingRepo) List(context.Context, *repo.NotificationSettingGetReq) ([]*model.NotificationSetting, error) {
	return r.settings, nil
}

func (r *fakeNotificationSettingRepo) Save(context.Context, *model.NotificationSetting) (*model.NotificationSetting, error) {
	panic("not implemented")
}

func (r *fakeNotificationSettingRepo) Upsert(context.Context, *model.NotificationSetting) (*model.NotificationSetting, error) {
	panic("not implemented")
}

var _ repo.NotificationTemplateRepo = (*fakeNotificationTemplateRepo)(nil)
var _ repo.NotificationSettingRepo = (*fakeNotificationSettingRepo)(nil)
