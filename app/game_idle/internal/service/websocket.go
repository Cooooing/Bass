package service

import (
	"common/pkg/constant"
	"context"
	"game_idle/internal/biz/usecase"
	"log/slog"
	stdhttp "net/http"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	mutex              sync.Mutex
	logger             *slog.Logger
	actionQueueUsecase *usecase.ActionQueueUsecase
	connections        map[*websocket.Conn]struct{}
	upgrader           websocket.Upgrader
	stop               context.CancelFunc
	running            bool
}

type WebSocketActionResultMessage struct {
	Type        string           `json:"type"`
	CharacterID int64            `json:"character_id"`
	ActionID    string           `json:"action_id"`
	Items       map[string]int64 `json:"items"`
}

func NewWebSocketService(
	logger *slog.Logger,
	actionQueueUsecase *usecase.ActionQueueUsecase,
) *WebSocketService {
	return &WebSocketService{
		logger:             logger,
		actionQueueUsecase: actionQueueUsecase,
		connections:        make(map[*websocket.Conn]struct{}),
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

func (s *WebSocketService) RegisterHttp(server *http.Server) {
	server.Handle("/ws", stdhttp.HandlerFunc(s.Handle))
	s.logger.Info("game idle websocket endpoint registered", slog.String(constant.LogFieldPath, "/ws"))
}

func (s *WebSocketService) Start(ctx context.Context) error {
	s.mutex.Lock()
	if s.running {
		s.mutex.Unlock()
		return nil
	}
	runCtx, cancel := context.WithCancel(ctx)
	s.stop = cancel
	s.running = true
	s.mutex.Unlock()

	go func() {
		for {
			select {
			case <-runCtx.Done():
				return
			case event := <-s.actionQueueUsecase.ResultEvents(runCtx):
				message := &WebSocketActionResultMessage{
					Type:        "action_result",
					CharacterID: event.CharacterID,
					ActionID:    event.ActionID,
					Items:       event.Items,
				}
				s.mutex.Lock()
				for conn := range s.connections {
					if err := conn.WriteJSON(message); err != nil {
						s.logger.Error("game idle websocket write failed", constant.LogFieldErr, err)
						conn.Close()
						delete(s.connections, conn)
					}
				}
				s.mutex.Unlock()
			}
		}
	}()

	return nil
}

func (s *WebSocketService) Stop(ctx context.Context) error {
	_ = ctx
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.stop != nil {
		s.stop()
		s.stop = nil
	}
	s.running = false
	for conn := range s.connections {
		conn.Close()
		delete(s.connections, conn)
	}
	return nil
}

func (s *WebSocketService) Handle(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.Error("game idle websocket upgrade failed", constant.LogFieldErr, err)
		return
	}
	s.mutex.Lock()
	s.connections[conn] = struct{}{}
	s.mutex.Unlock()
	s.logger.Info("game idle websocket client connected", constant.LogFieldAddress, r.RemoteAddr)
}
