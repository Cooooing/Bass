package usecase

import (
	"common/pkg/util"
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

// sseToken JWT 中携带的用户标识。
type sseToken struct {
	Id int64 `json:"id"`
}

// SSEUsecase SSE 连接管理业务逻辑。
type SSEUsecase struct {
	conf     *conf.Bootstrap
	log      *util.LogHelper
	registry repo.ConnectionRegistry
	natsSub  client.Subscriber
	nodeID   string
	tokenGen *jwt.TokenGenerator[sseToken]
	// writers 存储每个连接的 ResponseWriter，用于写入 SSE 数据。
	writers sync.Map // connID -> http.ResponseWriter
}

// NewSSEUsecase 创建 SSEUsecase 实例。
func NewSEEUsecase(
	conf *conf.Bootstrap,
	logger *slog.Logger,
	registry repo.ConnectionRegistry,
	natsSub client.Subscriber,
	nodeID string,
) *SSEUsecase {
	return &SSEUsecase{
		conf:     conf,
		log:      util.NewLogHelper(logger),
		registry: registry,
		natsSub:  natsSub,
		nodeID:   nodeID,
		tokenGen: jwt.NewTokenGenerator[sseToken](conf.Server.JwtSecret),
	}
}

// Connect 处理 SSE 客户端连接。
// 验证 token → 注册连接 → 持续推送心跳直到断开。
func (uc *SSEUsecase) Connect(ctx context.Context, token string, w http.ResponseWriter) {
	// 验证 JWT token
	tokenData, err := uc.tokenGen.Parse(token)
	if err != nil {
		uc.log.Warnf("SSE token 验证失败: %v", err)
		return
	}
	userID := tokenData.Id
	if userID == 0 {
		uc.log.Warn("SSE token 中用户 ID 为空")
		return
	}

	// 创建连接记录
	connID := uuid.New().String()
	conn := &model.Connection{
		ID:        connID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	// 注册到连接表
	if err := uc.registry.AddConnection(userID, conn); err != nil {
		uc.log.Errorf("注册 SSE 连接失败: %v", err)
		return
	}
	uc.writers.Store(connID, w)
	uc.log.Infof("SSE 连接已建立: user=%d conn=%s", userID, connID)

	defer func() {
		// 连接断开时清理资源
		uc.writers.Delete(connID)
		_ = uc.registry.RemoveConnection(userID, connID)
		uc.log.Infof("SSE 连接已断开: user=%d conn=%s", userID, connID)
	}()

	// 心跳定时器
	heartbeat := time.NewTicker(30 * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-heartbeat.C:
			// 发送 SSE 心跳保活
			if _, err := fmt.Fprint(w, ":keepalive\n\n"); err != nil {
				uc.log.Debugf("SSE 心跳写入失败: %v", err)
				return
			}
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
}

// GetConnectionCount 返回当前节点的 SSE 连接总数。
func (uc *SSEUsecase) GetConnectionCount() int64 {
	return uc.registry.GetConnectionCount()
}

// HandleNATSMessage 处理来自 NATS 的推送消息，写入目标用户的 SSE 连接。
func (uc *SSEUsecase) HandleNATSMessage(msg *client.Message) error {
	// 从消息中解析目标用户 ID
	var event struct {
		UserID  int64  `json:"user_id"`
		Type    int32  `json:"type"`
		Payload string `json:"payload"`
	}
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		return fmt.Errorf("解析 NATS 消息: %w", err)
	}

	// 获取目标用户的所有连接
	conns := uc.registry.GetConnections(event.UserID)
	if len(conns) == 0 {
		uc.log.Debugf("用户 %d 无在线连接，忽略消息", event.UserID)
		return nil
	}

	// 构造 SSE 事件数据
	sseData := fmt.Sprintf("event: %d\ndata: %s\n\n", event.Type, event.Payload)

	// 写入所有连接
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
			uc.log.Debugf("SSE 写入失败: user=%d conn=%s err=%v", event.UserID, conn.ID, err)
			// 写入失败，清理连接
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
