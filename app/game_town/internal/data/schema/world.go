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

// World 保存一个随机生成的文字世界。
type World struct{ ent.Schema }

func (World) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "worlds"}, entsql.WithComments(true)}
}
func (World) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("code").Comment("世界短码").MaxLen(32).NotEmpty(),
		field.String("name").Comment("世界名称").MaxLen(128).NotEmpty(),
		field.Text("description").Comment("创建描述"),
		field.String("scale").Comment("世界规模").MaxLen(32).NotEmpty(),
		field.String("status").Comment("世界状态").MaxLen(32).NotEmpty(),
		field.Int64("creator_player_id").Comment("创建玩家 ID"),
		field.Int64("default_location_id").Comment("默认地点 ID").Nillable().Optional(),
		field.Int64("agent_config_id").Comment("Agent 配置 ID").Nillable().Optional(),
		field.Int64("seed").Comment("生成种子"),
		field.JSON("generation_params", map[string]any{}).Comment("生成参数").Optional(),
		field.Text("generation_summary").Comment("生成摘要").Default(""),
	}
}
func (World) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}, utilent.SoftDeleteMixin{}}
}
func (World) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code").Unique().StorageKey("game_town_worlds_code_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("creator_player_id", "status").StorageKey("game_town_worlds_creator_status_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
func (World) Edges() []ent.Edge { return nil }
