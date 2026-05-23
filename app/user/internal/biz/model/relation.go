package model

import (
	"time"
	"user/internal/enum"
)

type Relation struct {
	// ID 是关系记录 ID。
	ID int64
	// Type 是关系类型。
	Type enum.RelationType
	// ActorID 是创建关系的账号 ID。
	ActorID int64
	// TargetID 是关系目标账号 ID。
	TargetID int64
	// CreatedAt 是记录创建时间。
	CreatedAt *time.Time
	// UpdatedAt 是记录更新时间。
	UpdatedAt *time.Time
}
