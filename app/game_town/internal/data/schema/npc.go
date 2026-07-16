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

// Npc 保存随机生成的 NPC 设定。
type Npc struct{ ent.Schema }

func (Npc) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "npcs"}, entsql.WithComments(true)}
}
func (Npc) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("world_id").Comment("世界 ID"),
		field.String("code").Comment("NPC 编码").MaxLen(64).NotEmpty(),
		field.String("name").Comment("NPC 名称").MaxLen(128).NotEmpty(),
		field.String("role").Comment("角色职责").MaxLen(128).NotEmpty(),
		field.Text("personality").Comment("性格"),
		field.Text("goal").Comment("目标"),
		field.Text("background").Comment("背景"),
		field.Int64("current_location_id").Comment("当前地点 ID"),
		field.String("state").Comment("NPC 状态").MaxLen(32).NotEmpty(),
		field.Text("system_prompt").Comment("系统提示词"),
		field.JSON("generated_profile", map[string]any{}).Comment("生成画像").Optional(),
		field.Bool("enabled").Comment("是否启用").Default(true),
	}
}
func (Npc) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}, utilent.SoftDeleteMixin{}}
}
func (Npc) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("world_id", "code").Unique().StorageKey("game_town_npcs_world_code_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("world_id", "current_location_id").StorageKey("game_town_npcs_world_location_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
func (Npc) Edges() []ent.Edge { return nil }
