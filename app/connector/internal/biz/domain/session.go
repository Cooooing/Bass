package domain

import (
	signalv1 "common/api/gen/signal/v1"
	"common/pkg/util"
	domainbase "connector/internal/biz/base"
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/samber/lo"
)

type SessionDomain struct {
	*domainbase.BaseDomain
	eventPool *util.EventPool

	sessionIds map[string]*websocket.Conn
	mu         sync.RWMutex
}

func NewSessionDomain(baseDomain *domainbase.BaseDomain, eventPool *util.EventPool) *SessionDomain {
	return &SessionDomain{
		BaseDomain: baseDomain,
		eventPool:  eventPool,
		sessionIds: map[string]*websocket.Conn{},
		mu:         sync.RWMutex{},
	}
}

func (d *SessionDomain) AddSessionId(sessionId string, conn *websocket.Conn) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.sessionIds[sessionId] = conn
}

func (d *SessionDomain) RemoveSessionId(sessionId string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.sessionIds, sessionId)
}

func (d *SessionDomain) GetSessionIds() []string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return lo.Keys(d.sessionIds)
}

func (d *SessionDomain) SendMessage(sessionIds []string, message any) error {
	if len(sessionIds) == 0 {
		return nil
	}
	d.mu.RLock()
	for _, sessionId := range sessionIds {
		if conn, ok := d.sessionIds[sessionId]; ok {
			return d.EventPool.Submit(func() {
				err := conn.WriteJSON(message)
				if err != nil {
					d.Log.Error("websocket [%s] send message error: %v", sessionId, err)
				}
			})
		}
	}
	d.mu.RUnlock()
	return nil
}

func (d *SessionDomain) RequestSessionId(ticket string) (string, error) {
	reply, err := d.SignalNodeClient.Online(context.Background(), &signalv1.SignalNodeOnlineRequest{Ticket: ticket})
	if err != nil {
		return "", err
	}
	return reply.SessionId, nil
}
