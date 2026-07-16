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

// AgentConfig 保存玩家可选的模型访问配置。
type AgentConfig struct{ ent.Schema }

func (AgentConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "agent_configs"}, entsql.WithComments(true)}
}
func (AgentConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("player_id").Comment("玩家 ID"),
		field.String("name").Comment("配置名称").MaxLen(64).NotEmpty(),
		field.String("provider").Comment("供应商").MaxLen(64).NotEmpty(),
		field.String("model").Comment("模型名称").MaxLen(128).NotEmpty(),
		field.String("base_url").Comment("接口地址").MaxLen(512).Default(""),
		field.String("api_key").Comment("接口密钥").MaxLen(1024).Default(""),
		field.Int32("timeout_seconds").Comment("超时秒数").Default(30),
		field.Bool("is_default").Comment("是否默认配置").Default(false),
		field.String("status").Comment("配置状态").MaxLen(32).Default("active"),
	}
}
func (AgentConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{utilent.TimeAuditMixin{}, utilent.SoftDeleteMixin{}}
}
func (AgentConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("player_id", "name").Unique().StorageKey("game_town_agent_configs_player_name_active_unique").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("player_id", "is_default").StorageKey("game_town_agent_configs_player_default_idx").Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
func (AgentConfig) Edges() []ent.Edge { return nil }
