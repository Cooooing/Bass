package domain

import (
	"context"
	"fmt"
	doaminbase "infra/internal/biz/base"

	"gopkg.in/gomail.v2"
)

type EmailDomain struct {
	*doaminbase.BaseDomain
}

func NewEmailDomain(base *doaminbase.BaseDomain) *EmailDomain {
	return &EmailDomain{
		BaseDomain: base,
	}
}

// Send 发送邮件
func (d *EmailDomain) Send(ctx context.Context, to []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", d.Conf.Server.Email.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(d.Conf.Server.Email.SmtpHost, int(d.Conf.Server.Email.SmtpPort), d.Conf.Server.Email.Username, d.Conf.Server.Email.Password)
	dialer.SSL = d.Conf.Server.Email.SmtpPort == 465

	err := dialer.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("send email error: %v", err)
	}
	d.Log.Infof(" %s send email to %v success", d.Conf.Server.Email.From, to)
	return nil
}
