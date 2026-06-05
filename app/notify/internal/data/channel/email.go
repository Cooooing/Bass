package channel

import (
	"context"
	"errors"
	"fmt"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"
	notifyenum "notify/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
	"gopkg.in/gomail.v2"
)

var _ bizchannel.EmailClient = (*EmailClient)(nil)

type EmailClient struct {
	conf *conf.Bootstrap
	log  *log.Helper
}

func NewEmailClient(conf *conf.Bootstrap, logger log.Logger) *EmailClient {
	return &EmailClient{conf: conf, log: log.NewHelper(logger)}
}

func (c *EmailClient) SendEmail(_ context.Context, req *bizchannel.EmailRequest) (*bizchannel.SendResult, error) {
	if req == nil || req.ToEmail == "" {
		return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusFailed}, nil
	}
	if c.conf == nil || c.conf.Server == nil || c.conf.Server.Email == nil || !c.conf.Server.Email.Enable {
		return nil, errors.New("email sender is disabled")
	}

	contentType := req.ContentType
	if contentType == "" {
		contentType = "text/html"
	}
	message := gomail.NewMessage()
	message.SetHeader("From", c.conf.Server.Email.From)
	message.SetHeader("To", req.ToEmail)
	message.SetHeader("Subject", req.Subject)
	message.SetBody(contentType, req.Body)

	dialer := gomail.NewDialer(
		c.conf.Server.Email.SmtpHost,
		int(c.conf.Server.Email.SmtpPort),
		c.conf.Server.Email.Username,
		c.conf.Server.Email.Password,
	)
	dialer.SSL = c.conf.Server.Email.SmtpPort == 465

	if err := dialer.DialAndSend(message); err != nil {
		return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusUnknown, ProviderResponse: new(fmt.Sprintf("send email: %v", err))}, nil
	}
	c.log.Infof("send email succeeded: to=%s", req.ToEmail)
	return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusSucceeded}, nil
}
