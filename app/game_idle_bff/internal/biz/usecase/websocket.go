package usecase

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonutil "common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
	"game_idle_bff/internal/config"
	"game_idle_bff/internal/enum"
	"log/slog"
	"sync"
	"time"
)

type WebSocketUsecase struct {
	logger           *slog.Logger
	eventRepo        repo.WebSocketEventRepo
	eventPool        *commonutil.EventPool
	characterUsecase *CharacterUsecase
	eventHandlers    WebSocketEventHandlers
	commandHandlers  WebSocketCommandHandlers
	pingInterval     time.Duration
	writeTimeout     time.Duration
	characters       map[int64]map[string]*WebSocketConnection
	sessions         map[string]*WebSocketConnection
	subscription     repo.WebSocketEventSubscription
	cancel           context.CancelFunc
	lock             sync.RWMutex
	running          bool
}

func NewWebSocketUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	eventRepo repo.WebSocketEventRepo,
	eventPool *commonutil.EventPool,
	characterUsecase *CharacterUsecase,
	eventHandlers WebSocketEventHandlers,
	commandHandlers WebSocketCommandHandlers,
) *WebSocketUsecase {
	pingInterval := 30 * time.Second
	if conf.GetWebsocket().GetPingInterval() != nil && conf.GetWebsocket().GetPingInterval().AsDuration() > 0 {
		pingInterval = conf.GetWebsocket().GetPingInterval().AsDuration()
	}
	writeTimeout := time.Second
	if conf.GetWebsocket().GetWriteTimeout() != nil && conf.GetWebsocket().GetWriteTimeout().AsDuration() > 0 {
		writeTimeout = conf.GetWebsocket().GetWriteTimeout().AsDuration()
	}
	return &WebSocketUsecase{
		logger:           logger,
		eventRepo:        eventRepo,
		eventPool:        eventPool,
		characterUsecase: characterUsecase,
		eventHandlers:    eventHandlers,
		commandHandlers:  commandHandlers,
		pingInterval:     pingInterval,
		writeTimeout:     writeTimeout,
		characters:       map[int64]map[string]*WebSocketConnection{},
		sessions:         map[string]*WebSocketConnection{},
	}
}

const webSocketConnectionSendBufferSize = 4

type WebSocketConnection struct {
	CharacterID int64
	SessionID   string
	Messages    chan *WebSocketSendMessage
	Closed      chan struct{}
	SendTimeout time.Duration
	closeOnce   sync.Once
}

type WebSocketSendMessage struct {
	TargetCharacterID int64
	TargetSessionID   string
	Broadcast         bool
	Type              enum.WebSocketMessageType
	Payload           any
	Close             bool
	SilentClose       bool
}

type CreateWebSocketSessionReq struct {
	UserID      int64
	CharacterID int64
}

func (u *WebSocketUsecase) CreateSession(ctx context.Context, req *CreateWebSocketSessionReq) (*model.WebSocketSession, error) {
	return u.characterUsecase.Online(ctx, &OnlineCharacterReq{
		UserID:      req.UserID,
		CharacterID: req.CharacterID,
	})
}

func (u *WebSocketUsecase) Ping(ctx context.Context, characterID int64, sessionID string) (*model.WebSocketSession, error) {
	return u.characterUsecase.Ping(ctx, &PingCharacterReq{
		CharacterID: characterID,
		SessionID:   sessionID,
	})
}

func (u *WebSocketUsecase) PingInterval(ctx context.Context) time.Duration {
	return u.pingInterval
}

func (u *WebSocketUsecase) WriteTimeout(ctx context.Context) time.Duration {
	return u.writeTimeout
}

func (u *WebSocketUsecase) Start(ctx context.Context) error {
	u.lock.Lock()
	if u.running {
		u.lock.Unlock()
		return nil
	}
	runCtx, cancel := context.WithCancel(ctx)
	subscription, err := u.eventRepo.Subscribe(runCtx, func(ctx context.Context, event *model.WebSocketEvent) error {
		err := u.eventPool.Submit(func() {
			result, err := u.HandleEvent(ctx, event)
			if err != nil {
				u.logger.Error("game idle bff websocket event handle failed", constant.LogFieldErr, err)
				return
			}
			if result == nil {
				return
			}
			u.Deliver(ctx, &WebSocketSendMessage{
				TargetCharacterID: result.TargetCharacterID,
				TargetSessionID:   result.TargetSessionID,
				Broadcast:         result.Broadcast,
				Type:              result.Type,
				Payload:           result.Payload,
				Close:             result.Close,
				SilentClose:       result.SilentClose,
			})
		})
		if err != nil {
			u.logger.Warn("game idle bff websocket event pool full", constant.LogFieldErr, err)
		}
		return nil
	})
	if err != nil {
		cancel()
		u.lock.Unlock()
		return err
	}
	u.subscription = subscription
	u.cancel = cancel
	u.running = true
	u.lock.Unlock()
	go u.consumeEvents(runCtx, subscription)
	return nil
}

func (u *WebSocketUsecase) Stop(ctx context.Context) error {
	u.lock.Lock()
	if u.cancel != nil {
		u.cancel()
	}
	if u.subscription != nil {
		if err := u.subscription.Unsubscribe(); err != nil {
			u.logger.Error("game idle bff websocket unsubscribe failed", constant.LogFieldErr, err)
		}
		u.subscription = nil
	}
	for _, connection := range u.sessions {
		connection.Close()
	}
	u.characters = map[int64]map[string]*WebSocketConnection{}
	u.sessions = map[string]*WebSocketConnection{}
	u.running = false
	u.lock.Unlock()
	return nil
}

func (u *WebSocketUsecase) Connect(ctx context.Context, characterID int64, sessionID string) *WebSocketConnection {
	connection := &WebSocketConnection{
		CharacterID: characterID,
		SessionID:   sessionID,
		Messages:    make(chan *WebSocketSendMessage, webSocketConnectionSendBufferSize),
		Closed:      make(chan struct{}),
		SendTimeout: u.writeTimeout,
	}
	u.lock.Lock()
	if old := u.sessions[sessionID]; old != nil {
		old.Close()
		if rows := u.characters[old.CharacterID]; rows != nil {
			delete(rows, old.SessionID)
			if len(rows) == 0 {
				delete(u.characters, old.CharacterID)
			}
		}
	}
	u.sessions[connection.SessionID] = connection
	if u.characters[characterID] == nil {
		u.characters[characterID] = map[string]*WebSocketConnection{}
	}
	u.characters[characterID][connection.SessionID] = connection
	u.lock.Unlock()
	return connection
}

func (u *WebSocketUsecase) Disconnect(ctx context.Context, connection *WebSocketConnection, timeout bool) {
	u.lock.Lock()
	if u.sessions[connection.SessionID] != connection {
		u.lock.Unlock()
		return
	}
	delete(u.sessions, connection.SessionID)
	if rows := u.characters[connection.CharacterID]; rows != nil {
		delete(rows, connection.SessionID)
		if len(rows) == 0 {
			delete(u.characters, connection.CharacterID)
		}
	}
	connection.Close()
	u.lock.Unlock()

	cleanupCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_ = u.characterUsecase.Offline(cleanupCtx, &OfflineCharacterReq{
		CharacterID: connection.CharacterID,
		SessionID:   connection.SessionID,
		Timeout:     timeout,
	})
}

func (c *WebSocketConnection) Close() {
	c.closeOnce.Do(func() {
		close(c.Closed)
	})
}

func (c *WebSocketConnection) Send(ctx context.Context, message *WebSocketSendMessage) bool {
	timer := time.NewTimer(c.SendTimeout)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return false
	case <-c.Closed:
		return false
	case c.Messages <- message:
		return true
	case <-timer.C:
		c.Close()
		return false
	}
}

func (u *WebSocketUsecase) HandleCommand(ctx context.Context, connection *WebSocketConnection, command *WebSocketCommand) {
	handler, ok := u.commandHandlers[command.Type]
	if !ok {
		connection.Send(ctx, &WebSocketSendMessage{
			Type: enum.WebSocketMessageTypeCommandFailed,
			Payload: &WebSocketCommandError{
				Message: apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT).Error(),
			},
		})
		return
	}
	err := u.eventPool.Submit(func() {
		if err := handler.Handle(ctx, &WebSocketCommandReq{
			CharacterID: connection.CharacterID,
			SessionID:   connection.SessionID,
			Connection:  connection,
			Command:     command,
		}); err != nil {
			connection.Send(ctx, &WebSocketSendMessage{
				Type: enum.WebSocketMessageTypeCommandFailed,
				Payload: &WebSocketCommandError{
					Message: err.Error(),
				},
			})
		}
	})
	if err != nil {
		connection.Send(ctx, &WebSocketSendMessage{
			Type: enum.WebSocketMessageTypeCommandFailed,
			Payload: &WebSocketCommandError{
				Message: err.Error(),
			},
		})
	}
}

func (u *WebSocketUsecase) consumeEvents(ctx context.Context, subscription repo.WebSocketEventSubscription) {
	<-ctx.Done()
	if err := subscription.Unsubscribe(); err != nil {
		u.logger.Error("game idle bff websocket unsubscribe failed", constant.LogFieldErr, err)
	}
}

func (u *WebSocketUsecase) HandleEvent(ctx context.Context, event *model.WebSocketEvent) (*WebSocketEventResult, error) {
	handler, ok := u.eventHandlers[event.Type]
	if !ok {
		return nil, nil
	}
	return handler.Handle(ctx, &WebSocketEventReq{
		Event: event,
	})
}

func (u *WebSocketUsecase) Deliver(ctx context.Context, message *WebSocketSendMessage) {
	connections := make([]*WebSocketConnection, 0)
	u.lock.RLock()
	if message.Broadcast {
		for _, connection := range u.sessions {
			connections = append(connections, connection)
		}
	} else if message.TargetSessionID != "" {
		if connection := u.sessions[message.TargetSessionID]; connection != nil {
			connections = append(connections, connection)
		}
	} else if message.TargetCharacterID > 0 {
		for _, connection := range u.characters[message.TargetCharacterID] {
			connections = append(connections, connection)
		}
	}
	u.lock.RUnlock()
	for _, connection := range connections {
		u.sendToConnection(ctx, connection, message)
	}
}

func (u *WebSocketUsecase) sendToConnection(ctx context.Context, connection *WebSocketConnection, message *WebSocketSendMessage) {
	if !connection.Send(ctx, message) {
		u.logger.Warn(
			"game idle bff websocket send failed",
			slog.Int64("character_id", connection.CharacterID),
			slog.String("session_id", connection.SessionID),
		)
	}
}
