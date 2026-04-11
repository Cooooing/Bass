package service

import (
	"common/pkg/util"
	"connector/internal/biz/domain"
	"connector/internal/biz/model"
	http2 "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
)

type WebsocketService struct {
	*BaseService
	sessionDomain *domain.SessionDomain
	eventPool     *util.EventPool
	upgrader      websocket.Upgrader
}

func NewWebsocketService(baseService *BaseService, sessionDomain *domain.SessionDomain, eventPool *util.EventPool) *WebsocketService {
	return &WebsocketService{
		BaseService:   baseService,
		sessionDomain: sessionDomain,
		eventPool:     eventPool,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// 生产环境请校验 origin
				return true
			},
		},
	}
}

func (s *WebsocketService) Handler() http2.HandlerFunc {
	return func(w http.ResponseWriter, r *http2.Request) {
		// 升级 Websocket
		conn, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			s.Log.Errorf("websocket upgrade failed: %v", err)
			return
		}
		defer func() {
			_ = conn.Close()
		}()

		ticket := r.URL.Query().Get("ticket")
		if ticket == "" {
			_ = conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "invalid ticket"),
				time.Now().Add(time.Second),
			)
			return
		}

		sessionId, err := s.sessionDomain.RequestSessionId(ticket)
		if err != nil {
			s.Log.Errorf("request session id error: %v", err)
			_ = conn.WriteControl(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "invalid ticket"),
				time.Now().Add(time.Second),
			)
			return
		}
		s.Log.Infof("websocket [%s] connected", sessionId)

		// 设置连接参数
		conn.SetReadLimit(1024 * 1024) // 1MB
		_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		conn.SetPongHandler(func(string) error {
			_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
			return nil
		})

		connection := model.NewConnection(conn)
		defer connection.Close()
		go connection.Listen()
		s.sessionDomain.AddSessionId(sessionId, connection)
		defer func() {
			s.sessionDomain.RemoveSessionId(sessionId)
		}()

		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				s.Log.Errorf("websocket read message error: %v", err)
				return
			}
			_ = conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		}
	}
}
