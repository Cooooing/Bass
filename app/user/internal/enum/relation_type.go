package enum

import (
	v1 "common/api/gen/user/v1"
	"common/pkg/enum"
)

type RelationType string

const (
	RelationTypeFollow RelationType = "follow"
	RelationTypeBlock  RelationType = "block"
)

var RelationTypeMap = enum.NewMapping[RelationType, v1.RelationType](map[RelationType]enum.Entry[RelationType, v1.RelationType]{
	RelationTypeFollow: {Proto: v1.RelationType_RELATION_TYPE_FOLLOW},
	RelationTypeBlock:  {Proto: v1.RelationType_RELATION_TYPE_BLOCK},
})
