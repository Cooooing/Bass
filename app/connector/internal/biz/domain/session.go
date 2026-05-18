package domain

import (
	"bytes"
	signalv1 "common/api/gen/signal/v1"
	"common/pkg/client"
	"common/pkg/util/server"
	"connector/internal/biz/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
)

type SessionDomain struct {
	log        *log.Helper
	httpClient *http.Client

	sessionIds map[string]*model.Connection
	mu         sync.RWMutex
}

func NewSessionDomain(
	logger log.Logger,
) *SessionDomain {
	return &SessionDomain{
		log:        log.NewHelper(logger),
		httpClient: client.NewHttpClient(),
		sessionIds: map[string]*model.Connection{},
		mu:         sync.RWMutex{},
	}
}

func (d *SessionDomain) AddSessionId(sessionId string, conn *model.Connection) {
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

func (d *SessionDomain) SendMessage(ctx context.Context, sessionIds []string, message any) error {
	if len(sessionIds) == 0 {
		return nil
	}
	d.mu.RLock()
	for _, sessionId := range sessionIds {
		if conn, ok := d.sessionIds[sessionId]; ok {
			conn.Send(message)
		}
	}
	d.mu.RUnlock()
	return nil
}

func (d *SessionDomain) RequestSessionId(ticket string) (string, error) {
	var sessionId string

	url := fmt.Sprintf("http://127.0.0.1:8000/api/signal/v1/node/online")

	param, err := json.Marshal(&signalv1.OnlineSignalNode_Request{Ticket: ticket})
	if err != nil {
		return sessionId, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(param))
	if err != nil {
		return sessionId, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Signal-NodeKey", "main")

	resp, err := d.httpClient.Do(req)
	if err != nil {
		return sessionId, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	data := &server.Result[*signalv1.OnlineSignalNode_Reply]{}
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return sessionId, err
	}

	if resp.StatusCode != http.StatusOK {
		return sessionId, fmt.Errorf("failed to get session from %s %s", url, data.Msg)
	}

	sessionId = data.Data.SessionId
	return sessionId, nil
}
