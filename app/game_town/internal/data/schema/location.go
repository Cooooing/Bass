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

type Location struct {
	ent.Schema
}

func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "locations"},
		entsql.WithComments(true),
	}
}

func (Location) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.String("code").MaxLen(64),
		field.String("name").MaxLen(128),
		field.Text("description"),
		field.Enum("status").Values(enum.LocationStatusMap.EnumValues()...).Default(string(enum.LocationStatusActive)),
		field.Int64("controlling_faction_id").Optional().Nillable(),
		field.JSON("environment_tags", []string{}).Optional(),
		field.JSON("attributes", map[string]any{}).Optional(),
		field.Bool("accessible").Default(true),
		field.Int64("version").Default(1),
		field.Int32("sort").Default(0),
	}
}

func (Location) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "code").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("world_id", "sort"),
	}
}
