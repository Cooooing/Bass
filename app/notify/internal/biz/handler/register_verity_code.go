package handler

import (
	"bytes"
	"common/pkg/cutil/handlerchain"
	commonModel "common/pkg/model"
	"context"
	"html/template"
	"notify/internal/biz/base"
)

type RegisterVerifyCode struct {
	*handlerchain.BaseHandler[*commonModel.Notification]
	emailDomain *base.EmailDomain
}

func NewRegisterVerifyCode(emailDomain *base.EmailDomain) *RegisterVerifyCode {
	return &RegisterVerifyCode{
		BaseHandler: &handlerchain.BaseHandler[*commonModel.Notification]{Name: "register_verify_code"},
		emailDomain: emailDomain,
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

	// 发送邮件
	err = h.emailDomain.Send([]string{data.Meta.RegisterVerifyCode.Email}, data.Title, data.ContentRender)
	if err != nil {
		return nil, err
	}

	return h.BaseHandler.Next(ctx, data)
}

func (h *RegisterVerifyCode) SetNext(next handlerchain.Handler[*commonModel.Notification]) {
	h.BaseHandler.SetNext(next)
}

func (h *RegisterVerifyCode) Name() string {
	return h.BaseHandler.GetName()
}
