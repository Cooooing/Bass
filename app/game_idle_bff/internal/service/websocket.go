package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
	"log/slog"
	"net"
	stdhttp "net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/gorilla/websocket"
)

const webSocketPath = "/ws"
const webSocketReadLimit = 64 * 1024

type WebSocketService struct {
	logger           *slog.Logger
	webSocketUsecase *usecase.WebSocketUsecase
	upgrader         websocket.Upgrader
}

type WebSocketResponse struct {
	Type    enum.WebSocketMessageType `json:"type"`
	Payload any                       `json:"payload,omitempty"`
}

type WebSocketRequest struct {
	Type    enum.WebSocketMessageType `json:"type"`
	Payload json.RawMessage           `json:"payload,omitempty"`
}

func NewWebSocketService(
	logger *slog.Logger,
	webSocketUsecase *usecase.WebSocketUsecase,
) *WebSocketService {
	return &WebSocketService{
		logger:           logger,
		webSocketUsecase: webSocketUsecase,
		upgrader: websocket.Upgrader{
			HandshakeTimeout: 5 * time.Second,
			CheckOrigin: func(*stdhttp.Request) bool {
				return true
			},
		},
	}
}

func (s *WebSocketService) RegisterGrpc(*grpc.Server) {
}

func (s *WebSocketService) RegisterHttp(hs *http.Server) {
	hs.Handle(webSocketPath, stdhttp.HandlerFunc(s.Handle))
	s.logger.Info("game idle bff websocket endpoint registered", slog.String(constant.LogFieldPath, webSocketPath))
}

func (s *WebSocketService) Handle(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	ctx := r.Context()
	characterID, err := strconv.ParseInt(r.URL.Query().Get("character_id"), 10, 64)
	if err != nil {
		stdhttp.Error(w, err.Error(), stdhttp.StatusUnauthorized)
		return
	}
	sessionID := r.URL.Query().Get("session_id")
	if sessionID == "" {
		stdhttp.Error(w, "websocket session invalid", stdhttp.StatusUnauthorized)
		return
	}
	if _, err = s.webSocketUsecase.Ping(ctx, characterID, sessionID); err != nil {
		s.logger.Error("game idle bff websocket session check failed", constant.LogFieldErr, err)
		stdhttp.Error(w, err.Error(), stdhttp.StatusUnauthorized)
		return
	}
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error("game idle bff websocket upgrade failed", constant.LogFieldErr, err)
		return
	}
	defer conn.Close()

	connCtx, cancel := context.WithCancel(context.WithoutCancel(ctx))
	defer cancel()

	var timeout atomic.Bool
	connection := s.webSocketUsecase.Connect(connCtx, characterID, sessionID)
	defer s.webSocketUsecase.Disconnect(connCtx, connection, timeout.Load())

	conn.SetReadLimit(webSocketReadLimit)
	_ = conn.SetReadDeadline(time.Now().Add(s.webSocketUsecase.PingInterval(connCtx) + s.webSocketUsecase.WriteTimeout(connCtx)))
	conn.SetPongHandler(func(string) error {
		if _, err := s.webSocketUsecase.Ping(connCtx, characterID, sessionID); err != nil {
			timeout.Store(true)
			cancel()
			return nil
		}
		_ = conn.SetReadDeadline(time.Now().Add(s.webSocketUsecase.PingInterval(connCtx) + s.webSocketUsecase.WriteTimeout(connCtx)))
		return nil
	})

	go func() {
		s.writePump(connCtx, conn, connection, cancel, &timeout)
	}()

	s.readPump(connCtx, conn, connection, cancel, &timeout)
}

func (s *WebSocketService) readPump(
	ctx context.Context,
	conn *websocket.Conn,
	connection *usecase.WebSocketConnection,
	cancel context.CancelFunc,
	timeout *atomic.Bool,
) {
	defer cancel()
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				timeout.Store(true)
			}
			return
		}
		command := &WebSocketRequest{}
		if err := json.Unmarshal(data, command); err != nil {
			s.webSocketUsecase.SendCommandFailed(ctx, connection, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT))
			continue
		}
		s.webSocketUsecase.HandleCommand(ctx, connection, command.Type, command.Payload)
	}
}

func (s *WebSocketService) writePump(
	ctx context.Context,
	conn *websocket.Conn,
	connection *usecase.WebSocketConnection,
	cancel context.CancelFunc,
	timeout *atomic.Bool,
) {
	ticker := time.NewTicker(s.webSocketUsecase.PingInterval(ctx))
	defer ticker.Stop()
	defer conn.Close()
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		case <-connection.Closed:
			return
		case message, ok := <-connection.Messages:
			if !ok {
				return
			}
			if !message.SilentClose {
				_ = conn.SetWriteDeadline(time.Now().Add(s.webSocketUsecase.WriteTimeout(ctx)))
				if err := conn.WriteJSON(&WebSocketResponse{
					Type:    message.Type,
					Payload: message.Payload,
				}); err != nil {
					return
				}
			}
			if message.Close {
				_ = conn.SetWriteDeadline(time.Now().Add(s.webSocketUsecase.WriteTimeout(ctx)))
				_ = conn.WriteControl(
					websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
					time.Now().Add(s.webSocketUsecase.WriteTimeout(ctx)),
				)
				return
			}
		case <-ticker.C:
			_ = conn.SetWriteDeadline(time.Now().Add(s.webSocketUsecase.WriteTimeout(ctx)))
			if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(s.webSocketUsecase.WriteTimeout(ctx))); err != nil {
				timeout.Store(true)
				return
			}
		}
	}
}
