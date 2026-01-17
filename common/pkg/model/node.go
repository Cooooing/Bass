package model

import signalv1 "common/api/signal/v1"

// Node 节点
type Node struct {
	ID          int64               `json:"id,omitempty"`           // ID
	OwnerID     *int64              `json:"owner_id,omitempty"`     // 节点拥有者 ID
	Name        string              `json:"name,omitempty"`         // 节点名称
	Description *string             `json:"description,omitempty"`  // 节点描述
	Secret      string              `json:"secret,omitempty"`       // 节点密钥
	CallbackURL string              `json:"callback_url,omitempty"` // 节点回调地址
	Status      signalv1.NodeStatus `json:"status,omitempty"`       // 节点状态

	Online             bool    `json:"online,omitempty"`              // 节点是否在线
	CurrentConnections int64   `json:"current_connections,omitempty"` // 当前连接数
	PingMs             int64   `json:"ping_ms,omitempty"`             // 节点 ping 耗时
	PowCostMs          int64   `json:"pow_cost_ms,omitempty"`         // 节点 PoW 耗时
	Weight             float64 `json:"weight,omitempty"`              // 节点权重
	Score              float64 `json:"score,omitempty"`               // 节点得分
	LastPingTime       *int64  `json:"last_ping_time,omitempty"`      // 最后一次 ping 时间
}
