package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/bff_bbs/v1/user/enum"
)

type RelationType string

const (
	RelationTypeFollow RelationType = "follow" // follow
	RelationTypeBlock  RelationType = "block"  // block
)

var RelationTypeMap = commonenum.NewMapping[RelationType, v1.RelationType](map[RelationType]commonenum.Entry[RelationType, v1.RelationType]{
	RelationTypeFollow: {Proto: v1.RelationType_RELATION_TYPE_FOLLOW},
	RelationTypeBlock:  {Proto: v1.RelationType_RELATION_TYPE_BLOCK},
})

func (e RelationType) String() string {
	return string(e)
}
