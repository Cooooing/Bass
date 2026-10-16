package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

const (
	// DTMBarrierTableName 是 DTM 子事务屏障表名。
	DTMBarrierTableName = "barrier"
	// DTMBarrierConstraintName 是 DTM PostgreSQL 方言要求的唯一约束名。
	DTMBarrierConstraintName = "uniq_barrier"
)

// DTMBarrierMixin 定义 DTM 子事务屏障表结构。
type DTMBarrierMixin struct {
	mixin.Schema
}

func (DTMBarrierMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: DTMBarrierTableName},
		entsql.WithComments(true),
	}
}

func (DTMBarrierMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("trans_type").Comment("事务类型"),
		field.String("gid").Comment("全局事务 ID"),
		field.String("branch_id").Comment("分支事务 ID"),
		field.String("op").Comment("分支操作"),
		field.String("barrier_id").Comment("屏障序号"),
		field.String("reason").Comment("操作原因"),
	}
}

func (DTMBarrierMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("gid", "branch_id", "op", "barrier_id").Unique().StorageKey(DTMBarrierConstraintName),
	}
}
