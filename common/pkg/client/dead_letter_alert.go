package client

import (
	"common/pkg/constant"
	"common/proto/gen/common"
	"context"
	"fmt"
	"log/slog"
	"time"
)

const defaultDeadLetterAlertDedupWindow = time.Hour

type DeadLetterAlert struct {
	Service   string
	Source    string
	EventID   string
	EventType string
	Subject   string
	Count     int32
	LastError string
	UpdatedAt *time.Time
}
type DeadLetterAlertClient struct {
	logger      *slog.Logger
	redisClient *RedisClient
	larkClient  *LarkWebhookClient
}

func NewDeadLetterAlertClient(
	logger *slog.Logger,
	redisClient *RedisClient,
	larkClient *LarkWebhookClient,
) *DeadLetterAlertClient {
	return &DeadLetterAlertClient{
		logger:      logger,
		redisClient: redisClient,
		larkClient:  larkClient,
	}
}

func (c *DeadLetterAlertClient) Alert(ctx context.Context, deadLetterConf *common.Event_DeadLetter, alertConf *common.Alert, alert *DeadLetterAlert) error {
	if c == nil || alert == nil || alert.Service == "" || alert.Source == "" || alert.EventID == "" {
		return nil
	}
	dedupWindow := defaultDeadLetterAlertDedupWindow
	if deadLetterConf != nil && deadLetterConf.GetAlertDedupWindow() != nil && deadLetterConf.GetAlertDedupWindow().AsDuration() > 0 {
		dedupWindow = deadLetterConf.GetAlertDedupWindow().AsDuration()
	}
	if c.redisClient != nil && c.redisClient.Client != nil {
		ok, err := c.redisClient.Client.SetNX(ctx, constant.GetKeyDeadLetterAlert(alert.Service, alert.Source, alert.EventID), "1", dedupWindow).Result()
		if err != nil || !ok {
			return err
		}
	}
	c.logger.ErrorContext(ctx, "dead letter detected", constant.LogFieldService, alert.Service, constant.LogFieldSource, alert.Source, constant.LogFieldEventID, alert.EventID, constant.LogFieldEventType, alert.EventType, constant.LogFieldSubject, alert.Subject, constant.LogFieldCount, alert.Count, constant.LogFieldLastError, alert.LastError)
	DeadLetterAlertsTotal.WithLabelValues(alert.Service, alert.Source, alert.EventType, alert.Subject).Inc()
	if deadLetterConf == nil || !deadLetterConf.GetEnableAlert() || alertConf == nil || alertConf.GetLarkWebhook() == nil || alertConf.GetLarkWebhook().GetToken() == "" || c.larkClient == nil {
		return nil
	}
	lark := alertConf.GetLarkWebhook()
	timeout := time.Duration(0)
	if lark.GetTimeout() != nil {
		timeout = lark.GetTimeout().AsDuration()
	}
	return c.larkClient.SendText(ctx, &LarkWebhookRequest{
		BaseURL: lark.GetBaseUrl(),
		Token:   lark.GetToken(),
		Secret:  lark.GetSecret(),
		Timeout: timeout,
		Text:    c.text(alert),
	})
}

func (c *DeadLetterAlertClient) text(alert *DeadLetterAlert) string {
	updatedAt := ""
	if alert.UpdatedAt != nil {
		updatedAt = alert.UpdatedAt.Format(time.RFC3339)
	}
	return fmt.Sprintf(
		"Dead letter detected\nservice: %s\nsource: %s\nevent_id: %s\nevent_type: %s\nsubject: %s\ncount: %d\nupdated_at: %s\nlast_error: %s",
		alert.Service,
		alert.Source,
		alert.EventID,
		alert.EventType,
		alert.Subject,
		alert.Count,
		updatedAt,
		alert.LastError,
	)
}
