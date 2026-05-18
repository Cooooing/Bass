package sender

import (
	v1 "common/api/gen/notify/v1"
	"context"
	"fmt"
	"notify/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/gomail.v2"
)

// SmtpSender 邮件发送器（SMTP）
type SmtpSender struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewSmtpSender(conf *conf.Bootstrap, logger log.Logger) *SmtpSender {
	return &SmtpSender{
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (s *SmtpSender) Channel() v1.NotificationChannel {
	return v1.NotificationChannel_NOTIFICATION_CHANNEL_EMAIL
}

func (s *SmtpSender) Send(ctx context.Context, req *SendRequest) error {
	to := req.UserInfo.Email
	if to == "" {
		return nil
	}
	m := gomail.NewMessage()
	m.SetHeader("From", s.conf.Server.Email.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", req.Title)
	m.SetBody("text/html", req.Content)

	dialer := gomail.NewDialer(
		s.conf.Server.Email.SmtpHost,
		int(s.conf.Server.Email.SmtpPort),
		s.conf.Server.Email.Username,
		s.conf.Server.Email.Password,
	)
	dialer.SSL = s.conf.Server.Email.SmtpPort == 465

	if err := dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("send email error: %v", err)
	}
	s.log.Infof("send email to %s success", to)
	return nil
}
