package model

import (
	v1 "common/api/gen/user/v1"
	"time"
)

type UserRelation struct {
	// 用户关系ID
	ID int64
	// 关系类型 0-follow 1-block
	Type v1.UserRelationType
	// 关系发起者
	ActorID int64
	// 关系目标用户
	TargetID int64
	// 创建时间
	CreatedAt *time.Time
	// 更新时间
	UpdatedAt *time.Time
}
