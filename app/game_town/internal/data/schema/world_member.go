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

// WorldMember 保存玩家在某个世界内的位置和角色。
type WorldMember struct{ ent.Schema }

func (WorldMember) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "world_members"}, entsql.WithComments(true)}
}
func (WorldMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("player_id").Comment("玩家 ID"),
		field.Int64("current_location_id").Comment("当前地点 ID"),
		field.String("role").Comment("成员角色").MaxLen(32).NotEmpty(),
		field.Time("joined_at").Comment("加入时间"),
		field.Time("last_seen_at").Comment("最近活跃时间"),
	}
}
func (WorldMember) Mixin() []ent.Mixin { return []ent.Mixin{utilent.TimeAuditMixin{}} }
func (WorldMember) Indexes() []ent.Index {
	return []ent.Index{index.Fields("world_id", "player_id").Unique().StorageKey("game_town_world_members_world_player_unique"), index.Fields("player_id").StorageKey("game_town_world_members_player_idx")}
}
func (WorldMember) Edges() []ent.Edge { return nil }
