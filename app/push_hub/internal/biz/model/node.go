package model

import "time"

// NodeInfo 推送节点信息。
type NodeInfo struct {
	NodeID          string    `json:"node_id"`
	Address         string    `json:"address"`
	ConnectionCount int64     `json:"connection_count"`
	Status          int32     `json:"status"`
	RegisteredAt    time.Time `json:"registered_at"`
	LastHeartbeatAt time.Time `json:"last_heartbeat_at"`
}
