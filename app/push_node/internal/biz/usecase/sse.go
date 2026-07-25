package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"common/pkg/client"
	"common/pkg/util/jwt"
	"push_node/internal/biz/model"
	"push_node/internal/biz/repo"
	"push_node/internal/config"

	"github.com/google/uuid"
)

type sseToken struct {
	Id int64 `json:"id"`
}

type SSEUsecase struct {
	conf     *config.Bootstrap
	log      *slog.Logger
	registry repo.ConnectionRegistry
	natsSub  client.Subscriber
	nodeID   string
	tokenGen *jwt.TokenGenerator[sseToken]
	writers  sync.Map
}

func NewSEEUsecase(
	conf *config.Bootstrap,
	logger *slog.Logger,
	registry repo.ConnectionRegistry,
	natsSub client.Subscriber,
	nodeID string,
) *SSEUsecase {
	return &SSEUsecase{
		conf:     conf,
		log:      logger,
		registry: registry,
		natsSub:  natsSub,
		nodeID:   nodeID,
		tokenGen: jwt.NewTokenGenerator[sseToken](conf.PushNode.JwtSecret),
	}
}

type ConnectReq struct {
	Token  string
	Writer http.ResponseWriter
}

func (uc *SSEUsecase) Connect(
	ctx context.Context,
	req *ConnectReq,
) error {
	tokenData, err := uc.tokenGen.Parse(req.Token)
	if err != nil {
		uc.log.Warn("SSE token validation failed")
		return nil
	}
	userID := tokenData.Id
	if userID == 0 {
		uc.log.Warn("SSE token user id is empty")
		return nil
	}

	connID := uuid.New().String()
	conn := &model.Connection{
		ID:        connID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	if err := uc.registry.AddConnection(ctx, &repo.AddConnectionReq{
		UserID:     userID,
		Connection: conn,
	}); err != nil {
		uc.log.Error(fmt.Sprintf("register SSE connection failed: err=%v", err))
		return nil
	}
	uc.writers.Store(connID, req.Writer)
	uc.log.Info(fmt.Sprintf("SSE connection established: user_id=%d conn_id=%s", userID, connID))

	defer func() {
		uc.writers.Delete(connID)
		_ = uc.registry.RemoveConnection(ctx, &repo.RemoveConnectionReq{
			UserID:       userID,
			ConnectionID: connID,
		})
		uc.log.Info(fmt.Sprintf("SSE connection closed: user_id=%d conn_id=%s", userID, connID))
	}()

	heartbeat := time.NewTicker(30 * time.Second)
	defer heartbeat.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-heartbeat.C:
			if _, err := fmt.Fprint(req.Writer, ":keepalive\n\n"); err != nil {
				uc.log.Debug(fmt.Sprintf("SSE heartbeat write failed: err=%v", err))
				return nil
			}
			if f, ok := req.Writer.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
}

func (uc *SSEUsecase) GetConnectionCount(
	ctx context.Context,
) (int64, error) {
	return uc.registry.GetConnectionCount(ctx)
}

func (uc *SSEUsecase) HandleNATSMessage(
	ctx context.Context,
	message *client.Message,
) error {
	_ = ctx
	var event struct {
		UserID  int64  `json:"user_id"`
		Type    int32  `json:"type"`
		Payload string `json:"payload"`
	}
	if err := json.Unmarshal(message.Data, &event); err != nil {
		return fmt.Errorf("parse NATS message: %w", err)
	}

	conns, err := uc.registry.GetConnections(ctx, event.UserID)
	if err != nil {
		return err
	}
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
			_ = uc.registry.RemoveConnection(ctx, &repo.RemoveConnectionReq{
				UserID:       event.UserID,
				ConnectionID: conn.ID,
			})
			continue
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
	return nil
}
