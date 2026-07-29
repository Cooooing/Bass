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

// WorldMember 定义玩家账号在某个世界里的角色状态。
type WorldMember struct {
	ent.Schema
}

func (WorldMember) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameTown.String() + "world_members",
		},
		entsql.WithComments(true),
	}
}

func (WorldMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.Int64("player_id").Comment("玩家账号 ID"),
		field.Int64("current_location_id").Comment("当前地点 ID"),
		field.Enum("role").Values(gameenum.WorldMemberRoleMap.EnumValues()...).Default(gameenum.WorldMemberRoleMember.String()).Comment("世界成员角色"),
		field.String("character_preference").Comment("玩家提交的角色倾向").MaxRuneLen(512).Default(""),
		field.String("character_name").Comment("世界内角色名").MaxRuneLen(128).Default(""),
		field.String("character_background").Comment("世界裁决后的角色背景").MaxRuneLen(1024).Default(""),
		field.String("character_goal").Comment("当前角色目标").MaxRuneLen(512).Default(""),
		field.JSON("character_traits", []string{}).Comment("玩家可见角色标签").Optional(),
		field.Bool("character_ready").Comment("角色是否生成完成").Default(false),
		field.Time("joined_at").Comment("加入时间"),
		field.Time("last_seen_at").Comment("最后活跃时间"),
	}
}

func (WorldMember) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
	}
}

func (WorldMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "player_id").
			Unique().
			StorageKey("game_town_world_members_world_player_unique"),
		index.Fields("player_id").
			StorageKey("game_town_world_members_player_idx"),
	}
}
