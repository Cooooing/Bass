package handler

import (
	"bytes"
	infrav1 "common/gen/infra/v1"
	v1 "common/gen/notify/v1"
	commonModel "common/pkg/model"
	"common/pkg/util/handlerchain"
	"context"
	"fmt"
	"html/template"
	domainbase "notify/internal/biz/base"
)

type RegisterVerifyCode struct {
	*domainbase.BaseDomain
	*handlerchain.BaseHandler[*commonModel.Notification]
}

func NewRegisterVerifyCode(baseDomain *domainbase.BaseDomain) *RegisterVerifyCode {
	return &RegisterVerifyCode{
		BaseDomain:  baseDomain,
		BaseHandler: &handlerchain.BaseHandler[*commonModel.Notification]{Name: "register_verify_code"},
	}
}

func (h *RegisterVerifyCode) Handle(ctx context.Context, data *commonModel.Notification) (*commonModel.Notification, error) {

	// 按模板渲染
	data.Meta.RegisterVerifyCode.ExpireMinutes = int(data.Meta.RegisterVerifyCode.Expire.Minutes())
	tpl, err := template.New(data.UUID).Parse(data.Content)
	if err != nil {
		return nil, err
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data.Meta); err != nil {
		return nil, err
	}
	data.ContentRender = buf.String()

	switch data.Channel {
	case v1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL:
		// 发送邮件
		_, err = h.Infra.Email.Send(ctx, &infrav1.SendEmailRequest{
			Title:   data.Title,
			Content: data.ContentRender,
			To:      []string{data.Meta.RegisterVerifyCode.Email},
		})
		if err != nil {
			return nil, err
		}
	case v1.NotificationChannel_NOTIFICATION_CHANNEL_SMS:
		// 发送短信
		_, err = h.Infra.Sms.Send(ctx, &infrav1.SendSmsRequest{
			Phone:  []string{data.Meta.RegisterVerifyCode.Phone},
			Params: []string{data.Meta.RegisterVerifyCode.Code},
		})
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("channel %s not supported", data.Channel)
	}

	return h.BaseHandler.Next(ctx, data)
}

func (h *RegisterVerifyCode) SetNext(next handlerchain.Handler[*commonModel.Notification]) {
	h.BaseHandler.SetNext(next)
}

func (h *RegisterVerifyCode) Name() string {
	return h.BaseHandler.GetName()
}
