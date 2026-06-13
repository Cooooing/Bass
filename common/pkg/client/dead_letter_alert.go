package client

import (
	"context"
	"fmt"
	"time"

	"common/pkg/constant"
	"common/proto/gen/common"

	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const defaultDeadLetterAlertDedupWindow = time.Hour

// DeadLetterAlert 描述一条需要告警的死信记录。
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

// DeadLetterAlertClient 负责死信告警去重、日志、指标和 Lark 通知。
type DeadLetterAlertClient struct {
	log         *log.Helper
	redisClient *RedisClient
	larkClient  *LarkWebhookClient
	counter     metric.Int64Counter
}

func NewDeadLetterAlertClient(logger log.Logger, redisClient *RedisClient, larkClient *LarkWebhookClient) *DeadLetterAlertClient {
	counter, _ := otel.Meter("common.dead_letter").Int64Counter(
		"dead_letter_alert_total",
		metric.WithDescription("Total number of deduplicated dead letter alerts."),
	)
	return &DeadLetterAlertClient{
		log:         log.NewHelper(logger),
		redisClient: redisClient,
		larkClient:  larkClient,
		counter:     counter,
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
	c.log.Errorf("dead letter detected: service=%s source=%s event_id=%s event_type=%s subject=%s count=%d err=%s",
		alert.Service, alert.Source, alert.EventID, alert.EventType, alert.Subject, alert.Count, alert.LastError)
	if c.counter != nil {
		c.counter.Add(ctx, 1, metric.WithAttributes(
			attribute.String("service", alert.Service),
			attribute.String("source", alert.Source),
			attribute.String("event_type", alert.EventType),
			attribute.String("subject", alert.Subject),
		))
	}
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
