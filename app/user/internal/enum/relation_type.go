package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type RelationType string

const (
	RelationTypeFollow RelationType = "follow" // 关注对方
	RelationTypeBlock  RelationType = "block"  // 拉黑对方
)

var RelationTypeMap = commonenum.NewMapping[RelationType, v1.RelationType](map[RelationType]commonenum.Entry[RelationType, v1.RelationType]{
	RelationTypeFollow: {Proto: v1.RelationType_RELATION_TYPE_FOLLOW},
	RelationTypeBlock:  {Proto: v1.RelationType_RELATION_TYPE_BLOCK},
})

func (e RelationType) String() string {
	return string(e)
}
