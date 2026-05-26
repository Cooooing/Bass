package channel

import (
	"context"
	"errors"
	"fmt"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/gomail.v2"
)

// EmailChannel 通过 SMTP 发送邮件。
type EmailChannel struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewEmailChannel(conf *conf.Bootstrap, logger log.Logger) *EmailChannel {
	return &EmailChannel{
		conf: conf,
		log:  log.NewHelper(logger),
	}
}

func (c *EmailChannel) Send(_ context.Context, req *bizchannel.SendReq) error {
	if req == nil || req.Target == "" {
		return nil
	}
	if c.conf == nil || c.conf.Server == nil || c.conf.Server.Email == nil || !c.conf.Server.Email.Enable {
		return errors.New("email sender is disabled")
	}

	message := gomail.NewMessage()
	message.SetHeader("From", c.conf.Server.Email.From)
	message.SetHeader("To", req.Target)
	message.SetHeader("Subject", req.Title)
	message.SetBody("text/html", req.Content)

	dialer := gomail.NewDialer(
		c.conf.Server.Email.SmtpHost,
		int(c.conf.Server.Email.SmtpPort),
		c.conf.Server.Email.Username,
		c.conf.Server.Email.Password,
	)
	dialer.SSL = c.conf.Server.Email.SmtpPort == 465

	if err := dialer.DialAndSend(message); err != nil {
		return fmt.Errorf("send email: %w", err)
	}
	c.log.Infof("send email to %s", req.Target)
	return nil
}
