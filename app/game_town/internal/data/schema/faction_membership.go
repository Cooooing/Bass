package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	gameenum "game_town/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type FactionMembership struct {
	ent.Schema
}

func (FactionMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "faction_memberships",
		},
		entsql.WithComments(true),
	}
}

func (FactionMembership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (FactionMembership) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.Int64("faction_id"),
		field.Enum("member_type").Values(gameenum.EntityTypeMap.EnumValues()...),
		field.Int64("member_id"),
		field.String("role").MaxRuneLen(128).Default("member"),
		field.JSON("reputation", map[string]float64{}).Optional(),
		field.JSON("tags", []string{}).Optional(),
		field.Time("joined_at"),
		field.Time("left_at").Nillable().Optional(),
	}
}

func (FactionMembership) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "faction_id", "member_type", "member_id").Unique().StorageKey("game_town_faction_memberships_member_unique").Annotations(entsql.IndexWhere("left_at IS NULL")),
		index.Fields("world_id", "member_type", "member_id").StorageKey("game_town_faction_memberships_member_idx"),
	}
}
