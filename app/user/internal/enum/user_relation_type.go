package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type UserRelationType string

const (
	UserRelationTypeFollow UserRelationType = "follow"
	UserRelationTypeBlock  UserRelationType = "block"
)

var UserRelationTypeMap = enum.NewMapping[UserRelationType, v1.UserRelationType](map[UserRelationType]enum.Entry[UserRelationType, v1.UserRelationType]{
	UserRelationTypeFollow: {Proto: v1.UserRelationType_USER_RELATION_TYPE_FOLLOW},
	UserRelationTypeBlock:  {Proto: v1.UserRelationType_USER_RELATION_TYPE_BLOCK},
})
