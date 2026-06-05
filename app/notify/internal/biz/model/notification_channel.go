package model

import (
	commonenum "common/pkg/enum"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationStationMessage struct {
	ID         int64
	EventID    string
	EventType  commonenum.EventType
	ReceiverID int64
	Title      string
	Content    string
	Status     notifyenum.NotificationChannelStatus
	ReadAt     *time.Time
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

type NotificationEmailDelivery struct {
	ID                int64
	EventID           string
	EventType         commonenum.EventType
	ReceiverID        *int64
	ToEmail           string
	Subject           string
	Body              string
	ContentType       string
	Status            notifyenum.NotificationChannelStatus
	AttemptCount      int32
	LastAttemptAt     *time.Time
	ProviderMessageID *string
	ProviderResponse  *string
	SentAt            *time.Time
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}

type NotificationTencentSMSDelivery struct {
	ID                 int64
	EventID            string
	EventType          commonenum.EventType
	ReceiverID         *int64
	Phone              string
	SMSSDKAppID        string
	SignName           string
	ProviderTemplateID string
	TemplateParams     []string
	Status             notifyenum.NotificationChannelStatus
	AttemptCount       int32
	LastAttemptAt      *time.Time
	ProviderRequestID  *string
	ProviderCode       *string
	ProviderMessage    *string
	SentAt             *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type NotificationLarkWebhookDelivery struct {
	ID            int64
	EventID       string
	EventType     commonenum.EventType
	WebhookID     string
	RequestBody   string
	Status        notifyenum.NotificationChannelStatus
	AttemptCount  int32
	LastAttemptAt *time.Time
	HTTPStatus    *int
	ResponseBody  *string
	SentAt        *time.Time
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
