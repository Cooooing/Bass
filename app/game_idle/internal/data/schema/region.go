package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Region 定义挂机游戏区域，区域只作为前端展示分类，不承载行动类型语义。
type Region struct {
	ent.Schema
}

func (Region) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "regions",
		},
		entsql.WithComments(true),
	}
}

func (Region) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("区域编码").MaxLen(64).NotEmpty().Immutable().Unique(),
		field.String("name").Comment("区域名称").MaxRuneLen(128).NotEmpty(),
		field.String("description").Comment("区域描述").MaxRuneLen(1024).Default(""),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (Region) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Region) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("enabled", "sort").
			StorageKey("game_idle_regions_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (Region) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("actions", Action.Type),
	}
}
