package repo

import (
	"context"
	"push_hub/internal/biz/model"
)

// NodeRegistry 节点注册表接口，定义节点注册、心跳、用户映射和离线事件存储。
type NodeRegistry interface {
	RegisterNode(ctx context.Context, req *RegisterNodeReq) (*RegisterNodeResponse, error)
	UpdateHeartbeat(ctx context.Context, req *UpdateHeartbeatReq) (*UpdateHeartbeatResponse, error)
	GetNode(ctx context.Context, req *GetNodeReq) (*GetNodeResponse, error)
	ListNodes(ctx context.Context, req *ListNodesReq) (*ListNodesResponse, error)
	RemoveNode(ctx context.Context, req *RemoveNodeReq) (*RemoveNodeResponse, error)
	MapUserToNode(ctx context.Context, req *MapUserToNodeReq) (*MapUserToNodeResponse, error)
	UnmapUserFromNode(ctx context.Context, req *UnmapUserFromNodeReq) (*UnmapUserFromNodeResponse, error)
	GetUserNodes(ctx context.Context, req *GetUserNodesReq) (*GetUserNodesResponse, error)
	GetAllOnlineNodes(ctx context.Context, req *GetAllOnlineNodesReq) (*GetAllOnlineNodesResponse, error)
	SaveOfflineEvent(ctx context.Context, req *SaveOfflineEventReq) (*SaveOfflineEventResponse, error)
	GetOfflineEvents(ctx context.Context, req *GetOfflineEventsReq) (*GetOfflineEventsResponse, error)
	ClearOfflineEvents(ctx context.Context, req *ClearOfflineEventsReq) (*ClearOfflineEventsResponse, error)
}

type RegisterNodeReq struct {
	NodeID  string
	Address string
}

type RegisterNodeResponse struct{}

type UpdateHeartbeatReq struct {
	NodeID          string
	ConnectionCount int64
}

type UpdateHeartbeatResponse struct{}

type GetNodeReq struct {
	NodeID string
}

type GetNodeResponse struct {
	Row *model.NodeInfo
}

type ListNodesReq struct{}

type ListNodesResponse struct {
	Rows []*model.NodeInfo
}

type RemoveNodeReq struct {
	NodeID string
}

type RemoveNodeResponse struct{}

type MapUserToNodeReq struct {
	UserID int64
	NodeID string
}

type MapUserToNodeResponse struct{}

type UnmapUserFromNodeReq struct {
	UserID int64
	NodeID string
}

type UnmapUserFromNodeResponse struct{}

type GetUserNodesReq struct {
	UserID int64
}

type GetUserNodesResponse struct {
	NodeIDs []string
}

type GetAllOnlineNodesReq struct{}

type GetAllOnlineNodesResponse struct {
	Rows []*model.NodeInfo
}

type SaveOfflineEventReq struct {
	UserID int64
	Event  *model.PushEvent
}

type SaveOfflineEventResponse struct{}

type GetOfflineEventsReq struct {
	UserID int64
}

type GetOfflineEventsResponse struct {
	Rows []*model.PushEvent
}

type ClearOfflineEventsReq struct {
	UserID int64
}

type ClearOfflineEventsResponse struct{}
