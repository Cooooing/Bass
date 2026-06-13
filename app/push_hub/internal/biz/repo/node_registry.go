package repo

import (
	"context"
	"push_hub/internal/biz/model"
)

// NodeRegistry 节点注册表接口，定义节点注册、心跳、用户映射和离线事件存储。
type NodeRegistry interface {
	// RegisterNode 注册推送节点。
	RegisterNode(ctx context.Context, nodeID, address string) error
	// UpdateHeartbeat 更新节点心跳和连接数。
	UpdateHeartbeat(ctx context.Context, nodeID string, connectionCount int64) error
	// GetNode 获取单个节点信息。
	GetNode(ctx context.Context, nodeID string) (*model.NodeInfo, error)
	// ListNodes 列出所有节点。
	ListNodes(ctx context.Context) ([]*model.NodeInfo, error)
	// RemoveNode 移除节点。
	RemoveNode(ctx context.Context, nodeID string) error
	// MapUserToNode 建立用户到节点的映射。
	MapUserToNode(ctx context.Context, userID int64, nodeID string) error
	// UnmapUserFromNode 移除用户到节点的映射。
	UnmapUserFromNode(ctx context.Context, userID int64, nodeID string) error
	// GetUserNodes 获取用户关联的所有节点 ID。
	GetUserNodes(ctx context.Context, userID int64) ([]string, error)
	// GetAllOnlineNodes 获取所有在线节点。
	GetAllOnlineNodes(ctx context.Context) ([]*model.NodeInfo, error)
	// SaveOfflineEvent 保存离线事件。
	SaveOfflineEvent(ctx context.Context, userID int64, event *model.PushEvent) error
	// GetOfflineEvents 获取用户离线事件。
	GetOfflineEvents(ctx context.Context, userID int64) ([]*model.PushEvent, error)
	// ClearOfflineEvents 清除用户离线事件。
	ClearOfflineEvents(ctx context.Context, userID int64) error
}
