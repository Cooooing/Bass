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

type Claim struct {
	ent.Schema
}

func (Claim) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "claims",
		},
		entsql.WithComments(true),
	}
}

func (Claim) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (Claim) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id"),
		field.Int64("origin_event_id").Optional().Nillable(),
		field.Enum("subject_type").Values(enum.EntityTypeMap.EnumValues()...),
		field.Int64("subject_id"),
		field.String("predicate").MaxLen(128),
		field.JSON("object", map[string]any{}),
		field.Enum("truth").Values(enum.ClaimTruthMap.EnumValues()...).Default(string(enum.ClaimTruthUnknown)),
	}
}

func (Claim) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "subject_type", "subject_id"),
	}
}
