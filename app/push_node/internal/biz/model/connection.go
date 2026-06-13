package model

import "time"

// Connection 表示一个 SSE 客户端连接。
type Connection struct {
	// ID 连接唯一标识。
	ID string `json:"id"`
	// UserID 关联的用户 ID。
	UserID int64 `json:"user_id"`
	// CreatedAt 连接创建时间。
	CreatedAt time.Time `json:"created_at"`
}
