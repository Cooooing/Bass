package repo

import (
	"context"
	"push_node/internal/biz/model"
)

// ConnectionRegistry 连接注册表接口，管理本节点上的 SSE 连接。
type ConnectionRegistry interface {
	AddConnection(ctx context.Context, req *AddConnectionReq) error
	RemoveConnection(ctx context.Context, req *RemoveConnectionReq) error
	GetConnections(ctx context.Context, userID int64) ([]*model.Connection, error)
	GetConnectionCount(ctx context.Context) (int64, error)
	GetAllUserIDs(ctx context.Context) ([]int64, error)
}

type AddConnectionReq struct {
	UserID     int64
	Connection *model.Connection
}

type RemoveConnectionReq struct {
	UserID       int64
	ConnectionID string
}
