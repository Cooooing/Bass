package domain

import (
	"bytes"
	v1 "common/api/gen/notify/v1"
	"common/pkg/client"
	commonModel "common/pkg/model"
	"context"
	"encoding/json"
	"html/template"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/ent"
	"notify/internal/data/ent/gen"
)

type EventHandler struct {
	*domainbase.BaseDomain
	natsClient          *client.NatsClient
	emailDomain         *EmailDomain
	smsDoamin           *TencentSmsDomain
	messageMetaRepo     repo.NotificationMetaRepo
	messageRecordRepo   repo.NotificationRecordRepo
	messageTemplateRepo repo.NotificationTemplateRepo

	ctx    context.Context
	cancel context.CancelFunc
}

func NewEventHandler(
	base *domainbase.BaseDomain,
	natsClient *client.NatsClient,
	emailDomain *EmailDomain,
	smsDoamin *TencentSmsDomain,
	messageMetaRepo repo.NotificationMetaRepo,
	messageRecordRepo repo.NotificationRecordRepo,
	messageTemplateRepo repo.NotificationTemplateRepo,
) (*EventHandler, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())
	h := &EventHandler{
		BaseDomain:          base,
		natsClient:          natsClient,
		messageMetaRepo:     messageMetaRepo,
		messageRecordRepo:   messageRecordRepo,
		messageTemplateRepo: messageTemplateRepo,
		ctx:                 ctx,
		cancel:              cancel,
	}
	return h, h.CleanUp, nil
}

func (h *EventHandler) Handle() {
	_, err := h.natsClient.QueueSubscribe(h.ctx, "notify.>", "msg_center", func(ctx context.Context, msg *client.Message) error {
		h.Log.Infof("received NATS message: subject=%s", msg.Subject)

		n := &commonModel.Notification{}
		if err := json.Unmarshal(msg.Data, n); err != nil {
			h.Log.Errorf("unmarshal failed: %v", err)
			return err
		}

		// 查模板
		channels := n.Channels
		if len(channels) == 0 {
			h.Log.Warnf("no channel specified")
			return nil
		}

		templates, err := h.messageTemplateRepo.GetCache(ctx, n.Type, channels)
		if err != nil {
			h.Log.Errorf("get template failed: %v", err)
			return err
		}

		// 逐模板处理
		var lastErr error
		saved := false
		for _, tpl := range templates {
			title := h.renderContent(tpl.Title, n.Meta)
			content := h.renderContent(tpl.Content, n.Meta)

			// 落库（仅存一次）
			if !saved {
				if err := h.saveMessage(ctx, n, title, content); err != nil {
					h.Log.Errorf("save message failed: err=%v", err)
					lastErr = err
					continue
				}
				saved = true
			}

			// 渠道分发
			h.dispatch(ctx, tpl.Channel, n, title, content)
		}

		return lastErr
	})
	if err != nil {
		h.Log.Errorf("subscribe failed: %v", err)
	}
	h.Log.Info("NATS consumer started: notify.>")
}

func (h *EventHandler) renderContent(tplStr string, meta commonModel.Meta) string {
	tpl, err := template.New("").Parse(tplStr)
	if err != nil {
		h.Log.Errorf("parse template failed: %v", err)
		return tplStr
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, meta); err != nil {
		h.Log.Errorf("execute template failed: %v", err)
		return tplStr
	}
	return buf.String()
}

func (h *EventHandler) saveMessage(ctx context.Context, n *commonModel.Notification, title, content string) error {
	return ent.WithTx(ctx, h.Db, func(tx *gen.Client) error {
		meta, err := h.messageMetaRepo.Save(ctx, tx, &model.NotificationMeta{
			NotificationMeta: &gen.NotificationMeta{
				UUID:             n.UUID,
				NotificationType: n.Type.String(),
				SenderID:         n.SenderId,
				Meta:             n.Meta,
				Title:            title,
				Content:          content,
				IsGlobal:         n.Global,
				Status:           v1.NotificationStatus_NOTIFICATION_STATUS_NORMAL.String(),
			},
		})
		if err != nil {
			return err
		}

		// 非全局消息，为每个接收者存 record
		if !n.Global && len(n.ReceiverIds) > 0 {
			records := make([]*model.NotificationRecord, 0, len(n.ReceiverIds))
			for _, rid := range n.ReceiverIds {
				records = append(records, &model.NotificationRecord{
					NotificationRecord: &gen.NotificationRecord{
						NotificationID: meta.ID,
						ReceiverID:     rid,
					},
				})
			}
			_, err = h.messageRecordRepo.Saves(ctx, tx, records)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (h *EventHandler) dispatch(ctx context.Context, channel string, n *commonModel.Notification, title, content string) {
	switch channel {
	case v1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL.String():
		for _, rid := range n.ReceiverIds {
			err := h.emailDomain.Send(ctx,
				[]string{}, // Todo 需要从用户服务获取邮箱
				title,
				content,
			)
			if err != nil {
				h.Log.Errorf("email dispatch failed: receiver=%d err=%v", rid, err)
			}
		}
	case v1.NotificationChannel_NOTIFICATION_CHANNEL_SMS.String():
		if n.Global {
			return
		}
		for _, rid := range n.ReceiverIds {
			_ = rid
			err := h.smsDoamin.Send(ctx,
				[]string{}, // Todo 需要从用户服务获取手机号
				[]string{content},
			)
			if err != nil {
				h.Log.Errorf("sms dispatch failed: receiver=%d err=%v", rid, err)
			}
		}
	case v1.NotificationChannel_NOTIFICATION_CHANNEL_WEBSITE.String():
		// 已落库，无需额外操作
	}
}

func (h *EventHandler) CleanUp() {
	h.cancel()
	if h.natsClient != nil {
		_ = h.natsClient.Close()
	}
}
