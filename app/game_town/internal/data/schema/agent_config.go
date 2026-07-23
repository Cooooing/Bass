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

// AgentConfig 定义模型调用配置。
type AgentConfig struct {
	ent.Schema
}

func (AgentConfig) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: constant.TablePrefixGameTown.String() + "agent_configs"},
		entsql.WithComments(true),
	}
}

func (AgentConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.String("name").Comment("配置名称").MaxRuneLen(64).NotEmpty(),
		field.Enum("provider").Values(gameenum.AgentProviderMap.EnumValues()...).Comment("模型提供方"),
		field.String("base_url").Comment("服务地址").MaxLen(512).NotEmpty(),
		field.String("model").Comment("模型名称").MaxLen(128).NotEmpty(),
		field.String("secret_env").Comment("密钥环境变量名").MaxLen(128).Default(""),
		field.Int32("timeout_seconds").Comment("请求超时秒数").Default(60),
	}
}

func (AgentConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (AgentConfig) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Unique().
			StorageKey("game_town_agent_configs_name_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
	}
}
