package schema

import (
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
)

// Barrier 定义 DTM 子事务屏障表。
type Barrier struct{ ent.Schema }

func (Barrier) Mixin() []ent.Mixin { return []ent.Mixin{utilent.DTMBarrierMixin{}} }
