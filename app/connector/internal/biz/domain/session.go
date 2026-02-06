package domain

import (
	"bytes"
	signalv1 "common/api/signal/v1"
	"common/pkg"
	"common/pkg/client"
	"common/pkg/cutil/collections/dict"
	"common/pkg/util"
	domainbase "connector/internal/biz/base"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

type SessionDomain struct {
	*domainbase.BaseDomain
	eventPool  *util.EventPool
	sessionIds dict.Map[string, *websocket.Conn]
	httpClient *http.Client
}

func NewSessionDomain(baseDomain *domainbase.BaseDomain, eventPool *util.EventPool) *SessionDomain {
	return &SessionDomain{
		BaseDomain: baseDomain,
		eventPool:  eventPool,
		httpClient: client.NewHttpClient(),
		sessionIds: dict.NewSafeMap[string, *websocket.Conn](0),
	}
}

func (d *SessionDomain) AddSessionId(sessionId string, conn *websocket.Conn) {
	d.sessionIds.Set(sessionId, conn)
}

func (d *SessionDomain) RemoveSessionId(sessionId string) {
	d.sessionIds.Remove(sessionId)
}

func (d *SessionDomain) GetSessionIds() []string {
	return d.sessionIds.Keys()
}

func (d *SessionDomain) SendMessage(sessionIds []string, message any) error {
	for _, sessionId := range sessionIds {
		if conn, ok := d.sessionIds.Get(sessionId); ok {
			return d.EventPool.Submit(func() {
				err := conn.WriteJSON(message)
				if err != nil {
					d.Log.Error("websocket [%s] send message error: %v", sessionId, err)
				}
			})
		}
	}
	return nil
}

func (d *SessionDomain) RequestSessionId(ticket string) (string, error) {
	var sessionId string

	url := fmt.Sprintf("http://127.0.0.1:8000/api/signal/v1/node/online")

	param, err := json.Marshal(&signalv1.SignalNodeOnlineRequest{Ticket: ticket})
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
			d.Log.Errorf("failed to close body: %v", err)
		}
	}(resp.Body)

	data := &pkg.Result[*signalv1.SignalNodeOnlineReply]{}
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
