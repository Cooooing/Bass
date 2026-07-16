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

// Player 保存文字世界内的玩家身份。
type Player struct{ ent.Schema }

func (Player) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "players"}, entsql.WithComments(true)}
}
func (Player) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("玩家登录名").MaxLen(64).NotEmpty(),
		field.String("display_name").Comment("展示名称").MaxLen(64).NotEmpty(),
		field.String("status").Comment("玩家状态").MaxLen(32).NotEmpty(),
	}
}
func (Player) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}, utilent.SoftDeleteMixin{}}
}
func (Player) Indexes() []ent.Index {
	return []ent.Index{index.Fields("name").Unique().StorageKey("game_town_players_name_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL"))}
}
func (Player) Edges() []ent.Edge { return nil }
