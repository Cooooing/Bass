package model

type ClientEvent struct {
	Name    string
	Payload any
}

type NotificationEventPayload struct {
	NotificationID int64 `json:"notification_id"`
}

type PrivateChatMessageEventPayload struct {
	MessageID   int64  `json:"message_id"`
	SenderID    int64  `json:"sender_id"`
	ReceiverID  int64  `json:"receiver_id"`
	MessageType string `json:"message_type"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
}

type GlobalChatMessageEventPayload struct {
	MessageID   int64  `json:"message_id"`
	SenderID    int64  `json:"sender_id"`
	MessageType string `json:"message_type"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
}

type GlobalChatBarrageEventPayload struct {
	BarrageID int64  `json:"barrage_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type GlobalChatTopicEventPayload struct {
	TopicID   int64  `json:"topic_id"`
	UserID    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
