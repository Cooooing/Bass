package schema

import (
	"common/pkg/constant"
	utilent "common/pkg/util/ent"
	gameenum "game_idle/internal/enum"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// CharacterAbility 定义角色能力进度。
type CharacterAbility struct {
	ent.Schema
}

func (CharacterAbility) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "character_abilities",
		},
		entsql.WithComments(true),
	}
}

func (CharacterAbility) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("character_id").Comment("角色 ID"),
		field.Enum("ability_id").
			Values(gameenum.AbilityValues()...).
			Comment("能力编码"),
		field.Int32("level").Comment("当前等级").Positive().Default(1),
		field.Int64("exp").Comment("累计经验").NonNegative().Default(0),
	}
}

func (CharacterAbility) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (CharacterAbility) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("character_id", "ability_id").
			Unique().
			StorageKey("game_idle_character_abilities_character_ability_unique"),
		index.Fields("ability_id", "level").
			StorageKey("game_idle_character_abilities_ability_level_idx"),
	}
}

func (CharacterAbility) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("character", Character.Type).
			Ref("abilities").
			Field("character_id").
			Unique().
			Required(),
	}
}
