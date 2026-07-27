package model

import (
	commonenum "common/pkg/enum"
	notifyenum "notify/internal/enum"
	"time"
)

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
