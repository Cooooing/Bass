package repo

import (
	"sync"
	"time"

	"push_node/internal/biz/model"
)

const connectionKeepAlive = 5 * time.Minute

type ConnectionRegistryRepo struct {
	mu          sync.RWMutex
	connections map[int64][]*model.Connection
}

func NewConnectionRegistryRepo() *ConnectionRegistryRepo {
	return &ConnectionRegistryRepo{connections: make(map[int64][]*model.Connection)}
}

func (r *ConnectionRegistryRepo) AddConnection(userID int64, conn *model.Connection) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.connections[userID] = append(r.connections[userID], conn)
	return nil
}

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

func (r *ConnectionRegistryRepo) GetConnections(userID int64) []*model.Connection {
	r.mu.RLock()
	defer r.mu.RUnlock()
	conns := r.connections[userID]
	if conns == nil {
		return nil
	}
	result := make([]*model.Connection, len(conns))
	copy(result, conns)
	return result
}

func (r *ConnectionRegistryRepo) GetConnectionCount() int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var count int64
	for _, conns := range r.connections {
		count += int64(len(conns))
	}
	return count
}

func (r *ConnectionRegistryRepo) GetAllUserIDs() []int64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	userIDs := make([]int64, 0, len(r.connections))
	for uid := range r.connections {
		userIDs = append(userIDs, uid)
	}
	return userIDs
}

func (r *ConnectionRegistryRepo) CleanupStaleConnections() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	var removed int
	now := time.Now()
	for userID, conns := range r.connections {
		filtered := make([]*model.Connection, 0, len(conns))
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
	return removed
}
