package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"common/pkg/client"
	"common/pkg/util/jwt"
	"push_node/internal/biz/model"
	"push_node/internal/biz/repo"
	"push_node/internal/conf"

	"github.com/google/uuid"
	"log/slog"
)

type sseToken struct {
	Id int64 `json:"id"`
}

type SSEUsecase struct {
	conf     *conf.Bootstrap
	log      *slog.Logger
	registry repo.ConnectionRegistry
	natsSub  client.Subscriber
	nodeID   string
	tokenGen *jwt.TokenGenerator[sseToken]
	writers  sync.Map
}

func NewSEEUsecase(conf *conf.Bootstrap, logger *slog.Logger, registry repo.ConnectionRegistry, natsSub client.Subscriber, nodeID string) *SSEUsecase {
	return &SSEUsecase{
		conf:     conf,
		log:      logger,
		registry: registry,
		natsSub:  natsSub,
		nodeID:   nodeID,
		tokenGen: jwt.NewTokenGenerator[sseToken](conf.Server.JwtSecret),
	}
}

func (uc *SSEUsecase) Connect(ctx context.Context, token string, w http.ResponseWriter) {
	tokenData, err := uc.tokenGen.Parse(token)
	if err != nil {
		uc.log.Warn(fmt.Sprintf("SSE token validation failed: err=%v", err))
		return
	}
	userID := tokenData.Id
	if userID == 0 {
		uc.log.Warn("SSE token user id is empty")
		return
	}

	connID := uuid.New().String()
	conn := &model.Connection{ID: connID, UserID: userID, CreatedAt: time.Now()}
	if err := uc.registry.AddConnection(userID, conn); err != nil {
		uc.log.Error(fmt.Sprintf("register SSE connection failed: err=%v", err))
		return
	}
	uc.writers.Store(connID, w)
	uc.log.Info(fmt.Sprintf("SSE connection established: user_id=%d conn_id=%s", userID, connID))

	defer func() {
		uc.writers.Delete(connID)
		_ = uc.registry.RemoveConnection(userID, connID)
		uc.log.Info(fmt.Sprintf("SSE connection closed: user_id=%d conn_id=%s", userID, connID))
	}()

	heartbeat := time.NewTicker(30 * time.Second)
	defer heartbeat.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-heartbeat.C:
			if _, err := fmt.Fprint(w, ":keepalive\n\n"); err != nil {
				uc.log.Debug(fmt.Sprintf("SSE heartbeat write failed: err=%v", err))
				return
			}
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
}

func (uc *SSEUsecase) GetConnectionCount() int64 {
	return uc.registry.GetConnectionCount()
}

func (uc *SSEUsecase) HandleNATSMessage(msg *client.Message) error {
	var event struct {
		UserID  int64  `json:"user_id"`
		Type    int32  `json:"type"`
		Payload string `json:"payload"`
	}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		return fmt.Errorf("parse NATS message: %w", err)
	}

	conns := uc.registry.GetConnections(event.UserID)
	if len(conns) == 0 {
		uc.log.Debug(fmt.Sprintf("user has no online connection, message ignored: user_id=%d", event.UserID))
		return nil
	}

	sseData := fmt.Sprintf("event: %d\ndata: %s\n\n", event.Type, event.Payload)
	for _, conn := range conns {
		wVal, ok := uc.writers.Load(conn.ID)
		if !ok {
			continue
		}
		w, ok := wVal.(http.ResponseWriter)
		if !ok {
			continue
		}
		if _, err := fmt.Fprint(w, sseData); err != nil {
			uc.log.Debug(fmt.Sprintf("SSE write failed: user_id=%d conn_id=%s err=%v", event.UserID, conn.ID, err))
			uc.writers.Delete(conn.ID)
			_ = uc.registry.RemoveConnection(event.UserID, conn.ID)
			continue
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
	return nil
}
