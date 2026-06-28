package channel

import (
	"common/pkg/util"
	"context"
	"errors"
	"fmt"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"
	notifyenum "notify/internal/enum"

	"gopkg.in/gomail.v2"
	"log/slog"
)

var _ bizchannel.EmailClient = (*EmailClient)(nil)

type EmailClient struct {
	conf *conf.Bootstrap
	log  *util.LogHelper
}

func NewEmailClient(conf *conf.Bootstrap, logger *slog.Logger) *EmailClient {
	return &EmailClient{conf: conf, log: util.NewLogHelper(logger)}
}

func (c *EmailClient) SendEmail(_ context.Context, req *bizchannel.EmailRequest) (*bizchannel.SendResult, error) {
	if req == nil || req.ToEmail == "" {
		return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusFailed}, nil
	}
	if c.conf == nil || c.conf.Server == nil || c.conf.Server.Email == nil || !c.conf.Server.Email.Enable {
		return nil, errors.New("email sender is disabled")
	}
	email := c.conf.Server.Email

	contentType := req.ContentType
	if contentType == "" {
		contentType = "text/html"
	}
	message := gomail.NewMessage()
	message.SetAddressHeader("From", email.FromEmail, email.FromName)
	message.SetHeader("To", req.ToEmail)
	message.SetHeader("Subject", req.Subject)
	message.SetBody(contentType, req.Body)

	dialer := gomail.NewDialer(
		email.SmtpHost,
		int(email.SmtpPort),
		email.Username,
		email.Password,
	)
	dialer.SSL = email.SmtpPort == 465

	if err := dialer.DialAndSend(message); err != nil {
		return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusUnknown, ProviderResponse: new(fmt.Sprintf("send email: %v", err))}, nil
	}
	c.log.Infof("send email succeeded: %s", req.ToEmail)
	return &bizchannel.SendResult{Status: notifyenum.NotificationChannelStatusSucceeded}, nil
}
