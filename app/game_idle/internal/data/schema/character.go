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

// Character 定义挂机游戏角色。
type Character struct {
	ent.Schema
}

func (Character) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: constant.TablePrefixGameIdle.String() + "characters",
		},
		entsql.WithComments(true),
	}
}

func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Immutable().Unique(),
		field.Int64("user_id").Comment("用户 ID"),
		field.Int32("slot").Comment("角色槽位，单用户最多三个").Positive().Default(1),
		field.String("name").Comment("角色名称").MaxRuneLen(32).NotEmpty(),
		field.String("name_key").Comment("小写角色名称，用于大小写不敏感唯一约束").MaxLen(32).NotEmpty(),
		field.Int32("action_queue_capacity").Comment("行动队列容量上限").Positive().Default(3),
		field.Int64("max_offline_seconds").Comment("最大离线收益结算秒数").Positive().Default(28800),
		field.Enum("status").
			Values(gameenum.CharacterStatusValues()...).
			Default(gameenum.CharacterStatusActive.String()).
			Comment("角色状态"),
	}
}

func (Character) Mixin() []ent.Mixin {
	return []ent.Mixin{
		utilent.TimeAuditMixin{},
		utilent.SoftDeleteMixin{},
	}
}

func (Character) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "slot").
			Unique().
			StorageKey("game_idle_characters_user_slot_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("name_key").
			Unique().
			StorageKey("game_idle_characters_name_key_active_unique").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("user_id").
			StorageKey("game_idle_characters_user_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("status").
			StorageKey("game_idle_characters_status_idx").
			Annotations(entsql.IndexWhere("deleted_at IS NULL")),
		index.Fields("deleted_at"),
	}
}

func (Character) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", CharacterItem.Type),
		edge.To("abilities", CharacterAbility.Type),
	}
}
