package usecase

import (
	"common/pkg/apperror"
	commonclient "common/pkg/client"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
	"game_idle_bff/internal/config"
	"game_idle_bff/internal/enum"
	"game_idle_bff/internal/errormessage"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

type WebSocketUsecase struct {
	logger           *slog.Logger
	eventRepo        repo.WebSocketEventRepo
	ticketRepo       repo.WebSocketTicketRepo
	workerPool       *commonclient.WorkerPool
	characterUsecase *CharacterUsecase
	eventHandlers    WebSocketEventHandlers
	commandHandlers  WebSocketCommandHandlers
	pingInterval     time.Duration
	writeTimeout     time.Duration
	ticketTTL        time.Duration
	sendBufferSize   int
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
	ticketRepo repo.WebSocketTicketRepo,
	workerPool *commonclient.WorkerPool,
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
		ticketRepo:       ticketRepo,
		workerPool:       workerPool,
		characterUsecase: characterUsecase,
		eventHandlers:    eventHandlers,
		commandHandlers:  commandHandlers,
		pingInterval:     pingInterval,
		writeTimeout:     writeTimeout,
		ticketTTL:        2 * time.Minute,
		sendBufferSize:   4,
		characters:       map[int64]map[string]*WebSocketConnection{},
		sessions:         map[string]*WebSocketConnection{},
	}
}

type WebSocketConnection struct {
	CharacterID  int64
	Ticket       string
	SessionID    string
	ExpiresIn    time.Duration
	Online       bool
	Disconnected bool
	Messages     chan *WebSocketSendMessage
	Closed       chan struct{}
	SendTimeout  time.Duration
	closeOnce    sync.Once
	lock         sync.RWMutex
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

type CreateWebSocketTicketReq struct {
	UserID      int64
	CharacterID int64
}

func (u *WebSocketUsecase) CreateTicket(ctx context.Context, req *CreateWebSocketTicketReq) (*model.WebSocketTicket, error) {
	characters, err := u.characterUsecase.List(ctx, &ListCharacterReq{
		UserID:      req.UserID,
		CharacterID: req.CharacterID,
	})
	if err != nil {
		return nil, err
	}
	if len(characters) == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	ticket := &model.WebSocketTicket{
		CharacterID:       req.CharacterID,
		Ticket:            uuid.NewString(),
		RemainingDuration: u.ticketTTL,
	}
	if err = u.ticketRepo.Save(ctx, ticket.CharacterID, ticket.Ticket, u.ticketTTL); err != nil {
		return nil, err
	}
	return ticket, nil
}

func (u *WebSocketUsecase) ConsumeTicket(ctx context.Context, characterID int64, ticket string) (*model.WebSocketTicket, error) {
	ok, err := u.ticketRepo.Consume(ctx, characterID, ticket)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_SESSION_INVALID)
	}
	return &model.WebSocketTicket{
		CharacterID:       characterID,
		Ticket:            ticket,
		RemainingDuration: 0,
	}, nil
}

func (u *WebSocketUsecase) Ping(ctx context.Context, characterID int64, sessionID string) (*model.CharacterOnlineSession, error) {
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
		err := u.workerPool.Submit(func() {
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
	connections := make([]*WebSocketConnection, 0, len(u.sessions))
	seen := map[*WebSocketConnection]struct{}{}
	for _, connection := range u.sessions {
		seen[connection] = struct{}{}
		connections = append(connections, connection)
	}
	for _, rows := range u.characters {
		for _, connection := range rows {
			if _, ok := seen[connection]; !ok {
				seen[connection] = struct{}{}
				connections = append(connections, connection)
			}
		}
	}
	u.running = false
	u.lock.Unlock()
	for _, connection := range connections {
		u.Disconnect(ctx, connection, false)
	}
	return nil
}

func (u *WebSocketUsecase) Connect(ctx context.Context, characterID int64, ticket string) *WebSocketConnection {
	connection := &WebSocketConnection{
		CharacterID: characterID,
		Ticket:      ticket,
		Messages:    make(chan *WebSocketSendMessage, u.sendBufferSize),
		Closed:      make(chan struct{}),
		SendTimeout: u.writeTimeout,
	}
	u.lock.Lock()
	if u.characters[characterID] == nil {
		u.characters[characterID] = map[string]*WebSocketConnection{}
	}
	if old := u.characters[characterID][connection.Ticket]; old != nil {
		old.Close()
	}
	u.characters[characterID][connection.Ticket] = connection
	u.lock.Unlock()
	return connection
}

func (u *WebSocketUsecase) Disconnect(ctx context.Context, connection *WebSocketConnection, timeout bool) {
	sessionID, online := connection.MarkDisconnected()
	characterKey := connection.Ticket
	if online {
		characterKey = sessionID
	}
	u.lock.Lock()
	if online && u.sessions[sessionID] != connection {
		u.lock.Unlock()
		return
	}
	if online {
		delete(u.sessions, sessionID)
	}
	if rows := u.characters[connection.CharacterID]; rows != nil {
		delete(rows, characterKey)
		if len(rows) == 0 {
			delete(u.characters, connection.CharacterID)
		}
	}
	connection.Close()
	u.lock.Unlock()

	if online {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = u.characterUsecase.Offline(cleanupCtx, &OfflineCharacterReq{
			CharacterID: connection.CharacterID,
			SessionID:   sessionID,
			Timeout:     timeout,
		})
	}
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

func (c *WebSocketConnection) OnlineSession() (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.SessionID, c.Online && c.SessionID != ""
}

func (c *WebSocketConnection) OnlineSessionInfo() (string, time.Duration, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.SessionID, c.ExpiresIn, c.Online && c.SessionID != ""
}

func (c *WebSocketConnection) MarkDisconnected() (string, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.Disconnected {
		return c.SessionID, false
	}
	c.Disconnected = true
	return c.SessionID, c.Online && c.SessionID != ""
}

func (u *WebSocketUsecase) BindOnline(ctx context.Context, connection *WebSocketConnection, session *model.CharacterOnlineSession) bool {
	connection.lock.Lock()
	if connection.Disconnected {
		connection.lock.Unlock()
		return false
	}
	oldSessionID := connection.SessionID
	connection.SessionID = session.SessionID
	connection.ExpiresIn = session.RemainingDuration
	connection.Online = true
	connection.lock.Unlock()

	u.lock.Lock()
	if oldSessionID != "" {
		delete(u.sessions, oldSessionID)
	}
	if old := u.sessions[session.SessionID]; old != nil && old != connection {
		old.Close()
	}
	u.sessions[session.SessionID] = connection
	if rows := u.characters[connection.CharacterID]; rows != nil {
		delete(rows, connection.Ticket)
		rows[session.SessionID] = connection
	}
	u.lock.Unlock()
	return true
}

func (u *WebSocketUsecase) HandleCommand(ctx context.Context, connection *WebSocketConnection, commandType enum.WebSocketMessageType, payloadData json.RawMessage) {
	handler, ok := u.commandHandlers[commandType]
	if !ok {
		u.SendCommandFailed(ctx, connection, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT))
		return
	}
	if commandType != enum.WebSocketMessageTypeCharacterOnline {
		if _, online := connection.OnlineSession(); !online {
			u.SendCommandFailed(ctx, connection, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_SESSION_INVALID))
			return
		}
	}
	payload := handler.Payload()
	if payload != nil {
		if len(payloadData) == 0 {
			u.SendCommandFailed(ctx, connection, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT))
			return
		}
		if err := (protojson.UnmarshalOptions{}).Unmarshal(payloadData, payload); err != nil {
			u.SendCommandFailed(ctx, connection, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT))
			return
		}
	}
	err := u.workerPool.Submit(func() {
		sessionID, _ := connection.OnlineSession()
		if err := handler.Handle(ctx, &WebSocketCommandReq{
			CharacterID: connection.CharacterID,
			SessionID:   sessionID,
			Connection:  connection,
			Payload:     payload,
			BindOnline:  u.BindOnline,
		}); err != nil {
			u.SendCommandFailed(ctx, connection, err)
		}
	})
	if err != nil {
		u.SendCommandFailed(ctx, connection, err)
	}
}

func (u *WebSocketUsecase) SendCommandFailed(ctx context.Context, connection *WebSocketConnection, err error) {
	code, message := errormessage.ResolveError(err)
	connection.Send(ctx, &WebSocketSendMessage{
		Type: enum.WebSocketMessageTypeCommandFailed,
		Payload: &WebSocketCommandError{
			Code:    int(code),
			Message: message,
		},
	})
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
			slog.String("session_id", message.TargetSessionID),
		)
	}
}
