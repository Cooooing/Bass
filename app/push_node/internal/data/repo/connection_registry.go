package repo

import (
	"context"
	"sync"
	"time"

	"push_node/internal/biz/model"
	bizrepo "push_node/internal/biz/repo"
)

const connectionKeepAlive = 5 * time.Minute

type ConnectionRegistryRepo struct {
	mu          sync.RWMutex
	connections map[int64][]*model.Connection
}

func NewConnectionRegistryRepo() *ConnectionRegistryRepo {
	return &ConnectionRegistryRepo{connections: make(map[int64][]*model.Connection)}
}

func (r *ConnectionRegistryRepo) AddConnection(ctx context.Context, req *bizrepo.AddConnectionReq) (*bizrepo.AddConnectionResponse, error) {
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()
	r.connections[req.UserID] = append(r.connections[req.UserID], req.Connection)
	return &bizrepo.AddConnectionResponse{}, nil
}

func (r *ConnectionRegistryRepo) RemoveConnection(ctx context.Context, req *bizrepo.RemoveConnectionReq) (*bizrepo.RemoveConnectionResponse, error) {
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()
	conns := r.connections[req.UserID]
	filtered := make([]*model.Connection, 0, len(conns))
	for _, c := range conns {
		if c.ID != req.ConnectionID {
			filtered = append(filtered, c)
		}
	}
	if len(filtered) == 0 {
		delete(r.connections, req.UserID)
	} else {
		r.connections[req.UserID] = filtered
	}
	return &bizrepo.RemoveConnectionResponse{}, nil
}

func (r *ConnectionRegistryRepo) GetConnections(ctx context.Context, req *bizrepo.GetConnectionsReq) (*bizrepo.GetConnectionsResponse, error) {
	_ = ctx
	r.mu.RLock()
	defer r.mu.RUnlock()
	conns := r.connections[req.UserID]
	if conns == nil {
		return &bizrepo.GetConnectionsResponse{}, nil
	}
	result := make([]*model.Connection, len(conns))
	copy(result, conns)
	return &bizrepo.GetConnectionsResponse{Rows: result}, nil
}

func (r *ConnectionRegistryRepo) GetConnectionCount(ctx context.Context, req *bizrepo.GetConnectionCountReq) (*bizrepo.GetConnectionCountResponse, error) {
	_ = ctx
	_ = req
	r.mu.RLock()
	defer r.mu.RUnlock()
	var count int64
	for _, conns := range r.connections {
		count += int64(len(conns))
	}
	return &bizrepo.GetConnectionCountResponse{Count: count}, nil
}

func (r *ConnectionRegistryRepo) GetAllUserIDs(ctx context.Context, req *bizrepo.GetAllUserIDsReq) (*bizrepo.GetAllUserIDsResponse, error) {
	_ = ctx
	_ = req
	r.mu.RLock()
	defer r.mu.RUnlock()
	userIDs := make([]int64, 0, len(r.connections))
	for uid := range r.connections {
		userIDs = append(userIDs, uid)
	}
	return &bizrepo.GetAllUserIDsResponse{UserIDs: userIDs}, nil
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
