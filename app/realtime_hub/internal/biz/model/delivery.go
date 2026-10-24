package model

import (
	"time"

	"realtime_hub/internal/enum"
)

type Delivery struct {
	UserID  int64
	Kind    enum.DeliveryKind
	Payload DeliveryPayload
}

type NodeMessage struct {
	Delivery        *Delivery
	CloseConnection *CloseConnection
	ConnectionBound *ConnectionBound
}

type CloseConnection struct {
	ConnectionID string
	Reason       string
}

type ConnectionBound struct {
	ConnectionID string
}

type EvictedConnection struct {
	NodeID       string
	ConnectionID string
}

// DeliveryPayload 是按照投递类别承载的领域下行内容。
type DeliveryPayload interface {
	deliveryPayload()
}

// NotificationDelivery 是站内信变化的下行内容。
type NotificationDelivery struct {
	NotificationID int64
}

func (NotificationDelivery) deliveryPayload() {}

// PrivateChatMessageDelivery 是私聊消息变化的下行内容。
type PrivateChatMessageDelivery struct {
	MessageID   int64
	SenderID    int64
	ReceiverID  int64
	MessageType int32
	Content     string
	CreatedAt   time.Time
}

func (PrivateChatMessageDelivery) deliveryPayload() {}

// GlobalChatMessageDelivery 是全局聊天室消息变化的下行内容。
type GlobalChatMessageDelivery struct {
	MessageID   int64
	SenderID    int64
	MessageType int32
	Content     string
	CreatedAt   time.Time
}

func (GlobalChatMessageDelivery) deliveryPayload() {}

type GlobalChatBarrageDelivery struct {
	BarrageID int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}

func (GlobalChatBarrageDelivery) deliveryPayload() {}
