package repo

import "push_node/internal/biz/model"

// ConnectionRegistry 连接注册表接口，管理本节点上的 SSE 连接。
type ConnectionRegistry interface {
	// AddConnection 添加用户 SSE 连接。
	AddConnection(userID int64, conn *model.Connection) error
	// RemoveConnection 移除指定连接。
	RemoveConnection(userID int64, connID string) error
	// GetConnections 获取用户的所有连接。
	GetConnections(userID int64) []*model.Connection
	// GetConnectionCount 获取当前总连接数。
	GetConnectionCount() int64
	// GetAllUserIDs 获取所有在线用户 ID 列表。
	GetAllUserIDs() []int64
}
