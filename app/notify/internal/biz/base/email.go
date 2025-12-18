package base

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailDomain struct {
	*BaseDomain
}

func NewEmailDomain(base *BaseDomain) *EmailDomain {
	return &EmailDomain{
		BaseDomain: base,
	}
}

// Send 发送邮件
func (d *EmailDomain) Send(to []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", d.Conf.Email.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(d.Conf.Email.SmtpHost, int(d.Conf.Email.SmtpPort), d.Conf.Email.Username, d.Conf.Email.Password)
	dialer.SSL = d.Conf.Email.SmtpPort == 465

	err := dialer.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("send email error: %v", err)
	}
	d.Log.Infof(" %s send email to %v success", d.Conf.Email.From, to)
	return nil
}
