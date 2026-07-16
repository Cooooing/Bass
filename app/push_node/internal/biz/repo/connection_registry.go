package repo

import (
	"context"
	"push_node/internal/biz/model"
)

// ConnectionRegistry 连接注册表接口，管理本节点上的 SSE 连接。
type ConnectionRegistry interface {
	AddConnection(ctx context.Context, req *AddConnectionReq) (*AddConnectionResponse, error)
	RemoveConnection(ctx context.Context, req *RemoveConnectionReq) (*RemoveConnectionResponse, error)
	GetConnections(ctx context.Context, req *GetConnectionsReq) (*GetConnectionsResponse, error)
	GetConnectionCount(ctx context.Context, req *GetConnectionCountReq) (*GetConnectionCountResponse, error)
	GetAllUserIDs(ctx context.Context, req *GetAllUserIDsReq) (*GetAllUserIDsResponse, error)
}

type AddConnectionReq struct {
	UserID     int64
	Connection *model.Connection
}

type AddConnectionResponse struct{}

type RemoveConnectionReq struct {
	UserID       int64
	ConnectionID string
}

type RemoveConnectionResponse struct{}

type GetConnectionsReq struct {
	UserID int64
}

type GetConnectionsResponse struct {
	Rows []*model.Connection
}

type GetConnectionCountReq struct{}

type GetConnectionCountResponse struct {
	Count int64
}

type GetAllUserIDsReq struct{}

type GetAllUserIDsResponse struct {
	UserIDs []int64
}
