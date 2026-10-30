package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"realtime_hub/internal/biz/model"
	"realtime_hub/internal/biz/repo"
	"realtime_hub/internal/enum"

	"github.com/google/uuid"
)

type RealtimeUsecase struct {
	log      *slog.Logger
	registry repo.Registry
	control  repo.ControlPublisher
	mu       sync.RWMutex
	nodes    map[string]map[string]*nodeSession
}

type nodeSession struct {
	ctx      context.Context
	messages chan *model.NodeMessage
}

func NewRealtimeUsecase(log *slog.Logger, registry repo.Registry, control repo.ControlPublisher) *RealtimeUsecase {
	return &RealtimeUsecase{log: log, registry: registry, control: control, nodes: map[string]map[string]*nodeSession{}}
}

func (u *RealtimeUsecase) OpenSession(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string) (<-chan *model.NodeMessage, error) {
	if nodeID == "" || len(realms) == 0 {
		return nil, fmt.Errorf("node session is invalid")
	}
	if err := u.registry.RegisterNode(ctx, nodeID, group, realms); err != nil {
		return nil, err
	}
	session := &nodeSession{ctx: ctx, messages: make(chan *model.NodeMessage, 128)}
	u.mu.Lock()
	for _, realm := range realms {
		if u.nodes[realm] == nil {
			u.nodes[realm] = map[string]*nodeSession{}
		}
		u.nodes[realm][nodeID] = session
	}
	u.mu.Unlock()
	return session.messages, nil
}

func (u *RealtimeUsecase) CloseSession(ctx context.Context, nodeID string, realms []string) {
	_ = u.registry.RemoveNode(context.WithoutCancel(ctx), nodeID, realms)
	u.removeLocal(nodeID, realms)
}

func (u *RealtimeUsecase) RenewSession(ctx context.Context, nodeID string, group enum.NodeGroup, realms []string) error {
	valid, err := u.registry.RenewNode(ctx, nodeID, group, realms)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("node lease requires a new stream")
	}
	return nil
}

func (u *RealtimeUsecase) BindConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) error {
	evicted, err := u.registry.BindConnection(ctx, nodeID, realm, userID, connectionID)
	if err != nil {
		return err
	}
	for _, connection := range evicted {
		if u.DeliverCloseConnection(ctx, realm, connection.NodeID, connection.ConnectionID, "connection limit exceeded") {
			continue
		}
		if err := u.control.PublishCloseConnection(ctx, realm, connection.NodeID, connection.ConnectionID, "connection limit exceeded"); err != nil {
			u.log.WarnContext(ctx, "publish realtime close connection control failed", "realm", realm, "node_id", connection.NodeID, "connection_id", connection.ConnectionID, "error", err)
		}
	}
	u.DeliverConnectionBound(ctx, realm, nodeID, connectionID)
	return nil
}

func (u *RealtimeUsecase) UnbindConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) error {
	return u.registry.UnbindConnection(ctx, nodeID, realm, userID, connectionID)
}

func (u *RealtimeUsecase) RenewConnection(ctx context.Context, nodeID, realm string, userID int64, connectionID string) error {
	return u.registry.RenewConnection(ctx, nodeID, realm, userID, connectionID)
}

func (u *RealtimeUsecase) QueryOnline(ctx context.Context, realm string, userIDs []int64) (map[int64]bool, error) {
	if realm == "" {
		return nil, fmt.Errorf("realtime realm is required")
	}
	return u.registry.QueryOnline(ctx, realm, userIDs)
}

func (u *RealtimeUsecase) ListOnlineUsers(ctx context.Context, realm string) ([]int64, error) {
	if realm == "" {
		return nil, fmt.Errorf("realtime realm is required")
	}
	return u.registry.ListOnlineUsers(ctx, realm)
}

func (u *RealtimeUsecase) IssueTicket(ctx context.Context, realm string, userID int64) (string, error) {
	ticket := uuid.NewString()
	return ticket, u.registry.IssueTicket(ctx, realm, ticket, userID)
}

func (u *RealtimeUsecase) ConsumeTicket(ctx context.Context, realm, ticket string) (int64, error) {
	return u.registry.ConsumeTicket(ctx, realm, ticket)
}

// DeliverNotification 将站内信变化路由到指定用户的本机 node stream。
func (u *RealtimeUsecase) DeliverNotification(ctx context.Context, realm string, userID, notificationID int64) error {
	return u.Deliver(ctx, realm, userID, enum.DeliveryKindNotification, model.NotificationDelivery{NotificationID: notificationID})
}

// DeliverPrivateChatMessage 将私聊变化路由到双方当前的所有本机连接。
func (u *RealtimeUsecase) DeliverPrivateChatMessage(ctx context.Context, realm string, senderID, receiverID, messageID int64, messageType int32, content string, createdAt time.Time) error {
	payload := model.PrivateChatMessageDelivery{MessageID: messageID, SenderID: senderID, ReceiverID: receiverID, MessageType: messageType, Content: content, CreatedAt: createdAt}
	if err := u.Deliver(ctx, realm, senderID, enum.DeliveryKindPrivateChatMessage, payload); err != nil {
		return err
	}
	if receiverID == senderID {
		return nil
	}
	return u.Deliver(ctx, realm, receiverID, enum.DeliveryKindPrivateChatMessage, payload)
}

// DeliverGlobalChatMessage 将全局聊天室变化广播给实时域中的每个本机 node。
func (u *RealtimeUsecase) DeliverGlobalChatMessage(realm string, messageID, senderID int64, messageType int32, content string, createdAt time.Time) {
	message := &model.NodeMessage{Delivery: &model.Delivery{Kind: enum.DeliveryKindGlobalChatMessage, Payload: model.GlobalChatMessageDelivery{MessageID: messageID, SenderID: senderID, MessageType: messageType, Content: content, CreatedAt: createdAt}}}
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, node := range u.nodes[realm] {
		select {
		case node.messages <- message:
		default:
		}
	}
}

// DeliverGlobalChatBarrage 将全局聊天室弹幕广播给实时域中的每个本机 node。
func (u *RealtimeUsecase) DeliverGlobalChatBarrage(realm string, barrageID, userID int64, content string, createdAt time.Time) {
	message := &model.NodeMessage{Delivery: &model.Delivery{
		Kind: enum.DeliveryKindGlobalChatBarrage,
		Payload: model.GlobalChatBarrageDelivery{
			BarrageID: barrageID,
			UserID:    userID,
			Content:   content,
			CreatedAt: createdAt,
		},
	}}
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, node := range u.nodes[realm] {
		select {
		case node.messages <- message:
		default:
		}
	}
}

// DeliverGlobalChatTopic 将话题变化广播给实时域中的每个本机 node。
func (u *RealtimeUsecase) DeliverGlobalChatTopic(realm string, topicID, userID int64, content string, createdAt time.Time) {
	message := &model.NodeMessage{Delivery: &model.Delivery{
		Kind: enum.DeliveryKindGlobalChatTopic,
		Payload: model.GlobalChatTopicDelivery{
			TopicID:   topicID,
			UserID:    userID,
			Content:   content,
			CreatedAt: createdAt,
		},
	}}
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, node := range u.nodes[realm] {
		select {
		case node.messages <- message:
		default:
		}
	}
}

func (u *RealtimeUsecase) Deliver(ctx context.Context, realm string, userID int64, kind enum.DeliveryKind, payload model.DeliveryPayload) error {
	if payload == nil {
		return fmt.Errorf("delivery payload is required")
	}
	nodeIDs, err := u.registry.UserNodes(ctx, realm, userID)
	if err != nil {
		return err
	}
	message := &model.NodeMessage{Delivery: &model.Delivery{UserID: userID, Kind: kind, Payload: payload}}
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, nodeID := range nodeIDs {
		if node := u.nodes[realm][nodeID]; node != nil {
			select {
			case node.messages <- message:
			default:
			}
		}
	}
	return nil
}

func (u *RealtimeUsecase) DeliverCloseConnection(ctx context.Context, realm, nodeID, connectionID, reason string) bool {
	u.mu.RLock()
	node := u.nodes[realm][nodeID]
	u.mu.RUnlock()
	if node == nil {
		return false
	}
	select {
	case node.messages <- &model.NodeMessage{CloseConnection: &model.CloseConnection{ConnectionID: connectionID, Reason: reason}}:
		return true
	case <-ctx.Done():
	case <-node.ctx.Done():
	}
	return false
}

func (u *RealtimeUsecase) DeliverConnectionBound(ctx context.Context, realm, nodeID, connectionID string) {
	u.mu.RLock()
	node := u.nodes[realm][nodeID]
	u.mu.RUnlock()
	if node == nil {
		return
	}
	select {
	case node.messages <- &model.NodeMessage{ConnectionBound: &model.ConnectionBound{ConnectionID: connectionID}}:
	case <-ctx.Done():
	case <-node.ctx.Done():
	}
}

func (u *RealtimeUsecase) removeLocal(nodeID string, realms []string) {
	u.mu.Lock()
	defer u.mu.Unlock()
	for _, realm := range realms {
		delete(u.nodes[realm], nodeID)
	}
}
