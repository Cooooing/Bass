package model

type ClientEvent struct {
	Name    string
	Payload any
}

type NotificationEventPayload struct {
	NotificationID int64 `json:"notification_id"`
}
