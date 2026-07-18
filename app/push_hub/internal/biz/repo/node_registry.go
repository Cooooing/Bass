package repo

import (
	"context"
	"push_hub/internal/biz/model"
)

// NodeRegistry 节点注册表接口，定义节点注册、心跳、用户映射和离线事件存储。
type NodeRegistry interface {
	RegisterNode(ctx context.Context, req *RegisterNodeReq) error
	UpdateHeartbeat(ctx context.Context, req *UpdateHeartbeatReq) error
	GetNode(ctx context.Context, nodeID string) (*model.NodeInfo, error)
	ListNodes(ctx context.Context) ([]*model.NodeInfo, error)
	RemoveNode(ctx context.Context, nodeID string) error
	MapUserToNode(ctx context.Context, req *MapUserToNodeReq) error
	UnmapUserFromNode(ctx context.Context, req *UnmapUserFromNodeReq) error
	GetUserNodes(ctx context.Context, userID int64) ([]string, error)
	GetAllOnlineNodes(ctx context.Context) ([]*model.NodeInfo, error)
	SaveOfflineEvent(ctx context.Context, req *SaveOfflineEventReq) error
	GetOfflineEvents(ctx context.Context, userID int64) ([]*model.PushEvent, error)
	ClearOfflineEvents(ctx context.Context, userID int64) error
}

type RegisterNodeReq struct {
	NodeID  string
	Address string
}

type UpdateHeartbeatReq struct {
	NodeID          string
	ConnectionCount int64
}

type MapUserToNodeReq struct {
	UserID int64
	NodeID string
}

type UnmapUserFromNodeReq struct {
	UserID int64
	NodeID string
}

type SaveOfflineEventReq struct {
	UserID int64
	Event  *model.PushEvent
}
