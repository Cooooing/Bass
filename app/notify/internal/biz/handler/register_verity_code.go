package handler

import (
	"bytes"
	v1 "common/api/notify/v1"
	"common/pkg/cutil/handlerchain"
	commonModel "common/pkg/model"
	"context"
	"fmt"
	"html/template"
	"notify/internal/biz/infra"
)

type RegisterVerifyCode struct {
	*handlerchain.BaseHandler[*commonModel.Notification]
	emailDomain      *infra.EmailDomain
	tencentSmsDomain *infra.TencentSmsDomain
}

func NewRegisterVerifyCode(emailDomain *infra.EmailDomain, tencentSmsDomain *infra.TencentSmsDomain) *RegisterVerifyCode {
	return &RegisterVerifyCode{
		BaseHandler:      &handlerchain.BaseHandler[*commonModel.Notification]{Name: "register_verify_code"},
		emailDomain:      emailDomain,
		tencentSmsDomain: tencentSmsDomain,
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
		err = h.emailDomain.Send([]string{data.Meta.RegisterVerifyCode.Email}, data.Title, data.ContentRender)
		if err != nil {
			return nil, err
		}
	case v1.NotificationChannel_NOTIFICATION_CHANNEL_SMS:
		// 发送短信
		err = h.tencentSmsDomain.SendSms(ctx, []string{data.Meta.RegisterVerifyCode.Phone}, []string{data.Meta.RegisterVerifyCode.Code})
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
