package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type RelationType string

const (
	RelationTypeFollow RelationType = "follow"
	RelationTypeBlock  RelationType = "block"
)

var RelationTypeMap = commonenum.NewMapping[RelationType, v1.RelationType](map[RelationType]commonenum.Entry[RelationType, v1.RelationType]{
	RelationTypeFollow: {Proto: v1.RelationType_RELATION_TYPE_FOLLOW},
	RelationTypeBlock:  {Proto: v1.RelationType_RELATION_TYPE_BLOCK},
})
