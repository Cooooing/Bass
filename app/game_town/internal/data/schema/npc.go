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

type Npc struct {
	ent.Schema
}

func (Npc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "npcs"},
		entsql.WithComments(true),
	}
}

func (Npc) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Npc) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.String("code").MaxLen(64),
		field.String("name").MaxLen(128),
		field.String("role").MaxLen(128),
		field.String("species").MaxLen(64),
		field.Text("personality"),
		field.Text("goal"),
		field.Text("background"),
		field.Int64("current_location_id"),
		field.Text("system_prompt"),
		field.Text("context_summary"),
		field.Enum("life_status").Values(enum.NpcLifeStatusMap.EnumValues()...).Default(string(enum.NpcLifeStatusAlive)),
		field.JSON("state_tags", []string{}).Optional(),
		field.JSON("attributes", map[string]any{}).Optional(),
		field.Time("birth_world_time").Optional().Nillable(),
		field.Time("death_world_time").Optional().Nillable(),
		field.Time("next_decision_at").Optional().Nillable(),
		field.Time("last_planned_at").Optional().Nillable(),
		field.Int64("version").Default(1),
	}
}

func (Npc) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "code").Unique().Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("world_id", "current_location_id"),
	}
}
