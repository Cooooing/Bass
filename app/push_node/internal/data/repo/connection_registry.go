package repo

import (
	"fmt"
	"sync"
	"time"

	"push_node/internal/biz/model"
)

// ConnectionRegistryRepo 基于内存的 SSE 连接注册表实现。
type ConnectionRegistryRepo struct {
	mu          sync.RWMutex
	connections map[int64][]*model.Connection // userID -> 连接列表
}

// NewConnectionRegistryRepo 创建内存连接注册表。
func NewConnectionRegistryRepo() *ConnectionRegistryRepo {
	return &ConnectionRegistryRepo{
		connections: make(map[int64][]*model.Connection),
	}
}

// AddConnection 添加用户 SSE 连接。
func (r *ConnectionRegistryRepo) AddConnection(userID int64, conn *model.Connection) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.connections[userID] = append(r.connections[userID], conn)
	return nil
}

// RemoveConnection 移除指定连接。
func (r *ConnectionRegistryRepo) RemoveConnection(userID int64, connID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	conns := r.connections[userID]
	filtered := make([]*model.Connection, 0, len(conns))
	for _, c := range conns {
		if c.ID != connID {
			filtered = append(filtered, c)
		}
	}

	if len(filtered) == 0 {
		delete(r.connections, userID)
	} else {
		r.connections[userID] = filtered
	}

	return nil
}

// GetConnections 获取用户的所有连接。
func (r *ConnectionRegistryRepo) GetConnections(userID int64) []*model.Connection {
	r.mu.RLock()
	defer r.mu.RUnlock()

	conns := r.connections[userID]
	if conns == nil {
		return nil
	}
	// 返回副本，避免外部修改
	result := make([]*model.Connection, len(conns))
	copy(result, conns)
	return result
}

// GetConnectionCount 获取当前总连接数。
func (r *ConnectionRegistryRepo) GetConnectionCount() int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var count int64
	for _, conns := range r.connections {
		count += int64(len(conns))
	}
	return count
}

// GetAllUserIDs 获取所有在线用户 ID 列表。
func (r *ConnectionRegistryRepo) GetAllUserIDs() []int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()

	userIDs := make([]int64, 0, len(r.connections))
	for uid := range r.connections {
		userIDs = append(userIDs, uid)
	}
	return userIDs
}

// 连接保活时间，用于清理长时间无更新的连接。
const connectionKeepAlive = 5 * time.Minute

// 清理过期连接（由心跳检查触发时使用）。
func (r *ConnectionRegistryRepo) CleanupStaleConnections() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	var removed int
	now := time.Now()
	for userID, conns := range r.connections {
		var filtered []*model.Connection
		for _, c := range conns {
			if now.Sub(c.CreatedAt) < connectionKeepAlive {
				filtered = append(filtered, c)
			} else {
				removed++
			}
		}
		if len(filtered) == 0 {
			delete(r.connections, userID)
		} else {
			r.connections[userID] = filtered
		}
	}
	if removed > 0 {
		fmt.Printf("清理过期连接: %d\n", removed)
	}
	return removed
}
