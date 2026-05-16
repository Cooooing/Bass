package domain

import (
	"bytes"
	"common/api/gen/common/enums"
	"common/pkg/client"
	"common/pkg/enum"
	"context"
	"html/template"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent/gen"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type EventHandler struct {
	*domainbase.BaseDomain
	natsClient               *client.NatsClient
	notificationTemplateRepo repo.NotificationTemplateRepo
	emailDomain              *EmailDomain
	smsDomain                *TencentSmsDomain

	ctx    context.Context
	cancel context.CancelFunc
}

func NewEventHandler(
	base *domainbase.BaseDomain,
	natsClient *client.NatsClient,
	notificationTemplateRepo repo.NotificationTemplateRepo,
	emailDomain *EmailDomain,
	smsDomain *TencentSmsDomain,
) *EventHandler {
	return &EventHandler{
		BaseDomain:               base,
		natsClient:               natsClient,
		notificationTemplateRepo: notificationTemplateRepo,
		emailDomain:              emailDomain,
		smsDomain:                smsDomain,
	}
}

// Start 实现 transport.Server，由 Kratos 调用
func (h *EventHandler) Start(ctx context.Context) error {
	h.ctx, h.cancel = context.WithCancel(ctx)
	h.initTemplates()
	for _, subject := range []string{"content.>", "user.>"} {
		_, err := h.natsClient.Subscribe(h.ctx, subject, h.handleMessage)
		if err != nil {
			h.Log.Errorf("subscribe %s failed: %v", subject, err)
			continue
		}
	}
	return nil
}

// Stop 实现 transport.Server，由 Kratos 调用
func (h *EventHandler) Stop(_ context.Context) error {
	h.cancel()
	if h.natsClient != nil {
		return h.natsClient.Close()
	}
	return nil
}

func (h *EventHandler) initTemplates() {
	type tpl struct {
		eventType enum.EventType
		channel   enum.NotificationChannel
		title     string
		content   string
	}
	defaults := []tpl{
		{enum.EventTypeArticlePublished, enum.NotificationChannelStation, "新文章发布", "{{.senderName}} 发布了文章「{{.title}}」"},
		{enum.EventTypeArticleLiked, enum.NotificationChannelStation, "文章被点赞", "{{.senderName}} 点赞了你的文章「{{.title}}」"},
		{enum.EventTypeArticleThanked, enum.NotificationChannelStation, "文章被感谢", "{{.senderName}} 感谢了你的文章「{{.title}}」"},
		{enum.EventTypeArticleCollected, enum.NotificationChannelStation, "文章被收藏", "{{.senderName}} 收藏了你的文章「{{.title}}」"},
		{enum.EventTypeArticleWatched, enum.NotificationChannelStation, "文章被关注", "{{.senderName}} 关注了你的文章「{{.title}}」"},
		{enum.EventTypeCommentPublished, enum.NotificationChannelStation, "收到新评论", "{{.senderName}} 评论了你的文章"},
		{enum.EventTypeCommentLiked, enum.NotificationChannelStation, "评论被点赞", "{{.senderName}} 点赞了你的评论"},
		{enum.EventTypeUserFollowCreated, enum.NotificationChannelStation, "新增关注", "{{.senderName}} 关注了你"},
	}
	ctx := context.Background()
	for _, t := range defaults {
		_, err := h.notificationTemplateRepo.Save(ctx, h.Db, &model.NotificationTemplate{
			NotificationTemplate: &gen.NotificationTemplate{
				EventType: t.eventType,
				Channel:   t.channel,
				Title:     t.title,
				Content:   t.content,
				Enable:    true,
			},
		})
		if err != nil {
			h.Log.Errorf("init template failed: type=%v channel=%v err=%v", t.eventType, t.channel, err)
		}
	}
	h.Log.Infof("default templates initialized (%d)", len(defaults))
}

func (h *EventHandler) handleMessage(ctx context.Context, msg *client.Message) error {
	var event enums.Event
	if err := proto.Unmarshal(msg.Data, &event); err != nil {
		h.Log.Errorf("unmarshal event failed: subject=%s err=%v", msg.Subject, err)
		return nil
	}

	h.Log.Infof("received event: type=%v receivers=%d subject=%s",
		event.Type, len(event.ReceiverIds), msg.Subject)

	if len(event.ReceiverIds) == 0 {
		h.Log.Debugf("no receivers, skip: type=%v", event.Type)
		return nil
	}

	// TODO: 存 notification 记录

	// 按模板渠道投递
	tplVars := h.extractPayload(&event)
	templates, err := h.notificationTemplateRepo.GetTemplates(ctx, event.Type)
	if err != nil {
		h.Log.Errorf("get templates failed: type=%v err=%v", event.Type, err)
		return nil
	}
	for _, tpl := range templates {
		title := h.render(tpl.Title, tplVars)
		content := h.render(tpl.Content, tplVars)

		h.Log.Infof("deliver to %d receivers channel=%s, title=%s, content=%s",
			len(event.ReceiverIds), tpl.Channel, title, content)

		// TODO: 按渠道投递（站内信/邮件/短信）
		_ = title
		_ = content
	}

	return nil
}

func (h *EventHandler) extractPayload(msg proto.Message) proto.Message {
	ref := msg.ProtoReflect()
	desc := ref.Descriptor()
	for i := 0; i < desc.Fields().Len(); i++ {
		fd := desc.Fields().Get(i)
		if fd.ContainingOneof() == nil {
			continue
		}
		if !ref.Has(fd) {
			continue
		}
		if fd.Kind() == protoreflect.MessageKind {
			return ref.Get(fd).Message().Interface()
		}
	}
	return nil
}

func (h *EventHandler) render(tplStr string, variables any) string {
	tpl, err := template.New("").Parse(tplStr)
	if err != nil {
		h.Log.Errorf("parse template failed: %v", err)
		return tplStr
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, variables); err != nil {
		h.Log.Errorf("execute template failed: %v", err)
		return tplStr
	}
	return buf.String()
}
