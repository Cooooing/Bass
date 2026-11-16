package model

import (
	"time"

	notifyenum "notify/internal/enum"
)

type NotificationTencentSMSDelivery struct {
	ID                 int64
	EventID            string
	NotificationType   notifyenum.NotificationType
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
