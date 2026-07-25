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

// Player 定义游戏玩家。
type Player struct {
	ent.Schema
}

func (Player) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "players",
		},
		entsql.WithComments(true),
	}
}

func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("玩家唯一名称").MaxRuneLen(64).NotEmpty(),
		field.String("display_name").Comment("显示名称").MaxRuneLen(64).NotEmpty(),
		field.Enum("status").Values(gameenum.PlayerStatusMap.EnumValues()...).Default(string(gameenum.PlayerStatusActive)).Comment("玩家状态"),
	}
}

func (Player) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Player) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique().
			StorageKey("game_town_players_name_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
