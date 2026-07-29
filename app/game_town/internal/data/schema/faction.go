package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	"game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Faction struct {
	ent.Schema
}

func (Faction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "factions",
		},
		entsql.WithComments(true),
	}
}

func (Faction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Faction) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.String("code").MaxLen(64),
		field.String("name").MaxLen(128),
		field.Text("description"),
		field.Text("public_goal"),
		field.Enum("status").Values(enum.FactionStatusMap.EnumValues()...).Default(enum.FactionStatusActive.String()),
		field.JSON("attributes", map[string]any{}).Optional(),
		field.Int64("version").Default(1),
	}
}

func (Faction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "code").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
