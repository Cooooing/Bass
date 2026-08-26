package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Formula 定义 Lua 数值公式配置。
type Formula struct {
	ent.Schema
}

func (Formula) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "formulas",
		},
		entsql.WithComments(true),
	}
}

func (Formula) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Comment("公式编码").MaxLen(64).NotEmpty().Immutable().Unique(),
		field.String("name").Comment("公式名称").MaxRuneLen(128).NotEmpty(),
		field.Text("expression").Comment("Lua 表达式"),
		field.String("description").Comment("公式描述").MaxRuneLen(1024).Default(""),
		field.Bool("enabled").Comment("是否启用").Default(true),
		field.Int32("sort").Comment("排序值").Default(0),
	}
}

func (Formula) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Formula) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("enabled", "sort").
			StorageKey("game_idle_formulas_enabled_sort_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}
