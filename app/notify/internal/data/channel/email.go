package channel

import (
	"context"
	"fmt"

	bizrepo "notify/internal/biz/repo"
	"notify/internal/config"
	notifyenum "notify/internal/enum"

	"log/slog"

	"gopkg.in/gomail.v2"
)

var _ bizrepo.EmailClient = (*EmailClient)(nil)

type EmailClient struct {
	conf *config.Bootstrap
	log  *slog.Logger
}

func NewEmailClient(
	conf *config.Bootstrap,
	logger *slog.Logger,
) *EmailClient {
	return &EmailClient{
		conf: conf,
		log:  logger,
	}
}

func (c *EmailClient) SendEmail(ctx context.Context, req *bizrepo.EmailRequest) (*bizrepo.EmailSendResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if req == nil || req.ToEmail == "" {
		return &bizrepo.EmailSendResult{
			Status: notifyenum.NotificationChannelStatusFailed,
		}, nil
	}
	if c.conf == nil || c.conf.GetBusiness() == nil || c.conf.GetBusiness().GetEmail() == nil || !c.conf.GetBusiness().GetEmail().GetEnable() {
		return &bizrepo.EmailSendResult{
			Status:       notifyenum.NotificationChannelStatusSucceeded,
			ProviderResp: new("email channel disabled, skipped provider call"),
		}, nil
	}
	email := c.conf.GetBusiness().GetEmail()

	contentType := req.ContentType
	if contentType == "" {
		contentType = "text/html"
	}
	message := gomail.NewMessage()
	message.SetAddressHeader("From", email.GetFromEmail(), email.GetFromName())
	message.SetHeader("To", req.ToEmail)
	message.SetHeader("Subject", req.Subject)
	message.SetBody(contentType, req.Body)

	dialer := gomail.NewDialer(email.GetSmtpHost(), int(email.GetSmtpPort()), email.GetUsername(), email.GetPassword())
	dialer.SSL = email.GetSmtpPort() == 465

	if err := dialer.DialAndSend(message); err != nil {
		return &bizrepo.EmailSendResult{
			Status:       notifyenum.NotificationChannelStatusUnknown,
			ProviderResp: new(fmt.Sprintf("send email: %v", err)),
		}, nil
	}
	c.log.Info("send email succeeded: to=" + req.ToEmail)
	return &bizrepo.EmailSendResult{
		Status: notifyenum.NotificationChannelStatusSucceeded,
	}, nil
}
