package model

import (
	"time"
	"user/internal/enum"
)

type UserRelation struct {
	// 用户关系ID
	ID int64
	// 关系类型
	Type enum.UserRelationType
	// 关系发起者
	ActorID int64
	// 关系目标用户
	TargetID int64
	// 创建时间
	CreatedAt *time.Time
	// 更新时间
	UpdatedAt *time.Time
}
