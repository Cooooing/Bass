package schema

import (
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
)

// Barrier 记录经济域的分布式事务屏障。
type Barrier struct {
	ent.Schema
}

func (Barrier) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.DTMBarrierMixin{}}
}
