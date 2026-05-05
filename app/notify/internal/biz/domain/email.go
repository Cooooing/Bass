package domain

import (
	"context"
	"fmt"
	domainbase "notify/internal/biz/base"

	"gopkg.in/gomail.v2"
)

type EmailDomain struct {
	*domainbase.BaseDomain
}

func NewEmailDomain(base *domainbase.BaseDomain) *EmailDomain {
	return &EmailDomain{
		BaseDomain: base,
	}
}

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
	d.Log.Infof("send email from %s to %v success", d.Conf.Server.Email.From, to)
	return nil
}
