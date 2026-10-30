package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	imv1enum "common/proto/gen/im/v1/enum"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"
	realtimehubv1enum "common/proto/gen/realtime_hub/v1/enum"
	"realtime_node/internal/biz/model"
	"realtime_node/internal/biz/repo"
	"realtime_node/internal/enum"

	"github.com/google/uuid"
)

type RealtimeUsecase struct {
	log      *slog.Logger
	hub      repo.Hub
	realm    string
	mu       sync.RWMutex
	users    map[int64]map[string]*connection
	byID     map[string]*connection
	outbound chan *realtimehubv1.NodeToHub
	runCtx   context.Context
	cancel   context.CancelFunc
	stream   *streamState
}

type connection struct {
	userID     int64
	deliveries chan *realtimehubv1.Delivery
	cancel     context.CancelFunc
	bound      chan struct{}
	boundOnce  sync.Once
}

type streamState struct {
	done chan struct{}
}

func NewRealtimeUsecase(log *slog.Logger, hub repo.Hub) *RealtimeUsecase {
	return &RealtimeUsecase{
		log:      log,
		hub:      hub,
		realm:    "bbs",
		users:    map[int64]map[string]*connection{},
		byID:     map[string]*connection{},
		outbound: make(chan *realtimehubv1.NodeToHub, 512),
	}
}

func (u *RealtimeUsecase) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	u.mu.Lock()
	u.runCtx = ctx
	u.cancel = cancel
	u.mu.Unlock()
	go u.maintainStream(ctx)
	return nil
}

func (u *RealtimeUsecase) Stop(context.Context) error {
	u.mu.RLock()
	cancel := u.cancel
	u.mu.RUnlock()
	if cancel != nil {
		cancel()
	}
	return nil
}

func (u *RealtimeUsecase) Authenticate(ctx context.Context, ticket string) (int64, error) {
	return u.hub.ConsumeTicket(ctx, u.realm, ticket)
}

func (u *RealtimeUsecase) Connect(ctx context.Context, userID int64, connectionID string, deliveries chan *realtimehubv1.Delivery, cancel context.CancelFunc) error {
	if cancel == nil {
		return fmt.Errorf("connection cancel is required")
	}
	clientConnection := &connection{userID: userID, deliveries: deliveries, cancel: cancel, bound: make(chan struct{})}
	u.mu.Lock()
	stream := u.stream
	if stream == nil {
		u.mu.Unlock()
		return fmt.Errorf("realtime hub stream is unavailable")
	}
	if u.users[userID] == nil {
		u.users[userID] = map[string]*connection{}
	}
	u.users[userID][connectionID] = clientConnection
	u.byID[connectionID] = clientConnection
	u.mu.Unlock()
	bind := &realtimehubv1.NodeToHub{Payload: &realtimehubv1.NodeToHub_BindConnection{BindConnection: &realtimehubv1.BindConnection{Realm: u.realm, UserId: userID, ConnectionId: connectionID}}}
	if err := u.enqueue(ctx, bind); err != nil {
		u.removeConnection(userID, connectionID)
		return err
	}
	select {
	case <-clientConnection.bound:
		return nil
	case <-ctx.Done():
		u.Disconnect(userID, connectionID)
		return ctx.Err()
	case <-stream.done:
		u.removeConnection(userID, connectionID)
		return fmt.Errorf("realtime hub stream ended before connection binding completed")
	}
}

func (u *RealtimeUsecase) Disconnect(userID int64, connectionID string) {
	if !u.removeConnection(userID, connectionID) {
		return
	}
	unbind := &realtimehubv1.NodeToHub{Payload: &realtimehubv1.NodeToHub_UnbindConnection{UnbindConnection: &realtimehubv1.UnbindConnection{Realm: u.realm, UserId: userID, ConnectionId: connectionID}}}
	u.mu.RLock()
	ctx := u.runCtx
	u.mu.RUnlock()
	if ctx != nil {
		_ = u.enqueue(ctx, unbind)
	}
}

// RenewConnection 续期仍可写入 SSE 的客户端连接租约。
func (u *RealtimeUsecase) RenewConnection(ctx context.Context, userID int64, connectionID string) error {
	u.mu.RLock()
	connection := u.byID[connectionID]
	u.mu.RUnlock()
	if connection == nil || connection.userID != userID {
		return fmt.Errorf("realtime connection is unavailable")
	}
	renew := &realtimehubv1.NodeToHub{Payload: &realtimehubv1.NodeToHub_RenewConnection{RenewConnection: &realtimehubv1.RenewConnection{Realm: u.realm, UserId: userID, ConnectionId: connectionID}}}
	return u.enqueue(ctx, renew)
}

func (u *RealtimeUsecase) maintainStream(ctx context.Context) {
	for ctx.Err() == nil {
		if err := u.runStream(ctx); err != nil && ctx.Err() == nil {
			u.log.WarnContext(ctx, "realtime hub stream ended; retrying", "retry_after", time.Second, "error", err)
			timer := time.NewTimer(time.Second)
			select {
			case <-ctx.Done():
				timer.Stop()
			case <-timer.C:
			}
		}
	}
}

func (u *RealtimeUsecase) runStream(ctx context.Context) error {
	stream, err := u.hub.Connect(ctx)
	if err != nil {
		return err
	}
	nodeID := "realtime-node:" + uuid.NewString()
	group, ok := enum.NodeGroupMap.ToProto(enum.NodeGroupInternal)
	if !ok {
		return fmt.Errorf("node group is invalid")
	}
	if err := stream.Send(&realtimehubv1.NodeToHub{Payload: &realtimehubv1.NodeToHub_RegisterNode{RegisterNode: &realtimehubv1.RegisterNode{NodeId: nodeID, Realms: []string{u.realm}, Group: group}}}); err != nil {
		return err
	}
	u.log.InfoContext(ctx, "realtime hub stream connected", "node_id", nodeID, "realm", u.realm)
	state := &streamState{done: make(chan struct{})}
	u.mu.Lock()
	u.stream = state
	u.mu.Unlock()
	defer func() {
		u.mu.Lock()
		if u.stream == state {
			u.stream = nil
			close(state.done)
		}
		u.mu.Unlock()
	}()
	errs := make(chan error, 1)
	go u.receiveMessages(ctx, stream, errs)
	u.mu.RLock()
	messages := make([]*realtimehubv1.NodeToHub, 0, len(u.byID))
	for connectionID, connection := range u.byID {
		messages = append(messages, &realtimehubv1.NodeToHub{Payload: &realtimehubv1.NodeToHub_BindConnection{BindConnection: &realtimehubv1.BindConnection{Realm: u.realm, UserId: connection.userID, ConnectionId: connectionID}}})
	}
	u.mu.RUnlock()
	for _, message := range messages {
		if err := stream.Send(message); err != nil {
			return err
		}
	}
	heartbeats := time.NewTicker(10 * time.Second)
	defer heartbeats.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-errs:
			return err
		case <-heartbeats.C:
			if err := stream.Send(&realtimehubv1.NodeToHub{Payload: &realtimehubv1.NodeToHub_Heartbeat{Heartbeat: &realtimehubv1.Heartbeat{}}}); err != nil {
				return err
			}
		case message := <-u.outbound:
			if bind := message.GetBindConnection(); bind != nil {
				u.mu.RLock()
				connection := u.byID[bind.GetConnectionId()]
				u.mu.RUnlock()
				if connection == nil || connection.userID != bind.GetUserId() {
					continue
				}
			}
			if err := stream.Send(message); err != nil {
				return err
			}
		}
	}
}

func (u *RealtimeUsecase) receiveMessages(ctx context.Context, stream realtimehubv1.RealtimeHubService_ConnectClient, errs chan<- error) {
	for {
		message, err := stream.Recv()
		if err != nil {
			errs <- err
			return
		}
		if ctx.Err() == nil {
			u.handleHubMessage(message)
		}
	}
}

func (u *RealtimeUsecase) handleHubMessage(message *realtimehubv1.HubToNode) {
	if delivery := message.GetDelivery(); delivery != nil {
		u.dispatch(delivery)
		return
	}
	if closeConnection := message.GetControl().GetCloseConnection(); closeConnection != nil {
		u.mu.RLock()
		connection := u.byID[closeConnection.GetConnectionId()]
		u.mu.RUnlock()
		if connection != nil {
			connection.cancel()
		}
		return
	}
	if bound := message.GetControl().GetConnectionBound(); bound != nil {
		u.mu.RLock()
		connection := u.byID[bound.GetConnectionId()]
		u.mu.RUnlock()
		if connection != nil {
			connection.boundOnce.Do(func() { close(connection.bound) })
		}
	}
}

func (u *RealtimeUsecase) dispatch(delivery *realtimehubv1.Delivery) {
	u.mu.RLock()
	connections := make([]chan *realtimehubv1.Delivery, 0)
	if delivery.GetKind() == realtimehubv1enum.DeliveryKind_DELIVERY_KIND_GLOBAL_CHAT_MESSAGE || delivery.GetKind() == realtimehubv1enum.DeliveryKind_DELIVERY_KIND_GLOBAL_CHAT_BARRAGE || delivery.GetKind() == realtimehubv1enum.DeliveryKind_DELIVERY_KIND_GLOBAL_CHAT_TOPIC {
		for _, userConnections := range u.users {
			for _, connection := range userConnections {
				connections = append(connections, connection.deliveries)
			}
		}
	} else {
		connections = make([]chan *realtimehubv1.Delivery, 0, len(u.users[delivery.GetUserId()]))
		for _, connection := range u.users[delivery.GetUserId()] {
			connections = append(connections, connection.deliveries)
		}
	}
	u.mu.RUnlock()
	for _, deliveries := range connections {
		select {
		case deliveries <- delivery:
		default:
		}
	}
}

func (*RealtimeUsecase) ClientEvent(delivery *realtimehubv1.Delivery) (*model.ClientEvent, bool) {
	if delivery == nil {
		return nil, false
	}
	kind, ok := enum.DeliveryKindMap.ToEnum(delivery.GetKind())
	if !ok {
		return nil, false
	}
	switch kind {
	case enum.DeliveryKindNotification:
		notification := delivery.GetNotification()
		if notification == nil || notification.GetNotificationId() <= 0 {
			return nil, false
		}
		return &model.ClientEvent{Name: string(kind), Payload: model.NotificationEventPayload{NotificationID: notification.GetNotificationId()}}, true
	case enum.DeliveryKindPrivateChatMessage:
		message := delivery.GetPrivateChatMessage()
		if message == nil || message.GetMessageId() <= 0 || message.GetSenderId() <= 0 || message.GetReceiverId() <= 0 {
			return nil, false
		}
		if message.GetMessageType() == imv1enum.IMMessageType_IM_MESSAGE_TYPE_UNSPECIFIED {
			return nil, false
		}
		createdAt := ""
		if message.GetCreatedAt() != nil {
			createdAt = message.GetCreatedAt().AsTime().Format(time.RFC3339Nano)
		}
		return &model.ClientEvent{Name: string(kind), Payload: model.PrivateChatMessageEventPayload{MessageID: message.GetMessageId(), SenderID: message.GetSenderId(), ReceiverID: message.GetReceiverId(), MessageType: message.GetMessageType().String(), Content: message.GetContent(), CreatedAt: createdAt}}, true
	case enum.DeliveryKindGlobalChatMessage:
		message := delivery.GetGlobalChatMessage()
		if message == nil || message.GetMessageId() <= 0 {
			return nil, false
		}
		if message.GetMessageType() == imv1enum.IMMessageType_IM_MESSAGE_TYPE_UNSPECIFIED || message.GetSenderId() <= 0 {
			return nil, false
		}
		createdAt := ""
		if message.GetCreatedAt() != nil {
			createdAt = message.GetCreatedAt().AsTime().Format(time.RFC3339Nano)
		}
		return &model.ClientEvent{Name: string(kind), Payload: model.GlobalChatMessageEventPayload{MessageID: message.GetMessageId(), SenderID: message.GetSenderId(), MessageType: message.GetMessageType().String(), Content: message.GetContent(), CreatedAt: createdAt}}, true
	case enum.DeliveryKindGlobalChatBarrage:
		barrage := delivery.GetGlobalChatBarrage()
		if barrage == nil || barrage.GetBarrageId() <= 0 || barrage.GetUserId() <= 0 || barrage.GetContent() == "" {
			return nil, false
		}
		createdAt := ""
		if barrage.GetCreatedAt() != nil {
			createdAt = barrage.GetCreatedAt().AsTime().Format(time.RFC3339Nano)
		}
		return &model.ClientEvent{Name: string(kind), Payload: model.GlobalChatBarrageEventPayload{BarrageID: barrage.GetBarrageId(), UserID: barrage.GetUserId(), Content: barrage.GetContent(), CreatedAt: createdAt}}, true
	case enum.DeliveryKindGlobalChatTopic:
		topic := delivery.GetGlobalChatTopic()
		if topic == nil || topic.GetTopicId() <= 0 || topic.GetUserId() <= 0 || topic.GetContent() == "" {
			return nil, false
		}
		createdAt := ""
		if topic.GetCreatedAt() != nil {
			createdAt = topic.GetCreatedAt().AsTime().Format(time.RFC3339Nano)
		}
		return &model.ClientEvent{Name: string(kind), Payload: model.GlobalChatTopicEventPayload{TopicID: topic.GetTopicId(), UserID: topic.GetUserId(), Content: topic.GetContent(), CreatedAt: createdAt}}, true
	default:
		return nil, false
	}
}

func (u *RealtimeUsecase) enqueue(ctx context.Context, message *realtimehubv1.NodeToHub) error {
	if ctx == nil {
		return fmt.Errorf("realtime node context is required")
	}
	u.mu.RLock()
	runCtx := u.runCtx
	u.mu.RUnlock()
	if runCtx == nil {
		return fmt.Errorf("realtime node is not running")
	}
	select {
	case u.outbound <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-runCtx.Done():
		return context.Cause(runCtx)
	}
}

func (u *RealtimeUsecase) removeConnection(userID int64, connectionID string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	connection := u.byID[connectionID]
	if connection == nil || connection.userID != userID {
		return false
	}
	delete(u.byID, connectionID)
	delete(u.users[userID], connectionID)
	if len(u.users[userID]) == 0 {
		delete(u.users, userID)
	}
	return true
}
